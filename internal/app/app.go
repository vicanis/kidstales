package app

import (
	"context"
	"kidstales/internal/cache/sqlite"
	"kidstales/internal/server"
	"log"

	"golang.org/x/sync/errgroup"
)

type App struct {
	server *server.Server
}

func NewApp(ctx context.Context) (*App, error) {
	return &App{
		server: server.NewServer(ctx),
	}, nil
}

func (a *App) Start(ctx context.Context) error {
	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return a.server.Start(egCtx)
	})

	eg.Go(func() error {
		err := sqlite.StartCleaner(egCtx)
		log.Printf("cleaner failed: %v", err)
		return nil
	})

	return eg.Wait()
}

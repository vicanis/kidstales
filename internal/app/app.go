package app

import (
	"context"
	"kidstales/internal/cache/sqlite"
	"kidstales/internal/server"

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
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return a.server.Start(ctx)
	})

	eg.Go(func() error {
		return sqlite.StartCleaner(ctx)
	})

	return eg.Wait()
}

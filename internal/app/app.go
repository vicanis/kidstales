package app

import (
	"context"
	"kidstales/internal/server"
	"log"
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
	log.Printf("starting server")

	return a.server.Start()
}

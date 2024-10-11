package app

import (
	"context"
	"kidstales/internal/server"
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
	return a.server.Start()
}

package main

import (
	"context"
	"kidstales/internal/app"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("app init failed: %v", err)
	}

	log.Printf("starting application")

	err = app.Start(ctx)
	if err != nil {
		log.Printf("app: %v", err)
	}
}

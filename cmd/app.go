package main

import (
	"context"
	"kidstales/internal/app"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("app init failed: %v", err)
	}

	log.Printf("starting application")

	err = app.Start(ctx)
	if err != nil {
		log.Printf("app error: %v", err)
	}
}

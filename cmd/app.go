package main

import (
	"context"
	"kidstales/internal/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Настройка канала для получения сигналов
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan // Ожидание сигнала
		cancel()     // Отмена контекста
	}()

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

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"go-hot-reload-linux/internal/application"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found - loading from environment")
	}
	cfg := application.NewConfigurationBuilder().
		WithPort(os.Getenv("PORT")).
		Build()
	app := application.NewApp(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err = app.Start(ctx)
	if err == nil {
		fmt.Println("Shutting down server: %w", err)
	}
}

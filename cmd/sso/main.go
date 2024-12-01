package main

import (
	"fmt"
	"gRPC-auth/internal/config"
	"log/slog"
	"os"
)

	const (
		envLocal = "local"
		envDev = "dev"
		envProd = "prod"
	)

func main() {
	// инициализировать объект конфига
	cfg := config.MustLoad()

	fmt.Println(cfg)


	// инициализировать логгер


	// инициализация приложения (app)

	// запустить gRPC-сервер приложения
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}
}
package main

import (
	"fmt"
	"gRPC-auth/internal/config"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// инициализировать объект конфига
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// инициализировать логгер
	log := setupLogger(cfg.Env)

	log.Info("starting app",
		slog.String("env", cfg.Env),
		slog.Any("cfg", cfg),
		slog.Int("port", cfg.GRPC.Port),
	)

	log.Debug("debug message")

	log.Error("error message")

	log.Warn("warning message")

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
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

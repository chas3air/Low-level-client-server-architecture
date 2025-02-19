package main

import (
	"client/internal/app"
	usersservice "client/internal/service"
	"client/internal/storage/server"
	"client/pkg/config"
	"client/pkg/lib/logger"
	"log/slog"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	log.Info("application config", slog.Any("config:", cfg))

	storage := server.New(log, cfg.Host, cfg.Port)
	userService := usersservice.New(log, storage)

	application := app.New(log, userService, cfg.Port, cfg.ExpirationTime)

	application.Start()

	log.Info("Exit")
}

package main

import (
	"log/slog"
	"os"
	"os/signal"
	"server/internal/app"
	"server/pkg/config"
	"server/pkg/lib/logger"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	logger := logger.SetupLogger(cfg.Env)

	logger.Info("starting application", slog.Any("config:", cfg))

	application := app.New(logger, cfg.Grpc.Port)

	go func() {
		application.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	application.GRPCServer.Stop()
	logger.Info("application stopped")
}

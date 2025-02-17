package app

import (
	"log/slog"
	grpcapp "server/internal/app/grpc"
	"server/internal/services/usersmanager"
	"server/internal/storage/mock"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, port int) *App {
	storage := mock.New(log)
	usersmanager := usersmanager.New(log, storage)

	grpcapp := grpcapp.New(log, usersmanager, port)
	return &App{
		GRPCServer: grpcapp,
	}
}

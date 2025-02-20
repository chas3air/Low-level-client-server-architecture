package app

import (
	"log/slog"
	grpcapp "server/internal/app/grpc"
	"server/internal/domain/interfaces"
	"server/internal/services/usersmanager"
	"server/internal/storage/mock"
	psql "server/internal/storage/postgres"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, port int) *App {
	//storage := mock.New(log)
	var storage interfaces.Storage
	storage, err := psql.New("psql", "postgres", "123", 5432, "psql", "Users", log)
	if err != nil {
		storage = mock.New(log)
	}
	usersmanager := usersmanager.New(log, storage)

	grpcapp := grpcapp.New(log, usersmanager, port)
	return &App{
		GRPCServer: grpcapp,
	}
}

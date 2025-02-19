package server

import (
	"client/internal/domain/models"
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type ServerUsersStorage struct {
	log   *slog.Logger
	users map[uuid.UUID]models.User
}

func New(log *slog.Logger) *ServerUsersStorage {
	return &ServerUsersStorage{
		log:   log,
		users: make(map[uuid.UUID]models.User),
	}
}

// GetUsers implements interfaces.ServerUserFetcher.
func (s ServerUsersStorage) GetUsers(context.Context) ([]models.User, error) {
	return nil, nil
}

// GetUserById implements interfaces.ServerUserFetcher.
func (s ServerUsersStorage) GetUserById(context.Context, uuid.UUID) (models.User, error) {
	panic("unimplemented")
}

// GetUserByEmail implements interfaces.ServerUserFetcher.
func (s ServerUsersStorage) GetUserByEmail(context.Context, string) (models.User, error) {
	panic("unimplemented")
}

// Insert implements interfaces.ServerUserFetcher.
func (s ServerUsersStorage) Insert(context.Context, models.User) error {
	panic("unimplemented")
}

// Update implements interfaces.ServerUserFetcher.
func (s ServerUsersStorage) Update(context.Context, uuid.UUID, models.User) error {
	panic("unimplemented")
}

// Delete implements interfaces.ServerUserFetcher.
func (s ServerUsersStorage) Delete(context.Context, uuid.UUID) (models.User, error) {
	panic("unimplemented")
}

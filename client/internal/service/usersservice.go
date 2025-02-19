package usersservice

import (
	storage "client/internal/domain/interfaces/server"
	"client/internal/domain/models"
	"client/pkg/lib/logger/sl"
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

type UserService struct {
	log     *slog.Logger
	storage storage.ServerUserFetcher
}

func New(log *slog.Logger, storage storage.ServerUserFetcher) *UserService {
	return &UserService{
		log:     log,
		storage: storage,
	}
}

// GetUsers implements server.ServerUserFetcher.
func (u *UserService) GetUsers(ctx context.Context) ([]models.User, error) {
	const op = "service.getUsers"
	log := u.log.With(
		slog.String("op", op),
	)

	users, err := u.storage.GetUsers(ctx)
	if err != nil {
		log.Warn("failed to fetch users", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return users, nil
}

// GetUserById implements server.ServerUserFetcher.
func (u *UserService) GetUserById(context.Context, uuid.UUID) (models.User, error) {
	panic("unimplemented")
}

// GetUserByEmail implements server.ServerUserFetcher.
func (u *UserService) GetUserByEmail(context.Context, string) (models.User, error) {
	panic("unimplemented")
}

// Insert implements server.ServerUserFetcher.
func (u *UserService) Insert(context.Context, models.User) error {
	panic("unimplemented")
}

// Update implements server.ServerUserFetcher.
func (u *UserService) Update(context.Context, uuid.UUID, models.User) error {
	panic("unimplemented")
}

// Delete implements server.ServerUserFetcher.
func (u *UserService) Delete(context.Context, uuid.UUID) (models.User, error) {
	panic("unimplemented")
}

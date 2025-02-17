package usersmanager

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"server/internal/domain/interfaces"
	"server/internal/domain/models"
	"server/internal/storage"
	"server/pkg/lib/logger/sl"

	"github.com/google/uuid"
)

type UsersManager struct {
	log     *slog.Logger
	storage interfaces.Storage
}

var ErrInvalidCredentials = errors.New("invalid credentials")

func New(log *slog.Logger, storage interfaces.Storage) *UsersManager {
	return &UsersManager{
		log:     log,
		storage: storage,
	}
}

// GetUsers implements interfaces.Storage.
func (u *UsersManager) GetUsers(ctx context.Context) ([]models.User, error) {
	const op = "services.usersmanager.getUsers"
	log := u.log.With(slog.String("operation", op))

	users, err := u.storage.GetUsers(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			log.Warn("%s: %w", op, ErrInvalidCredentials)

			return nil, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		log.Error("Failed to retrieve users", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return users, nil
}

// GetUserById implements interfaces.Storage.
func (u *UsersManager) GetUserById(ctx context.Context, id uuid.UUID) (models.User, error) {
	const op = "services.usersmanager.getUserById"
	log := u.log.With(slog.String("operation", op))

	user, err := u.storage.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			log.Warn("%s: %w", op, ErrInvalidCredentials)

			return models.User{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		log.Error("Failed to retrieve user by id", sl.Err(err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

// GetUsersByEmail implements interfaces.Storage.
func (u *UsersManager) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	const op = "services.usersmanager.getUserByEmail"
	log := u.log.With(slog.String("operation", op))

	user, err := u.storage.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			log.Warn("%s: %w", op, ErrInvalidCredentials)

			return models.User{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		log.Error("Failed to retrieve user by email", sl.Err(err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

// Insert implements interfaces.Storage.
func (u *UsersManager) Insert(ctx context.Context, user models.User) error {
	const op = "services.usersmanager.insert"
	log := u.log.With(slog.String("operation", op))

	err := u.storage.Insert(ctx, user)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			log.Warn("%s: %w", op, ErrInvalidCredentials)

			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		log.Error("Failed to insert user", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Update implements interfaces.Storage.
func (u *UsersManager) Update(ctx context.Context, id uuid.UUID, user models.User) error {
	const op = "services.usermanager.update"
	log := u.log.With(slog.String("op", op))

	err := u.storage.Update(ctx, id, user)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			log.Warn("%s: %w", op, ErrInvalidCredentials)

			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		log.Error("Failed to update user:", sl.Err(err))
		return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	return nil
}

// Delete implements interfaces.Storage.
func (u *UsersManager) Delete(ctx context.Context, id uuid.UUID) (models.User, error) {
	const op = "services.usermanager.delete"
	log := u.log.With(slog.String("op", op))

	user, err := u.storage.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			log.Warn("%s: %w", op, ErrInvalidCredentials)

			return models.User{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		log.Error("Failed to update user:", sl.Err(err))
		return models.User{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	return user, nil
}

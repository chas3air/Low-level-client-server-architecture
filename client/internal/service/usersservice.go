package usersservice

import (
	storage "client/internal/domain/interfaces/server"
	"client/internal/domain/models"
	storage_errors "client/internal/storage"
	"client/pkg/lib/logger/sl"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

type UserService struct {
	log     *slog.Logger
	storage storage.ServerUserFetcher
}

var ErrInvalidCredentials = errors.New("invalid credentials")

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
func (u *UserService) GetUserById(ctx context.Context, uid uuid.UUID) (models.User, error) {
	const op = "services.userManager.GetUserById"
	log := u.log.With(slog.String("operation", op))

	user, err := u.storage.GetUserById(ctx, uid)
	if err != nil {
		if errors.Is(err, storage_errors.ErrUserNotFound) {
			log.Warn("User not found", sl.Err(err), slog.String("userId", uid.String()), slog.String("error", err.Error()))

			return models.User{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		log.Error("Failed to retrieve user by id", sl.Err(err), slog.String("userId", uid.String()), slog.String("error", err.Error()))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Successfully retrieved user", slog.Any("additional info", user), slog.String("error", "nil"))
	return user, nil
}

// GetUserByEmail implements server.ServerUserFetcher.
func (u *UserService) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	const op = "services.userManager.GetUserByEmail"
	log := u.log.With(slog.String("operation", op))

	user, err := u.storage.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, storage_errors.ErrUserNotFound) {
			log.Warn("User not found", sl.Err(err), slog.String("email", email), slog.String("error", err.Error()))

			return models.User{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		log.Error("Failed to retrieve user by email", sl.Err(err), slog.String("email", email), slog.String("error", err.Error()))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Successfully retrieved user", slog.Any("additional info", user), slog.String("error", "nil"))
	return user, nil
}

// Insert implements server.ServerUserFetcher.
func (u *UserService) Insert(ctx context.Context, user models.User) error {
	const op = "services.userManager.Insert"
	log := u.log.With(slog.String("operation", op))

	err := u.storage.Insert(ctx, user)
	if err != nil {
		if errors.Is(err, storage_errors.ErrUserExists) {
			log.Warn("User already exists", sl.Err(err), slog.Any("additional info", user), slog.String("error", err.Error()))

			return fmt.Errorf("%s: %s", op, "user already exists")
		}

		log.Error("Failed to insert user", sl.Err(err), slog.Any("additional info", user), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("User inserted successfully", slog.Any("additional info", []map[string]interface{}{
		{"user": user},
	}), slog.String("error", "nil"))
	return nil
}

// Update implements server.ServerUserFetcher.
func (u *UserService) Update(ctx context.Context, uid uuid.UUID, user models.User) error {
	const op = "services.userManager.Update"
	log := u.log.With(slog.String("operation", op))

	err := u.storage.Update(ctx, uid, user)
	if err != nil {
		if errors.Is(err, storage_errors.ErrUserNotFound) {
			log.Warn("User not found", sl.Err(err), slog.String("userId", uid.String()), slog.String("error", err.Error()))

			return fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		log.Error("Failed to update user", sl.Err(err), slog.String("userId", uid.String()), slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("User updated successfully", slog.Any("additional info", []map[string]interface{}{
		{"user": user},
	}), slog.String("error", "nil"))
	return nil
}

// Delete implements server.ServerUserFetcher.
func (u *UserService) Delete(ctx context.Context, uid uuid.UUID) (models.User, error) {
	const op = "services.userManager.Delete"
	log := u.log.With(slog.String("operation", op))

	user, err := u.storage.Delete(ctx, uid)
	if err != nil {
		if errors.Is(err, storage_errors.ErrUserNotFound) {
			log.Warn("User not found", sl.Err(err), slog.String("userId", uid.String()), slog.String("error", err.Error()))

			return models.User{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		log.Error("Failed to delete user by id", sl.Err(err), slog.String("userId", uid.String()), slog.String("error", err.Error()))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("User deleted successfully", slog.Any("additional info", []map[string]interface{}{
		{"user": user},
	}), slog.String("error", "nil"))
	return user, nil
}

package usersmanager

import (
	"log/slog"
	"server/internal/domain/interfaces"
	"server/internal/domain/models"

	"github.com/google/uuid"
)

type UsersManager struct {
	log     *slog.Logger
	storage interfaces.Storage
}

func New(log *slog.Logger, storage interfaces.Storage) *UsersManager {
	return &UsersManager{
		log:     log,
		storage: storage,
	}
}

// GetUsers implements interfaces.Storage.
func (u *UsersManager) GetUsers() ([]models.User, error) {
	return u.storage.GetUsers()
}

// GetUserById implements interfaces.Storage.
func (u *UsersManager) GetUserById(id uuid.UUID) (models.User, error) {
	return u.storage.GetUserById(id)
}

// GetUsersByEmail implements interfaces.Storage.
func (u *UsersManager) GetUsersByEmail(email string) (models.User, error) {
	return u.storage.GetUsersByEmail(email)
}

// Insert implements interfaces.Storage.
func (u *UsersManager) Insert(user models.User) (models.User, error) {
	return u.storage.Insert(user)
}

// Update implements interfaces.Storage.
func (u *UsersManager) Update(id uuid.UUID, user models.User) (models.User, error) {
	return u.storage.Update(id, user)
}

// Delete implements interfaces.Storage.
func (u *UsersManager) Delete(id uuid.UUID) (models.User, error) {
	return u.storage.Delete(id)
}

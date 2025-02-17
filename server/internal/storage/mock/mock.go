package mock

import (
	"log/slog"
	"server/internal/domain/models"

	"github.com/google/uuid"
)

type MockStorage struct {
	users []models.User
	log   *slog.Logger
}

func New(log *slog.Logger) *MockStorage {
	return &MockStorage{
		users: make([]models.User, 0, 5),
		log:   log,
	}
}

// GetUsers implements interfaces.Storage.
func (m *MockStorage) GetUsers() ([]models.User, error) {
	panic("unimplemented")
}

// GetUserById implements interfaces.Storage.
func (m *MockStorage) GetUserById(id uuid.UUID) (models.User, error) {
	panic("unimplemented")
}

// GetUsersByEmail implements interfaces.Storage.
func (m *MockStorage) GetUserByEmail(email string) (models.User, error) {
	panic("unimplemented")
}

// Insert implements interfaces.Storage.
func (m *MockStorage) Insert(user models.User) error {
	panic("unimplemented")
}

// Update implements interfaces.Storage.
func (m *MockStorage) Update(id uuid.UUID, user models.User) error {
	panic("unimplemented")
}

// Delete implements interfaces.Storage.
func (m *MockStorage) Delete(id uuid.UUID) (models.User, error) {
	panic("unimplemented")
}

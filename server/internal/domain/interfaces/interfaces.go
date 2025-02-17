package interfaces

import (
	"server/internal/domain/models"

	"github.com/google/uuid"
)

type Storage interface {
	GetUsers() ([]models.User, error)
	GetUserById(id uuid.UUID) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	Insert(user models.User) error
	Update(id uuid.UUID, user models.User) error
	Delete(id uuid.UUID) (models.User, error)
}

type UsersManager interface {
	GetUsers() ([]models.User, error)
	GetUserById(id uuid.UUID) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	Insert(user models.User) error
	Update(id uuid.UUID, user models.User) error
	Delete(id uuid.UUID) (models.User, error)
}

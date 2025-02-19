package userservice

import (
	"client/internal/domain/models"
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	GetUsers(context.Context) ([]models.User, error)
	GetUserById(context.Context, uuid.UUID) (models.User, error)
	GetUserByEmail(context.Context, string) (models.User, error)
	Insert(context.Context, models.User) error
	Update(context.Context, uuid.UUID, models.User) error
	Delete(context.Context, uuid.UUID) (models.User, error)
}

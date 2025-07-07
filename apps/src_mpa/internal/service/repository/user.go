package repository

import (
	"src/internal/entities"

	"github.com/google/uuid"
)

type IUserRepository interface {
	Update(user *entities.User) error
	Create(user *entities.User) error
	GetUserByID(userID uuid.UUID) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
}

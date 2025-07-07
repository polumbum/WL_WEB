package repository

import (
	"src/internal/domain"

	"github.com/google/uuid"
)

type IUserRepository interface {
	Update(user *domain.User) error
	Create(user *domain.User) error
	GetUserByID(userID uuid.UUID) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	Delete(id uuid.UUID) error
}

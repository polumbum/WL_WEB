package domain

import (
	"src/internal/constants"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Email    string
	Password string
	Role     constants.UserRole
	RoleID   uuid.UUID
}

package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Email    string `gorm:"type:TEXT;check:check_valid_email, (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Z|a-z]{2,}$')"`
	Password string
	Role     string
	RoleID   uuid.UUID
}

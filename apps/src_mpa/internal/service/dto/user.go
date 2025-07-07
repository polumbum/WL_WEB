package dto

import (
	"src/internal/constants"
	"src/internal/entities"

	"github.com/google/uuid"
)

type LoginUserReq struct {
	Email    string
	Password string
}

type RegisterUserReq struct {
	Email    string
	Password string
	Role     constants.UserRole
	RoleID   uuid.UUID
}

func (req *RegisterUserReq) Copy(user *entities.User) {
	user.Email = req.Email
	user.Password = req.Password
	user.Role = req.Role
	user.RoleID = req.RoleID
}

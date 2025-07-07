package dto

import (
	"net/http"

	"github.com/google/uuid"
)

// User model info
// @Description Пользователь
type User struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	//Password string `json:"password"`
	Role   string    `json:"role"`
	RoleID uuid.UUID `json:"role_id"`
}

// User model info
// @Description Пользователь
type UserReq struct {
	*User
} // @name UserReq

// Login
// @Description Вход в аккаунт
type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
} // @name LoginUserReq

// Update
// @Description Обновить пароль и почту
type UpdateUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
} // @name UpdateUserReq

// Register
// @Description Создание пользователя
type RegisterUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
} // @name RegisterUserReq

// User
// @Description Пользователь
type UserResp struct {
	*User
} // @name UserResp

func NewUserResp(u *User) *UserResp {
	resp := &UserResp{User: u}

	return resp
}

func (rd *UserResp) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

type LoginResp struct {
	Token string `json:"token"`
	//ExpiresIn time.Duration `json:"expires_in"`
	*User `json:"user"`
} // @name LoginResp

func NewLoginResp(u *User, token string) *LoginResp {
	resp := &LoginResp{
		User:  u,
		Token: token,
	}

	return resp
}

func (rd *LoginResp) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

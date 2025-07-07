package service

import (
	"errors"
	"src/internal/entities"
	"src/internal/service/dto"
	"src/internal/service/repository"

	"github.com/google/uuid"
)

type IUserService interface {
	Login(req *dto.LoginUserReq) (*entities.User, error)
	Register(req *dto.RegisterUserReq) (*entities.User, error)
	GetUserByID(id uuid.UUID) (*entities.User, error)
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Login(req *dto.LoginUserReq) (*entities.User, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	if user.Password != req.Password {
		return nil, ErrWrongPassword
	}

	return user, nil
}

func (s *UserService) Register(req *dto.RegisterUserReq) (*entities.User, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	if user != nil {
		return nil, ErrAlreadyRegistered
	}

	var newUser entities.User
	req.Copy(&newUser)

	err = s.repo.Create(&newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (s *UserService) GetUserByID(id uuid.UUID) (*entities.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

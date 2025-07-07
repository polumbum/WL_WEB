package service

import (
	"errors"
	"log"
	dataaccess "src/internal/data_access"
	"src/internal/domain"
	"src/internal/service/repository"

	"github.com/google/uuid"
)

type IUserService interface {
	Login(u *domain.User) (*domain.User, error)
	Register(u *domain.User) (*domain.User, error)
	GetUserByID(id uuid.UUID) (*domain.User, error)
	Delete(id uuid.UUID) error
	Update(u *domain.User) (*domain.User, error)
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Update(u *domain.User) (*domain.User, error) {
	if u == nil {
		return nil, ErrNilRef
	}
	err := s.repo.Update(u)
	if err != nil {
		log.Println("upd", err)
		return nil, err
	}

	return u, nil
}

func (s *UserService) Login(u *domain.User) (*domain.User, error) {
	user, err := s.repo.GetUserByEmail(u.Email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if user == nil {
		log.Println(err)
		return nil, ErrUserNotFound
	}

	if user.Password != u.Password {
		log.Println(err)
		return nil, ErrWrongPassword
	}

	return user, nil
}

func (s *UserService) Register(u *domain.User) (*domain.User, error) {
	user, err := s.repo.GetUserByEmail(u.Email)
	if err != nil && !errors.Is(err, dataaccess.ErrNotFound) {
		log.Println(err)
		return nil, err
	}

	if user != nil {
		log.Println(err)
		return nil, ErrAlreadyRegistered
	}

	u.ID = uuid.New()

	err = s.repo.Create(u)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return u, nil
}

func (s *UserService) GetUserByID(id uuid.UUID) (*domain.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

func (s *UserService) Delete(id uuid.UUID) error {
	err := s.repo.Delete(id)
	log.Println(err)
	return err
}

package dataaccess

import (
	"errors"
	"src/internal/entities"
	"src/internal/service"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
type IUserRepository interface {
	Update(user *entities.User) error
	Create(user *entities.User) error
	GetUserByID(userID uuid.UUID) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
}
*/

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Update(user *entities.User) error {
	err := r.db.Save(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return service.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *UserRepository) Create(user *entities.User) error {
	user.ID = uuid.New()
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByID(userID uuid.UUID) (*entities.User, error) {
	user := &entities.User{}
	err := r.db.Where("id = ?", userID).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	err := r.db.Where("email = ?", email).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, err
	}
	return user, nil
}

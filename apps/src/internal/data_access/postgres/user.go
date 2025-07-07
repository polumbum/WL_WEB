package dataaccess

import (
	"errors"
	"log"
	"src/internal/converters"
	dataaccess "src/internal/data_access"
	"src/internal/data_access/models"
	"src/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Update(user *domain.User) error {
	model, err := converters.NewUserConverter().ToModel(user)
	if err != nil {
		return err
	}

	err = r.db.Save(model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dataaccess.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *UserRepository) Create(user *domain.User) error {
	model, err := converters.NewUserConverter().ToModel(user)
	if err != nil {
		return err
	}

	model.ID = uuid.New()
	err = r.db.Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByID(userID uuid.UUID) (
	*domain.User,
	error,
) {
	user := &models.User{}
	err := r.db.Where("id = ?", userID).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}

	domain, err := converters.NewUserConverter().ToDomain(user)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (r *UserRepository) GetUserByEmail(email string) (
	*domain.User,
	error,
) {
	user := &models.User{}
	err := r.db.Where("email = ?", email).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}

	domain, err := converters.NewUserConverter().ToDomain(user)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	err := r.db.Delete(&models.User{}, id).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

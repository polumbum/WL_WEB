package dataaccess

import (
	"errors"
	"src/internal/converters"
	dataaccess "src/internal/data_access"
	"src/internal/data_access/models"
	"src/internal/domain"
	"src/internal/service"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompAccessRepository struct {
	db *gorm.DB
}

func NewCompAccessRepository(db *gorm.DB) *CompAccessRepository {
	return &CompAccessRepository{
		db: db,
	}
}

func (r *CompAccessRepository) Update(ad *domain.CompAccess) (
	*domain.CompAccess,
	error,
) {
	model, err := converters.NewAccessConverter().ToModel(ad)
	if err != nil {
		return nil, err
	}

	err = r.db.Save(model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}
	return ad, nil
}

func (r *CompAccessRepository) Create(ad *domain.CompAccess) error {
	model, err := converters.NewAccessConverter().ToModel(ad)
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

func (r *CompAccessRepository) GetAccessBySmID(sportsmanID uuid.UUID) (
	*domain.CompAccess,
	error,
) {
	ca := &models.CompAccess{}
	err := r.db.Where("sportsman_id = ?", sportsmanID).First(ca).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrAccessNotFound
		}
		return nil, err
	}

	domain, err := converters.NewAccessConverter().ToDomain(ca)
	if err != nil {
		return nil, err
	}
	return domain, nil
}

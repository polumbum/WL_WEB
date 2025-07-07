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

type ADopingRepository struct {
	db *gorm.DB
}

func NewADopingRepository(db *gorm.DB) *ADopingRepository {
	return &ADopingRepository{
		db: db,
	}
}

func (r *ADopingRepository) Update(ad *domain.Antidoping) (
	*domain.Antidoping,
	error,
) {
	adModel, err := converters.NewADopingConverter().ToModel(ad)
	if err != nil {
		return nil, err
	}
	err = r.db.Save(adModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}
	return ad, nil
}

func (r *ADopingRepository) Create(ad *domain.Antidoping) error {
	adModel, err := converters.NewADopingConverter().ToModel(ad)
	if err != nil {
		return err
	}
	adModel.ID = uuid.New()
	err = r.db.Create(adModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ADopingRepository) GetADopingBySmID(sportsmanID uuid.UUID) (
	*domain.Antidoping,
	error,
) {
	ad := &models.Antidoping{}
	err := r.db.Where("sportsman_id = ?", sportsmanID).First(ad).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrADopingNotFound
		}
		return nil, err
	}
	adDomain, err := converters.NewADopingConverter().
		ToDomain(ad)
	if err != nil {
		return nil, err
	}

	return adDomain, nil
}

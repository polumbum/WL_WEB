package dataaccess

import (
	"errors"
	"src/internal/entities"
	"src/internal/service"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
type IADopingRepository interface {
	Update(ad *entities.Antidoping) error
	Create(ad *entities.Antidoping) error
	GetADopingBySmID(sportsmanID uuid.UUID) (*entities.Antidoping, error)
}
*/

type ADopingRepository struct {
	db *gorm.DB
}

func NewADopingRepository(db *gorm.DB) *ADopingRepository {
	return &ADopingRepository{
		db: db,
	}
}

func (r *ADopingRepository) Update(ad *entities.Antidoping) error {
	err := r.db.Save(ad).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return service.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *ADopingRepository) Create(ad *entities.Antidoping) error {
	ad.ID = uuid.New()
	err := r.db.Create(ad).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ADopingRepository) GetADopingBySmID(sportsmanID uuid.UUID) (*entities.Antidoping, error) {
	ad := &entities.Antidoping{}
	err := r.db.Where("sportsman_id = ?", sportsmanID).First(ad).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrADopingNotFound
		}
		return nil, err
	}
	return ad, nil
}

package dataaccess

import (
	"errors"
	"src/internal/entities"
	"src/internal/service"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
type ICompAccessRepository interface {
	Update(ca *entities.CompAccess) error
	Create(ca *entities.CompAccess) error
	GetAccessBySmID(sportsmanID uuid.UUID) (*entities.CompAccess, error)
}
*/

type CompAccessRepository struct {
	db *gorm.DB
}

func NewCompAccessRepository(db *gorm.DB) *CompAccessRepository {
	return &CompAccessRepository{
		db: db,
	}
}

func (r *CompAccessRepository) Update(ad *entities.CompAccess) error {
	err := r.db.Save(ad).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return service.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *CompAccessRepository) Create(ad *entities.CompAccess) error {
	ad.ID = uuid.New()
	err := r.db.Create(ad).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CompAccessRepository) GetAccessBySmID(sportsmanID uuid.UUID) (*entities.CompAccess,
	error) {
	ca := &entities.CompAccess{}
	err := r.db.Where("sportsman_id = ?", sportsmanID).First(ca).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrAccessNotFound
		}
		return nil, err
	}
	return ca, nil
}

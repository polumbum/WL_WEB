package dataaccess

import (
	"errors"
	"src/internal/entities"
	"src/internal/service"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SportsmanRepository struct {
	db *gorm.DB
}

func NewSportsmanRepository(db *gorm.DB) *SportsmanRepository {
	return &SportsmanRepository{
		db: db,
	}
}

func (r *SportsmanRepository) Update(sportsman *entities.Sportsman) error {
	err := r.db.Save(sportsman).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return service.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *SportsmanRepository) Create(sportsman *entities.Sportsman) error {
	sportsman.ID = uuid.New()
	err := r.db.Create(sportsman).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SportsmanRepository) ListSportsmen() ([]*entities.Sportsman, error) {
	sportsmen := []*entities.Sportsman{}
	err := r.db.Find(&sportsmen).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, err
	}
	return sportsmen, nil
}

func (r *SportsmanRepository) GetSportsmanByID(sportsmanID uuid.UUID) (*entities.Sportsman, error) {
	sportsman := &entities.Sportsman{}
	err := r.db.Where("id = ?", sportsmanID).First(sportsman).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, err
	}
	return sportsman, nil
}

func (r *SportsmanRepository) ListResults(sportsmanID uuid.UUID) ([]*entities.Result, error) {
	results := []*entities.Result{}
	err := r.db.Where("sportsman_id = ?", sportsmanID).Find(&results).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, err
	}
	return results, nil
}

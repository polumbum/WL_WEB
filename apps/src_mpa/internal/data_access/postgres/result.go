package dataaccess

import (
	"errors"
	"src/internal/entities"
	"src/internal/service"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
type IResultRepository interface {
	Update(result *entities.Result) error
	Create(result *entities.Result) error
	ListResults() ([]*entities.Result, error)
	ListSportsmanResults(sportsmanID uuid.UUID) ([]*entities.Result, error)
	GetResultByID(resultID uuid.UUID) (*entities.Result, error)
}
*/

type ResultRepository struct {
	db *gorm.DB
}

func NewResultRepository(db *gorm.DB) *ResultRepository {
	return &ResultRepository{
		db: db,
	}
}

func (r *ResultRepository) Update(result *entities.Result) error {
	//err := r.db.Save(result).Error
	err := r.db.Model(&entities.Result{}).
		Where("sportsman_id = ? AND competition_id = ?", result.SportsmanID, result.CompetitionID).
		Updates(map[string]interface{}{
			"weight_category": result.WeightCategory,
			"snatch":          result.Snatch,
			"clean_and_jerk":  result.CleanAndJerk,
			"place":           result.Place,
		}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return service.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *ResultRepository) Create(result *entities.Result) error {
	err := r.db.Create(result).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ResultRepository) ListResults() ([]*entities.Result, error) {
	results := []*entities.Result{}
	err := r.db.Find(&results).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, err
	}
	return results, nil
}

func (r *ResultRepository) GetResultByID(smID, compID uuid.UUID) (*entities.Result, error) {
	result := &entities.Result{}
	err := r.db.Where("sportsman_id = ? AND competition_id = ?", smID, compID).First(result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, err
	}
	return result, nil
}

func (r *ResultRepository) ListSportsmanResults(sportsmanID uuid.UUID) ([]*entities.Result, error) {
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

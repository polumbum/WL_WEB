package dataaccess

import (
	"errors"
	"src/internal/converters"
	dataaccess "src/internal/data_access"
	"src/internal/data_access/models"
	"src/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResultRepository struct {
	db *gorm.DB
}

func NewResultRepository(db *gorm.DB) *ResultRepository {
	return &ResultRepository{
		db: db,
	}
}

func (r *ResultRepository) Update(result *domain.Result) error {
	//err := r.db.Save(result).Error
	err := r.db.Model(&models.Result{}).
		Where("sportsman_id = ? AND competition_id = ?", result.SportsmanID, result.CompetitionID).
		Updates(map[string]interface{}{
			"weight_category": result.WeightCategory,
			"snatch":          result.Snatch,
			"clean_and_jerk":  result.CleanAndJerk,
			"place":           result.Place,
		}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dataaccess.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *ResultRepository) Create(result *domain.Result) error {
	model, err := converters.NewResultConverter().ToModel(result)
	if err != nil {
		return err
	}
	err = r.db.Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ResultRepository) ListResults() ([]*domain.Result, error) {
	results := []*models.Result{}
	err := r.db.Find(&results).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}

	domain := []*domain.Result{}
	conv := converters.NewResultConverter()
	for _, item := range results {
		r, err := conv.ToDomain(item)
		if err != nil {
			return nil, err
		}
		domain = append(domain, r)
	}
	return domain, nil
}

func (r *ResultRepository) GetResultByID(smID, compID uuid.UUID) (
	*domain.Result,
	error,
) {
	result := &models.Result{}
	err := r.db.Where("sportsman_id = ? AND competition_id = ?", smID, compID).First(result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}

	domain, err := converters.NewResultConverter().ToDomain(result)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (r *ResultRepository) ListSportsmanResults(sportsmanID uuid.UUID) (
	[]*domain.Result,
	error,
) {
	results := []*models.Result{}
	err := r.db.Where("sportsman_id = ?", sportsmanID).Find(&results).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}

	domain := []*domain.Result{}
	conv := converters.NewResultConverter()
	for _, item := range results {
		r, err := conv.ToDomain(item)
		if err != nil {
			return nil, err
		}
		domain = append(domain, r)
	}
	return domain, nil
}

func (r *ResultRepository) ListCompResults(compID uuid.UUID) (
	[]*domain.Result,
	error,
) {
	results := []*models.Result{}
	err := r.db.Where("competition_id = ?", compID).Find(&results).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}

	domain := []*domain.Result{}
	conv := converters.NewResultConverter()
	for _, item := range results {
		r, err := conv.ToDomain(item)
		if err != nil {
			return nil, err
		}
		domain = append(domain, r)
	}
	return domain, nil
}

func (r *ResultRepository) ListCoachResults(coachID uuid.UUID) (
	[]*domain.Result,
	error,
) {
	results := []*models.Result{}

	err := r.db.Table("results").
		Select("results.*").
		Joins("JOIN sportsmen_coaches ON sportsmen_coaches.sportsman_id = results.sportsman_id").
		Where("sportsmen_coaches.coach_id = ?", coachID).
		Find(&results).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}

	domain := []*domain.Result{}
	conv := converters.NewResultConverter()
	for _, item := range results {
		r, err := conv.ToDomain(item)
		if err != nil {
			return nil, err
		}
		domain = append(domain, r)
	}
	return domain, nil
}

package dataaccess

import (
	"errors"
	"src/internal/entities"
	"src/internal/service"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CoachRepository struct {
	db *gorm.DB
}

func NewCoachRepository(db *gorm.DB) *CoachRepository {
	return &CoachRepository{
		db: db,
	}
}

func (r *CoachRepository) Update(coach *entities.Coach) error {
	err := r.db.Save(coach).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return service.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *CoachRepository) Create(coach *entities.Coach) error {
	coach.ID = uuid.New()
	err := r.db.Create(coach).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CoachRepository) ListCoaches() ([]*entities.Coach, error) {
	coaches := []*entities.Coach{}
	err := r.db.Find(&coaches).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, err
	}
	return coaches, nil
}

func (r *CoachRepository) GetCoachByID(coachID uuid.UUID) (*entities.Coach, error) {
	coach := &entities.Coach{}
	err := r.db.Where("id = ?", coachID).First(coach).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, err
	}
	return coach, nil
}

func (r *CoachRepository) ListSportsmen(coachID uuid.UUID) ([]*entities.Sportsman, error) {
	sportsmen := []*entities.Sportsman{}

	err := r.db.Table("sportsmen").
		Joins("INNER JOIN sportsmen_coaches ON sportsmen.id = sportsmen_coaches.sportsman_id").
		Joins("INNER JOIN coaches ON coaches.id = sportsmen_coaches.coach_id AND coaches.id = ?",
			coachID).
		Find(&sportsmen).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, err
	}
	return sportsmen, nil
}

func (r *CoachRepository) AddSportsman(coachID, smID uuid.UUID) (*entities.SportsmenCoach, error) {
	rec := &entities.SportsmenCoach{
		SportsmanID: smID,
		CoachID:     coachID,
	}
	err := r.db.Create(rec).Error
	if err != nil {
		return nil, err
	}
	return rec, nil
}

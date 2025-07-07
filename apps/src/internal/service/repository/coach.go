package repository

import (
	"src/internal/domain"

	"github.com/google/uuid"
)

type ICoachRepository interface {
	Update(coach *domain.Coach) error
	Create(coach *domain.Coach) error
	ListCoaches(
		page int,
		batch int,
		sort string,
		filter string,
	) ([]*domain.Coach, error)
	GetCoachByID(coachID uuid.UUID) (*domain.Coach, error)
	ListSportsmen(
		coachID uuid.UUID,
		page int,
		batch int,
		sort string,
		filter string,
	) ([]*domain.Sportsman, error)
	AddSportsman(coachID, smID uuid.UUID) (*domain.SportsmenCoach, error)
	RemoveSportsman(coachID, smID uuid.UUID) error
	Delete(coachID uuid.UUID) error
}

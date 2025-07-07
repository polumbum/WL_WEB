package repository

import (
	"src/internal/entities"

	"github.com/google/uuid"
)

type ICoachRepository interface {
	Update(coach *entities.Coach) error
	Create(coach *entities.Coach) error
	ListCoaches() ([]*entities.Coach, error)
	GetCoachByID(coachID uuid.UUID) (*entities.Coach, error)
	ListSportsmen(coachID uuid.UUID) ([]*entities.Sportsman, error)
	AddSportsman(coachID uuid.UUID, smID uuid.UUID) (*entities.SportsmenCoach, error)
}

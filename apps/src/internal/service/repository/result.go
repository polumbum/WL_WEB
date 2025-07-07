package repository

import (
	"src/internal/domain"

	"github.com/google/uuid"
)

type IResultRepository interface {
	Update(result *domain.Result) error
	Create(result *domain.Result) error
	ListResults() ([]*domain.Result, error)
	ListSportsmanResults(sportsmanID uuid.UUID) ([]*domain.Result, error)
	ListCompResults(compID uuid.UUID) ([]*domain.Result, error)
	ListCoachResults(coachID uuid.UUID) ([]*domain.Result, error)
	GetResultByID(smID, compID uuid.UUID) (*domain.Result, error)
}

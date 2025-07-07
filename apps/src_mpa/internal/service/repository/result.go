package repository

import (
	"src/internal/entities"

	"github.com/google/uuid"
)

type IResultRepository interface {
	Update(result *entities.Result) error
	Create(result *entities.Result) error
	ListResults() ([]*entities.Result, error)
	ListSportsmanResults(sportsmanID uuid.UUID) ([]*entities.Result, error)
	GetResultByID(smID, compID uuid.UUID) (*entities.Result, error)
}

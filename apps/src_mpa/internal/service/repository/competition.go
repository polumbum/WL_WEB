package repository

import (
	"src/internal/entities"

	"github.com/google/uuid"
)

type ICompetitionRepository interface {
	Update(competition *entities.Competition) error
	Create(competition *entities.Competition) error
	ListCompetitions() ([]*entities.Competition, error)
	GetCompetitionByID(competitionID uuid.UUID) (*entities.Competition, error)
	RegisterSportsman(compApplication *entities.CompApplication) error
	DeleteRegistration(smID, compID uuid.UUID) error
	GetUpcoming(smID uuid.UUID) ([]*entities.Competition, error)
	ListUpcoming() ([]*entities.Competition, error)
}

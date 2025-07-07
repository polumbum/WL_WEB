package repository

import (
	"src/internal/domain"

	"github.com/google/uuid"
)

type ICompetitionRepository interface {
	Update(competition *domain.Competition) error
	Create(competition *domain.Competition) error
	ListCompetitions(
		page int,
		batch int,
		sort string,
		filter string,
	) ([]*domain.Competition, error)
	GetCompetitionByID(competitionID uuid.UUID) (*domain.Competition, error)
	RegisterSportsman(compApplication *domain.CompApplication) error
	DeleteRegistration(smID, compID uuid.UUID) error
	//GetUpcoming(smID uuid.UUID) ([]*domain.Competition, error)
	//ListUpcoming() ([]*domain.Competition, error)
	Delete(id uuid.UUID) error
	ListByOrgID(id uuid.UUID) ([]*domain.Competition, error)
}

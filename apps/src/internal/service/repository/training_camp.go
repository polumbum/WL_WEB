package repository

import (
	"src/internal/domain"

	"github.com/google/uuid"
)

type ITCampRepository interface {
	Update(trainingCamp *domain.TCamp) error
	Create(trainingCamp *domain.TCamp) error
	ListTCamps(
		page int,
		batch int,
		sort string,
		filter string,
	) ([]*domain.TCamp, error)
	GetTCampByID(trainingCampID uuid.UUID) (*domain.TCamp, error)
	RegisterSportsman(tCampApplication *domain.TCampApplication) error
	DeleteRegistration(smID, tCampID uuid.UUID) error
	GetUpcoming(smID uuid.UUID) ([]*domain.TCamp, error)
	ListUpcoming() ([]*domain.TCamp, error)
	Delete(id uuid.UUID) error
	ListByOrgID(id uuid.UUID) ([]*domain.TCamp, error)
}

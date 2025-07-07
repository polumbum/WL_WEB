package repository

import (
	"src/internal/entities"

	"github.com/google/uuid"
)

type ITCampRepository interface {
	Update(trainingCamp *entities.TCamp) error
	Create(trainingCamp *entities.TCamp) error
	ListTCamps() ([]*entities.TCamp, error)
	GetTCampByID(trainingCampID uuid.UUID) (*entities.TCamp, error)
	RegisterSportsman(tCampApplication *entities.TCampApplication) error
	DeleteRegistration(smID, tCampID uuid.UUID) error
	GetUpcoming(smID uuid.UUID) ([]*entities.TCamp, error)
	ListUpcoming() ([]*entities.TCamp, error)
}

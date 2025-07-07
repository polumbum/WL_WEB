package repository

import (
	"src/internal/entities"

	"github.com/google/uuid"
)

type IADopingRepository interface {
	Update(ad *entities.Antidoping) error
	Create(ad *entities.Antidoping) error
	GetADopingBySmID(sportsmanID uuid.UUID) (*entities.Antidoping, error)
}

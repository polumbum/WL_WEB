package repository

import (
	"src/internal/domain"

	"github.com/google/uuid"
)

type IADopingRepository interface {
	Update(ad *domain.Antidoping) (*domain.Antidoping, error)
	Create(ad *domain.Antidoping) error
	GetADopingBySmID(sportsmanID uuid.UUID) (*domain.Antidoping, error)
}

package repository

import (
	"src/internal/entities"

	"github.com/google/uuid"
)

type ICompAccessRepository interface {
	Update(ca *entities.CompAccess) error
	Create(ca *entities.CompAccess) error
	GetAccessBySmID(sportsmanID uuid.UUID) (*entities.CompAccess, error)
}

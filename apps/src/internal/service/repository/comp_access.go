package repository

import (
	"src/internal/domain"

	"github.com/google/uuid"
)

type ICompAccessRepository interface {
	Update(ca *domain.CompAccess) (*domain.CompAccess, error)
	Create(ca *domain.CompAccess) error
	GetAccessBySmID(sportsmanID uuid.UUID) (*domain.CompAccess, error)
}

package repository

import (
	"src/internal/domain"

	"github.com/google/uuid"
)

type ISportsmanRepository interface {
	Update(sportsman *domain.Sportsman) (*domain.Sportsman, error)
	Create(sportsman *domain.Sportsman) error
	ListSportsmen(page int,
		batch int,
		sort string,
		filter string,
	) ([]*domain.Sportsman, error)
	GetSportsmanByID(sportsmanID uuid.UUID) (*domain.Sportsman, error)
	ListResults(sportsmanID uuid.UUID) ([]*domain.Result, error)
	Delete(sportsmanID uuid.UUID) error
}

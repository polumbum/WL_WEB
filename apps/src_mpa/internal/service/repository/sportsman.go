package repository

import (
	"src/internal/entities"

	"github.com/google/uuid"
)

type ISportsmanRepository interface {
	Update(sportsman *entities.Sportsman) error
	Create(sportsman *entities.Sportsman) error
	ListSportsmen() ([]*entities.Sportsman, error)
	GetSportsmanByID(sportsmanID uuid.UUID) (*entities.Sportsman, error)
	ListResults(sportsmanID uuid.UUID) ([]*entities.Result, error)
}

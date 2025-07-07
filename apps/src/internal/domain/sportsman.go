package domain

import (
	"time"

	"src/internal/constants"

	"github.com/google/uuid"
)

type Sportsman struct {
	ID             uuid.UUID
	Surname        string
	Name           string
	Patronymic     string
	Birthday       time.Time
	SportsCategory constants.SportsCategoryT
	Gender         constants.GenderT
	MoscowTeam     bool
}

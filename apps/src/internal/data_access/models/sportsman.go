package models

import (
	"time"

	"github.com/google/uuid"
)

type Sportsman struct {
	ID             uuid.UUID
	Surname        string
	Name           string
	Patronymic     string
	Birthday       time.Time
	SportsCategory string
	Gender         bool
	MoscowTeam     bool
}

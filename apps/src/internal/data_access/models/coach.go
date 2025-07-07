package models

import (
	"time"

	"github.com/google/uuid"
)

type Coach struct {
	ID         uuid.UUID
	Surname    string
	Name       string
	Patronymic string
	Experience int // years
	Birthday   time.Time
	Gender     bool
}

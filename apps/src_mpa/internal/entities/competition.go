package entities

import (
	"time"

	"src/internal/constants"

	"github.com/google/uuid"
)

type Competition struct {
	ID                uuid.UUID
	Name              string
	City              string
	Address           string
	BegDate           time.Time
	EndDate           time.Time
	Age               constants.AgeCategoryT // age category
	MinSportsCategory constants.SportsCategoryT
	Antidoping        bool
}

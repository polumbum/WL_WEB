package models

import (
	"time"

	"github.com/google/uuid"
)

type Competition struct {
	ID                uuid.UUID
	Name              string
	City              string
	Address           string
	BegDate           time.Time
	EndDate           time.Time
	Age               string
	MinSportsCategory string
	Antidoping        bool
	OrgID             uuid.UUID
}

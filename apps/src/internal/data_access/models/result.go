package models

import (
	"github.com/google/uuid"
)

type Result struct {
	SportsmanID    uuid.UUID
	CompetitionID  uuid.UUID
	WeightCategory int
	Snatch         int
	CleanAndJerk   int
	Place          int
}

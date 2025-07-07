package models

import (
	"github.com/google/uuid"
)

type CompApplication struct {
	SportsmanID       uuid.UUID
	CompetitionID     uuid.UUID
	WeightCategory    int
	StartSnatch       int
	StartCleanAndJerk int
}

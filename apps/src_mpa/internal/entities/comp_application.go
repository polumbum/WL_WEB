package entities

import (
	"src/internal/constants"

	"github.com/google/uuid"
)

type CompApplication struct {
	SportsmanID       uuid.UUID
	CompetitionID     uuid.UUID
	WeightCategory    constants.WeightCategoryT
	StartSnatch       int
	StartCleanAndJerk int
}

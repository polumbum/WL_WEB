package domain

import (
	"src/internal/constants"

	"github.com/google/uuid"
)

type Result struct {
	SportsmanID    uuid.UUID
	CompetitionID  uuid.UUID
	WeightCategory constants.WeightCategoryT
	Snatch         int
	CleanAndJerk   int
	Place          int
}

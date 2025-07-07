package domain

import (
	"github.com/google/uuid"
)

type SportsmenCoach struct {
	SportsmanID uuid.UUID
	CoachID     uuid.UUID
}

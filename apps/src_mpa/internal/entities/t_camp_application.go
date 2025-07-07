package entities

import (
	"github.com/google/uuid"
)

type TCampApplication struct {
	SportsmanID uuid.UUID
	TCampID     uuid.UUID
}

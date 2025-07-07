package models

import (
	"github.com/google/uuid"
)

type TCampApplication struct {
	SportsmanID uuid.UUID
	TCampID     uuid.UUID
}

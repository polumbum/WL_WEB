package domain

import (
	"time"

	"github.com/google/uuid"
)

type Antidoping struct {
	ID          uuid.UUID
	SportsmanID uuid.UUID
	Validity    time.Time
}

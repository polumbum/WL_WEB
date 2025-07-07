package entities

import (
	"time"

	"github.com/google/uuid"
)

type CompAccess struct {
	ID          uuid.UUID
	SportsmanID uuid.UUID
	Validity    time.Time
	Institution string
}

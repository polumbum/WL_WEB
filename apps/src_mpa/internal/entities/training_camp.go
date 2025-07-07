package entities

import (
	"time"

	"github.com/google/uuid"
)

type TCamp struct {
	ID      uuid.UUID
	City    string
	Address string
	BegDate time.Time
	EndDate time.Time
}

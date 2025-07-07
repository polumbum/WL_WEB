package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateADopingReq struct {
	SmID     uuid.UUID
	Validity time.Time
}

type UpdateADopingReq struct {
	SmID     uuid.UUID
	Validity time.Time
}

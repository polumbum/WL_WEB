package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateAccessReq struct {
	SmID        uuid.UUID
	Validity    time.Time
	Institution string
}

type UpdateAccessReq struct {
	SmID        uuid.UUID
	Validity    time.Time
	Institution string
}

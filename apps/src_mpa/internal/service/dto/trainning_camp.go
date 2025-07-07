package dto

import (
	"time"

	"src/internal/entities"

	"github.com/google/uuid"
)

type UpdateTCampReq struct {
	ID      uuid.UUID
	City    string
	Address string
	BegDate time.Time
	EndDate time.Time
}

type CreateTCampReq struct {
	City    string
	Address string
	BegDate time.Time
	EndDate time.Time
}

type RegForTCampReq struct {
	TCampID     uuid.UUID
	SportsmanID uuid.UUID
}

func (req *UpdateTCampReq) Copy(camp *entities.TCamp) {
	updateStringField(req.City, &camp.City)
	updateStringField(req.Address, &camp.Address)
	updateTimeField(req.BegDate, &camp.BegDate)
	updateTimeField(req.EndDate, &camp.EndDate)
}

func (req *CreateTCampReq) Copy(camp *entities.TCamp) {
	camp.City = req.City
	camp.Address = req.Address
	camp.BegDate = req.BegDate
	camp.EndDate = req.EndDate
}

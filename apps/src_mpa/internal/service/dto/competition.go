package dto

import (
	"time"

	"src/internal/constants"
	"src/internal/entities"

	"github.com/google/uuid"
)

type UpdateCompReq struct {
	ID                uuid.UUID
	Name              string
	City              string
	Address           string
	BegDate           time.Time
	EndDate           time.Time
	Age               constants.AgeCategoryT // age category
	MinSportsCategory constants.SportsCategoryT
}

type CreateCompReq struct {
	Name              string
	City              string
	Address           string
	BegDate           time.Time
	EndDate           time.Time
	Age               constants.AgeCategoryT // age category
	MinSportsCategory constants.SportsCategoryT
	Antidoping        bool
}

type RegForCompReq struct {
	CompetitionID     uuid.UUID
	SportsmanID       uuid.UUID
	WeighCategory     constants.WeightCategoryT
	StartSnatch       int
	StartCleanAndJerk int
}

func (req *UpdateCompReq) Copy(comp *entities.Competition) {
	updateStringField(req.Name, &comp.Name)
	updateStringField(req.City, &comp.City)
	updateStringField(req.Address, &comp.Address)
	updateTimeField(req.BegDate, &comp.BegDate)
	updateTimeField(req.EndDate, &comp.EndDate)
	updateStringField(string(req.Age), (*string)(&comp.Age))
	updateStringField(string(req.MinSportsCategory), (*string)(&comp.MinSportsCategory))
}

func (req *CreateCompReq) Copy(comp *entities.Competition) {
	comp.Name = req.Name
	comp.City = req.City
	comp.Address = req.Address
	comp.BegDate = req.BegDate
	comp.EndDate = req.EndDate
	comp.Age = req.Age
	comp.MinSportsCategory = req.MinSportsCategory
	comp.Antidoping = req.Antidoping
}

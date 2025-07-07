package dto

import (
	"src/internal/constants"
	"src/internal/entities"

	"github.com/google/uuid"
)

type UpdateResultReq struct {
	SportsmanID    uuid.UUID
	CompetitionID  uuid.UUID
	WeightCategory constants.WeightCategoryT
	Snatch         int
	CleanAndJerk   int
	Place          int
}

type CreateResultReq struct {
	SportsmanID    uuid.UUID
	CompetitionID  uuid.UUID
	WeightCategory constants.WeightCategoryT
	Snatch         int
	CleanAndJerk   int
	Place          int
}

func (req *UpdateResultReq) Copy(res *entities.Result) {
	updateIntField(int(req.WeightCategory), (*int)(&res.WeightCategory))
	if req.Snatch > constants.MinWeight && res.Snatch != req.Snatch {
		res.Snatch = req.Snatch
	}
	if req.CleanAndJerk > constants.MinWeight && res.CleanAndJerk != req.CleanAndJerk {
		res.CleanAndJerk = req.CleanAndJerk
	}
	updateIntField(req.Place, &res.Place)
}

func (req *CreateResultReq) Copy(res *entities.Result) {
	res.WeightCategory = req.WeightCategory
	res.CompetitionID = req.CompetitionID
	res.SportsmanID = req.SportsmanID
	res.Snatch = req.Snatch
	res.CleanAndJerk = req.CleanAndJerk
	res.Place = req.Place
}

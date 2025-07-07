package dto

import (
	"time"

	"src/internal/constants"
	"src/internal/entities"

	"github.com/google/uuid"
)

type UpdateCoachReq struct {
	ID         uuid.UUID
	Surname    string
	Name       string
	Patronymic string
	Experience int // years
	Birthday   time.Time
	Gender     *constants.GenderT
}

type CreateCoachReq struct {
	Surname    string
	Name       string
	Patronymic string
	Experience int // years
	Birthday   time.Time
	Gender     constants.GenderT
}

func (req *UpdateCoachReq) Copy(coach *entities.Coach) {
	updateStringField(req.Surname, &coach.Surname)
	updateStringField(req.Name, &coach.Name)
	updateStringField(req.Patronymic, &coach.Patronymic)
	updateTimeField(req.Birthday, &coach.Birthday)
	updateIntField(req.Experience, &coach.Experience)
	updateBoolField((*bool)(req.Gender), (*bool)(&coach.Gender))
}

func (req *CreateCoachReq) Copy(coach *entities.Coach) {
	coach.Surname = req.Surname
	coach.Name = req.Name
	coach.Patronymic = req.Patronymic
	coach.Birthday = req.Birthday
	coach.Experience = req.Experience
	coach.Gender = req.Gender
}

package dto

import (
	"time"

	"src/internal/constants"
	"src/internal/entities"

	"github.com/google/uuid"
)

type UpdateSportsmanReq struct {
	ID             uuid.UUID
	SportsCategory constants.SportsCategoryT
	MoscowTeam     *bool
}

type CreateSportsmanReq struct {
	Surname        string
	Name           string
	Patronymic     string
	Birthday       time.Time
	Gender         constants.GenderT
	SportsCategory constants.SportsCategoryT
	MoscowTeam     bool
}

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

type CreateADopingReq struct {
	SmID     uuid.UUID
	Validity time.Time
}

type UpdateADopingReq struct {
	SmID     uuid.UUID
	Validity time.Time
}

func UpdToCreateAdReq(r *UpdateADopingReq) *CreateADopingReq {
	req := &CreateADopingReq{
		SmID:     r.SmID,
		Validity: r.Validity,
	}
	return req
}

func UpdToCreateCaReq(r *UpdateAccessReq) *CreateAccessReq {
	req := &CreateAccessReq{
		SmID:        r.SmID,
		Validity:    r.Validity,
		Institution: r.Institution,
	}
	return req
}

func (req *CreateAccessReq) Copy(ca *entities.CompAccess) {
	ca.SportsmanID = req.SmID
	ca.Validity = req.Validity
	ca.Institution = req.Institution
}

func (req *UpdateAccessReq) Copy(ca *entities.CompAccess) {
	updateTimeField(req.Validity, &ca.Validity)
	updateStringField(req.Institution, &ca.Institution)
}

func (req *CreateADopingReq) Copy(ca *entities.Antidoping) {
	ca.SportsmanID = req.SmID
	ca.Validity = req.Validity
}

func (req *UpdateADopingReq) Copy(ad *entities.Antidoping) {
	updateTimeField(req.Validity, &ad.Validity)
}

func (req *UpdateSportsmanReq) Copy(sportsman *entities.Sportsman) {
	updateBoolField(req.MoscowTeam, &sportsman.MoscowTeam)
	updateStringField(string(req.SportsCategory),
		(*string)(&sportsman.SportsCategory))
}

func (req *CreateSportsmanReq) Copy(sportsman *entities.Sportsman) {
	sportsman.Surname = req.Surname
	sportsman.Name = req.Name
	sportsman.Patronymic = req.Patronymic
	sportsman.Birthday = req.Birthday
	sportsman.MoscowTeam = req.MoscowTeam
	sportsman.SportsCategory = req.SportsCategory
	sportsman.Gender = req.Gender
}

package converters

import (
	"src/internal/data_access/models"
	"src/internal/domain"
)

type ISmCoachesConverter interface {
	ToDomain(model *models.SportsmenCoach) (
		*domain.SportsmenCoach,
		error,
	)
	ToModel(domain *domain.SportsmenCoach) (
		*models.SportsmenCoach,
		error,
	)
}

type SmCoachesConverter struct{}

func NewSmCoachesConverter() *SmCoachesConverter {
	return &SmCoachesConverter{}
}

func (c *SmCoachesConverter) ToDomain(model *models.SportsmenCoach) (
	*domain.SportsmenCoach,
	error,
) {
	if model == nil {
		return nil, ErrConvert
	}
	return &domain.SportsmenCoach{
		SportsmanID: model.SportsmanID,
		CoachID:     model.CoachID,
	}, nil
}

func (c *SmCoachesConverter) ToModel(domain *domain.SportsmenCoach) (
	*models.SportsmenCoach,
	error,
) {
	if domain == nil {
		return nil, ErrConvert
	}
	return &models.SportsmenCoach{
		SportsmanID: domain.SportsmanID,
		CoachID:     domain.CoachID,
	}, nil
}

package converters

import (
	"src/internal/constants"
	"src/internal/data_access/models"
	"src/internal/domain"
	"src/internal/dto"
)

type IResultConverter interface {
	ToDomain(model *models.Result) (*domain.Result,
		error,
	)
	ToModel(domain *domain.Result) (*models.Result,
		error,
	)
	ToDTO(domain *domain.Result) (*dto.Result,
		error,
	)
	FromCreateReq(req *dto.CreateResultReq) (
		*domain.Result,
		error,
	)
}

type ResultConverter struct{}

func NewResultConverter() *ResultConverter {
	return &ResultConverter{}
}

func (c *ResultConverter) ToDomain(model *models.Result) (
	*domain.Result,
	error,
) {
	if model == nil {
		return nil, ErrConvert
	}
	return &domain.Result{
		SportsmanID:    model.SportsmanID,
		CompetitionID:  model.CompetitionID,
		WeightCategory: constants.WeightCategoryT(model.WeightCategory),
		Snatch:         model.Snatch,
		CleanAndJerk:   model.CleanAndJerk,
		Place:          model.Place,
	}, nil
}

func (c *ResultConverter) ToModel(domain *domain.Result) (
	*models.Result,
	error,
) {
	if domain == nil {
		return nil, ErrConvert
	}
	return &models.Result{
		SportsmanID:    domain.SportsmanID,
		CompetitionID:  domain.CompetitionID,
		WeightCategory: int(domain.WeightCategory),
		Snatch:         domain.Snatch,
		CleanAndJerk:   domain.CleanAndJerk,
		Place:          domain.Place,
	}, nil
}

func (c *ResultConverter) ToDTO(domain *domain.Result) (*dto.Result,
	error,
) {
	if domain == nil {
		return nil, ErrConvert
	}
	return &dto.Result{
		SportsmanID:    domain.SportsmanID,
		CompetitionID:  domain.CompetitionID,
		WeightCategory: int(domain.WeightCategory),
		Snatch:         domain.Snatch,
		CleanAndJerk:   domain.CleanAndJerk,
		Place:          domain.Place,
	}, nil
}

func (c *ResultConverter) FromCreateReq(req *dto.CreateResultReq) (
	*domain.Result,
	error,
) {
	if req == nil {
		return nil, ErrConvert
	}
	return &domain.Result{
		SportsmanID:    req.SportsmanID,
		CompetitionID:  req.CompetitionID,
		WeightCategory: constants.WeightCategoryT(req.WeightCategory),
		Snatch:         req.Snatch,
		CleanAndJerk:   req.CleanAndJerk,
		Place:          req.Place,
	}, nil
}

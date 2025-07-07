package converters

import (
	"src/internal/data_access/models"
	"src/internal/domain"
	"src/internal/dto"
)

type ITCampConverter interface {
	ToDomain(model *models.TCamp) (
		*domain.TCamp,
		error,
	)
	ToModel(domain *domain.TCamp) (
		*models.TCamp,
		error,
	)
	ToDTO(domain *domain.Competition) (
		*dto.TCamp,
		error,
	)
	FromCreateReq(req *dto.CreateTCampReq) (
		*domain.TCamp,
		error,
	)
}

type TCampConverter struct{}

func NewTCampConverter() *TCampConverter {
	return &TCampConverter{}
}

func (c *TCampConverter) ToDomain(model *models.TCamp) (
	*domain.TCamp,
	error,
) {
	if model == nil {
		return nil, ErrConvert
	}
	return &domain.TCamp{
		ID:      model.ID,
		City:    model.City,
		Address: model.Address,
		BegDate: model.BegDate,
		EndDate: model.EndDate,
		OrgID:   model.OrgID,
	}, nil
}

func (c *TCampConverter) ToModel(domain *domain.TCamp) (
	*models.TCamp,
	error,
) {
	if domain == nil {
		return nil, ErrConvert
	}
	return &models.TCamp{
		ID:      domain.ID,
		City:    domain.City,
		Address: domain.Address,
		BegDate: domain.BegDate,
		EndDate: domain.EndDate,
		OrgID:   domain.OrgID,
	}, nil
}

func (c *TCampConverter) ToDTO(domain *domain.TCamp) (
	*dto.TCamp,
	error,
) {
	if domain == nil {
		return nil, ErrConvert
	}
	return &dto.TCamp{
		ID:      domain.ID,
		City:    domain.City,
		Address: domain.Address,
		BegDate: domain.BegDate,
		EndDate: domain.EndDate,
	}, nil
}

func (c *TCampConverter) FromCreateReq(req *dto.CreateTCampReq) (
	*domain.TCamp,
	error,
) {
	if req == nil {
		return nil, ErrConvert
	}
	return &domain.TCamp{
		City:    req.City,
		Address: req.Address,
		BegDate: req.BegDate,
		EndDate: req.EndDate,
	}, nil
}

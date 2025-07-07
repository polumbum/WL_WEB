package converters

import (
	"src/internal/constants"
	"src/internal/data_access/models"
	"src/internal/domain"
	"src/internal/dto"
)

type ICompConverter interface {
	ToDomain(cModel *models.Competition) (*domain.Competition,
		error,
	)
	ToModel(cDomain *domain.Competition) (*models.Competition,
		error,
	)
	ToDTO(cDomain *domain.Competition) (*dto.Competition,
		error,
	)
	FromCreateReq(req *dto.CreateCompReq) (
		*domain.Competition,
		error,
	)
}

type CompConverter struct{}

func NewCompConverter() *CompConverter {
	return &CompConverter{}
}

func (c *CompConverter) ToDomain(cModel *models.Competition) (
	*domain.Competition,
	error,
) {
	if cModel == nil {
		return nil, ErrConvert
	}
	return &domain.Competition{
		ID:                cModel.ID,
		Name:              cModel.Name,
		City:              cModel.City,
		Address:           cModel.Address,
		BegDate:           cModel.BegDate,
		EndDate:           cModel.EndDate,
		Age:               constants.AgeCategoryT(cModel.Age),
		MinSportsCategory: constants.SportsCategoryT(cModel.MinSportsCategory),
		Antidoping:        cModel.Antidoping,
		OrgID:             cModel.OrgID,
	}, nil
}

func (c *CompConverter) ToModel(cDomain *domain.Competition) (
	*models.Competition,
	error,
) {
	if cDomain == nil {
		return nil, ErrConvert
	}
	return &models.Competition{
		ID:                cDomain.ID,
		Name:              cDomain.Name,
		City:              cDomain.City,
		Address:           cDomain.Address,
		BegDate:           cDomain.BegDate,
		EndDate:           cDomain.EndDate,
		Age:               string(cDomain.Age),
		MinSportsCategory: string(cDomain.MinSportsCategory),
		Antidoping:        cDomain.Antidoping,
		OrgID:             cDomain.OrgID,
	}, nil
}

func (c *CompConverter) ToDTO(cDomain *domain.Competition) (*dto.Competition,
	error,
) {
	if cDomain == nil {
		return nil, ErrConvert
	}
	return &dto.Competition{
		ID:                cDomain.ID,
		Name:              cDomain.Name,
		City:              cDomain.City,
		Address:           cDomain.Address,
		BegDate:           cDomain.BegDate,
		EndDate:           cDomain.EndDate,
		Age:               string(cDomain.Age),
		MinSportsCategory: string(cDomain.MinSportsCategory),
		Antidoping:        cDomain.Antidoping,
	}, nil
}

func (c *CompConverter) FromCreateReq(req *dto.CreateCompReq) (
	*domain.Competition,
	error,
) {
	if req == nil {
		return nil, ErrConvert
	}
	return &domain.Competition{
		Name:              req.Name,
		City:              req.City,
		Address:           req.Address,
		BegDate:           req.BegDate,
		EndDate:           req.EndDate,
		Age:               constants.AgeCategoryT(req.Age),
		MinSportsCategory: constants.SportsCategoryT(req.MinSportsCategory),
		Antidoping:        req.Antidoping,
		OrgID:             req.OrgID,
	}, nil
}

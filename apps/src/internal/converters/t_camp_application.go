package converters

import (
	"src/internal/data_access/models"
	"src/internal/domain"
	"src/internal/dto"

	"github.com/google/uuid"
)

type ITCApplConverter interface {
	ToDomain(model *models.TCampApplication) (
		*domain.TCampApplication,
		error,
	)
	ToModel(domain *domain.TCampApplication) (
		*models.TCampApplication,
		error,
	)
	ToDTO(domain *domain.TCampApplication) (
		*dto.TCampApplication,
		error,
	)
}

type TCApplConverter struct{}

func NewTCApplConverter() *TCApplConverter {
	return &TCApplConverter{}
}

func (c *TCApplConverter) ToDomain(model *models.TCampApplication) (
	*domain.TCampApplication,
	error,
) {
	if model == nil {
		return nil, ErrConvert
	}
	return &domain.TCampApplication{
		SportsmanID: model.SportsmanID,
		TCampID:     model.TCampID,
	}, nil
}

func (c *TCApplConverter) ToModel(domain *domain.TCampApplication) (
	*models.TCampApplication,
	error,
) {
	if domain == nil {
		return nil, ErrConvert
	}
	return &models.TCampApplication{
		SportsmanID: domain.SportsmanID,
		TCampID:     domain.TCampID,
	}, nil
}

func (c *TCApplConverter) ToDTO(domain *domain.TCampApplication) (
	*dto.TCampApplication,
	error,
) {
	if domain == nil {
		return nil, ErrConvert
	}
	return &dto.TCampApplication{
		SmID:    domain.SportsmanID,
		TCampID: domain.TCampID,
	}, nil
}

func (c *TCApplConverter) FromRegisterReq(
	campID uuid.UUID,
	req *dto.RegForTCampReq,
) (
	*domain.TCampApplication,
	error,
) {
	if req == nil {
		return nil, ErrConvert
	}
	return &domain.TCampApplication{
		SportsmanID: req.SmID,
		TCampID:     campID,
	}, nil
}

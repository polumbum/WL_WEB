package converters

import (
	"src/internal/constants"
	"src/internal/data_access/models"
	"src/internal/domain"
	"src/internal/dto"

	"github.com/google/uuid"
)

type ICompApplConverter interface {
	ToDomain(caModel *models.CompApplication) (*domain.CompApplication,
		error)
	ToModel(caDomain *domain.CompApplication) (*models.CompApplication,
		error)
	ToDTO(domain *domain.CompApplication) (*dto.CompApplication,
		error,
	)
	FromRegisterReq(compID uuid.UUID, req *dto.RegForCompReq) (
		*domain.CompApplication,
		error,
	)
}

type CompApplConverter struct{}

func NewCompApplConverter() *CompApplConverter {
	return &CompApplConverter{}
}

func (c *CompApplConverter) ToDomain(caModel *models.CompApplication) (
	*domain.CompApplication,
	error,
) {
	if caModel == nil {
		return nil, ErrConvert
	}
	return &domain.CompApplication{
		SportsmanID:       caModel.SportsmanID,
		CompetitionID:     caModel.CompetitionID,
		WeightCategory:    constants.WeightCategoryT(caModel.WeightCategory),
		StartSnatch:       caModel.StartSnatch,
		StartCleanAndJerk: caModel.StartCleanAndJerk,
	}, nil
}

func (c *CompApplConverter) ToModel(caDomain *domain.CompApplication) (
	*models.CompApplication,
	error,
) {
	if caDomain == nil {
		return nil, ErrConvert
	}
	return &models.CompApplication{
		SportsmanID:       caDomain.SportsmanID,
		CompetitionID:     caDomain.CompetitionID,
		WeightCategory:    int(caDomain.WeightCategory),
		StartSnatch:       caDomain.StartSnatch,
		StartCleanAndJerk: caDomain.StartCleanAndJerk,
	}, nil
}

func (c *CompApplConverter) ToDTO(domain *domain.CompApplication) (*dto.CompApplication,
	error,
) {
	if domain == nil {
		return nil, ErrConvert
	}
	return &dto.CompApplication{
		SmID:              domain.SportsmanID,
		CompID:            domain.CompetitionID,
		WeightCategory:    int(domain.WeightCategory),
		StartSnatch:       domain.StartSnatch,
		StartCleanAndJerk: domain.StartCleanAndJerk,
	}, nil
}

func (c *CompApplConverter) FromRegisterReq(
	compID uuid.UUID,
	req *dto.RegForCompReq,
) (
	*domain.CompApplication,
	error,
) {
	if req == nil {
		return nil, ErrConvert
	}
	return &domain.CompApplication{
		SportsmanID:       req.SmID,
		CompetitionID:     compID,
		WeightCategory:    constants.WeightCategoryT(req.WeighCategory),
		StartSnatch:       req.StartSnatch,
		StartCleanAndJerk: req.StartCleanAndJerk,
	}, nil
}

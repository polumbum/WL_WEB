package converters

import (
	"src/internal/data_access/models"
	"src/internal/domain"
	"src/internal/dto"
)

type IAccessConverter interface {
	ToDomain(caModel *models.CompAccess) (*domain.CompAccess,
		error)
	ToModel(caDomain *domain.CompAccess) (*models.CompAccess,
		error)
	FromUpdateSmReq(ca *domain.CompAccess, req *dto.UpdateSportsmanReq) (
		*domain.CompAccess,
		error,
	)
	FromUpdateReq(ca *domain.CompAccess, req *dto.UpdateAccessReq) (
		*domain.CompAccess,
		error,
	)
	FromCreateReq(req *dto.CreateAccessReq) (
		*domain.CompAccess,
		error,
	)

	/*ToUpdateDTO(smID uuid.UUID,
	updateReq *payload.UpdateSportsmanReq) (*dto.UpdateAccessReq, error)*/
}

type AccessConverter struct{}

func NewAccessConverter() *AccessConverter {
	return &AccessConverter{}
}

func (c *AccessConverter) ToDomain(caModel *models.CompAccess) (
	*domain.CompAccess,
	error,
) {
	if caModel == nil {
		return nil, ErrConvert
	}
	return &domain.CompAccess{
		ID:          caModel.ID,
		SportsmanID: caModel.SportsmanID,
		Validity:    caModel.Validity,
		Institution: caModel.Institution,
	}, nil
}

func (c *AccessConverter) ToModel(caDomain *domain.CompAccess) (
	*models.CompAccess,
	error,
) {
	if caDomain == nil {
		return nil, ErrConvert
	}
	return &models.CompAccess{
		ID:          caDomain.ID,
		SportsmanID: caDomain.SportsmanID,
		Validity:    caDomain.Validity,
		Institution: caDomain.Institution,
	}, nil
}

func (c *AccessConverter) FromUpdateSmReq(
	ca *domain.CompAccess,
	req *dto.UpdateSportsmanReq,
) (
	*domain.CompAccess,
	error,
) {
	if req == nil || ca == nil {
		return nil, ErrConvert
	}
	if !req.Access_validity.IsZero() {
		ca.Validity = req.Access_validity
	}
	if req.Access_institution != "" {
		ca.Institution = req.Access_institution
	}

	return ca, nil
}

func (c *AccessConverter) FromUpdateReq(
	ca *domain.CompAccess,
	req *dto.UpdateAccessReq,
) (
	*domain.CompAccess,
	error,
) {
	if req == nil || ca == nil {
		return nil, ErrConvert
	}
	ca.Validity = req.Validity
	ca.Institution = req.Institution
	return ca, nil
}

func (c *AccessConverter) FromCreateReq(req *dto.CreateAccessReq) (
	*domain.CompAccess,
	error,
) {
	if req == nil {
		return nil, ErrConvert
	}
	return &domain.CompAccess{
		SportsmanID: req.SmID,
		Validity:    req.Validity,
		Institution: req.Institution,
	}, nil
}

/*
func (c *AccessConverter) ToUpdateDTO(smID uuid.UUID,
	updateReq *payload.UpdateSportsmanReq) (*dto.UpdateAccessReq, error) {
	if updateReq == nil {
		return nil, ErrConvert
	}
	return &dto.UpdateAccessReq{
		SmID:        smID,
		Validity:    updateReq.Access_validity,
		Institution: updateReq.Access_institution,
	}, nil
}
*/

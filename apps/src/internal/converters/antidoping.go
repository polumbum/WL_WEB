package converters

import (
	"src/internal/data_access/models"
	"src/internal/domain"
	"src/internal/dto"
)

type IAntidopingConverter interface {
	ToDomain(adModel *models.Antidoping) (*domain.Antidoping,
		error)
	ToModel(adDomain *domain.Antidoping) (*models.Antidoping,
		error)
	FromUpdateSmReq(ad *domain.Antidoping, req *dto.UpdateSportsmanReq) (
		*domain.Antidoping,
		error,
	)
	FromUpdateReq(ad *domain.Antidoping, req *dto.UpdateADopingReq) (
		*domain.Antidoping,
		error,
	)
	FromCreateReq(req *dto.CreateADopingReq) (
		*domain.Antidoping,
		error,
	)

	/*ToUpdateDTO(smID uuid.UUID,
	updateReq *payload.UpdateSportsmanReq) (*dto.UpdateADopingReq, error)*/
}

type ADopingConverter struct{}

func NewADopingConverter() *ADopingConverter {
	return &ADopingConverter{}
}

func (c *ADopingConverter) ToDomain(adModel *models.Antidoping) (
	*domain.Antidoping,
	error,
) {
	if adModel == nil {
		return nil, ErrConvert
	}
	return &domain.Antidoping{
		ID:          adModel.ID,
		SportsmanID: adModel.SportsmanID,
		Validity:    adModel.Validity,
	}, nil
}

func (c *ADopingConverter) ToModel(adDomain *domain.Antidoping) (
	*models.Antidoping,
	error,
) {
	if adDomain == nil {
		return nil, ErrConvert
	}
	return &models.Antidoping{
		ID:          adDomain.ID,
		SportsmanID: adDomain.SportsmanID,
		Validity:    adDomain.Validity,
	}, nil
}

func (c *ADopingConverter) FromUpdateSmReq(
	ad *domain.Antidoping,
	req *dto.UpdateSportsmanReq,
) (
	*domain.Antidoping,
	error,
) {
	if ad == nil || req == nil {
		return nil, ErrConvert
	}
	if !req.Adoping_validity.IsZero() {
		ad.Validity = req.Adoping_validity
	}

	return ad, nil
}

func (c *ADopingConverter) FromUpdateReq(
	ad *domain.Antidoping,
	req *dto.UpdateADopingReq,
) (
	*domain.Antidoping,
	error,
) {
	if ad == nil || req == nil {
		return nil, ErrConvert
	}
	ad.Validity = req.Validity
	return ad, nil
}

func (c *ADopingConverter) FromCreateReq(req *dto.CreateADopingReq) (
	*domain.Antidoping,
	error,
) {
	if req == nil {
		return nil, ErrConvert
	}
	return &domain.Antidoping{
		SportsmanID: req.SmID,
		Validity:    req.Validity,
	}, nil
}

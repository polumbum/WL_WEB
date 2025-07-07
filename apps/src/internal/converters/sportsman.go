package converters

import (
	"src/internal/constants"
	"src/internal/data_access/models"
	"src/internal/domain"
	"src/internal/dto"
)

type ISportsmanConverter interface {
	ToDomain(model *models.Sportsman) (*domain.Sportsman,
		error,
	)
	ToModel(domain *domain.Sportsman) (*models.Sportsman,
		error,
	)
	ToDTO(domain *domain.Sportsman) (*dto.Sportsman,
		error,
	)
	FromDTO(sm *dto.Sportsman) (*domain.Sportsman,
		error,
	)
	FromUpdateReq(sm *dto.Sportsman, req *dto.UpdateSportsmanReq) (
		*domain.Sportsman,
		error,
	)
	FromCreateReq(req *dto.CreateSportsmanReq) (
		*domain.Sportsman,
		error,
	)
}

type SportsmanConverter struct{}

func NewSportsmanConverter() *SportsmanConverter {
	return &SportsmanConverter{}
}

func (c *SportsmanConverter) ToDomain(model *models.Sportsman) (
	*domain.Sportsman,
	error,
) {
	if model == nil {
		return nil, ErrConvert
	}
	return &domain.Sportsman{
		ID:             model.ID,
		Surname:        model.Surname,
		Name:           model.Name,
		Patronymic:     model.Patronymic,
		Birthday:       model.Birthday,
		SportsCategory: constants.SportsCategoryT(model.SportsCategory),
		Gender:         constants.GenderT(model.Gender),
		MoscowTeam:     model.MoscowTeam,
	}, nil
}

func (c *SportsmanConverter) ToModel(domain *domain.Sportsman) (
	*models.Sportsman,
	error,
) {
	if domain == nil {
		return nil, ErrConvert
	}
	return &models.Sportsman{
		ID:             domain.ID,
		Surname:        domain.Surname,
		Name:           domain.Name,
		Patronymic:     domain.Patronymic,
		Birthday:       domain.Birthday,
		SportsCategory: string(domain.SportsCategory),
		Gender:         bool(domain.Gender),
		MoscowTeam:     domain.MoscowTeam,
	}, nil
}

func (c *SportsmanConverter) ToDTO(domain *domain.Sportsman) (
	*dto.Sportsman,
	error,
) {
	if domain == nil {
		return nil, ErrConvert
	}
	return &dto.Sportsman{
		ID:             domain.ID,
		Surname:        domain.Surname,
		Name:           domain.Name,
		Patronymic:     domain.Patronymic,
		Birthday:       domain.Birthday,
		SportsCategory: string(domain.SportsCategory),
		Gender:         bool(domain.Gender),
		MoscowTeam:     domain.MoscowTeam,
	}, nil
}

func (c *SportsmanConverter) FromDTO(sm *dto.Sportsman) (
	*domain.Sportsman,
	error,
) {
	if sm == nil {
		return nil, ErrConvert
	}
	return &domain.Sportsman{
		ID:             sm.ID,
		Surname:        sm.Surname,
		Name:           sm.Name,
		Patronymic:     sm.Patronymic,
		Birthday:       sm.Birthday,
		SportsCategory: constants.SportsCategoryT(sm.SportsCategory),
		Gender:         constants.GenderT(sm.Gender),
		MoscowTeam:     sm.MoscowTeam,
	}, nil
}

func (c *SportsmanConverter) FromUpdateReq(
	sm *dto.Sportsman,
	req *dto.UpdateSportsmanReq,
) (
	*domain.Sportsman,
	error,
) {
	if req == nil || sm == nil {
		return nil, ErrConvert
	}
	sm.MoscowTeam = req.MoscowTeam
	sm.SportsCategory = req.SportsCategory

	return c.FromDTO(sm)
}

func (c *SportsmanConverter) FromCreateReq(req *dto.CreateSportsmanReq) (
	*domain.Sportsman,
	error,
) {
	if req == nil {
		return nil, ErrConvert
	}
	return &domain.Sportsman{
		Surname:        req.Surname,
		Name:           req.Name,
		Patronymic:     req.Patronymic,
		Birthday:       req.Birthday,
		MoscowTeam:     req.MoscowTeam,
		SportsCategory: constants.SportsCategoryT(req.SportsCategory),
		Gender:         constants.GenderT(req.Gender),
	}, nil
}

package converters

import (
	"src/internal/constants"
	"src/internal/data_access/models"
	"src/internal/domain"
	"src/internal/dto"
)

type ICoachConverter interface {
	ToDomain(coachModel *models.Coach) (*domain.Coach,
		error,
	)
	ToModel(coachDomain *domain.Coach) (*models.Coach,
		error,
	)
	ToDTO(coachDomain *domain.Coach) (*dto.Coach,
		error,
	)
	FromUpdateReq(coach *domain.Coach, req *dto.UpdateCoachReq) (
		*domain.Coach,
		error,
	)
	FromCreateReq(req *dto.CreateCoachReq) (
		*domain.Coach,
		error,
	)
}

type CoachConverter struct{}

func NewCoachConverter() *CoachConverter {
	return &CoachConverter{}
}

func (c *CoachConverter) ToDomain(coachModel *models.Coach) (
	*domain.Coach,
	error,
) {
	if coachModel == nil {
		return nil, ErrConvert
	}
	return &domain.Coach{
		ID:         coachModel.ID,
		Surname:    coachModel.Surname,
		Name:       coachModel.Name,
		Patronymic: coachModel.Patronymic,
		Experience: coachModel.Experience,
		Birthday:   coachModel.Birthday,
		Gender:     constants.GenderT(coachModel.Gender),
	}, nil
}

func (c *CoachConverter) ToModel(coachDomain *domain.Coach) (
	*models.Coach,
	error,
) {
	if coachDomain == nil {
		return nil, ErrConvert
	}
	return &models.Coach{
		ID:         coachDomain.ID,
		Surname:    coachDomain.Surname,
		Name:       coachDomain.Name,
		Patronymic: coachDomain.Patronymic,
		Experience: coachDomain.Experience,
		Birthday:   coachDomain.Birthday,
		Gender:     bool(coachDomain.Gender),
	}, nil
}

func (c *CoachConverter) ToDTO(coachDomain *domain.Coach) (*dto.Coach,
	error,
) {
	if coachDomain == nil {
		return nil, ErrConvert
	}
	return &dto.Coach{
		ID:         coachDomain.ID,
		Surname:    coachDomain.Surname,
		Name:       coachDomain.Name,
		Patronymic: coachDomain.Patronymic,
		Experience: coachDomain.Experience,
		Birthday:   coachDomain.Birthday,
		Gender:     bool(coachDomain.Gender),
	}, nil
}

func (c *CoachConverter) FromUpdateReq(
	coach *domain.Coach,
	req *dto.UpdateCoachReq,
) (
	*domain.Coach,
	error,
) {
	if req == nil || coach == nil {
		return nil, ErrConvert
	}
	coach.Surname = req.Surname
	coach.Name = req.Name
	coach.Patronymic = req.Patronymic
	coach.Birthday = req.Birthday
	coach.Experience = req.Experience
	coach.Gender = constants.GenderT(req.Gender)
	return coach, nil
}

func (c *CoachConverter) FromCreateReq(req *dto.CreateCoachReq) (
	*domain.Coach,
	error,
) {
	if req == nil {
		return nil, ErrConvert
	}
	return &domain.Coach{
		Surname:    req.Surname,
		Name:       req.Name,
		Patronymic: req.Patronymic,
		Birthday:   req.Birthday,
		Experience: req.Experience,
		Gender:     constants.GenderT(req.Gender),
	}, nil
}

/*func (c *CoachConverter) ToModel(cEntity *domain.Coach) (*payload.Coach,
	error) {
	if cEntity == nil {
		return nil, ErrConvert
	}
	return &payload.Coach{
		ID:         cEntity.ID,
		Surname:    cEntity.Surname,
		Name:       cEntity.Name,
		Patronymic: cEntity.Patronymic,
		Birthday:   cEntity.Birthday,
		Experience: cEntity.Experience,
		Gender:     bool(cEntity.Gender),
	}, nil
}

func (c *CoachConverter) ToEntity(cModel *payload.Coach) (*domain.Coach,
	error) {
	if cModel == nil {
		return nil, ErrConvert
	}
	return &domain.Coach{
		ID:         cModel.ID,
		Surname:    cModel.Surname,
		Name:       cModel.Name,
		Patronymic: cModel.Patronymic,
		Experience: cModel.Experience,
		Birthday:   cModel.Birthday,
		Gender:     constants.GenderT(cModel.Gender),
	}, nil
}
*/

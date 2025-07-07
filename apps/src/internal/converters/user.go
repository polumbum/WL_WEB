package converters

import (
	"src/internal/constants"
	"src/internal/data_access/models"
	"src/internal/domain"
	"src/internal/dto"
)

type IUserConverter interface {
	ToDomain(model *models.User) (
		*domain.User,
		error,
	)
	ToModel(domain *domain.User) (
		*models.User,
		error,
	)
	ToDTO(domain *domain.User) (
		*dto.User,
		error,
	)
	FromRegisterReq(req *dto.RegisterUserReq) (
		*domain.User,
		error,
	)
	FromUpdateReq(
		user *dto.User,
		req *dto.UpdateUserReq,
	) (
		*domain.User,
		error,
	)
	LoginUserReq(req *dto.LoginUserReq) (
		*domain.User,
		error,
	)
}

type UserConverter struct{}

func NewUserConverter() *UserConverter {
	return &UserConverter{}
}

func (c *UserConverter) ToDomain(model *models.User) (
	*domain.User,
	error,
) {
	if model == nil {
		return nil, ErrConvert
	}
	return &domain.User{
		ID:       model.ID,
		Email:    model.Email,
		Password: model.Password,
		Role:     constants.UserRole(model.Role),
		RoleID:   model.RoleID,
	}, nil
}

func (c *UserConverter) ToModel(domain *domain.User) (
	*models.User,
	error,
) {
	if domain == nil {
		return nil, ErrConvert
	}
	return &models.User{
		ID:       domain.ID,
		Email:    domain.Email,
		Password: domain.Password,
		Role:     string(domain.Role),
		RoleID:   domain.RoleID,
	}, nil
}

func (c *UserConverter) ToDTO(domain *domain.User) (
	*dto.User,
	error,
) {
	if domain == nil {
		return nil, ErrConvert
	}
	return &dto.User{
		ID:     domain.ID,
		Email:  domain.Email,
		Role:   string(domain.Role),
		RoleID: domain.RoleID,
	}, nil
}

func (c *UserConverter) FromDTO(sm *dto.User) (
	*domain.User,
	error,
) {
	if sm == nil {
		return nil, ErrConvert
	}
	return &domain.User{
		ID:     sm.ID,
		Email:  sm.Email,
		Role:   constants.UserRole(sm.Role),
		RoleID: sm.RoleID,
	}, nil
}

func (c *UserConverter) FromRegisterReq(req *dto.RegisterUserReq) (
	*domain.User,
	error,
) {
	if req == nil {
		return nil, ErrConvert
	}
	return &domain.User{
		Email:    req.Email,
		Password: req.Password,
		Role:     constants.UserRole(req.Role),
	}, nil
}

func (c *UserConverter) FromUpdateReq(
	user *dto.User,
	req *dto.UpdateUserReq,
) (
	*domain.User,
	error,
) {
	if req == nil || user == nil {
		return nil, ErrConvert
	}
	user.Email = req.Email
	return c.FromDTO(user)
}

func (c *UserConverter) LoginUserReq(req *dto.LoginUserReq) (
	*domain.User,
	error,
) {
	if req == nil {
		return nil, ErrConvert
	}
	return &domain.User{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}

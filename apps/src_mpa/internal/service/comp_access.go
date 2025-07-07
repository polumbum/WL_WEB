package service

import (
	"src/internal/entities"
	"src/internal/service/dto"
	"src/internal/service/repository"

	"github.com/google/uuid"
)

type IAccessService interface {
	Create(req *dto.CreateAccessReq) (*entities.CompAccess, error)
	Update(req *dto.UpdateAccessReq) (*entities.CompAccess, error)
	GetAccessByID(sportsmanID uuid.UUID) (*entities.CompAccess, error)
}

type AccessService struct {
	repo repository.ICompAccessRepository
}

func NewAccessService(
	repo repository.ICompAccessRepository,
) *AccessService {
	return &AccessService{
		repo: repo,
	}
}

func (s *AccessService) Create(req *dto.CreateAccessReq) (*entities.CompAccess, error) {
	var ca entities.CompAccess
	req.Copy(&ca)
	err := s.repo.Create(&ca)
	if err != nil {
		return nil, err
	}

	return &ca, nil
}

func (s *AccessService) Update(req *dto.UpdateAccessReq) (*entities.CompAccess, error) {
	ca, err := s.repo.GetAccessBySmID(req.SmID)
	if err != nil {
		newReq := dto.UpdToCreateCaReq(req)
		return s.Create(newReq)
	}

	req.Copy(ca)

	err = s.repo.Update(ca)
	if err != nil {
		return nil, err
	}

	return ca, nil
}

func (s *AccessService) GetAccessByID(sportsmanID uuid.UUID) (*entities.CompAccess, error) {
	ca, err := s.repo.GetAccessBySmID(sportsmanID)
	if err != nil {
		return nil, err
	}

	return ca, nil
}

package service

import (
	"src/internal/domain"
	"src/internal/service/repository"

	"github.com/google/uuid"
)

type IAccessService interface {
	GetAccessByID(smID uuid.UUID) (*domain.CompAccess, error)
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

func (s *AccessService) GetAccessByID(smID uuid.UUID) (*domain.CompAccess, error) {
	ca, err := s.repo.GetAccessBySmID(smID)
	if err != nil {
		return nil, err
	}

	return ca, nil
}

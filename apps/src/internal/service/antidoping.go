package service

import (
	"src/internal/domain"
	"src/internal/service/repository"

	"github.com/google/uuid"
)

type IADopingService interface {
	GetADopingByID(smID uuid.UUID) (*domain.Antidoping, error)
}

type ADopingService struct {
	repo repository.IADopingRepository
}

func NewADopingService(
	repo repository.IADopingRepository,
) *ADopingService {
	return &ADopingService{
		repo: repo,
	}
}

func (s *ADopingService) GetADopingByID(smID uuid.UUID) (*domain.Antidoping, error) {
	ad, err := s.repo.GetADopingBySmID(smID)
	if err != nil {
		return nil, err
	}

	return ad, nil
}

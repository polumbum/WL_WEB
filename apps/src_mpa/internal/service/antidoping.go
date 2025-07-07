package service

import (
	"src/internal/entities"
	"src/internal/service/dto"
	"src/internal/service/repository"

	"github.com/google/uuid"
)

type IADopingService interface {
	Create(req *dto.CreateADopingReq) (*entities.Antidoping, error)
	Update(req *dto.UpdateADopingReq) (*entities.Antidoping, error)
	GetADopingByID(sportsmanID uuid.UUID) (*entities.Antidoping, error)
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

func (s *ADopingService) Create(req *dto.CreateADopingReq) (*entities.Antidoping, error) {
	var ad entities.Antidoping
	req.Copy(&ad)
	err := s.repo.Create(&ad)
	if err != nil {
		return nil, err
	}

	return &ad, nil
}

func (s *ADopingService) Update(req *dto.UpdateADopingReq) (*entities.Antidoping, error) {
	ad, err := s.repo.GetADopingBySmID(req.SmID)
	if err != nil {
		newReq := dto.UpdToCreateAdReq(req)
		return s.Create(newReq)
	}

	req.Copy(ad)

	err = s.repo.Update(ad)
	if err != nil {
		return nil, err
	}

	return ad, nil
}

func (s *ADopingService) GetADopingByID(sportsmanID uuid.UUID) (*entities.Antidoping, error) {
	ad, err := s.repo.GetADopingBySmID(sportsmanID)
	if err != nil {
		return nil, err
	}

	return ad, nil
}

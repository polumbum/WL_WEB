package service

import (
	"src/internal/entities"
	"src/internal/service/dto"
	"src/internal/service/repository"

	"github.com/google/uuid"
)

type IResultService interface {
	Update(req *dto.UpdateResultReq) (*entities.Result, error)
	Create(req *dto.CreateResultReq) (*entities.Result, error)
	ListResults() ([]*entities.Result, error)
	GetResultByID(smID, compID uuid.UUID) (*entities.Result, error)
}

type ResultService struct {
	repo repository.IResultRepository
}

func NewResultService(repo repository.IResultRepository) *ResultService {
	return &ResultService{repo: repo}
}

func (s *ResultService) Update(req *dto.UpdateResultReq) (*entities.Result, error) {
	res, err := s.repo.GetResultByID(req.SportsmanID, req.CompetitionID)
	if err != nil {
		return nil, err
	}
	req.Copy(res)
	err = s.repo.Update(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ResultService) Create(req *dto.CreateResultReq) (*entities.Result, error) {
	var res entities.Result
	req.Copy(&res)
	err := s.repo.Create(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *ResultService) ListResults() ([]*entities.Result, error) {
	results, err := s.repo.ListResults()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *ResultService) GetResultByID(smID, compID uuid.UUID) (*entities.Result, error) {
	res, err := s.repo.GetResultByID(smID, compID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

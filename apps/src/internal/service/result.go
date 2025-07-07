package service

import (
	"src/internal/domain"
	"src/internal/service/repository"

	"github.com/google/uuid"
)

type IResultService interface {
	//Update(req *dto.UpdateResultReq) (*domain.Result, error)
	Create(res *domain.Result) (*domain.Result, error)
	ListResults() ([]*domain.Result, error)
	GetResultByID(smID, compID uuid.UUID) (*domain.Result, error)
}

type ResultService struct {
	repo repository.IResultRepository
}

func NewResultService(repo repository.IResultRepository) *ResultService {
	return &ResultService{repo: repo}
}

/*
func (s *ResultService) Update(req *dto.UpdateResultReq) (*domain.Result, error) {
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
*/

func (s *ResultService) Create(res *domain.Result) (
	*domain.Result,
	error,
) {
	err := s.repo.Create(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ResultService) ListResults() ([]*domain.Result, error) {
	results, err := s.repo.ListResults()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *ResultService) GetResultByID(smID, compID uuid.UUID) (*domain.Result, error) {
	res, err := s.repo.GetResultByID(smID, compID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

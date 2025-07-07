package service

import (
	"src/internal/entities"
	"src/internal/service/dto"
	"src/internal/service/repository"

	"github.com/google/uuid"
)

type ICoachService interface {
	Update(req *dto.UpdateCoachReq) (*entities.Coach, error)
	Create(req *dto.CreateCoachReq) (*entities.Coach, error)
	ListCoaches() ([]*entities.Coach, error)
	GetCoachByID(coachID uuid.UUID) (*entities.Coach, error)
	ListSportsmen(coachID uuid.UUID) ([]*entities.Sportsman, error)
	AddSportsman(coachID uuid.UUID, smID uuid.UUID) (*entities.SportsmenCoach, error)
}

type CoachService struct {
	repo repository.ICoachRepository
}

func NewCoachService(repo repository.ICoachRepository) *CoachService {
	return &CoachService{repo: repo}
}

func (s *CoachService) Update(req *dto.UpdateCoachReq) (*entities.Coach, error) {
	coach, err := s.repo.GetCoachByID(req.ID)
	if err != nil {
		return nil, err
	}
	req.Copy(coach)
	err = s.repo.Update(coach)
	if err != nil {
		return nil, err
	}

	return coach, nil
}

func (s *CoachService) Create(req *dto.CreateCoachReq) (*entities.Coach, error) {
	var coach entities.Coach
	req.Copy(&coach)
	err := s.repo.Create(&coach)
	if err != nil {
		return nil, err
	}

	return &coach, nil
}

func (s *CoachService) ListCoaches() ([]*entities.Coach, error) {
	coaches, err := s.repo.ListCoaches()
	if err != nil {
		return nil, err
	}
	return coaches, nil
}

func (s *CoachService) ListSportsmen(coachID uuid.UUID) ([]*entities.Sportsman, error) {
	sportsmen, err := s.repo.ListSportsmen(coachID)

	if err != nil {
		return nil, err
	}
	return sportsmen, nil
}

func (s *CoachService) GetCoachByID(coachID uuid.UUID) (*entities.Coach, error) {
	coach, err := s.repo.GetCoachByID(coachID)
	if err != nil {
		return nil, err
	}

	return coach, nil
}

func (s *CoachService) AddSportsman(coachID, smID uuid.UUID) (*entities.SportsmenCoach, error) {
	res, err := s.repo.AddSportsman(coachID, smID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

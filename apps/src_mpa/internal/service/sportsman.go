package service

import (
	"os"
	"src/internal/entities"
	"src/internal/service/dto"
	"src/internal/service/repository"
	"time"

	"github.com/google/uuid"
)

type ISportsmanService interface {
	Update(req *dto.UpdateSportsmanReq) (*entities.Sportsman, error)
	Create(req *dto.CreateSportsmanReq) (*entities.Sportsman, error)
	ListSportsmen() ([]*entities.Sportsman, error)
	GetSportsmanByID(sportsmanID uuid.UUID) (*entities.Sportsman, error)
	ListResults(sportsmanID uuid.UUID) ([]*entities.Result, error)
}

type SportsmanService struct {
	repo        repository.ISportsmanRepository
	repoResults repository.IResultRepository
}

func NewSportsmanService(
	repo repository.ISportsmanRepository,
	repoResults repository.IResultRepository,
) *SportsmanService {
	return &SportsmanService{
		repo:        repo,
		repoResults: repoResults,
	}
}

func (s *SportsmanService) Update(req *dto.UpdateSportsmanReq) (*entities.Sportsman, error) {
	sportsman, err := s.repo.GetSportsmanByID(req.ID)
	if err != nil {
		return nil, err
	}

	req.Copy(sportsman)

	err = s.repo.Update(sportsman)
	if err != nil {
		return nil, err
	}

	return sportsman, nil
}

func (s *SportsmanService) Create(req *dto.CreateSportsmanReq) (*entities.Sportsman, error) {
	limPath := os.Getenv("LIM_FILE_PATH")
	config, err := LoadLimits(limPath)
	if err != nil {
		return nil, err
	}
	minAge := config.Limitations.MinAge
	age := GetAge(req.Birthday)
	if age < minAge {
		return nil, ErrYoung
	}

	var sportsman entities.Sportsman
	req.Copy(&sportsman)
	err = s.repo.Create(&sportsman)
	if err != nil {
		return nil, err
	}

	return &sportsman, nil
}

func (s *SportsmanService) ListSportsmen() ([]*entities.Sportsman, error) {
	sportsmen, err := s.repo.ListSportsmen()
	if err != nil {
		return nil, err
	}

	return sportsmen, nil
}

func (s *SportsmanService) GetSportsmanByID(sportsmanID uuid.UUID) (*entities.Sportsman, error) {
	sportsman, err := s.repo.GetSportsmanByID(sportsmanID)
	if err != nil {
		return nil, err
	}

	return sportsman, nil
}

func (s *SportsmanService) ListResults(sportsmanID uuid.UUID) ([]*entities.Result, error) {
	results, err := s.repoResults.ListSportsmanResults(sportsmanID)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func GetAge(birthday time.Time) int {
	currentTime := time.Now()
	diff := currentTime.Sub(birthday)
	age := int(diff.Hours() / 24 / 365)

	return age
}

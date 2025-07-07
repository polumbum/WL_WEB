package service

import (
	"os"
	"src/internal/domain"
	"src/internal/service/repository"
	"time"

	"github.com/google/uuid"
)

type ISportsmanService interface {
	Update(sm *domain.Sportsman) (*domain.Sportsman, error)
	Create(sm *domain.Sportsman) (*domain.Sportsman, error)
	ListSportsmen(
		page int,
		batch int,
		sort string,
		filter string,
	) ([]*domain.Sportsman, error)
	GetSportsmanByID(sportsmanID uuid.UUID) (*domain.Sportsman, error)
	ListResults(sportsmanID uuid.UUID) ([]*domain.Result, error)
	Delete(sportsmanID uuid.UUID) error
}

type SportsmanService struct {
	repo        repository.ISportsmanRepository
	repoResults repository.IResultRepository
	repoAD      repository.IADopingRepository
	repoCA      repository.ICompAccessRepository
}

func NewSportsmanService(
	repo repository.ISportsmanRepository,
	repoResults repository.IResultRepository,
	repoAD repository.IADopingRepository,
	repoCA repository.ICompAccessRepository,
) *SportsmanService {
	return &SportsmanService{
		repo:        repo,
		repoResults: repoResults,
		repoAD:      repoAD,
		repoCA:      repoCA,
	}
}

func (s *SportsmanService) Update(sm *domain.Sportsman) (*domain.Sportsman, error) {
	sm, err := s.repo.GetSportsmanByID(sm.ID)
	if err != nil {
		return nil, err
	}

	ad, err := s.repoAD.GetADopingBySmID(sm.ID)
	if err != nil {
		ad = &domain.Antidoping{SportsmanID: sm.ID}
	}
	ca, err := s.repoCA.GetAccessBySmID(sm.ID)
	if err != nil {
		ca = &domain.CompAccess{SportsmanID: sm.ID}
	}

	sm, err = s.repo.Update(sm)
	if err != nil {
		return nil, err
	}
	_, err = s.repoAD.Update(ad)
	if err != nil {
		return nil, err
	}
	_, err = s.repoCA.Update(ca)
	if err != nil {
		return nil, err
	}

	return sm, nil
}

func (s *SportsmanService) Create(sm *domain.Sportsman) (
	*domain.Sportsman,
	error,
) {
	limPath := os.Getenv("LIM_FILE_PATH")
	config, err := LoadLimits(limPath)
	if err != nil {
		return nil, err
	}
	minAge := config.Limitations.MinAge
	age := GetAge(sm.Birthday)
	if age < minAge {
		return nil, ErrYoung
	}

	sm.ID = uuid.New()
	err = s.repo.Create(sm)
	if err != nil {
		return nil, err
	}

	return sm, nil
}

func (s *SportsmanService) ListSportsmen(
	page int,
	batch int,
	sort string,
	filter string,
) ([]*domain.Sportsman, error) {
	sportsmen, err := s.repo.ListSportsmen(page,
		batch,
		sort,
		filter)
	if err != nil {
		return nil, err
	}

	return sportsmen, nil
}

func (s *SportsmanService) GetSportsmanByID(sportsmanID uuid.UUID) (*domain.Sportsman, error) {
	sportsman, err := s.repo.GetSportsmanByID(sportsmanID)
	if err != nil {
		return nil, err
	}

	return sportsman, nil
}

func (s *SportsmanService) ListResults(sportsmanID uuid.UUID) ([]*domain.Result, error) {
	results, err := s.repoResults.ListSportsmanResults(sportsmanID)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *SportsmanService) Delete(smID uuid.UUID) error {
	err := s.repo.Delete(smID)
	return err
}

func GetAge(birthday time.Time) int {
	currentTime := time.Now()
	diff := currentTime.Sub(birthday)
	age := int(diff.Hours() / 24 / 365)

	return age
}

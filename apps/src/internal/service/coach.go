package service

import (
	"log"
	"src/internal/domain"
	"src/internal/service/repository"

	"github.com/google/uuid"
)

type ICoachService interface {
	Update(coach *domain.Coach) (*domain.Coach, error)
	Create(coach *domain.Coach) (*domain.Coach, error)
	ListCoaches(
		page int,
		batch int,
		sort string,
		filter string,
	) ([]*domain.Coach, error)
	GetCoachByID(coachID uuid.UUID) (*domain.Coach, error)
	ListSportsmen(
		coachID uuid.UUID,
		page int,
		batch int,
		sort string,
		fullname string,
	) ([]*domain.Sportsman, error)
	AddSportsman(coachID uuid.UUID, smID uuid.UUID) (
		*domain.SportsmenCoach,
		error,
	)
	RemoveSportsman(coachID uuid.UUID, smID uuid.UUID) error
	ListResults(coachID uuid.UUID) (
		[]*domain.Sportsman,
		[]*domain.Result,
		error,
	)
	Delete(coachID uuid.UUID) error
}

type CoachService struct {
	repo        repository.ICoachRepository
	repoResults repository.IResultRepository
	repoSm      repository.ISportsmanRepository
}

func NewCoachService(
	repo repository.ICoachRepository,
	repoResults repository.IResultRepository,
	repoSm repository.ISportsmanRepository,
) *CoachService {
	return &CoachService{
		repo:        repo,
		repoResults: repoResults,
		repoSm:      repoSm,
	}
}

func (s *CoachService) Update(coach *domain.Coach) (*domain.Coach, error) {
	err := s.repo.Update(coach)
	if err != nil {
		return nil, err
	}

	return coach, nil
}

func (s *CoachService) Create(coach *domain.Coach) (*domain.Coach, error) {
	coach.ID = uuid.New()
	err := s.repo.Create(coach)
	if err != nil {
		return nil, err
	}

	return coach, nil
}

func (s *CoachService) ListCoaches(
	page int,
	batch int,
	sort string,
	filter string,
) ([]*domain.Coach, error) {
	coaches, err := s.repo.ListCoaches(page,
		batch,
		sort,
		filter)
	if err != nil {
		return nil, err
	}
	return coaches, nil
}

func (s *CoachService) ListSportsmen(
	coachID uuid.UUID,
	page int,
	batch int,
	sort string,
	fullname string,
) ([]*domain.Sportsman, error) {
	sportsmen, err := s.repo.ListSportsmen(
		coachID,
		page,
		batch,
		sort,
		fullname,
	)

	if err != nil {
		return nil, err
	}
	return sportsmen, nil
}

func (s *CoachService) GetCoachByID(coachID uuid.UUID) (*domain.Coach, error) {
	coach, err := s.repo.GetCoachByID(coachID)
	if err != nil {
		return nil, err
	}

	return coach, nil
}

func (s *CoachService) AddSportsman(coachID, smID uuid.UUID) (*domain.SportsmenCoach, error) {
	res, err := s.repo.AddSportsman(coachID, smID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}

func (s *CoachService) ListResults(coachID uuid.UUID) (
	[]*domain.Sportsman,
	[]*domain.Result,
	error,
) {
	results, err := s.repoResults.ListCoachResults(coachID)
	if err != nil {
		return nil, nil, err
	}

	sm := []*domain.Sportsman{}
	for _, item := range results {
		s, err := s.repoSm.GetSportsmanByID(item.SportsmanID)
		if err != nil {
			return nil, nil, err
		}
		sm = append(sm, s)
	}

	return sm, results, nil
}

func (s *CoachService) RemoveSportsman(coachID, smID uuid.UUID) error {
	err := s.repo.RemoveSportsman(coachID, smID)
	if err != nil {
		return err
	}
	return nil
}

func (s *CoachService) Delete(coachID uuid.UUID) error {
	err := s.repo.Delete(coachID)
	return err
}

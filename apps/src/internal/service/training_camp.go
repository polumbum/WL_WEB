package service

import (
	"log"
	"src/internal/domain"
	"src/internal/service/repository"

	"github.com/google/uuid"
)

type ITCampService interface {
	Create(tc *domain.TCamp) (*domain.TCamp, error)
	ListTCamps(
		page int,
		batch int,
		sort string,
		filter string,
	) ([]*domain.TCamp, error)
	GetTCampByID(id uuid.UUID) (*domain.TCamp, error)
	RegisterSportsman(tc *domain.TCampApplication) (
		*domain.TCampApplication,
		error,
	)
	CancelRegistration(smID, tCampID uuid.UUID) error
	GetUpcoming(smID uuid.UUID) ([]*domain.TCamp, error)
	ListUpcoming() ([]*domain.TCamp, error)
	Delete(id uuid.UUID) error
	ListByOrgID(id uuid.UUID) ([]*domain.TCamp, error)
}

type TCampService struct {
	repo          repository.ITCampRepository
	repoSportsman repository.ISportsmanRepository
}

func NewTCampService(
	repo repository.ITCampRepository,
	repoSportsman repository.ISportsmanRepository,
) *TCampService {
	return &TCampService{
		repo:          repo,
		repoSportsman: repoSportsman,
	}
}

/*
func (s *TCampService) Update(req *dto.UpdateTCampReq) (*domain.TCamp, error) {
	camp, err := s.repo.GetTCampByID(req.ID)
	if err != nil {
		return nil, err
	}

	req.Copy(camp)
	err = s.repo.Update(camp)
	if err != nil {
		return nil, err
	}

	return camp, nil
}
*/

func (s *TCampService) Create(tc *domain.TCamp) (*domain.TCamp, error) {
	tc.ID = uuid.New()

	err := s.repo.Create(tc)
	if err != nil {
		return nil, err
	}

	return tc, nil
}

func (s *TCampService) ListTCamps(
	page int,
	batch int,
	sort string,
	filter string,
) ([]*domain.TCamp, error) {
	camps, err := s.repo.ListTCamps(
		page,
		batch,
		sort,
		filter,
	)
	if err != nil {
		return nil, err
	}

	return camps, nil
}

func (s *TCampService) ListByOrgID(id uuid.UUID) (
	[]*domain.TCamp,
	error,
) {
	camps, err := s.repo.ListByOrgID(id)
	if err != nil {
		return nil, err
	}

	return camps, nil
}

func (s *TCampService) GetTCampByID(tCampID uuid.UUID) (*domain.TCamp, error) {
	camp, err := s.repo.GetTCampByID(tCampID)
	if err != nil {
		return nil, err
	}

	return camp, nil
}

func (s *TCampService) RegisterSportsman(tc *domain.TCampApplication) (
	*domain.TCampApplication,
	error,
) {
	_, err := s.repoSportsman.GetSportsmanByID(tc.SportsmanID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = s.repo.GetTCampByID(tc.TCampID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = s.repo.RegisterSportsman(tc)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return tc, nil
}

func (s *TCampService) CancelRegistration(smID, tCampID uuid.UUID) error {
	err := s.repo.DeleteRegistration(smID, tCampID)
	if err != nil {
		return err
	}

	return nil
}

func (s *TCampService) GetUpcoming(smID uuid.UUID) ([]*domain.TCamp, error) {
	tCamps, err := s.repo.GetUpcoming(smID)
	if err != nil {
		return nil, err
	}

	return tCamps, nil
}

func (s *TCampService) ListUpcoming() ([]*domain.TCamp, error) {
	tCamps, err := s.repo.ListUpcoming()
	if err != nil {
		return nil, err
	}

	return tCamps, nil
}

func (s *TCampService) Delete(id uuid.UUID) error {
	err := s.repo.Delete(id)
	return err
}

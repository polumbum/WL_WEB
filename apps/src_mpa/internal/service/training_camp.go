package service

import (
	"src/internal/entities"
	"src/internal/service/dto"
	"src/internal/service/repository"

	"github.com/google/uuid"
)

type ITCampService interface {
	Update(req *dto.UpdateTCampReq) (*entities.TCamp, error)
	Create(req *dto.CreateTCampReq) (*entities.TCamp, error)
	ListTCamps() ([]*entities.TCamp, error)
	GetTCampByID(TCampID uuid.UUID) (*entities.TCamp, error)
	RegisterSportsman(req *dto.RegForTCampReq) (*entities.TCampApplication, error)
	CancelRegistration(smID, tCampID uuid.UUID) error
	GetUpcoming(smID uuid.UUID) ([]*entities.TCamp, error)
	ListUpcoming() ([]*entities.TCamp, error)
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

func (s *TCampService) Update(req *dto.UpdateTCampReq) (*entities.TCamp, error) {
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

func (s *TCampService) Create(req *dto.CreateTCampReq) (*entities.TCamp, error) {
	var camp entities.TCamp
	req.Copy(&camp)
	err := s.repo.Create(&camp)
	if err != nil {
		return nil, err
	}

	return &camp, nil
}

func (s *TCampService) ListTCamps() ([]*entities.TCamp, error) {
	camps, err := s.repo.ListTCamps()
	if err != nil {
		return nil, err
	}

	return camps, nil
}

func (s *TCampService) GetTCampByID(tCampID uuid.UUID) (*entities.TCamp, error) {
	camp, err := s.repo.GetTCampByID(tCampID)
	if err != nil {
		return nil, err
	}

	return camp, nil
}

func (s *TCampService) RegisterSportsman(req *dto.RegForTCampReq) (*entities.TCampApplication,
	error,
) {
	_, err := s.repoSportsman.GetSportsmanByID(req.SportsmanID)
	if err != nil {
		return nil, err
	}
	_, err = s.repo.GetTCampByID(req.TCampID)
	if err != nil {
		return nil, err
	}
	tCampApplication := entities.TCampApplication{
		TCampID:     req.TCampID,
		SportsmanID: req.SportsmanID,
	}

	err = s.repo.RegisterSportsman(&tCampApplication)
	if err != nil {
		return nil, err
	}
	return &tCampApplication, nil
}

func (s *TCampService) CancelRegistration(smID, tCampID uuid.UUID) error {
	err := s.repo.DeleteRegistration(smID, tCampID)
	if err != nil {
		return err
	}

	return nil
}

func (s *TCampService) GetUpcoming(smID uuid.UUID) ([]*entities.TCamp, error) {
	tCamps, err := s.repo.GetUpcoming(smID)
	if err != nil {
		return nil, err
	}

	return tCamps, nil
}

func (s *TCampService) ListUpcoming() ([]*entities.TCamp, error) {
	tCamps, err := s.repo.ListUpcoming()
	if err != nil {
		return nil, err
	}

	return tCamps, nil
}

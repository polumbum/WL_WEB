package service

import (
	"src/internal/constants"
	"src/internal/entities"
	"src/internal/service/dto"
	"src/internal/service/repository"

	"github.com/google/uuid"
)

type ICompetitionService interface {
	Update(req *dto.UpdateCompReq) (*entities.Competition, error)
	Create(req *dto.CreateCompReq) (*entities.Competition, error)
	ListCompetitions() ([]*entities.Competition, error)
	GetCompetitionByID(competitionID uuid.UUID) (*entities.Competition, error)
	RegisterSportsman(req *dto.RegForCompReq) (*entities.CompApplication, error)
	CancelRegistration(smID, compID uuid.UUID) error
	GetUpcoming(smID uuid.UUID) ([]*entities.Competition, error)
	ListUpcoming() ([]*entities.Competition, error)
}

type CompetitionService struct {
	repo           repository.ICompetitionRepository
	repoSportsman  repository.ISportsmanRepository
	repoAntidoping repository.IADopingRepository
	repoAccess     repository.ICompAccessRepository
}

func NewCompetitionService(
	repo repository.ICompetitionRepository,
	repoSportsman repository.ISportsmanRepository,
	repoAntidoping repository.IADopingRepository,
	repoAccess repository.ICompAccessRepository,
) *CompetitionService {
	return &CompetitionService{
		repo:           repo,
		repoSportsman:  repoSportsman,
		repoAntidoping: repoAntidoping,
		repoAccess:     repoAccess,
	}
}

func (s *CompetitionService) Update(req *dto.UpdateCompReq) (*entities.Competition, error) {
	comp, err := s.repo.GetCompetitionByID(req.ID)
	if err != nil {
		return nil, err
	}
	req.Copy(comp)
	err = s.repo.Update(comp)
	if err != nil {
		return nil, err
	}

	return comp, nil
}

func (s *CompetitionService) Create(req *dto.CreateCompReq) (*entities.Competition, error) {
	var comp entities.Competition
	req.Copy(&comp)
	err := s.repo.Create(&comp)
	if err != nil {
		return nil, err
	}

	return &comp, nil
}

func (s *CompetitionService) ListCompetitions() ([]*entities.Competition, error) {
	comps, err := s.repo.ListCompetitions()
	if err != nil {
		return nil, err
	}

	return comps, nil
}

func (s *CompetitionService) GetCompetitionByID(competitionID uuid.UUID) (*entities.Competition,
	error,
) {
	comp, err := s.repo.GetCompetitionByID(competitionID)
	if err != nil {
		return nil, err
	}

	return comp, nil
}

func (s *CompetitionService) RegisterSportsman(req *dto.RegForCompReq) (*entities.CompApplication,
	error,
) {
	sportsman, err := s.repoSportsman.GetSportsmanByID(req.SportsmanID)
	if err != nil {
		return nil, err
	}

	ad, err := s.repoAntidoping.GetADopingBySmID(req.SportsmanID)
	if ad != nil && err != nil {
		return nil, err
	}

	ca, err := s.repoAccess.GetAccessBySmID(req.SportsmanID)
	if ca != nil && err != nil {
		return nil, err
	}

	competition, err := s.repo.GetCompetitionByID(req.CompetitionID)
	if err != nil {
		return nil, err
	}

	err = ValidateCompApplication(ad, ca, sportsman, competition)
	if err != nil {
		return nil, err
	}

	compApplication := entities.CompApplication{
		CompetitionID:     req.CompetitionID,
		SportsmanID:       req.SportsmanID,
		WeightCategory:    req.WeighCategory,
		StartSnatch:       req.StartSnatch,
		StartCleanAndJerk: req.StartCleanAndJerk,
	}
	err = s.repo.RegisterSportsman(&compApplication)
	if err != nil {
		return nil, err
	}

	return &compApplication, nil
}

func (s *CompetitionService) CancelRegistration(smID, compID uuid.UUID) error {
	err := s.repo.DeleteRegistration(smID, compID)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompetitionService) GetUpcoming(smID uuid.UUID) ([]*entities.Competition, error) {
	comps, err := s.repo.GetUpcoming(smID)
	if err != nil {
		return nil, err
	}

	return comps, nil
}

func (s *CompetitionService) ListUpcoming() ([]*entities.Competition, error) {
	comps, err := s.repo.ListUpcoming()
	if err != nil {
		return nil, err
	}

	return comps, nil
}

func ValidateCompApplication(
	ad *entities.Antidoping,
	ca *entities.CompAccess,
	sportsman *entities.Sportsman,
	comp *entities.Competition,
) error {
	if comp.Antidoping && (ad == nil || ad.Validity.Before(comp.BegDate)) { // Antidoping requirement
		return ErrAntidoping
	}

	if ca == nil || ca.Validity.Before(comp.BegDate) { // Access validity
		return ErrAccess
	}

	catRes := constants.CompareSportsCategory(&sportsman.SportsCategory, &comp.MinSportsCategory)
	if catRes < 0 {
		return ErrSportsCat
	}

	age := GetAge(sportsman.Birthday)
	//fmt.Println("age = ", age)
	ageRes := constants.ValidateAgeCategory(age, &comp.Age)
	//fmt.Println("ageRes = ", ageRes)
	if !ageRes {
		//fmt.Println("!ageRes = ", !ageRes)
		return ErrAgeCat
	}
	return nil
}

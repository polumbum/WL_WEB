package service

import (
	"log"
	"src/internal/constants"
	"src/internal/domain"
	"src/internal/service/repository"

	"github.com/google/uuid"
)

type ICompetitionService interface {
	Create(comp *domain.Competition) (*domain.Competition, error)
	ListCompetitions(
		page int,
		batch int,
		sort string,
		filter string,
	) ([]*domain.Competition, error)
	GetCompetitionByID(id uuid.UUID) (*domain.Competition, error)
	RegisterSportsman(ca *domain.CompApplication) (
		*domain.CompApplication,
		error)
	CancelRegistration(smID, compID uuid.UUID) error
	ListCompsByRes([]*domain.Result) ([]*domain.Competition,
		error)
	ListResults(id uuid.UUID) ([]*domain.Result, error)
	Delete(id uuid.UUID) error
	ListByOrgID(id uuid.UUID) ([]*domain.Competition, error)
}

type CompetitionService struct {
	repo           repository.ICompetitionRepository
	repoSportsman  repository.ISportsmanRepository
	repoAntidoping repository.IADopingRepository
	repoAccess     repository.ICompAccessRepository
	repoRes        repository.IResultRepository
}

func NewCompetitionService(
	repo repository.ICompetitionRepository,
	repoSportsman repository.ISportsmanRepository,
	repoAntidoping repository.IADopingRepository,
	repoAccess repository.ICompAccessRepository,
	repoRes repository.IResultRepository,
) *CompetitionService {
	return &CompetitionService{
		repo:           repo,
		repoSportsman:  repoSportsman,
		repoAntidoping: repoAntidoping,
		repoAccess:     repoAccess,
		repoRes:        repoRes,
	}
}

/*
func (s *CompetitionService) Update(req *dto.UpdateCompReq) (*domain.Competition, error) {
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
*/

func (s *CompetitionService) Create(comp *domain.Competition) (
	*domain.Competition,
	error,
) {
	comp.ID = uuid.New()

	err := s.repo.Create(comp)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return comp, nil
}

func (s *CompetitionService) ListCompetitions(
	page int,
	batch int,
	sort string,
	filter string,
) ([]*domain.Competition, error) {
	comps, err := s.repo.ListCompetitions(
		page,
		batch,
		sort,
		filter,
	)
	if err != nil {
		return nil, err
	}

	return comps, nil
}

func (s *CompetitionService) ListByOrgID(id uuid.UUID) (
	[]*domain.Competition,
	error,
) {
	comps, err := s.repo.ListByOrgID(id)
	if err != nil {
		return nil, err
	}

	return comps, nil
}

func (s *CompetitionService) GetCompetitionByID(competitionID uuid.UUID) (
	*domain.Competition,
	error,
) {
	comp, err := s.repo.GetCompetitionByID(competitionID)
	if err != nil {
		return nil, err
	}

	return comp, nil
}

func (s *CompetitionService) RegisterSportsman(appl *domain.CompApplication) (
	*domain.CompApplication,
	error,
) {
	sm, err := s.repoSportsman.GetSportsmanByID(appl.SportsmanID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ad, err := s.repoAntidoping.GetADopingBySmID(appl.SportsmanID)
	if ad != nil && err != nil {
		log.Println(err)
		return nil, err
	}

	ca, err := s.repoAccess.GetAccessBySmID(appl.SportsmanID)
	if ca != nil && err != nil {
		log.Println(err)
		return nil, err
	}

	comp, err := s.repo.GetCompetitionByID(appl.CompetitionID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = ValidateCompApplication(ad, ca, sm, comp)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = s.repo.RegisterSportsman(appl)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return appl, nil
}

func (s *CompetitionService) CancelRegistration(smID, compID uuid.UUID) error {
	err := s.repo.DeleteRegistration(smID, compID)
	if err != nil {
		return err
	}

	return nil
}

/*
func (s *CompetitionService) GetUpcoming(smID uuid.UUID) (
	[]*domain.Competition,
	error,
) {
	comps, err := s.repo.GetUpcoming(smID)
	if err != nil {
		return nil, err
	}

	return comps, nil
}

func (s *CompetitionService) ListUpcoming() ([]*domain.Competition, error) {
	comps, err := s.repo.ListUpcoming()
	if err != nil {
		return nil, err
	}

	return comps, nil
}*/

func (s *CompetitionService) ListCompsByRes(res []*domain.Result) (
	[]*domain.Competition,
	error,
) {
	list := []*domain.Competition{}
	for _, item := range res {
		if item == nil {
			return nil, ErrNilRef
		}
		comp, err := s.repo.GetCompetitionByID(item.CompetitionID)
		if err != nil {
			return nil, err
		}
		list = append(list, comp)
	}
	return list, nil
}

func (s *CompetitionService) ListResults(id uuid.UUID) (
	[]*domain.Result,
	error,
) {
	res, err := s.repoRes.ListCompResults(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CompetitionService) Delete(id uuid.UUID) error {
	err := s.repo.Delete(id)
	return err
}

func ValidateCompApplication(
	ad *domain.Antidoping,
	ca *domain.CompAccess,
	sportsman *domain.Sportsman,
	comp *domain.Competition,
) error {
	if comp.Antidoping && (ad == nil || ad.Validity.Before(comp.BegDate)) { // Antidoping requirement
		return ErrAntidoping
	}

	if ca == nil || ca.Validity.Before(comp.BegDate) { // Access validity
		return ErrAccess
	}

	catRes := constants.CompareSportsCategory(
		&sportsman.SportsCategory,
		&comp.MinSportsCategory,
	)
	if catRes < 0 {
		return ErrSportsCat
	}

	age := GetAge(sportsman.Birthday)
	log.Println("age = ", age)
	log.Println("age comp = ", comp.Age)
	ageRes := constants.ValidateAgeCategory(age, &comp.Age)
	log.Println("ageRes = ", ageRes)
	if !ageRes {
		//fmt.Println("!ageRes = ", !ageRes)
		return ErrAgeCat
	}
	return nil
}

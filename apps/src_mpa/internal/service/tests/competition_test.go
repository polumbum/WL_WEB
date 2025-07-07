package service_test

import (
	"testing"
	"time"

	"src/internal/constants"
	"src/internal/entities"

	"src/internal/service"
	"src/internal/service/dto"
	"src/internal/service/repository/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type CompetitionServiceTestSuite struct {
	suite.Suite
	mockRepo          *mocks.ICompetitionRepository
	mockRepoSportsman *mocks.ISportsmanRepository
	mockRepoAccess    *mocks.ICompAccessRepository
	mockRepoADop      *mocks.IADopingRepository
	service           service.ICompetitionService
}

func (suite *CompetitionServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewICompetitionRepository(suite.T())
	suite.mockRepoSportsman = mocks.NewISportsmanRepository(suite.T())
	suite.mockRepoAccess = mocks.NewICompAccessRepository(suite.T())
	suite.mockRepoADop = mocks.NewIADopingRepository(suite.T())
	suite.service = service.NewCompetitionService(suite.mockRepo,
		suite.mockRepoSportsman,
		suite.mockRepoADop,
		suite.mockRepoAccess)
}

func TestCompetitionServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CompetitionServiceTestSuite))
}

// Update.
func (suite *CompetitionServiceTestSuite) TestUpdateGetCompetitionByIDFail() {
	req := &dto.UpdateCompReq{
		ID:                uuid.New(),
		Name:              "ABC",
		City:              "DEF",
		Address:           "GHI",
		BegDate:           time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategory1,
	}

	suite.mockRepo.On("GetCompetitionByID", req.ID).
		Return(nil, ErrServiceTest).Times(1)

	competition, err := suite.service.Update(req)

	suite.Nil(competition)
	suite.Error(err)
}

func (suite *CompetitionServiceTestSuite) TestUpdateFail() {
	id := uuid.New()

	req := &dto.UpdateCompReq{
		ID:                id,
		Name:              "ABC",
		City:              "DEF",
		Address:           "GHI",
		BegDate:           time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategory1,
	}

	expectedComp := &entities.Competition{
		ID:                id,
		Name:              "ABC",
		City:              "DEF",
		Address:           "GHI",
		BegDate:           time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategory1,
	}

	suite.mockRepo.On("GetCompetitionByID", req.ID).
		Return(expectedComp, nil).Times(1)

	suite.mockRepo.On("Update", expectedComp).
		Return(ErrServiceTest).Times(1)

	competition, err := suite.service.Update(req)

	suite.Nil(competition)
	suite.Error(err)
}

func (suite *CompetitionServiceTestSuite) TestUpdateSuccess() {
	id := uuid.New()

	req := &dto.UpdateCompReq{
		ID:                id,
		Name:              "ABC",
		City:              "DEF",
		Address:           "GHI",
		BegDate:           time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategory1,
	}

	expectedComp := &entities.Competition{
		ID:                id,
		Name:              "ABC",
		City:              "DEF",
		Address:           "GHI",
		BegDate:           time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategory1,
	}

	suite.mockRepo.On("GetCompetitionByID", req.ID).
		Return(expectedComp, nil).Times(1)

	suite.mockRepo.On("Update", expectedComp).
		Return(nil).Times(1)

	competition, err := suite.service.Update(req)

	suite.Equal(expectedComp, competition)
	suite.NoError(err)
}

// Create.
func (suite *CompetitionServiceTestSuite) TestCreateFail() {
	req := &dto.CreateCompReq{
		Name:              "ABC",
		City:              "DEF",
		Address:           "GHI",
		BegDate:           time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategory1,
	}

	expectedComp := &entities.Competition{
		Name:              "ABC",
		City:              "DEF",
		Address:           "GHI",
		BegDate:           time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategory1,
	}

	suite.mockRepo.On("Create", expectedComp).
		Return(ErrServiceTest).Times(1)

	competition, err := suite.service.Create(req)

	suite.Nil(competition)
	suite.Error(err)
}

func (suite *CompetitionServiceTestSuite) TestCreateSuccess() {
	req := &dto.CreateCompReq{
		Name:              "ABC",
		City:              "DEF",
		Address:           "GHI",
		BegDate:           time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategory1,
	}

	expectedComp := &entities.Competition{
		Name:              "ABC",
		City:              "DEF",
		Address:           "GHI",
		BegDate:           time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategory1,
	}

	suite.mockRepo.On("Create", expectedComp).
		Return(nil).Times(1)

	competition, err := suite.service.Create(req)

	suite.Equal(expectedComp, competition)
	suite.NoError(err)
}

// ListCompetitions.
func (suite *CompetitionServiceTestSuite) TestListCompetitionsFail() {
	suite.mockRepo.On("ListCompetitions").
		Return(nil, ErrServiceTest).Times(1)

	competitiones, err := suite.service.ListCompetitions()

	suite.Nil(competitiones)
	suite.Error(err)
}

func (suite *CompetitionServiceTestSuite) TestListCompetitionsSuccess() {
	id := uuid.New()

	expectedComp := &entities.Competition{
		ID:                id,
		Name:              "ABC",
		City:              "DEF",
		Address:           "GHI",
		BegDate:           time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategory1,
	}
	expectedComps := []*entities.Competition{expectedComp}

	suite.mockRepo.On("ListCompetitions").
		Return(expectedComps, nil).Times(1)

	competitions, err := suite.service.ListCompetitions()

	suite.Equal(expectedComps, competitions)
	suite.NoError(err)
}

// GetCompetitionByID.
func (suite *CompetitionServiceTestSuite) TestGetCompetitionByIDFail() {
	id := uuid.New()

	suite.mockRepo.On("GetCompetitionByID", id).
		Return(nil, ErrServiceTest).Times(1)

	competition, err := suite.service.GetCompetitionByID(id)

	suite.Nil(competition)
	suite.Error(err)
}

func (suite *CompetitionServiceTestSuite) TestGetCompetitionByIDSuccess() {
	id := uuid.New()

	expectedComp := &entities.Competition{
		ID:                id,
		Name:              "ABC",
		City:              "DEF",
		Address:           "GHI",
		BegDate:           time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategory1,
	}

	suite.mockRepo.On("GetCompetitionByID", id).
		Return(expectedComp, nil).Times(1)

	competition, err := suite.service.GetCompetitionByID(id)

	suite.Equal(expectedComp, competition)
	suite.NoError(err)
}

// RegisterSportsman.
func (suite *CompetitionServiceTestSuite) TestRegisterGetSportsmanByIDFail() {
	compID := uuid.New()
	sportsmanID := uuid.New()

	req := &dto.RegForCompReq{
		CompetitionID:     compID,
		SportsmanID:       sportsmanID,
		WeighCategory:     constants.WC59,
		StartSnatch:       60,
		StartCleanAndJerk: 70,
	}

	suite.mockRepoSportsman.On("GetSportsmanByID", req.SportsmanID).
		Return(nil, ErrServiceTest).Times(1)

	compApp, err := suite.service.RegisterSportsman(req)

	suite.Nil(compApp)
	suite.Error(err)
}

func (suite *CompetitionServiceTestSuite) TestRegisterGetCompetitionByIDFail() {
	compID := uuid.New()
	sportsmanID := uuid.New()
	caID := uuid.New()
	adID := uuid.New()

	req := &dto.RegForCompReq{
		CompetitionID:     compID,
		SportsmanID:       sportsmanID,
		WeighCategory:     constants.WC59,
		StartSnatch:       60,
		StartCleanAndJerk: 70,
	}

	expectedSportsman := &entities.Sportsman{
		ID:             sportsmanID,
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS,
		Gender:         true,
	}

	expectedADoping := &entities.Antidoping{
		ID:          adID,
		SportsmanID: sportsmanID,
		Validity:    time.Date(2025, time.November, 10, 0, 0, 0, 0, time.UTC),
	}

	expectedAccess := &entities.CompAccess{
		ID:          caID,
		SportsmanID: sportsmanID,
		Validity:    time.Date(2025, time.November, 10, 0, 0, 0, 0, time.UTC),
		Institution: "ABC",
	}

	suite.mockRepoSportsman.On("GetSportsmanByID", req.SportsmanID).
		Return(expectedSportsman, nil).Times(1)

	suite.mockRepoADop.On("GetADopingBySmID", req.SportsmanID).
		Return(expectedADoping, nil).Times(1)

	suite.mockRepoAccess.On("GetAccessBySmID", req.SportsmanID).
		Return(expectedAccess, nil).Times(1)

	suite.mockRepo.On("GetCompetitionByID", req.CompetitionID).
		Return(nil, ErrServiceTest).Times(1)

	compApp, err := suite.service.RegisterSportsman(req)

	suite.Nil(compApp)
	suite.Error(err)
}

func (suite *CompetitionServiceTestSuite) TestWrongSportsCat() {
	compID := uuid.New()
	sportsmanID := uuid.New()
	caID := uuid.New()
	adID := uuid.New()

	req := &dto.RegForCompReq{
		CompetitionID:     compID,
		SportsmanID:       sportsmanID,
		WeighCategory:     constants.WC59,
		StartSnatch:       60,
		StartCleanAndJerk: 70,
	}

	expectedSportsman := &entities.Sportsman{
		ID:             sportsmanID,
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(2002, time.November, 10, 0, 0, 0, 0, time.UTC), // 21 now
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS, // КМС
		Gender:         true,
	}

	expectedComp := &entities.Competition{
		ID:                compID,
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategoryMS, // МС
		Antidoping:        true,
	}

	expectedADoping := &entities.Antidoping{
		ID:          adID,
		SportsmanID: sportsmanID,
		Validity:    time.Date(2025, time.November, 10, 0, 0, 0, 0, time.UTC),
	}

	expectedAccess := &entities.CompAccess{
		ID:          caID,
		SportsmanID: sportsmanID,
		Validity:    time.Date(2025, time.November, 10, 0, 0, 0, 0, time.UTC),
		Institution: "ABC",
	}

	suite.mockRepoSportsman.On("GetSportsmanByID", req.SportsmanID).
		Return(expectedSportsman, nil).Times(1)

	suite.mockRepoADop.On("GetADopingBySmID", req.SportsmanID).
		Return(expectedADoping, nil).Times(1)

	suite.mockRepoAccess.On("GetAccessBySmID", req.SportsmanID).
		Return(expectedAccess, nil).Times(1)

	suite.mockRepo.On("GetCompetitionByID", req.CompetitionID).
		Return(expectedComp, nil).Times(1)

	compApp, err := suite.service.RegisterSportsman(req)

	suite.Nil(compApp)
	suite.Equal(service.ErrSportsCat, err)
}

func (suite *CompetitionServiceTestSuite) TestWrongAgeCat() {
	compID := uuid.New()
	sportsmanID := uuid.New()
	caID := uuid.New()
	adID := uuid.New()

	req := &dto.RegForCompReq{
		CompetitionID:     compID,
		SportsmanID:       sportsmanID,
		WeighCategory:     constants.WC59,
		StartSnatch:       60,
		StartCleanAndJerk: 70,
	}

	expectedSportsman := &entities.Sportsman{
		ID:             sportsmanID,
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS, // КМС
		Gender:         true,
	}

	expectedComp := &entities.Competition{
		ID:                compID,
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}

	expectedADoping := &entities.Antidoping{
		ID:          adID,
		SportsmanID: sportsmanID,
		Validity:    time.Date(2025, time.November, 10, 0, 0, 0, 0, time.UTC),
	}

	expectedAccess := &entities.CompAccess{
		ID:          caID,
		SportsmanID: sportsmanID,
		Validity:    time.Date(2025, time.November, 10, 0, 0, 0, 0, time.UTC),
		Institution: "ABC",
	}

	suite.mockRepoSportsman.On("GetSportsmanByID", req.SportsmanID).
		Return(expectedSportsman, nil).Times(1)

	suite.mockRepoADop.On("GetADopingBySmID", req.SportsmanID).
		Return(expectedADoping, nil).Times(1)

	suite.mockRepoAccess.On("GetAccessBySmID", req.SportsmanID).
		Return(expectedAccess, nil).Times(1)

	suite.mockRepo.On("GetCompetitionByID", req.CompetitionID).
		Return(expectedComp, nil).Times(1)

	compApp, err := suite.service.RegisterSportsman(req)

	suite.Nil(compApp)
	suite.Equal(service.ErrAgeCat, err)
}

func (suite *CompetitionServiceTestSuite) TestOkAgeCat() {
	compID := uuid.New()
	sportsmanID := uuid.New()
	caID := uuid.New()
	adID := uuid.New()

	req := &dto.RegForCompReq{
		CompetitionID:     compID,
		SportsmanID:       sportsmanID,
		WeighCategory:     constants.WC59,
		StartSnatch:       60,
		StartCleanAndJerk: 70,
	}

	expectedSportsman := &entities.Sportsman{
		ID:             sportsmanID,
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(2000, time.September, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS, // КМС
		Gender:         true,
	}

	expectedComp := &entities.Competition{
		ID:                compID,
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryMW,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}

	expectedADoping := &entities.Antidoping{
		ID:          adID,
		SportsmanID: sportsmanID,
		Validity:    time.Date(2025, time.November, 10, 0, 0, 0, 0, time.UTC),
	}

	expectedAccess := &entities.CompAccess{
		ID:          caID,
		SportsmanID: sportsmanID,
		Validity:    time.Date(2025, time.November, 10, 0, 0, 0, 0, time.UTC),
		Institution: "ABC",
	}

	expectedCompApp := &entities.CompApplication{
		CompetitionID:     req.CompetitionID,
		SportsmanID:       req.SportsmanID,
		WeightCategory:    req.WeighCategory,
		StartSnatch:       req.StartSnatch,
		StartCleanAndJerk: req.StartCleanAndJerk,
	}

	suite.mockRepoSportsman.On("GetSportsmanByID", req.SportsmanID).
		Return(expectedSportsman, nil).Times(1)

	suite.mockRepoADop.On("GetADopingBySmID", req.SportsmanID).
		Return(expectedADoping, nil).Times(1)

	suite.mockRepoAccess.On("GetAccessBySmID", req.SportsmanID).
		Return(expectedAccess, nil).Times(1)

	suite.mockRepo.On("GetCompetitionByID", req.CompetitionID).
		Return(expectedComp, nil).Times(1)

	suite.mockRepo.On("RegisterSportsman", expectedCompApp).
		Return(nil).Times(1)

	compApp, err := suite.service.RegisterSportsman(req)

	//suite.Equal(expectedCompApp, compApp)
	suite.NotNil(compApp)
	suite.NoError(err)
}

func (suite *CompetitionServiceTestSuite) TestNoAntidoping() {
	compID := uuid.New()
	sportsmanID := uuid.New()
	caID := uuid.New()

	req := &dto.RegForCompReq{
		CompetitionID:     compID,
		SportsmanID:       sportsmanID,
		WeighCategory:     constants.WC59,
		StartSnatch:       60,
		StartCleanAndJerk: 70,
	}

	expectedSportsman := &entities.Sportsman{
		ID:             sportsmanID,
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS, // КМС
		Gender:         true,
	}

	expectedComp := &entities.Competition{
		ID:                compID,
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}

	expectedAccess := &entities.CompAccess{
		ID:          caID,
		SportsmanID: sportsmanID,
		Validity:    time.Date(2025, time.November, 10, 0, 0, 0, 0, time.UTC),
		Institution: "ABC",
	}

	suite.mockRepoSportsman.On("GetSportsmanByID", req.SportsmanID).
		Return(expectedSportsman, nil).Times(1)

	suite.mockRepoADop.On("GetADopingBySmID", req.SportsmanID).
		Return(nil, ErrServiceTest).Times(1)

	suite.mockRepoAccess.On("GetAccessBySmID", req.SportsmanID).
		Return(expectedAccess, nil).Times(1)

	suite.mockRepo.On("GetCompetitionByID", req.CompetitionID).
		Return(expectedComp, nil).Times(1)

	compApp, err := suite.service.RegisterSportsman(req)

	suite.Nil(compApp)
	suite.Equal(service.ErrAntidoping, err)
}

func (suite *CompetitionServiceTestSuite) TestNoAccess() {
	compID := uuid.New()
	sportsmanID := uuid.New()
	adID := uuid.New()

	req := &dto.RegForCompReq{
		CompetitionID:     compID,
		SportsmanID:       sportsmanID,
		WeighCategory:     constants.WC59,
		StartSnatch:       60,
		StartCleanAndJerk: 70,
	}

	expectedSportsman := &entities.Sportsman{
		ID:             sportsmanID,
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS, // КМС
		Gender:         true,
	}

	expectedComp := &entities.Competition{
		ID:                compID,
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}

	expectedADoping := &entities.Antidoping{
		ID:          adID,
		SportsmanID: sportsmanID,
		Validity:    time.Date(2025, time.November, 10, 0, 0, 0, 0, time.UTC),
	}

	suite.mockRepoSportsman.On("GetSportsmanByID", req.SportsmanID).
		Return(expectedSportsman, nil).Times(1)

	suite.mockRepoADop.On("GetADopingBySmID", req.SportsmanID).
		Return(expectedADoping, nil).Times(1)

	suite.mockRepoAccess.On("GetAccessBySmID", req.SportsmanID).
		Return(nil, ErrServiceTest).Times(1)

	suite.mockRepo.On("GetCompetitionByID", req.CompetitionID).
		Return(expectedComp, nil).Times(1)

	compApp, err := suite.service.RegisterSportsman(req)

	suite.Nil(compApp)
	suite.Equal(service.ErrAccess, err)
}

func (suite *CompetitionServiceTestSuite) TestRegisterFail() {
	compID := uuid.New()
	sportsmanID := uuid.New()
	caID := uuid.New()
	adID := uuid.New()

	req := &dto.RegForCompReq{
		CompetitionID:     compID,
		SportsmanID:       sportsmanID,
		WeighCategory:     constants.WC59,
		StartSnatch:       60,
		StartCleanAndJerk: 70,
	}

	expectedSportsman := &entities.Sportsman{
		ID:             sportsmanID,
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS, // КМС
		Gender:         true,
	}

	expectedADoping := &entities.Antidoping{
		ID:          adID,
		SportsmanID: sportsmanID,
		Validity:    time.Date(2025, time.November, 10, 0, 0, 0, 0, time.UTC),
	}

	expectedAccess := &entities.CompAccess{
		ID:          caID,
		SportsmanID: sportsmanID,
		Validity:    time.Date(2025, time.November, 10, 0, 0, 0, 0, time.UTC),
		Institution: "ABC",
	}

	expectedComp := &entities.Competition{
		ID:                compID,
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryMW,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}

	expectedCompApp := &entities.CompApplication{
		CompetitionID:     req.CompetitionID,
		SportsmanID:       req.SportsmanID,
		WeightCategory:    req.WeighCategory,
		StartSnatch:       req.StartSnatch,
		StartCleanAndJerk: req.StartCleanAndJerk,
	}

	suite.mockRepoSportsman.On("GetSportsmanByID", req.SportsmanID).
		Return(expectedSportsman, nil).Times(1)

	suite.mockRepoADop.On("GetADopingBySmID", req.SportsmanID).
		Return(expectedADoping, nil).Times(1)

	suite.mockRepoAccess.On("GetAccessBySmID", req.SportsmanID).
		Return(expectedAccess, nil).Times(1)

	suite.mockRepo.On("GetCompetitionByID", req.CompetitionID).
		Return(expectedComp, nil).Times(1)

	suite.mockRepo.On("RegisterSportsman", expectedCompApp).
		Return(ErrServiceTest).Times(1)

	compApp, err := suite.service.RegisterSportsman(req)

	suite.Nil(compApp)
	suite.Error(err)
}

func (suite *CompetitionServiceTestSuite) TestRegisterSuccess() {
	compID := uuid.New()
	sportsmanID := uuid.New()
	caID := uuid.New()
	adID := uuid.New()

	req := &dto.RegForCompReq{
		CompetitionID:     compID,
		SportsmanID:       sportsmanID,
		WeighCategory:     constants.WC59,
		StartSnatch:       60,
		StartCleanAndJerk: 70,
	}

	expectedSportsman := &entities.Sportsman{
		ID:             sportsmanID,
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS, // КМС
		Gender:         true,
	}

	expectedADoping := &entities.Antidoping{
		ID:          adID,
		SportsmanID: sportsmanID,
		Validity:    time.Date(2025, time.November, 10, 0, 0, 0, 0, time.UTC),
	}

	expectedAccess := &entities.CompAccess{
		ID:          caID,
		SportsmanID: sportsmanID,
		Validity:    time.Date(2025, time.November, 10, 0, 0, 0, 0, time.UTC),
		Institution: "ABC",
	}

	expectedComp := &entities.Competition{
		ID:                compID,
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryMW,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}

	expectedCompApp := &entities.CompApplication{
		CompetitionID:     req.CompetitionID,
		SportsmanID:       req.SportsmanID,
		WeightCategory:    req.WeighCategory,
		StartSnatch:       req.StartSnatch,
		StartCleanAndJerk: req.StartCleanAndJerk,
	}

	suite.mockRepoSportsman.On("GetSportsmanByID", req.SportsmanID).
		Return(expectedSportsman, nil).Times(1)

	suite.mockRepoADop.On("GetADopingBySmID", req.SportsmanID).
		Return(expectedADoping, nil).Times(1)

	suite.mockRepoAccess.On("GetAccessBySmID", req.SportsmanID).
		Return(expectedAccess, nil).Times(1)

	suite.mockRepo.On("GetCompetitionByID", req.CompetitionID).
		Return(expectedComp, nil).Times(1)

	suite.mockRepo.On("RegisterSportsman", expectedCompApp).
		Return(nil).Times(1)

	compApp, err := suite.service.RegisterSportsman(req)

	suite.Equal(expectedCompApp, compApp)
	suite.NoError(err)
}

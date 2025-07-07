package service_test

import (
	"os"
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

type SportsmanServiceTestSuite struct {
	suite.Suite
	mockRepo       *mocks.ISportsmanRepository
	mockRepoResult *mocks.IResultRepository
	mockRepoAccess *mocks.ICompAccessRepository
	mockRepoADop   *mocks.IADopingRepository
	service        service.ISportsmanService
}

func (suite *SportsmanServiceTestSuite) SetupTest() {
	os.Setenv("LIM_FILE_PATH", "../limitations.json")
	suite.mockRepo = mocks.NewISportsmanRepository(suite.T())
	suite.mockRepoResult = mocks.NewIResultRepository(suite.T())
	suite.mockRepoAccess = mocks.NewICompAccessRepository(suite.T())
	suite.mockRepoADop = mocks.NewIADopingRepository(suite.T())
	suite.service = service.NewSportsmanService(suite.mockRepo,
		suite.mockRepoResult)
}

func TestSportsmanServiceTestSuite(t *testing.T) {
	suite.Run(t, new(SportsmanServiceTestSuite))
}

// Update.
func (suite *SportsmanServiceTestSuite) TestUpdateGetSportsmanByIDFail() {
	id := uuid.New()
	boolval := true
	boolp := &boolval
	req := &dto.UpdateSportsmanReq{
		ID: id,
		/*Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),*/
		MoscowTeam:     boolp,
		SportsCategory: constants.SportsCategoryCMS,
		//Gender:         true,
	}

	suite.mockRepo.On("GetSportsmanByID", req.ID).
		Return(nil, ErrServiceTest).Times(1)

	sportsman, err := suite.service.Update(req)

	suite.Nil(sportsman)
	suite.Error(err)
}

func (suite *SportsmanServiceTestSuite) TestUpdateFail() {
	id := uuid.New()
	boolval := true
	boolp := &boolval
	req := &dto.UpdateSportsmanReq{
		ID: id,
		/*Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),*/
		MoscowTeam:     boolp,
		SportsCategory: constants.SportsCategoryCMS,
		//Gender:         true,
	}

	expectedSportsman := &entities.Sportsman{
		ID:             req.ID,
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     false,
		SportsCategory: constants.SportsCategoryCMS,
		Gender:         true,
	}

	expectedSportsmanUpd := &entities.Sportsman{
		ID:             req.ID,
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true, // upd
		SportsCategory: constants.SportsCategoryCMS,
		Gender:         true,
	}

	suite.mockRepo.On("GetSportsmanByID", req.ID).
		Return(expectedSportsman, nil).Times(1)

	suite.mockRepo.On("Update", expectedSportsmanUpd).
		Return(ErrServiceTest).Times(1)

	sportsman, err := suite.service.Update(req)

	suite.Nil(sportsman)
	suite.Error(err)
}

func (suite *SportsmanServiceTestSuite) TestUpdateSuccess() {
	id := uuid.New()
	boolval := true
	boolp := &boolval
	req := &dto.UpdateSportsmanReq{
		ID: id,
		/*Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),*/
		MoscowTeam:     boolp,
		SportsCategory: constants.SportsCategoryCMS,
		//Gender:         true,
	}

	expectedSportsman := &entities.Sportsman{
		ID:             req.ID,
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     false,
		SportsCategory: constants.SportsCategoryCMS,
		Gender:         true,
	}

	expectedSportsmanUpd := &entities.Sportsman{
		ID:             req.ID,
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true, // upd
		SportsCategory: constants.SportsCategoryCMS,
		Gender:         true,
	}

	suite.mockRepo.On("GetSportsmanByID", req.ID).
		Return(expectedSportsman, nil).Times(1)

	suite.mockRepo.On("Update", expectedSportsmanUpd).
		Return(nil).Times(1)

	sportsman, err := suite.service.Update(req)

	suite.Equal(expectedSportsmanUpd, sportsman)
	suite.NoError(err)
}

// Create.
func (suite *SportsmanServiceTestSuite) TestCreateFail() {
	req := &dto.CreateSportsmanReq{
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS,
		Gender:         true,
	}

	expectedSportsman := &entities.Sportsman{
		Surname:        req.Surname,
		Name:           req.Name,
		Patronymic:     req.Patronymic,
		Birthday:       req.Birthday,
		MoscowTeam:     req.MoscowTeam,
		SportsCategory: req.SportsCategory,
		Gender:         req.Gender,
	}

	suite.mockRepo.On("Create", expectedSportsman).
		Return(ErrServiceTest).Times(1)

	sportsman, err := suite.service.Create(req)

	suite.Nil(sportsman)
	suite.Error(err)
}

func (suite *SportsmanServiceTestSuite) TestCreateSuccess() {
	req := &dto.CreateSportsmanReq{
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS,
		Gender:         true,
	}

	expectedSportsman := &entities.Sportsman{
		Surname:        req.Surname,
		Name:           req.Name,
		Patronymic:     req.Patronymic,
		Birthday:       req.Birthday,
		MoscowTeam:     req.MoscowTeam,
		SportsCategory: req.SportsCategory,
		Gender:         req.Gender,
	}

	suite.mockRepo.On("Create", expectedSportsman).
		Return(nil).Times(1)

	sportsman, err := suite.service.Create(req)

	suite.Equal(expectedSportsman, sportsman)
	suite.NoError(err)
}

// ListSportsmen.
func (suite *SportsmanServiceTestSuite) TestListSportsmenFail() {
	suite.mockRepo.On("ListSportsmen").
		Return(nil, ErrServiceTest).Times(1)

	sportsmen, err := suite.service.ListSportsmen()

	suite.Nil(sportsmen)
	suite.Error(err)
}

func (suite *SportsmanServiceTestSuite) TestListSportsmenSuccess() {
	id := uuid.New()

	expectedSportsman := &entities.Sportsman{
		ID:             id,
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS,
		Gender:         true,
	}
	expectedComps := []*entities.Sportsman{expectedSportsman}

	suite.mockRepo.On("ListSportsmen").
		Return(expectedComps, nil).Times(1)

	sportsmen, err := suite.service.ListSportsmen()

	suite.Equal(expectedComps, sportsmen)
	suite.NoError(err)
}

// GetSportsmanByID.
func (suite *SportsmanServiceTestSuite) TestGetSportsmanByIDFail() {
	id := uuid.New()

	suite.mockRepo.On("GetSportsmanByID", id).
		Return(nil, ErrServiceTest).Times(1)

	sportsman, err := suite.service.GetSportsmanByID(id)

	suite.Nil(sportsman)
	suite.Error(err)
}

func (suite *SportsmanServiceTestSuite) TestGetSportsmanByIDSuccess() {
	id := uuid.New()

	expectedSportsman := &entities.Sportsman{
		ID:             id,
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS,
		Gender:         true,
	}

	suite.mockRepo.On("GetSportsmanByID", id).
		Return(expectedSportsman, nil).Times(1)

	sportsman, err := suite.service.GetSportsmanByID(id)

	suite.Equal(expectedSportsman, sportsman)
	suite.NoError(err)
}

// ListResults.
func (suite *SportsmanServiceTestSuite) TestListResultsFail() {
	sportsmanID := uuid.New()

	suite.mockRepoResult.On("ListSportsmanResults", sportsmanID).
		Return(nil, ErrServiceTest).Times(1)

	sportsman, err := suite.service.ListResults(sportsmanID)

	suite.Nil(sportsman)
	suite.Error(err)
}

func (suite *SportsmanServiceTestSuite) TestListResultsSuccess() {
	sportsmanID := uuid.New()

	expectedResult := &entities.Result{
		WeightCategory: constants.WC59,
		CompetitionID:  uuid.New(),
		SportsmanID:    sportsmanID,
		Snatch:         70,
		CleanAndJerk:   96,
		Place:          1,
	}
	expectedResults := []*entities.Result{expectedResult}

	suite.mockRepoResult.On("ListSportsmanResults", sportsmanID).
		Return(expectedResults, nil).Times(1)

	results, err := suite.service.ListResults(sportsmanID)

	suite.Equal(expectedResults, results)
	suite.NoError(err)
}

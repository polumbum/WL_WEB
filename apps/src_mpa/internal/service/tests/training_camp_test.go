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

type TCampServiceTestSuite struct {
	suite.Suite
	mockRepo          *mocks.ITCampRepository
	mockRepoSportsman *mocks.ISportsmanRepository
	service           service.ITCampService
}

func (suite *TCampServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewITCampRepository(suite.T())
	suite.mockRepoSportsman = mocks.NewISportsmanRepository(suite.T())
	suite.service = service.NewTCampService(suite.mockRepo, suite.mockRepoSportsman)
}

func TestTCampServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TCampServiceTestSuite))
}

// Update.
func (suite *TCampServiceTestSuite) TestUpdateGetTCampByIDFail() {
	req := &dto.UpdateTCampReq{
		ID:      uuid.New(),
		City:    "DEF",
		Address: "GHI",
		BegDate: time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	suite.mockRepo.On("GetTCampByID", req.ID).
		Return(nil, ErrServiceTest).Times(1)

	trainingCamp, err := suite.service.Update(req)

	suite.Nil(trainingCamp)
	suite.Error(err)
}

func (suite *TCampServiceTestSuite) TestUpdateFail() {
	id := uuid.New()

	req := &dto.UpdateTCampReq{
		ID:      uuid.New(),
		City:    "DEF",
		Address: "GHI",
		BegDate: time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	expectedComp := &entities.TCamp{
		ID:      id,
		City:    "DEF",
		Address: "GHI",
		BegDate: time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	suite.mockRepo.On("GetTCampByID", req.ID).
		Return(expectedComp, nil).Times(1)

	suite.mockRepo.On("Update", expectedComp).
		Return(ErrServiceTest).Times(1)

	trainingCamp, err := suite.service.Update(req)

	suite.Nil(trainingCamp)
	suite.Error(err)
}

func (suite *TCampServiceTestSuite) TestUpdateSuccess() {
	id := uuid.New()

	req := &dto.UpdateTCampReq{
		ID:      uuid.New(),
		City:    "DEF",
		Address: "GHI",
		BegDate: time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	expectedComp := &entities.TCamp{
		ID:      id,
		City:    "DEF",
		Address: "GHI",
		BegDate: time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	suite.mockRepo.On("GetTCampByID", req.ID).
		Return(expectedComp, nil).Times(1)

	suite.mockRepo.On("Update", expectedComp).
		Return(nil).Times(1)

	trainingCamp, err := suite.service.Update(req)

	suite.Equal(expectedComp, trainingCamp)
	suite.NoError(err)
}

// Create.
func (suite *TCampServiceTestSuite) TestCreateFail() {
	req := &dto.CreateTCampReq{
		City:    "DEF",
		Address: "GHI",
		BegDate: time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	expectedComp := &entities.TCamp{
		City:    "DEF",
		Address: "GHI",
		BegDate: time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	suite.mockRepo.On("Create", expectedComp).
		Return(ErrServiceTest).Times(1)

	trainingCamp, err := suite.service.Create(req)

	suite.Nil(trainingCamp)
	suite.Error(err)
}

func (suite *TCampServiceTestSuite) TestCreateSuccess() {
	req := &dto.CreateTCampReq{
		City:    "DEF",
		Address: "GHI",
		BegDate: time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	expectedComp := &entities.TCamp{
		City:    "DEF",
		Address: "GHI",
		BegDate: time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	suite.mockRepo.On("Create", expectedComp).
		Return(nil).Times(1)

	trainingCamp, err := suite.service.Create(req)

	suite.Equal(expectedComp, trainingCamp)
	suite.NoError(err)
}

// ListTCamps.
func (suite *TCampServiceTestSuite) TestListTCampsFail() {
	suite.mockRepo.On("ListTCamps").
		Return(nil, ErrServiceTest).Times(1)

	trainingCampes, err := suite.service.ListTCamps()

	suite.Nil(trainingCampes)
	suite.Error(err)
}

func (suite *TCampServiceTestSuite) TestListTCampsSuccess() {
	id := uuid.New()

	expectedComp := &entities.TCamp{
		ID:      id,
		City:    "DEF",
		Address: "GHI",
		BegDate: time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
	}
	expectedComps := []*entities.TCamp{expectedComp}

	suite.mockRepo.On("ListTCamps").
		Return(expectedComps, nil).Times(1)

	trainingCamps, err := suite.service.ListTCamps()

	suite.Equal(expectedComps, trainingCamps)
	suite.NoError(err)
}

// GetTCampByID.
func (suite *TCampServiceTestSuite) TestGetTCampByIDFail() {
	id := uuid.New()

	suite.mockRepo.On("GetTCampByID", id).
		Return(nil, ErrServiceTest).Times(1)

	trainingCamp, err := suite.service.GetTCampByID(id)

	suite.Nil(trainingCamp)
	suite.Error(err)
}

func (suite *TCampServiceTestSuite) TestGetTCampByIDSuccess() {
	id := uuid.New()

	expectedComp := &entities.TCamp{
		ID:      id,
		City:    "DEF",
		Address: "GHI",
		BegDate: time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(1990, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	suite.mockRepo.On("GetTCampByID", id).
		Return(expectedComp, nil).Times(1)

	trainingCamp, err := suite.service.GetTCampByID(id)

	suite.Equal(expectedComp, trainingCamp)
	suite.NoError(err)
}

// RegisterSportsman.
func (suite *TCampServiceTestSuite) TestRegisterGetSportsmanByIDFail() {
	compID := uuid.New()
	sportsmanID := uuid.New()

	req := &dto.RegForTCampReq{
		TCampID:     compID,
		SportsmanID: sportsmanID,
	}

	suite.mockRepoSportsman.On("GetSportsmanByID", req.SportsmanID).
		Return(nil, ErrServiceTest).Times(1)

	compApp, err := suite.service.RegisterSportsman(req)

	suite.Nil(compApp)
	suite.Error(err)
}

func (suite *TCampServiceTestSuite) TestRegisterGetTCampByIDFail() {
	compID := uuid.New()
	sportsmanID := uuid.New()

	req := &dto.RegForTCampReq{
		TCampID:     compID,
		SportsmanID: sportsmanID,
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

	suite.mockRepoSportsman.On("GetSportsmanByID", req.SportsmanID).
		Return(expectedSportsman, nil).Times(1)

	suite.mockRepo.On("GetTCampByID", req.TCampID).
		Return(nil, ErrServiceTest).Times(1)

	compApp, err := suite.service.RegisterSportsman(req)

	suite.Nil(compApp)
	suite.Error(err)
}

func (suite *TCampServiceTestSuite) TestRegisterFail() {
	compID := uuid.New()
	sportsmanID := uuid.New()

	req := &dto.RegForTCampReq{
		TCampID:     compID,
		SportsmanID: sportsmanID,
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

	expectedComp := &entities.TCamp{
		ID:      compID,
		City:    "Moscow",
		Address: "Moscow Street",
		BegDate: time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	expectedCompApp := &entities.TCampApplication{
		TCampID:     req.TCampID,
		SportsmanID: req.SportsmanID,
	}

	suite.mockRepoSportsman.On("GetSportsmanByID", req.SportsmanID).
		Return(expectedSportsman, nil).Times(1)

	suite.mockRepo.On("GetTCampByID", req.TCampID).
		Return(expectedComp, nil).Times(1)

	suite.mockRepo.On("RegisterSportsman", expectedCompApp).
		Return(ErrServiceTest).Times(1)

	compApp, err := suite.service.RegisterSportsman(req)

	suite.Nil(compApp)
	suite.Error(err)
}

func (suite *TCampServiceTestSuite) TestRegisterSuccess() {
	compID := uuid.New()
	sportsmanID := uuid.New()

	req := &dto.RegForTCampReq{
		TCampID:     compID,
		SportsmanID: sportsmanID,
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

	expectedComp := &entities.TCamp{
		ID:      compID,
		City:    "Moscow",
		Address: "Moscow Street",
		BegDate: time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	expectedCompApp := &entities.TCampApplication{
		TCampID:     req.TCampID,
		SportsmanID: req.SportsmanID,
	}

	suite.mockRepoSportsman.On("GetSportsmanByID", req.SportsmanID).
		Return(expectedSportsman, nil).Times(1)

	suite.mockRepo.On("GetTCampByID", req.TCampID).
		Return(expectedComp, nil).Times(1)

	suite.mockRepo.On("RegisterSportsman", expectedCompApp).
		Return(nil).Times(1)

	compApp, err := suite.service.RegisterSportsman(req)

	suite.Equal(expectedCompApp, compApp)
	suite.NoError(err)
}

package service_test

import (
	"testing"

	"src/internal/constants"
	"src/internal/entities"
	"src/internal/service"
	"src/internal/service/dto"
	"src/internal/service/repository/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type ResultServiceTestSuite struct {
	suite.Suite
	mockRepo *mocks.IResultRepository
	service  service.IResultService
}

func (suite *ResultServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewIResultRepository(suite.T())
	suite.service = service.NewResultService(suite.mockRepo)
}

func TestResultServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ResultServiceTestSuite))
}

// Update.
func (suite *ResultServiceTestSuite) TestUpdateGetResultByIDFail() {
	smID := uuid.New()
	compID := uuid.New()

	req := &dto.UpdateResultReq{
		SportsmanID:    smID,
		CompetitionID:  compID,
		WeightCategory: constants.WC59,
		Snatch:         70,
		CleanAndJerk:   95,
		Place:          1,
	}

	suite.mockRepo.On("GetResultByID",
		req.SportsmanID, req.CompetitionID).
		Return(nil, ErrServiceTest).Times(1)

	result, err := suite.service.Update(req)

	suite.Nil(result)
	suite.Error(err)
}

func (suite *ResultServiceTestSuite) TestUpdateFail() {
	sportsmanID := uuid.New()
	compID := uuid.New()

	req := &dto.UpdateResultReq{
		SportsmanID:    sportsmanID,
		CompetitionID:  compID,
		WeightCategory: constants.WC59,
		Snatch:         70,
		CleanAndJerk:   95,
		Place:          1,
	}

	expectedRes := &entities.Result{
		CompetitionID:  compID,
		SportsmanID:    sportsmanID,
		WeightCategory: req.WeightCategory,
		Snatch:         65,
		CleanAndJerk:   90,
		Place:          1,
	}

	suite.mockRepo.On("GetResultByID",
		req.SportsmanID, req.CompetitionID).
		Return(expectedRes, nil).Times(1)

	suite.mockRepo.On("Update", expectedRes).
		Return(ErrServiceTest).Times(1)

	result, err := suite.service.Update(req)

	suite.Nil(result)
	suite.Error(err)
}

func (suite *ResultServiceTestSuite) TestUpdateSuccess() {
	sportsmanID := uuid.New()
	compID := uuid.New()

	req := &dto.UpdateResultReq{
		CompetitionID:  compID,
		SportsmanID:    sportsmanID,
		WeightCategory: constants.WC59,
		Snatch:         70,
		CleanAndJerk:   95,
		Place:          1,
	}

	expectedRes := &entities.Result{
		CompetitionID:  compID,
		SportsmanID:    sportsmanID,
		WeightCategory: req.WeightCategory,
		Snatch:         65,
		CleanAndJerk:   90,
		Place:          1,
	}

	expectedResUpd := &entities.Result{
		CompetitionID:  compID,
		SportsmanID:    sportsmanID,
		WeightCategory: req.WeightCategory,
		Snatch:         70,
		CleanAndJerk:   95,
		Place:          1,
	}

	suite.mockRepo.On("GetResultByID",
		req.SportsmanID, req.CompetitionID).
		Return(expectedRes, nil).Times(1)

	suite.mockRepo.On("Update", expectedResUpd).
		Return(nil).Times(1)

	result, err := suite.service.Update(req)

	suite.Equal(expectedResUpd, result)
	suite.NoError(err)
}

// Create.
func (suite *ResultServiceTestSuite) TestCreateFail() {
	sportsmanID := uuid.New()
	compID := uuid.New()

	req := &dto.CreateResultReq{
		WeightCategory: constants.WC59,
		CompetitionID:  compID,
		SportsmanID:    sportsmanID,
		Snatch:         70,
		CleanAndJerk:   95,
		Place:          1,
	}

	expectedResult := &entities.Result{
		WeightCategory: constants.WC59,
		CompetitionID:  req.CompetitionID,
		SportsmanID:    req.SportsmanID,
		Snatch:         70,
		CleanAndJerk:   95,
		Place:          1,
	}

	suite.mockRepo.On("Create", expectedResult).
		Return(ErrServiceTest).Times(1)

	result, err := suite.service.Create(req)

	suite.Nil(result)
	suite.Error(err)
}

func (suite *ResultServiceTestSuite) TestCreateSuccess() {
	sportsmanID := uuid.New()
	compID := uuid.New()

	req := &dto.CreateResultReq{
		WeightCategory: constants.WC59,
		CompetitionID:  compID,
		SportsmanID:    sportsmanID,
		Snatch:         70,
		CleanAndJerk:   96,
		Place:          1,
	}

	expectedResult := &entities.Result{
		WeightCategory: constants.WC59,
		CompetitionID:  compID,
		SportsmanID:    sportsmanID,
		Snatch:         70,
		CleanAndJerk:   96,
		Place:          1,
	}

	suite.mockRepo.On("Create", expectedResult).
		Return(nil).Times(1)

	result, err := suite.service.Create(req)

	suite.Equal(expectedResult, result)
	suite.NoError(err)
}

// ListResults.
func (suite *ResultServiceTestSuite) TestListResultsFail() {
	suite.mockRepo.On("ListResults").
		Return(nil, ErrServiceTest).Times(1)

	results, err := suite.service.ListResults()

	suite.Nil(results)
	suite.Error(err)
}

func (suite *ResultServiceTestSuite) TestListResultsSuccess() {
	sportsmanID := uuid.New()
	compID := uuid.New()

	expectedResult := &entities.Result{
		CompetitionID:  compID,
		SportsmanID:    sportsmanID,
		WeightCategory: constants.WC59,
		Snatch:         70,
		CleanAndJerk:   96,
		Place:          1,
	}
	expectedResults := []*entities.Result{expectedResult}

	suite.mockRepo.On("ListResults").
		Return(expectedResults, nil).Times(1)

	resultes, err := suite.service.ListResults()

	suite.Equal(expectedResults, resultes)
	suite.NoError(err)
}

// GetResultByID.
func (suite *ResultServiceTestSuite) TestGetResultByIDFail() {
	smID := uuid.New()
	compID := uuid.New()

	suite.mockRepo.On("GetResultByID", smID, compID).
		Return(nil, ErrServiceTest).Times(1)

	result, err := suite.service.GetResultByID(smID, compID)

	suite.Nil(result)
	suite.Error(err)
}

func (suite *ResultServiceTestSuite) TestGetResultByIDSuccess() {
	sportsmanID := uuid.New()
	compID := uuid.New()

	expectedResult := &entities.Result{
		CompetitionID:  compID,
		SportsmanID:    sportsmanID,
		WeightCategory: constants.WC59,
		Snatch:         70,
		CleanAndJerk:   96,
		Place:          1,
	}

	suite.mockRepo.On("GetResultByID",
		sportsmanID, compID).
		Return(expectedResult, nil).Times(1)

	result, err := suite.service.GetResultByID(sportsmanID, compID)

	suite.Equal(expectedResult, result)
	suite.NoError(err)
}

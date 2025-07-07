package service_test

import (
	"errors"
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

var ErrServiceTest = errors.New("error")

type CoachServiceTestSuite struct {
	suite.Suite
	mockRepo  *mocks.ICoachRepository
	coachServ service.ICoachService
}

func (suite *CoachServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewICoachRepository(suite.T())
	suite.coachServ = service.NewCoachService(suite.mockRepo)
}

func TestCoachServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CoachServiceTestSuite))
}

// Update.
func (suite *CoachServiceTestSuite) TestUpdateGetCoachByIDFail() {
	boolval := false
	boolp := &boolval
	req := &dto.UpdateCoachReq{
		ID:         uuid.New(),
		Surname:    "ABC",
		Name:       "DEF",
		Patronymic: "GHI",
		Birthday:   time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		Experience: 5,
		Gender:     (*constants.GenderT)(boolp),
	}

	suite.mockRepo.On("GetCoachByID", req.ID).
		Return(nil, ErrServiceTest).Times(1)

	coach, err := suite.coachServ.Update(req)

	suite.Nil(coach)
	suite.Error(err)
}

func (suite *CoachServiceTestSuite) TestUpdateFail() {
	id := uuid.New()
	boolval := false
	boolp := &boolval
	req := &dto.UpdateCoachReq{
		ID:         id,
		Surname:    "ABC",
		Name:       "DEF",
		Patronymic: "GHI",
		Birthday:   time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		Experience: 5,
		Gender:     (*constants.GenderT)(boolp),
	}

	expectedCoach := &entities.Coach{
		ID:         id,
		Surname:    "ABC",
		Name:       "DEF",
		Patronymic: "GHI",
		Birthday:   time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		Experience: 5,
		Gender:     false,
	}

	suite.mockRepo.On("GetCoachByID", req.ID).
		Return(expectedCoach, nil).Times(1)

	suite.mockRepo.On("Update", expectedCoach).
		Return(ErrServiceTest).Times(1)

	coach, err := suite.coachServ.Update(req)

	suite.Nil(coach)
	suite.Error(err)
}

func (suite *CoachServiceTestSuite) TestUpdateSuccess() {
	id := uuid.New()
	boolval := false
	boolp := &boolval
	req := &dto.UpdateCoachReq{
		ID:         id,
		Surname:    "ABC",
		Name:       "DEF",
		Patronymic: "GHI",
		Birthday:   time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		Experience: 5,
		Gender:     (*constants.GenderT)(boolp),
	}

	expectedCoach := &entities.Coach{
		ID:         id,
		Surname:    "ABC",
		Name:       "DEF",
		Patronymic: "GHI",
		Birthday:   time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		Experience: 5,
		Gender:     false,
	}

	suite.mockRepo.On("GetCoachByID", req.ID).
		Return(expectedCoach, nil).Times(1)

	suite.mockRepo.On("Update", expectedCoach).
		Return(nil).Times(1)

	coach, err := suite.coachServ.Update(req)

	suite.Equal(expectedCoach, coach)
	suite.NoError(err)
}

// Create.
func (suite *CoachServiceTestSuite) TestCreateFail() {
	req := &dto.CreateCoachReq{
		Surname:    "ABC",
		Name:       "DEF",
		Patronymic: "GHI",
		Birthday:   time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		Experience: 5,
		Gender:     false,
	}

	expectedCoach := &entities.Coach{
		Surname:    "ABC",
		Name:       "DEF",
		Patronymic: "GHI",
		Birthday:   time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		Experience: 5,
		Gender:     false,
	}

	suite.mockRepo.On("Create", expectedCoach).
		Return(ErrServiceTest).Times(1)

	coach, err := suite.coachServ.Create(req)

	suite.Nil(coach)
	suite.Error(err)
}

func (suite *CoachServiceTestSuite) TestCreateSuccess() {
	req := &dto.CreateCoachReq{
		Surname:    "ABC",
		Name:       "DEF",
		Patronymic: "GHI",
		Birthday:   time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		Experience: 5,
		Gender:     false,
	}

	expectedCoach := &entities.Coach{
		Surname:    "ABC",
		Name:       "DEF",
		Patronymic: "GHI",
		Birthday:   time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		Experience: 5,
		Gender:     false,
	}

	suite.mockRepo.On("Create", expectedCoach).
		Return(nil).Times(1)

	coach, err := suite.coachServ.Create(req)

	suite.Equal(expectedCoach, coach)
	suite.NoError(err)
}

// ListCoaches.
func (suite *CoachServiceTestSuite) TestListCoachesFail() {
	suite.mockRepo.On("ListCoaches").
		Return(nil, ErrServiceTest).Times(1)

	coaches, err := suite.coachServ.ListCoaches()

	suite.Nil(coaches)
	suite.Error(err)
}

func (suite *CoachServiceTestSuite) TestListCoachesSuccess() {
	id := uuid.New()

	expectedCoach := &entities.Coach{
		ID:         id,
		Surname:    "ABC",
		Name:       "DEF",
		Patronymic: "GHI",
		Birthday:   time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		Experience: 5,
		Gender:     false,
	}
	expectedCoaches := []*entities.Coach{expectedCoach}

	suite.mockRepo.On("ListCoaches").
		Return(expectedCoaches, nil).Times(1)

	coaches, err := suite.coachServ.ListCoaches()

	suite.Equal(expectedCoaches, coaches)
	suite.NoError(err)
}

// GetCoachByID.
func (suite *CoachServiceTestSuite) TestGetCoachByIDFail() {
	id := uuid.New()

	suite.mockRepo.On("GetCoachByID", id).
		Return(nil, ErrServiceTest).Times(1)

	coach, err := suite.coachServ.GetCoachByID(id)

	suite.Nil(coach)
	suite.Error(err)
}

func (suite *CoachServiceTestSuite) TestGetCoachByIDSuccess() {
	id := uuid.New()

	expectedCoach := &entities.Coach{
		ID:         id,
		Surname:    "ABC",
		Name:       "DEF",
		Patronymic: "GHI",
		Birthday:   time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		Experience: 5,
		Gender:     false,
	}

	suite.mockRepo.On("GetCoachByID", id).
		Return(expectedCoach, nil).Times(1)

	coach, err := suite.coachServ.GetCoachByID(id)

	suite.Equal(expectedCoach, coach)
	suite.NoError(err)
}

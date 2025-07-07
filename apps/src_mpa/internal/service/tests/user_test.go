package service_test

import (
	"testing"

	"src/internal/entities"
	"src/internal/service"
	"src/internal/service/dto"
	"src/internal/service/repository/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	mockRepo *mocks.IUserRepository
	service  service.IUserService
}

func (suite *UserServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewIUserRepository(suite.T())
	suite.service = service.NewUserService(suite.mockRepo)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

// Login.
func (suite *UserServiceTestSuite) TestLoginGetUserByEmailFail() {
	req := &dto.LoginUserReq{
		Email:    "abc@mail.ru",
		Password: "abc1234560+",
	}

	suite.mockRepo.On("GetUserByEmail", req.Email).
		Return(nil, ErrServiceTest).Times(1)

	user, err := suite.service.Login(req)

	suite.Nil(user)
	suite.Error(err)
}

func (suite *UserServiceTestSuite) TestUserNotFound() {
	req := &dto.LoginUserReq{
		Email:    "abc@mail.ru",
		Password: "abc1234560+",
	}

	suite.mockRepo.On("GetUserByEmail", req.Email).
		Return(nil, service.ErrUserNotFound).Times(1)

	user, err := suite.service.Login(req)

	suite.Nil(user)
	suite.Equal(service.ErrUserNotFound, err)
}

func (suite *UserServiceTestSuite) TestWrongPassword() {
	req := &dto.LoginUserReq{
		Email:    "abc@mail.ru",
		Password: "abc1234560+",
	}

	suite.mockRepo.On("GetUserByEmail", req.Email).
		Return(&entities.User{
			Email:    "abc@mail.ru",
			Password: "123",
		}, nil).Times(1)

	user, err := suite.service.Login(req)

	suite.Nil(user)
	suite.Equal(service.ErrWrongPassword, err)
}

func (suite *UserServiceTestSuite) TestLoginSuccess() {
	req := &dto.LoginUserReq{
		Email:    "abc@mail.ru",
		Password: "abc1234560+",
	}

	expectedUser := &entities.User{
		Email:    "abc@mail.ru",
		Password: "abc1234560+",
	}

	suite.mockRepo.On("GetUserByEmail", req.Email).
		Return(&entities.User{
			Email:    "abc@mail.ru",
			Password: "abc1234560+",
		}, nil).Times(1)

	user, err := suite.service.Login(req)

	suite.Equal(expectedUser, user)
	suite.NoError(err)
}

// Register.
func (suite *UserServiceTestSuite) TestRegisterGetUserByEmailFail() {
	req := &dto.RegisterUserReq{
		Email:    "abc@mail.ru",
		Password: "abc1234560+",
	}

	suite.mockRepo.On("GetUserByEmail", req.Email).
		Return(nil, ErrServiceTest).Times(1)

	user, err := suite.service.Register(req)

	suite.Nil(user)
	suite.Error(err)
}

func (suite *UserServiceTestSuite) TestAlreadyRegistered() {
	req := &dto.RegisterUserReq{
		Email:    "abc@mail.ru",
		Password: "abc1234560+",
	}

	suite.mockRepo.On("GetUserByEmail", req.Email).
		Return(&entities.User{
			Email:    "abc@mail.ru",
			Password: "abc1234560+",
		}, nil).Times(1)

	user, err := suite.service.Register(req)

	suite.Nil(user)
	suite.Equal(service.ErrAlreadyRegistered, err)
}

func (suite *UserServiceTestSuite) TestRegisterCreateFail() {
	req := &dto.RegisterUserReq{
		Email:    "abc@mail.ru",
		Password: "abc1234560+",
	}

	expectedUser := &entities.User{
		Email:    "abc@mail.ru",
		Password: "abc1234560+",
	}

	suite.mockRepo.On("GetUserByEmail", req.Email).
		Return(nil, nil).Times(1)

	suite.mockRepo.On("Create", expectedUser).
		Return(ErrServiceTest)

	user, err := suite.service.Register(req)

	suite.Nil(user)
	suite.Error(err)
}

func (suite *UserServiceTestSuite) TestRegisterSuccess() {
	req := &dto.RegisterUserReq{
		Email:    "abc@mail.ru",
		Password: "abc1234560+",
	}

	expectedUser := &entities.User{
		Email:    "abc@mail.ru",
		Password: "abc1234560+",
	}

	suite.mockRepo.On("GetUserByEmail", req.Email).
		Return(nil, nil).Times(1)

	suite.mockRepo.On("Create", expectedUser).
		Return(nil)

	user, err := suite.service.Register(req)

	suite.Equal(expectedUser, user)
	suite.NoError(err)
}

// GetUserByID.
func (suite *UserServiceTestSuite) TestGetUserByIDFail() {
	id := uuid.New()

	suite.mockRepo.On("GetUserByID", id).
		Return(nil, ErrServiceTest).Times(1)

	user, err := suite.service.GetUserByID(id)

	suite.Nil(user)
	suite.Error(err)
}

func (suite *UserServiceTestSuite) TestGetUserByIDSuccess() {
	id := uuid.New()

	expectedUser := &entities.User{
		Email:    "abc@mail.ru",
		Password: "abc1234560+",
	}

	suite.mockRepo.On("GetUserByID", id).
		Return(&entities.User{
			Email:    "abc@mail.ru",
			Password: "abc1234560+",
		}, nil).Times(1)

	user, err := suite.service.GetUserByID(id)

	suite.Equal(expectedUser, user)
	suite.NoError(err)
}

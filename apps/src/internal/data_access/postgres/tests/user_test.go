package service_test

import (
	"log"
	"testing"

	"src/internal/constants"
	"src/internal/data_access/connect"
	dataaccess "src/internal/data_access/postgres"
	"src/internal/domain"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserDATestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo *dataaccess.UserRepository
}

func (suite *UserDATestSuite) SetupSuite() {
	config, err := connect.LoadConfig("../../connect/config.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	config.Database.DBName = "test_WL" // Test DB!
	dbConnStr := config.Database.GetPostgresConnectionStr("postgres")
	db, err := gorm.Open(postgres.Open(dbConnStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	/*err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatal(err)
		return
	}*/

	suite.repo = dataaccess.NewUserRepository(db)
	suite.db = db
}

func TestUserDASuite(t *testing.T) {
	suite.Run(t, new(UserDATestSuite))
}

func (suite *UserDATestSuite) TearDownSuite() {
	/*err := suite.db.Migrator().DropTable(&domain.User{})
	if err != nil {
		log.Fatal(err)
		return
	}*/

	db, err := suite.db.DB()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	err = db.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (suite *UserDATestSuite) TearDownTest() {
	err := suite.db.Where("email <> ?", "").Delete(&domain.User{}).Error
	if err != nil {
		log.Fatal(err)
	}
}

// Create.
func (suite *UserDATestSuite) TestCreate() {
	user := &domain.User{
		Email:    "ABC@mail.ru",
		Password: "DEF",
		Role:     constants.UserRoleSportsman,
	}
	err := suite.repo.Create(user)
	suite.NoError(err)

	var fetchedUser domain.User
	result := suite.db.First(&fetchedUser, "id = ?", user.ID)
	suite.NoError(result.Error)

	suite.Equal(user.Email, fetchedUser.Email)
	suite.Equal(user.Password, fetchedUser.Password)
	suite.Equal(user.Role, fetchedUser.Role)
}

// Update.
func (suite *UserDATestSuite) TestUpdate() {
	user := &domain.User{
		Email:    "ABC@mail.ru",
		Password: "DEF",
		Role:     constants.UserRoleSportsman,
	}
	err := suite.repo.Create(user)
	suite.NoError(err)

	user.Email = "JKL@mail.ru"
	user.Password = "MNO"
	err = suite.repo.Update(user)
	suite.NoError(err)

	var updatedUser domain.User
	result := suite.db.First(&updatedUser, "id = ?", user.ID)
	suite.NoError(result.Error)

	suite.Equal("JKL@mail.ru", updatedUser.Email)
	suite.Equal("MNO", updatedUser.Password)
}

// GetUserByID.
func (suite *UserDATestSuite) TestGetUserByID() {
	user := &domain.User{
		Email:    "ABC@mail.ru",
		Password: "DEF",
		Role:     constants.UserRoleSportsman,
	}
	err := suite.repo.Create(user)
	suite.NoError(err)

	fetchedUser, err := suite.repo.GetUserByID(user.ID)
	suite.NoError(err)

	suite.Equal(user.Email, fetchedUser.Email)
	suite.Equal(user.Password, fetchedUser.Password)
	suite.Equal(user.Role, fetchedUser.Role)
}

// GetUserByEmail.
func (suite *UserDATestSuite) TestGetUserByEmail() {
	user := &domain.User{
		Email:    "ABC@mail.ru",
		Password: "DEF",
		Role:     constants.UserRoleSportsman,
	}
	err := suite.repo.Create(user)
	suite.NoError(err)

	fetchedUser, err := suite.repo.GetUserByEmail(user.Email)
	suite.NoError(err)

	suite.Equal(user.ID, fetchedUser.ID)
	suite.Equal(user.Password, fetchedUser.Password)
	suite.Equal(user.Role, fetchedUser.Role)
}

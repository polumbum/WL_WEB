package service_test

import (
	"log"
	"testing"
	"time"

	"src/internal/constants"
	"src/internal/data_access/connect"
	dataaccess "src/internal/data_access/postgres"
	"src/internal/entities"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CoachDATestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo *dataaccess.CoachRepository
}

func (suite *CoachDATestSuite) SetupSuite() {
	config, err := connect.LoadConfig("../../connect/config.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	config.Database.DBName = "test_WL" // Test DB!
	dbConnStr := config.Database.GetPostgresConnectionStr()
	db, err := gorm.Open(postgres.Open(dbConnStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	/*err = db.AutoMigrate(&entities.Coach{})
	if err != nil {
		log.Fatal(err)
		return
	}*/

	suite.repo = dataaccess.NewCoachRepository(db)
	suite.db = db
}

func TestCoachDASuite(t *testing.T) {
	suite.Run(t, new(CoachDATestSuite))
}

func (suite *CoachDATestSuite) TearDownSuite() {
	/*err := suite.db.Migrator().DropTable(&entities.Coach{})
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

func (suite *CoachDATestSuite) TearDownTest() {
	err := suite.db.Exec("DELETE FROM sportsmen_coaches").Error
	if err != nil {
		log.Fatal(err)
	}

	err = suite.db.Where("name <> ?", "").Delete(&entities.Coach{}).Error
	if err != nil {
		log.Fatal(err)
	}
}

// Create.
func (suite *CoachDATestSuite) TestCreate() {
	coach := &entities.Coach{
		Surname:    "ABC",
		Name:       "DEF",
		Patronymic: "GHI",
		Birthday:   time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		Experience: 5,
		Gender:     constants.Female,
	}
	err := suite.repo.Create(coach)
	suite.NoError(err)

	var fetchedCoach entities.Coach
	result := suite.db.First(&fetchedCoach, "id = ?", coach.ID)
	suite.NoError(result.Error)

	suite.Equal(coach.Surname, fetchedCoach.Surname)
	suite.Equal(coach.Name, fetchedCoach.Name)
	suite.Equal(coach.Patronymic, fetchedCoach.Patronymic)
	suite.Equal(coach.Birthday, fetchedCoach.Birthday)
	suite.Equal(coach.Experience, fetchedCoach.Experience)
	suite.Equal(coach.Gender, fetchedCoach.Gender)
}

// Update.
func (suite *CoachDATestSuite) TestUpdate() {
	coach := &entities.Coach{
		ID:         uuid.New(),
		Surname:    "ABC",
		Name:       "DEF",
		Patronymic: "GHI",
		Birthday:   time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		Experience: 5,
		Gender:     constants.Female,
	}
	err := suite.repo.Create(coach)
	suite.NoError(err)

	coach.Surname = "JKL"
	coach.Name = "MNO"
	coach.Patronymic = ""
	coach.Birthday = time.Date(1999, time.November, 10, 0, 0, 0, 0, time.UTC)
	coach.Experience = 6
	coach.Gender = constants.Male
	err = suite.repo.Update(coach)
	suite.NoError(err)

	var updatedCoach entities.Coach
	result := suite.db.First(&updatedCoach, "id = ?", coach.ID)
	suite.NoError(result.Error)

	suite.Equal("JKL", updatedCoach.Surname)
	suite.Equal("MNO", updatedCoach.Name)
	suite.Equal("", updatedCoach.Patronymic)
	suite.Equal(coach.Birthday, updatedCoach.Birthday)
	suite.Equal(6, updatedCoach.Experience)
	suite.Equal(constants.Male, updatedCoach.Gender)
}

// ListCoaches.
func (suite *CoachDATestSuite) TestListCoaches() {
	coach1 := &entities.Coach{Name: "Coach1"}
	coach2 := &entities.Coach{Name: "Coach2"}
	err := suite.repo.Create(coach1)
	suite.NoError(err)
	err = suite.repo.Create(coach2)
	suite.NoError(err)

	coaches, err := suite.repo.ListCoaches()
	suite.NoError(err)
	suite.NotNil(coaches)
	suite.Len(coaches, 2)

	suite.Equal("Coach1", coaches[0].Name)
	suite.Equal("Coach2", coaches[1].Name)
}

// GetCoachByID.
func (suite *CoachDATestSuite) TestGetCoachByID() {
	coach := &entities.Coach{
		Surname:    "ABC",
		Name:       "DEF",
		Patronymic: "GHI",
		Birthday:   time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		Experience: 5,
		Gender:     constants.Female,
	}
	err := suite.repo.Create(coach)
	suite.NoError(err)

	fetchedCoach, err := suite.repo.GetCoachByID(coach.ID)
	suite.NoError(err)

	suite.Equal(coach.Surname, fetchedCoach.Surname)
	suite.Equal(coach.Name, fetchedCoach.Name)
	suite.Equal(coach.Patronymic, fetchedCoach.Patronymic)
	suite.Equal(coach.Birthday, fetchedCoach.Birthday)
	suite.Equal(coach.Experience, fetchedCoach.Experience)
	suite.Equal(coach.Gender, fetchedCoach.Gender)
}

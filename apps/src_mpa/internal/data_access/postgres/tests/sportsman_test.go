package service_test

import (
	"log"
	"testing"
	"time"

	"src/internal/constants"
	"src/internal/data_access/connect"
	dataaccess "src/internal/data_access/postgres"
	"src/internal/entities"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SportsmanDATestSuite struct {
	suite.Suite
	db       *gorm.DB
	repo     *dataaccess.SportsmanRepository
	repoRes  *dataaccess.ResultRepository
	repoComp *dataaccess.CompetitionRepository
}

func (suite *SportsmanDATestSuite) SetupSuite() {
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
	/*err = db.AutoMigrate(&entities.Sportsman{})
	if err != nil {
		log.Fatal(err)
		return
	}

	err = db.AutoMigrate(&entities.Result{}, &entities.Competition{})
	if err != nil {
		log.Fatal(err)
		return
	}*/

	suite.repo = dataaccess.NewSportsmanRepository(db)
	suite.repoRes = dataaccess.NewResultRepository(db)
	suite.repoComp = dataaccess.NewCompetitionRepository(db)
	suite.db = db
}

func TestSportsmanDASuite(t *testing.T) {
	suite.Run(t, new(SportsmanDATestSuite))
}

func (suite *SportsmanDATestSuite) TearDownSuite() {
	/*err := suite.db.Migrator().DropTable(&entities.Sportsman{})
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

func (suite *SportsmanDATestSuite) TearDownTest() {
	/*err := suite.db.Exec("TRUNCATE TABLE sportsmen").Error
	if err != nil {
		log.Fatal(err)
	}*/
	err := suite.db.Exec("DELETE FROM results").Error
	if err != nil {
		log.Fatal(err)
	}

	err = suite.db.Where("name <> ?", "").Delete(&entities.Sportsman{}).Error
	if err != nil {
		log.Fatal(err)
	}
	/*err = suite.db.Exec("TRUNCATE TABLE results").Error
	if err != nil {
		log.Fatal(err)
	}*/
}

// Create.
func (suite *SportsmanDATestSuite) TestCreate() {
	sportsman := &entities.Sportsman{
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS,
		Gender:         constants.Female,
	}
	err := suite.repo.Create(sportsman)
	suite.NoError(err)

	var fetchedSportsman entities.Sportsman
	res := suite.db.First(&fetchedSportsman, "id = ?", sportsman.ID)
	suite.NoError(res.Error)

	suite.Equal(sportsman.Surname, fetchedSportsman.Surname)
	suite.Equal(sportsman.Name, fetchedSportsman.Name)
	suite.Equal(sportsman.Patronymic, fetchedSportsman.Patronymic)
	suite.Equal(sportsman.Birthday, fetchedSportsman.Birthday)
	suite.Equal(sportsman.SportsCategory, fetchedSportsman.SportsCategory)
	suite.Equal(sportsman.Gender, fetchedSportsman.Gender)
	suite.Equal(sportsman.MoscowTeam, fetchedSportsman.MoscowTeam)
}

// Update.
func (suite *SportsmanDATestSuite) TestUpdate() {
	sportsman := &entities.Sportsman{
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS,
		Gender:         constants.Female,
	}
	err := suite.repo.Create(sportsman)
	suite.NoError(err)

	sportsman.Surname = "JKL"
	sportsman.Name = "MNO"
	sportsman.Patronymic = ""
	sportsman.Birthday = time.Date(1999, time.November, 10, 0, 0, 0, 0, time.UTC)
	sportsman.MoscowTeam = false
	sportsman.Gender = constants.Male
	sportsman.SportsCategory = constants.SportsCategoryMS
	err = suite.repo.Update(sportsman)
	suite.NoError(err)

	var updatedSportsman entities.Sportsman
	res := suite.db.First(&updatedSportsman, "id = ?", sportsman.ID)
	suite.NoError(res.Error)

	suite.Equal("JKL", updatedSportsman.Surname)
	suite.Equal("MNO", updatedSportsman.Name)
	suite.Equal("", updatedSportsman.Patronymic)
	suite.Equal(sportsman.Birthday, updatedSportsman.Birthday)
	suite.False(updatedSportsman.MoscowTeam)
	suite.Equal(constants.Male, updatedSportsman.Gender)
	suite.Equal(constants.SportsCategoryMS, updatedSportsman.SportsCategory)
}

// ListSportsmen.
func (suite *SportsmanDATestSuite) TestListSportsmen() {
	sportsman1 := &entities.Sportsman{Name: "Sportsman1"}
	sportsman2 := &entities.Sportsman{Name: "Sportsman2"}
	err := suite.repo.Create(sportsman1)
	suite.NoError(err)
	err = suite.repo.Create(sportsman2)
	suite.NoError(err)

	sportsmen, err := suite.repo.ListSportsmen()
	suite.NoError(err)
	suite.NotNil(sportsmen)
	suite.Len(sportsmen, 2)

	suite.Equal("Sportsman1", sportsmen[0].Name)
	suite.Equal("Sportsman2", sportsmen[1].Name)
}

// GetSportsmanByID.
func (suite *SportsmanDATestSuite) TestGetSportsmanByID() {
	sportsman := &entities.Sportsman{
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS,
		Gender:         constants.Female,
	}
	err := suite.repo.Create(sportsman)
	suite.NoError(err)

	fetchedSportsman, err := suite.repo.GetSportsmanByID(sportsman.ID)
	suite.NoError(err)

	suite.Equal(sportsman.Surname, fetchedSportsman.Surname)
	suite.Equal(sportsman.Name, fetchedSportsman.Name)
	suite.Equal(sportsman.Patronymic, fetchedSportsman.Patronymic)
	suite.Equal(sportsman.Birthday, fetchedSportsman.Birthday)
	suite.Equal(sportsman.SportsCategory, fetchedSportsman.SportsCategory)
	suite.Equal(sportsman.Gender, fetchedSportsman.Gender)
	suite.Equal(sportsman.MoscowTeam, fetchedSportsman.MoscowTeam)
}

// ListResults.
func (suite *SportsmanDATestSuite) TestListResults() {
	sportsman := &entities.Sportsman{
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS, // КМС
		Gender:         true,
	}
	err := suite.repo.Create(sportsman)
	suite.NoError(err)
	competition1 := &entities.Competition{
		Name:              "ABC",
		City:              "DEF",
		Address:           "GHI",
		BegDate:           time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryMW,
		MinSportsCategory: constants.SportsCategory1,
		Antidoping:        true,
	}
	err = suite.repoComp.Create(competition1)
	suite.NoError(err)
	result1 := &entities.Result{
		SportsmanID:    sportsman.ID,
		CompetitionID:  competition1.ID,
		WeightCategory: constants.WC59,
		Snatch:         60,
		CleanAndJerk:   70,
		Place:          1,
	}
	err = suite.repoRes.Create(result1)
	suite.NoError(err)
	competition2 := &entities.Competition{
		Name:              "ABC",
		City:              "DEF",
		Address:           "GHI",
		BegDate:           time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryMW,
		MinSportsCategory: constants.SportsCategory1,
		Antidoping:        true,
	}
	err = suite.repoComp.Create(competition2)
	suite.NoError(err)
	result2 := &entities.Result{
		SportsmanID:    sportsman.ID,
		CompetitionID:  competition2.ID,
		WeightCategory: constants.WC59,
		Snatch:         80,
		CleanAndJerk:   90,
		Place:          1,
	}
	err = suite.repoRes.Create(result2)
	suite.NoError(err)
	sportsmen, err := suite.repo.ListResults(sportsman.ID)
	suite.NoError(err)
	suite.NoError(err)
	suite.NotNil(sportsmen)
	suite.Len(sportsmen, 2)
	suite.Equal(60, sportsmen[0].Snatch)
	suite.Equal(80, sportsmen[1].Snatch)
	suite.Equal(70, sportsmen[0].CleanAndJerk)
	suite.Equal(90, sportsmen[1].CleanAndJerk)
	suite.Equal(1, sportsmen[0].Place)
	suite.Equal(1, sportsmen[1].Place)
}

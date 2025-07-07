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

type ResultDATestSuite struct {
	suite.Suite
	db       *gorm.DB
	repo     *dataaccess.ResultRepository
	repoSM   *dataaccess.SportsmanRepository
	repoComp *dataaccess.CompetitionRepository
}

func (suite *ResultDATestSuite) SetupSuite() {
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

	/*err = db.AutoMigrate(&entities.Competition{}, &entities.Sportsman{})
	if err != nil {
		log.Fatal(err)
		return
	}

	err = db.AutoMigrate(&entities.Result{})
	if err != nil {
		log.Fatal(err)
		return
	}*/

	suite.repo = dataaccess.NewResultRepository(db)
	suite.repoSM = dataaccess.NewSportsmanRepository(db)
	suite.repoComp = dataaccess.NewCompetitionRepository(db)
	suite.db = db
}

func TestResultDASuite(t *testing.T) {
	suite.Run(t, new(ResultDATestSuite))
}

func (suite *ResultDATestSuite) TearDownSuite() {
	/*err := suite.db.Migrator().DropTable(&entities.Result{})
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

func (suite *ResultDATestSuite) TearDownTest() {
	err := suite.db.Exec("DELETE FROM competitions").Error
	if err != nil {
		log.Fatal(err)
	}

	err = suite.db.Exec("DELETE FROM results").Error
	if err != nil {
		log.Fatal(err)
	}
}

// Create.
func (suite *ResultDATestSuite) TestCreate() {
	sportsman := &entities.Sportsman{
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS, // КМС
		Gender:         true,
	}
	err := suite.repoSM.Create(sportsman)
	suite.NoError(err)

	competition := &entities.Competition{
		Name:              "ABC",
		City:              "DEF",
		Address:           "GHI",
		BegDate:           time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryMW,
		MinSportsCategory: constants.SportsCategory1,
		Antidoping:        true,
	}
	err = suite.repoComp.Create(competition)
	suite.NoError(err)

	result := &entities.Result{
		SportsmanID:    sportsman.ID,
		CompetitionID:  competition.ID,
		WeightCategory: constants.WC59,
		Snatch:         60,
		CleanAndJerk:   70,
		Place:          1,
	}
	err = suite.repo.Create(result)
	suite.NoError(err)

	var fetchedResult entities.Result
	res := suite.db.First(&fetchedResult,
		"sportsman_id = ? AND competition_id = ?",
		result.SportsmanID,
		result.CompetitionID)
	suite.NoError(res.Error)

	suite.Equal(result.SportsmanID, fetchedResult.SportsmanID)
	suite.Equal(result.CompetitionID, fetchedResult.CompetitionID)
	suite.Equal(result.WeightCategory, fetchedResult.WeightCategory)
	suite.Equal(result.Snatch, fetchedResult.Snatch)
	suite.Equal(result.CleanAndJerk, fetchedResult.CleanAndJerk)
	suite.Equal(result.Place, fetchedResult.Place)
}

// Update.
func (suite *ResultDATestSuite) TestUpdate() {
	sportsman := &entities.Sportsman{
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS, // КМС
		Gender:         true,
	}
	err := suite.repoSM.Create(sportsman)
	suite.NoError(err)

	competition := &entities.Competition{
		Name:              "ABC",
		City:              "DEF",
		Address:           "GHI",
		BegDate:           time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryMW,
		MinSportsCategory: constants.SportsCategory1,
		Antidoping:        true,
	}
	err = suite.repoComp.Create(competition)
	suite.NoError(err)

	result := &entities.Result{
		SportsmanID:    sportsman.ID,
		CompetitionID:  competition.ID,
		WeightCategory: constants.WC59,
		Snatch:         60,
		CleanAndJerk:   70,
		Place:          1,
	}
	err = suite.repo.Create(result)
	suite.NoError(err)

	result.WeightCategory = constants.WC64
	result.Snatch = 70
	result.CleanAndJerk = 80
	result.Place = 2
	err = suite.repo.Update(result)
	suite.NoError(err)

	var updatedResult entities.Result
	res := suite.db.First(&updatedResult,
		"sportsman_id = ? AND competition_id = ?",
		result.SportsmanID,
		result.CompetitionID)
	suite.NoError(res.Error)

	suite.Equal(constants.WC64, updatedResult.WeightCategory)
	suite.Equal(70, updatedResult.Snatch)
	suite.Equal(80, updatedResult.CleanAndJerk)
	suite.Equal(2, updatedResult.Place)
}

// ListResults.
func (suite *ResultDATestSuite) TestListResults() {
	sportsman1 := &entities.Sportsman{
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS, // КМС
		Gender:         true,
	}
	err := suite.repoSM.Create(sportsman1)
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
		SportsmanID:    sportsman1.ID,
		CompetitionID:  competition1.ID,
		WeightCategory: constants.WC59,
		Snatch:         60,
		CleanAndJerk:   70,
		Place:          1,
	}
	err = suite.repo.Create(result1)
	suite.NoError(err)
	sportsman2 := &entities.Sportsman{
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS, // КМС
		Gender:         true,
	}
	err = suite.repoSM.Create(sportsman2)
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
		SportsmanID:    sportsman2.ID,
		CompetitionID:  competition2.ID,
		WeightCategory: constants.WC59,
		Snatch:         60,
		CleanAndJerk:   70,
		Place:          1,
	}
	err = suite.repo.Create(result2)
	suite.NoError(err)
	results, err := suite.repo.ListResults()
	suite.NoError(err)
	suite.NotNil(results)
	suite.Len(results, 2)
	suite.Equal(sportsman1.ID, results[0].SportsmanID)
	suite.Equal(sportsman2.ID, results[1].SportsmanID)
}

// GetResultByID.
func (suite *ResultDATestSuite) TestGetResultByID() {
	sportsman := &entities.Sportsman{
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS, // КМС
		Gender:         true,
	}
	err := suite.repoSM.Create(sportsman)
	suite.NoError(err)

	competition := &entities.Competition{
		Name:              "ABC",
		City:              "DEF",
		Address:           "GHI",
		BegDate:           time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryMW,
		MinSportsCategory: constants.SportsCategory1,
		Antidoping:        true,
	}
	err = suite.repoComp.Create(competition)
	suite.NoError(err)

	result := &entities.Result{
		SportsmanID:    sportsman.ID,
		CompetitionID:  competition.ID,
		WeightCategory: constants.WC59,
		Snatch:         60,
		CleanAndJerk:   70,
		Place:          1,
	}
	err = suite.repo.Create(result)
	suite.NoError(err)

	fetchedResult, err := suite.repo.GetResultByID(result.SportsmanID,
		result.CompetitionID)
	suite.NoError(err)

	suite.Equal(result.SportsmanID, fetchedResult.SportsmanID)
	suite.Equal(result.CompetitionID, fetchedResult.CompetitionID)
	suite.Equal(result.WeightCategory, fetchedResult.WeightCategory)
	suite.Equal(result.Snatch, fetchedResult.Snatch)
	suite.Equal(result.CleanAndJerk, fetchedResult.CleanAndJerk)
	suite.Equal(result.Place, fetchedResult.Place)
}

// ListSportsmanResults.
func (suite *ResultDATestSuite) TestListSportsmanResults() {
	sportsman := &entities.Sportsman{
		Surname:        "ABC",
		Name:           "DEF",
		Patronymic:     "GHI",
		Birthday:       time.Date(1990, time.November, 10, 0, 0, 0, 0, time.UTC),
		MoscowTeam:     true,
		SportsCategory: constants.SportsCategoryCMS, // КМС
		Gender:         true,
	}
	err := suite.repoSM.Create(sportsman)
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
	err = suite.repo.Create(result1)
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
	err = suite.repo.Create(result2)
	suite.NoError(err)
	results, err := suite.repo.ListSportsmanResults(sportsman.ID)
	suite.NoError(err)
	suite.NoError(err)
	suite.NotNil(results)
	suite.Len(results, 2)
	suite.Equal(60, results[0].Snatch)
	suite.Equal(80, results[1].Snatch)
	suite.Equal(70, results[0].CleanAndJerk)
	suite.Equal(90, results[1].CleanAndJerk)
	suite.Equal(1, results[0].Place)
	suite.Equal(1, results[1].Place)
}

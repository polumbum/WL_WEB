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

type CompetitionDATestSuite struct {
	suite.Suite
	db     *gorm.DB
	repo   *dataaccess.CompetitionRepository
	repoSM *dataaccess.SportsmanRepository
}

func (suite *CompetitionDATestSuite) SetupSuite() {
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
	/*err = db.AutoMigrate(&entities.Competition{})
	if err != nil {
		log.Fatal(err)
		return
	}

	err = db.AutoMigrate(&entities.Sportsman{})
	if err != nil {
		log.Fatal(err)
		return
	}*/

	suite.repo = dataaccess.NewCompetitionRepository(db)
	suite.repoSM = dataaccess.NewSportsmanRepository(db)
	suite.db = db
}

func TestCompetitionDASuite(t *testing.T) {
	suite.Run(t, new(CompetitionDATestSuite))
}

func (suite *CompetitionDATestSuite) TearDownSuite() {
	/*err := suite.db.Migrator().DropTable(&entities.Competition{})
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

func (suite *CompetitionDATestSuite) TearDownTest() {
	err := suite.db.Exec("DELETE FROM results").Error
	if err != nil {
		log.Fatal(err)
	}

	err = suite.db.Exec("DELETE FROM comp_applications").Error
	if err != nil {
		log.Fatal(err)
	}

	err = suite.db.Where("name <> ?", "").Delete(&entities.Competition{}).Error
	if err != nil {
		log.Fatal(err)
	}
}

// Create.
func (suite *CompetitionDATestSuite) TestCreate() {
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
	err := suite.repo.Create(competition)
	suite.NoError(err)

	var fetchedCompetition entities.Competition
	result := suite.db.First(&fetchedCompetition, "id = ?", competition.ID)
	suite.NoError(result.Error)

	suite.Equal(competition.Name, fetchedCompetition.Name)
	suite.Equal(competition.City, fetchedCompetition.City)
	suite.Equal(competition.Address, fetchedCompetition.Address)
	suite.Equal(competition.BegDate, fetchedCompetition.BegDate)
	suite.Equal(competition.EndDate, fetchedCompetition.EndDate)
	suite.Equal(competition.Age, fetchedCompetition.Age)
	suite.Equal(competition.MinSportsCategory, fetchedCompetition.MinSportsCategory)
	suite.Equal(competition.Antidoping, fetchedCompetition.Antidoping)
}

// Update.
func (suite *CompetitionDATestSuite) TestUpdate() {
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
	err := suite.repo.Create(competition)
	suite.NoError(err)

	competition.Name = "JKL"
	competition.City = "MNO"
	competition.Address = "PQR"
	competition.BegDate = time.Date(2024, time.November, 11, 0, 0, 0, 0, time.UTC)
	competition.EndDate = time.Date(2024, time.November, 13, 0, 0, 0, 0, time.UTC)
	competition.MinSportsCategory = constants.SportsCategory2
	competition.Age = constants.AgeCategoryBG13_15
	competition.Antidoping = false
	err = suite.repo.Update(competition)
	suite.NoError(err)

	var updatedCompetition entities.Competition
	result := suite.db.First(&updatedCompetition, "id = ?", competition.ID)
	suite.NoError(result.Error)

	suite.Equal("JKL", updatedCompetition.Name)
	suite.Equal("MNO", updatedCompetition.City)
	suite.Equal("PQR", updatedCompetition.Address)
	suite.Equal(competition.BegDate, updatedCompetition.BegDate)
	suite.Equal(competition.EndDate, updatedCompetition.EndDate)
	suite.Equal(constants.SportsCategory2, updatedCompetition.MinSportsCategory)
	suite.Equal(constants.AgeCategoryBG13_15, updatedCompetition.Age)
	suite.False(updatedCompetition.Antidoping)
}

// ListCompetitions.
func (suite *CompetitionDATestSuite) TestListCompetitions() {
	competition1 := &entities.Competition{Name: "Competition1"}
	competition2 := &entities.Competition{Name: "Competition2"}
	err := suite.repo.Create(competition1)
	suite.NoError(err)
	err = suite.repo.Create(competition2)
	suite.NoError(err)

	competitions, err := suite.repo.ListCompetitions()
	suite.NoError(err)
	suite.NotNil(competitions)
	suite.Len(competitions, 2)

	suite.Equal("Competition1", competitions[0].Name)
	suite.Equal("Competition2", competitions[1].Name)
}

// GetCompetitionByID.
func (suite *CompetitionDATestSuite) TestGetCompetitionByID() {
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
	err := suite.repo.Create(competition)
	suite.NoError(err)

	fetchedCompetition, err := suite.repo.GetCompetitionByID(competition.ID)
	suite.NoError(err)

	suite.Equal(competition.Name, fetchedCompetition.Name)
	suite.Equal(competition.City, fetchedCompetition.City)
	suite.Equal(competition.Address, fetchedCompetition.Address)
	suite.Equal(competition.BegDate, fetchedCompetition.BegDate)
	suite.Equal(competition.EndDate, fetchedCompetition.EndDate)
	suite.Equal(competition.Age, fetchedCompetition.Age)
	suite.Equal(competition.MinSportsCategory, fetchedCompetition.MinSportsCategory)
	suite.Equal(competition.Antidoping, fetchedCompetition.Antidoping)
}

// RegisterSportsman.
func (suite *CompetitionDATestSuite) TestRegisterSportsman() {
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
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}
	err = suite.repo.Create(competition)
	suite.NoError(err)

	compApplication := &entities.CompApplication{
		SportsmanID:       sportsman.ID,
		CompetitionID:     competition.ID,
		WeightCategory:    constants.WC59,
		StartSnatch:       60,
		StartCleanAndJerk: 70,
	}

	err = suite.repo.RegisterSportsman(compApplication)
	suite.NoError(err)

	var createdCompApp entities.CompApplication
	result := suite.db.First(&createdCompApp,
		"sportsman_id = ? AND competition_id = ?",
		compApplication.SportsmanID,
		compApplication.CompetitionID)
	suite.NoError(result.Error)

	suite.Equal(sportsman.ID, createdCompApp.SportsmanID)
	suite.Equal(competition.ID, createdCompApp.CompetitionID)
	suite.Equal(constants.WC59, createdCompApp.WeightCategory)
	suite.Equal(60, createdCompApp.StartSnatch)
	suite.Equal(70, createdCompApp.StartCleanAndJerk)
}

// DeleteRegistration.
func (suite *CompetitionDATestSuite) TestDeleteRegistration() {
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
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryMW,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}
	err = suite.repo.Create(competition)
	suite.NoError(err)

	compApplication := &entities.CompApplication{
		SportsmanID:       sportsman.ID,
		CompetitionID:     competition.ID,
		WeightCategory:    constants.WC59,
		StartSnatch:       60,
		StartCleanAndJerk: 70,
	}

	err = suite.repo.RegisterSportsman(compApplication)
	suite.NoError(err)

	var createdCompApp entities.CompApplication
	result := suite.db.First(&createdCompApp,
		"sportsman_id = ? AND competition_id = ?",
		compApplication.SportsmanID,
		compApplication.CompetitionID)
	suite.NoError(result.Error)

	suite.Equal(sportsman.ID, createdCompApp.SportsmanID)
	suite.Equal(competition.ID, createdCompApp.CompetitionID)
	suite.Equal(constants.WC59, createdCompApp.WeightCategory)
	suite.Equal(60, createdCompApp.StartSnatch)
	suite.Equal(70, createdCompApp.StartCleanAndJerk)

	err = suite.repo.DeleteRegistration(createdCompApp.SportsmanID,
		createdCompApp.CompetitionID)
	suite.NoError(err)

	var notFoundCompApp *entities.CompApplication
	result = suite.db.First(notFoundCompApp,
		"sportsman_id = ? AND competition_id = ?",
		compApplication.SportsmanID,
		compApplication.CompetitionID)
	suite.Error(result.Error)
	suite.Nil(notFoundCompApp)
}

// GetUpcoming.
func (suite *CompetitionDATestSuite) TestGetUpcomingEmpty() {
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

	result, err := suite.repo.GetUpcoming(sportsman.ID)
	suite.NoError(err)

	suite.Equal(len(result), 0)
}

func (suite *CompetitionDATestSuite) TestGetUpcomingSuccess() {
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
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2050, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2050, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}
	err = suite.repo.Create(competition)
	suite.NoError(err)

	competition1 := &entities.Competition{
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2000, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2000, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}
	err = suite.repo.Create(competition1)
	suite.NoError(err)

	compApplication := &entities.CompApplication{
		SportsmanID:       sportsman.ID,
		CompetitionID:     competition.ID,
		WeightCategory:    constants.WC59,
		StartSnatch:       60,
		StartCleanAndJerk: 70,
	}

	err = suite.repo.RegisterSportsman(compApplication)
	suite.NoError(err)
	compApplication.CompetitionID = competition1.ID
	err = suite.repo.RegisterSportsman(compApplication)
	suite.NoError(err)

	result, err := suite.repo.GetUpcoming(sportsman.ID)
	suite.NoError(err)

	suite.Equal(len(result), 1)
	suite.Equal(result[0].Name, competition.Name)
	suite.Equal(result[0].City, competition.City)
	suite.Equal(result[0].Address, competition.Address)
	suite.Equal(result[0].BegDate, competition.BegDate)
	suite.Equal(result[0].EndDate, competition.EndDate)
	suite.Equal(result[0].Age, competition.Age)
	suite.Equal(result[0].MinSportsCategory, competition.MinSportsCategory)
	suite.Equal(result[0].Antidoping, competition.Antidoping)
}

func (suite *CompetitionDATestSuite) TestGetUpcomingSuccessFew() {
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
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2030, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2030, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}

	competition1 := &entities.Competition{
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2050, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2050, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}

	competition2 := &entities.Competition{
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2000, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2000, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}
	err = suite.repo.Create(competition1)
	suite.NoError(err)
	err = suite.repo.Create(competition)
	suite.NoError(err)
	err = suite.repo.Create(competition2)
	suite.NoError(err)

	compApplication := &entities.CompApplication{
		SportsmanID:       sportsman.ID,
		CompetitionID:     competition.ID,
		WeightCategory:    constants.WC59,
		StartSnatch:       60,
		StartCleanAndJerk: 70,
	}

	err = suite.repo.RegisterSportsman(compApplication)
	suite.NoError(err)
	compApplication.CompetitionID = competition1.ID
	err = suite.repo.RegisterSportsman(compApplication)
	suite.NoError(err)
	compApplication.CompetitionID = competition2.ID
	err = suite.repo.RegisterSportsman(compApplication)
	suite.NoError(err)

	result, err := suite.repo.GetUpcoming(sportsman.ID)
	suite.NoError(err)

	suite.Equal(len(result), 2)
	suite.Equal(result[0].Name, competition1.Name)
	suite.Equal(result[0].City, competition1.City)
	suite.Equal(result[0].Address, competition1.Address)
	suite.Equal(result[0].BegDate, competition1.BegDate)
	suite.Equal(result[0].EndDate, competition1.EndDate)
	suite.Equal(result[0].Age, competition1.Age)
	suite.Equal(result[0].MinSportsCategory, competition1.MinSportsCategory)
	suite.Equal(result[0].Antidoping, competition1.Antidoping)
	suite.Equal(result[1].Name, competition.Name)
	suite.Equal(result[1].City, competition.City)
	suite.Equal(result[1].Address, competition.Address)
	suite.Equal(result[1].BegDate, competition.BegDate)
	suite.Equal(result[1].EndDate, competition.EndDate)
	suite.Equal(result[1].Age, competition.Age)
	suite.Equal(result[1].MinSportsCategory, competition.MinSportsCategory)
	suite.Equal(result[1].Antidoping, competition.Antidoping)
}

// ListUpcoming.
func (suite *CompetitionDATestSuite) TestListUpcomingEmpty() {
	/*competition1 := &entities.Competition{
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2000, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2000, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}
	err := suite.repo.Create(competition1)
	suite.NoError(err)*/

	result, err := suite.repo.ListUpcoming()
	suite.NoError(err)

	suite.Equal(len(result), 0)
}

func (suite *CompetitionDATestSuite) TestListUpcomingSuccess() {
	competition := &entities.Competition{
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2030, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2030, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}
	err := suite.repo.Create(competition)
	suite.NoError(err)

	competition1 := &entities.Competition{
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2000, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2000, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}
	err = suite.repo.Create(competition1)
	suite.NoError(err)

	result, err := suite.repo.ListUpcoming()
	suite.NoError(err)

	suite.Equal(len(result), 1)
	suite.Equal(result[0].ID, competition.ID)
	suite.Equal(result[0].Name, competition.Name)
	suite.Equal(result[0].City, competition.City)
	suite.Equal(result[0].Address, competition.Address)
	suite.Equal(result[0].BegDate, competition.BegDate)
	suite.Equal(result[0].EndDate, competition.EndDate)
	suite.Equal(result[0].Age, competition.Age)
	suite.Equal(result[0].MinSportsCategory, competition.MinSportsCategory)
	suite.Equal(result[0].Antidoping, competition.Antidoping)
}

func (suite *CompetitionDATestSuite) TestListUpcomingSuccessFew() {
	competition := &entities.Competition{
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2030, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2030, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}
	err := suite.repo.Create(competition)
	suite.NoError(err)

	competition1 := &entities.Competition{
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2000, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2000, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}
	err = suite.repo.Create(competition1)
	suite.NoError(err)

	competition2 := &entities.Competition{
		Name:              "Competition",
		City:              "Moscow",
		Address:           "Moscow Street",
		BegDate:           time.Date(2050, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate:           time.Date(2050, time.November, 12, 0, 0, 0, 0, time.UTC),
		Age:               constants.AgeCategoryY21_23,
		MinSportsCategory: constants.SportsCategoryCMS, // КМС
		Antidoping:        true,
	}
	err = suite.repo.Create(competition2)
	suite.NoError(err)

	result, err := suite.repo.ListUpcoming()
	suite.NoError(err)

	suite.Equal(len(result), 2)
	suite.Equal(result[0].Name, competition.Name)
	suite.Equal(result[0].City, competition.City)
	suite.Equal(result[0].Address, competition.Address)
	suite.Equal(result[0].BegDate, competition.BegDate)
	suite.Equal(result[0].EndDate, competition.EndDate)
	suite.Equal(result[0].Age, competition.Age)
	suite.Equal(result[0].MinSportsCategory, competition.MinSportsCategory)
	suite.Equal(result[0].Antidoping, competition.Antidoping)
	suite.Equal(result[1].Name, competition2.Name)
	suite.Equal(result[1].City, competition2.City)
	suite.Equal(result[1].Address, competition2.Address)
	suite.Equal(result[1].BegDate, competition2.BegDate)
	suite.Equal(result[1].EndDate, competition2.EndDate)
	suite.Equal(result[1].Age, competition2.Age)
	suite.Equal(result[1].MinSportsCategory, competition2.MinSportsCategory)
	suite.Equal(result[1].Antidoping, competition2.Antidoping)
}

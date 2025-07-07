package service_test

import (
	"log"
	"testing"
	"time"

	"src/internal/constants"
	"src/internal/data_access/connect"
	dataaccess "src/internal/data_access/postgres"
	"src/internal/domain"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TCampDATestSuite struct {
	suite.Suite
	db     *gorm.DB
	repo   *dataaccess.TCampRepository
	repoSM *dataaccess.SportsmanRepository
}

func (suite *TCampDATestSuite) SetupSuite() {
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
	/*err = db.AutoMigrate(&domain.TCamp{})
	if err != nil {
		log.Fatal(err)
		return
	}

	err = db.AutoMigrate(&domain.TCampApplication{}, &domain.Sportsman{})
	if err != nil {
		log.Fatal(err)
		return
	}*/

	suite.repo = dataaccess.NewTCampRepository(db)
	suite.repoSM = dataaccess.NewSportsmanRepository(db)
	suite.db = db
}

func TestTCampDASuite(t *testing.T) {
	suite.Run(t, new(TCampDATestSuite))
}

func (suite *TCampDATestSuite) TearDownSuite() {
	/*err := suite.db.Migrator().DropTable(&domain.TCamp{})
	if err != nil {
		log.Fatal(err)
		return
	}
	*/
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

func (suite *TCampDATestSuite) TearDownTest() {
	err := suite.db.Where("city <> ?", "").Delete(&domain.TCamp{}).Error
	if err != nil {
		log.Fatal(err)
	}
}

// Create.
func (suite *TCampDATestSuite) TestCreate() {
	tCamp := &domain.TCamp{
		City:    "DEF",
		Address: "GHI",
		BegDate: time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
	}
	err := suite.repo.Create(tCamp)
	suite.NoError(err)

	var fetchedTCamp domain.TCamp
	result := suite.db.First(&fetchedTCamp, "id = ?", tCamp.ID)
	suite.NoError(result.Error)

	suite.Equal(tCamp.City, fetchedTCamp.City)
	suite.Equal(tCamp.Address, fetchedTCamp.Address)
	suite.Equal(tCamp.BegDate, fetchedTCamp.BegDate)
	suite.Equal(tCamp.EndDate, fetchedTCamp.EndDate)
}

// Update.
func (suite *TCampDATestSuite) TestUpdate() {
	tCamp := &domain.TCamp{
		City:    "DEF",
		Address: "GHI",
		BegDate: time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
	}
	err := suite.repo.Create(tCamp)
	suite.NoError(err)

	tCamp.City = "MNO"
	tCamp.Address = "PQR"
	tCamp.BegDate = time.Date(2024, time.November, 11, 0, 0, 0, 0, time.UTC)
	tCamp.EndDate = time.Date(2024, time.November, 13, 0, 0, 0, 0, time.UTC)
	err = suite.repo.Update(tCamp)
	suite.NoError(err)

	var updatedTCamp domain.TCamp
	result := suite.db.First(&updatedTCamp, "id = ?", tCamp.ID)
	suite.NoError(result.Error)

	suite.Equal("MNO", updatedTCamp.City)
	suite.Equal("PQR", updatedTCamp.Address)
	suite.Equal(tCamp.BegDate, updatedTCamp.BegDate)
	suite.Equal(tCamp.EndDate, updatedTCamp.EndDate)
}

// ListTCamps.
func (suite *TCampDATestSuite) TestListTCamps() {
	tCamp1 := &domain.TCamp{City: "TCamp1"}
	tCamp2 := &domain.TCamp{City: "TCamp2"}
	err := suite.repo.Create(tCamp1)
	suite.NoError(err)
	err = suite.repo.Create(tCamp2)
	suite.NoError(err)

	tCamps, err := suite.repo.ListTCamps(1, 10, "", "")
	suite.NoError(err)
	suite.NotNil(tCamps)
	suite.Len(tCamps, 2)

	suite.Equal("TCamp1", tCamps[0].City)
	suite.Equal("TCamp2", tCamps[1].City)
}

// GetTCampByID.
func (suite *TCampDATestSuite) TestGetTCampByID() {
	tCamp := &domain.TCamp{
		City:    "DEF",
		Address: "GHI",
		BegDate: time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
	}
	err := suite.repo.Create(tCamp)
	suite.NoError(err)

	fetchedTCamp, err := suite.repo.GetTCampByID(tCamp.ID)
	suite.NoError(err)

	suite.Equal(tCamp.City, fetchedTCamp.City)
	suite.Equal(tCamp.Address, fetchedTCamp.Address)
	suite.Equal(tCamp.BegDate, fetchedTCamp.BegDate)
	suite.Equal(tCamp.EndDate, fetchedTCamp.EndDate)
}

// RegisterSportsman.
func (suite *TCampDATestSuite) TestRegisterSportsman() {
	sportsman := &domain.Sportsman{
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

	tCamp := &domain.TCamp{
		City:    "Moscow",
		Address: "Moscow Street",
		BegDate: time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
	}
	err = suite.repo.Create(tCamp)
	suite.NoError(err)

	tCampApplication := &domain.TCampApplication{
		SportsmanID: sportsman.ID,
		TCampID:     tCamp.ID,
	}

	err = suite.repo.RegisterSportsman(tCampApplication)
	suite.NoError(err)

	var createdTCampApp domain.TCampApplication
	result := suite.db.First(&createdTCampApp, "sportsman_id = ? AND t_camp_id = ?",
		tCampApplication.SportsmanID,
		tCampApplication.TCampID)
	suite.NoError(result.Error)

	suite.Equal(sportsman.ID, createdTCampApp.SportsmanID)
	suite.Equal(tCamp.ID, createdTCampApp.TCampID)
}

// DeleteRegistration.
func (suite *TCampDATestSuite) TestDeleteRegistration() {
	sportsman := &domain.Sportsman{
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

	tCamp := &domain.TCamp{
		City:    "Moscow",
		Address: "Moscow Street",
		BegDate: time.Date(2024, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2024, time.November, 12, 0, 0, 0, 0, time.UTC),
	}
	err = suite.repo.Create(tCamp)
	suite.NoError(err)

	tCampApplication := &domain.TCampApplication{
		SportsmanID: sportsman.ID,
		TCampID:     tCamp.ID,
	}

	err = suite.repo.RegisterSportsman(tCampApplication)
	suite.NoError(err)

	var createdTCampApp domain.TCampApplication
	result := suite.db.First(&createdTCampApp,
		"sportsman_id = ? AND t_camp_id = ?",
		tCampApplication.SportsmanID,
		tCampApplication.TCampID)
	suite.NoError(result.Error)

	suite.Equal(sportsman.ID, createdTCampApp.SportsmanID)
	suite.Equal(tCamp.ID, createdTCampApp.TCampID)

	err = suite.repo.DeleteRegistration(createdTCampApp.SportsmanID,
		createdTCampApp.TCampID)
	suite.NoError(err)

	var notFoundTCampApp *domain.TCampApplication
	result = suite.db.First(notFoundTCampApp,
		"sportsman_id = ? AND t_camp_id = ?",
		tCampApplication.SportsmanID,
		tCampApplication.TCampID)
	suite.Error(result.Error)
	suite.Nil(notFoundTCampApp)
}

// GetUpcoming.
func (suite *TCampDATestSuite) TestGetUpcomingEmpty() {
	sportsman := &domain.Sportsman{
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

func (suite *TCampDATestSuite) TestGetUpcomingSuccess() {
	sportsman := &domain.Sportsman{
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

	tCamp := &domain.TCamp{
		City:    "Moscow",
		Address: "Moscow Street",
		BegDate: time.Date(2030, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2030, time.November, 12, 0, 0, 0, 0, time.UTC),
	}
	err = suite.repo.Create(tCamp)
	suite.NoError(err)

	tCamp1 := &domain.TCamp{
		City:    "Moscow",
		Address: "Moscow Street",
		BegDate: time.Date(2000, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2000, time.November, 12, 0, 0, 0, 0, time.UTC),
	}
	err = suite.repo.Create(tCamp1)
	suite.NoError(err)

	tCampApplication := &domain.TCampApplication{
		SportsmanID: sportsman.ID,
		TCampID:     tCamp.ID,
	}

	err = suite.repo.RegisterSportsman(tCampApplication)
	suite.NoError(err)
	tCampApplication.TCampID = tCamp1.ID
	err = suite.repo.RegisterSportsman(tCampApplication)
	suite.NoError(err)

	result, err := suite.repo.GetUpcoming(sportsman.ID)
	suite.NoError(err)

	suite.Equal(result[0].City, tCamp.City)
	suite.Equal(result[0].Address, tCamp.Address)
	suite.Equal(result[0].BegDate, tCamp.BegDate)
	suite.Equal(result[0].EndDate, tCamp.EndDate)
}

func (suite *TCampDATestSuite) TestGetUpcomingSuccessFew() {
	sportsman := &domain.Sportsman{
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

	tCamp := &domain.TCamp{
		City:    "Moscow",
		Address: "Moscow Street",
		BegDate: time.Date(2030, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2030, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	tCamp1 := &domain.TCamp{
		City:    "Moscow",
		Address: "Moscow Street",
		BegDate: time.Date(2050, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2050, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	tCamp2 := &domain.TCamp{
		City:    "Moscow",
		Address: "Moscow Street",
		BegDate: time.Date(2000, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2000, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	err = suite.repo.Create(tCamp)
	suite.NoError(err)
	err = suite.repo.Create(tCamp1)
	suite.NoError(err)
	err = suite.repo.Create(tCamp2)
	suite.NoError(err)

	tCampApplication := &domain.TCampApplication{
		SportsmanID: sportsman.ID,
		TCampID:     tCamp.ID,
	}

	err = suite.repo.RegisterSportsman(tCampApplication)
	suite.NoError(err)
	tCampApplication.TCampID = tCamp1.ID
	err = suite.repo.RegisterSportsman(tCampApplication)
	suite.NoError(err)
	tCampApplication.TCampID = tCamp2.ID
	err = suite.repo.RegisterSportsman(tCampApplication)
	suite.NoError(err)

	result, err := suite.repo.GetUpcoming(sportsman.ID)
	suite.NoError(err)

	suite.Equal(len(result), 2)
	suite.Equal(result[0].City, tCamp.City)
	suite.Equal(result[0].Address, tCamp.Address)
	suite.Equal(result[0].BegDate, tCamp.BegDate)
	suite.Equal(result[0].EndDate, tCamp.EndDate)
	suite.Equal(result[1].City, tCamp1.City)
	suite.Equal(result[1].Address, tCamp1.Address)
	suite.Equal(result[1].BegDate, tCamp1.BegDate)
	suite.Equal(result[1].EndDate, tCamp1.EndDate)
}

// ListUpcoming.
func (suite *TCampDATestSuite) TestListUpcomingEmpty() {
	tCamp1 := &domain.TCamp{
		City:    "Moscow",
		Address: "Moscow Street",
		BegDate: time.Date(2000, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2000, time.November, 12, 0, 0, 0, 0, time.UTC),
	}
	err := suite.repo.Create(tCamp1)
	suite.NoError(err)

	result, err := suite.repo.ListUpcoming()
	suite.NoError(err)

	suite.Equal(len(result), 0)
}

func (suite *TCampDATestSuite) TestListUpcomingSuccess() {
	tCamp := &domain.TCamp{
		City:    "Moscow",
		Address: "Moscow Street",
		BegDate: time.Date(2050, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2050, time.November, 12, 0, 0, 0, 0, time.UTC),
	}
	err := suite.repo.Create(tCamp)
	suite.NoError(err)

	tCamp1 := &domain.TCamp{
		City:    "Moscow",
		Address: "Moscow Street",
		BegDate: time.Date(2000, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2000, time.November, 12, 0, 0, 0, 0, time.UTC),
	}
	err = suite.repo.Create(tCamp1)
	suite.NoError(err)

	result, err := suite.repo.ListUpcoming()
	suite.NoError(err)

	suite.Equal(len(result), 1)
	suite.Equal(result[0].ID, tCamp.ID)
	suite.Equal(result[0].City, tCamp.City)
	suite.Equal(result[0].Address, tCamp.Address)
	suite.Equal(result[0].BegDate, tCamp.BegDate)
	suite.Equal(result[0].EndDate, tCamp.EndDate)
}

func (suite *TCampDATestSuite) TestListUpcomingSuccessFew() {
	tCamp := &domain.TCamp{
		City:    "Moscow",
		Address: "Moscow Street",
		BegDate: time.Date(2030, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2030, time.November, 12, 0, 0, 0, 0, time.UTC),
	}
	err := suite.repo.Create(tCamp)
	suite.NoError(err)

	tCamp1 := &domain.TCamp{
		City:    "Moscow",
		Address: "Moscow Street",
		BegDate: time.Date(2000, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2000, time.November, 12, 0, 0, 0, 0, time.UTC),
	}
	err = suite.repo.Create(tCamp1)
	suite.NoError(err)

	tCamp2 := &domain.TCamp{
		City:    "Moscow",
		Address: "Moscow Street",
		BegDate: time.Date(2050, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndDate: time.Date(2050, time.November, 12, 0, 0, 0, 0, time.UTC),
	}
	err = suite.repo.Create(tCamp2)
	suite.NoError(err)

	result, err := suite.repo.ListUpcoming()
	suite.NoError(err)

	suite.Equal(len(result), 2)
	suite.Equal(result[0].City, tCamp.City)
	suite.Equal(result[0].Address, tCamp.Address)
	suite.Equal(result[0].BegDate, tCamp.BegDate)
	suite.Equal(result[0].EndDate, tCamp.EndDate)
	suite.Equal(result[1].City, tCamp2.City)
	suite.Equal(result[1].Address, tCamp2.Address)
	suite.Equal(result[1].BegDate, tCamp2.BegDate)
	suite.Equal(result[1].EndDate, tCamp2.EndDate)
}

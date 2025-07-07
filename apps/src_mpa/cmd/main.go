package main

import (
	"log"
	"os"
	"src/internal/data_access/connect"
	dataaccess "src/internal/data_access/postgres"
	"src/internal/service"
	server "src/web"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	logPath := os.Getenv("LOG_FILE_PATH")
	configPath := os.Getenv("CONFIG_FILE_PATH")
	file, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logg := log.New(file, "LOG: ", log.Ldate|log.Ltime)
	log.SetOutput(file)
	gormLogger := logger.New(
		logg,
		logger.Config{},
	)
	config, err := connect.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	//config.Database.DBName = "test_WL" // TestDB
	dbConnStr := config.Database.GetPostgresConnectionStr()
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(dbConnStr), &gorm.Config{Logger: gormLogger})
	if err != nil {
		log.Fatal(err)
	}
	logg.Println("DB connected.")
	userRepo := dataaccess.NewUserRepository(db)
	coachRepo := dataaccess.NewCoachRepository(db)
	compRepo := dataaccess.NewCompetitionRepository(db)
	resRepo := dataaccess.NewResultRepository(db)
	sportsmanRepo := dataaccess.NewSportsmanRepository(db)
	tCampRepo := dataaccess.NewTCampRepository(db)
	cAccessRepo := dataaccess.NewCompAccessRepository(db)
	aDopRepo := dataaccess.NewADopingRepository(db)
	server := server.Server{
		DB:           db,
		Logger:       logg,
		UserService:  service.NewUserService(userRepo),
		CoachService: service.NewCoachService(coachRepo),
		CompService: service.NewCompetitionService(
			compRepo,
			sportsmanRepo,
			aDopRepo,
			cAccessRepo,
		),
		ResultService: service.NewResultService(resRepo),
		SportsmanService: service.NewSportsmanService(
			sportsmanRepo,
			resRepo,
		),
		TCampService:   service.NewTCampService(tCampRepo, sportsmanRepo),
		ADopingService: service.NewADopingService(aDopRepo),
		AccessService:  service.NewAccessService(cAccessRepo),
	}
	err = server.StartServer()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	err = sqlDB.Close()
	if err != nil {
		log.Fatal(err)
	}
}

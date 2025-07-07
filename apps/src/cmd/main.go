package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"src/internal/constants"
	"src/internal/data_access/connect"
	dataaccess "src/internal/data_access/postgres"
	"src/internal/http-server/handlers"
	"src/internal/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	//"gorm.io/gorm/logger"

	mw "src/internal/http-server/middleware"

	_ "src/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// @title  ФТАМ API
// @version 1.0
// @description Приложение для спортcменов, тренеров, организаторов и представителей ФТАМ - автоматизация деятельности ФТАМ.
func main() {
	logPath := os.Getenv("LOG_FILE_PATH")
	configPath := os.Getenv("CONFIG_FILE_PATH")
	userDB := os.Getenv("DB_USER")
	file, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logg := log.New(file, "LOG: ", log.Ldate|log.Ltime)
	log.SetOutput(file)
	/*gormLogger := logger.New(
		logg,
		logger.Config{},
	)*/

	/*log := setupLogger(envLocal)
	log = log.With(slog.String("env", envLocal))

	log.Info("initializing server")
	log.Debug("logger debug mode enabled")*/

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	config, err := connect.LoadConfig(configPath)
	if err != nil {
		logg.Fatal(err)
	}

	dbConnStr := config.Database.GetPostgresConnectionStr(userDB)
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(dbConnStr), &gorm.Config{})
	if err != nil {
		logg.Fatal(err)
	}
	logg.Print("DB connected.")
	userRepo := dataaccess.NewUserRepository(db)
	coachRepo := dataaccess.NewCoachRepository(db)
	compRepo := dataaccess.NewCompetitionRepository(db)
	resRepo := dataaccess.NewResultRepository(db)
	sportsmanRepo := dataaccess.NewSportsmanRepository(db)
	tCampRepo := dataaccess.NewTCampRepository(db)
	cAccessRepo := dataaccess.NewCompAccessRepository(db)
	aDopRepo := dataaccess.NewADopingRepository(db)

	serviceSm := service.NewSportsmanService(sportsmanRepo,
		resRepo,
		aDopRepo,
		cAccessRepo)
	serviceC := service.NewCoachService(
		coachRepo,
		resRepo,
		sportsmanRepo,
	)

	serviceComp := service.NewCompetitionService(
		compRepo,
		sportsmanRepo,
		aDopRepo,
		cAccessRepo,
		resRepo,
	)

	serviceTC := service.NewTCampService(
		tCampRepo,
		sportsmanRepo,
	)

	serviceU := service.NewUserService(
		userRepo,
	)

	smHandler := handlers.NewSportsmanHandler(
		serviceSm,
		serviceC,
		serviceComp,
	)

	coachHandler := handlers.NewCoachHandler(
		serviceC,
		serviceComp,
		serviceSm,
	)

	compHandler := handlers.NewCompHandler(
		serviceComp,
		serviceSm,
	)

	tCampHandler := handlers.NewTCampHandler(
		serviceTC,
		serviceSm,
	)

	userHandler := handlers.NewUserHandler(
		serviceU,
		serviceTC,
		serviceComp,
	)

	r.Route("/sportsmen", func(r chi.Router) {
		r.With(mw.JWT("")).
			With(mw.Paginate).
			With(mw.Sort).
			With(mw.FNameFilter).
			Get("/", smHandler.GetAllSportsmen)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(smHandler.SportsmanCtx)
			r.With(mw.JWT(string(constants.UserRoleSportsman))).
				Get("/", smHandler.GetSportsman)
			r.With(mw.JWT(string(constants.UserRoleSportsman))).
				Get("/results", smHandler.GetResults)
			r.With(mw.JWT(string(constants.UserRoleChiefSecretary))).
				Patch("/", smHandler.UpdateSportsman)
			//r.Delete("/", smHandler.DeleteSportsman) // DELETE /articles/123
			r.With(mw.JWT(string(constants.UserRoleChiefSecretary))).
				Post("/coach", smHandler.RegForCoach)
		})
	})

	r.Route("/coaches", func(r chi.Router) {
		r.With(mw.JWT("")).
			With(mw.Paginate).
			With(mw.Sort).
			With(mw.FNameFilter).
			Get("/", coachHandler.GetAllCoaches)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(coachHandler.CoachCtx)
			r.With(mw.JWT(string(constants.UserRoleCoach))).
				Get("/", coachHandler.GetCoach)
			//r.Delete("/", coachHandler.DeleteCoach)
			r.With(mw.JWT(string(constants.UserRoleCoach))).
				Get("/sportsmen/results", coachHandler.GetResults)
			r.With(mw.JWT(string(constants.UserRoleCoach))).
				With(mw.Paginate).
				With(mw.Sort).
				With(mw.FNameFilter).
				Get("/sportsmen", coachHandler.GetSportsmen)
			r.With(mw.JWT(string(constants.UserRoleChiefSecretary))).
				Route("/sportsmen/{sm_id}", func(r chi.Router) {
					r.Use(coachHandler.SportsmanCtx)
					r.Delete("/", coachHandler.RemoveSportsman)
				})
		})
	})

	r.Route("/competitions", func(r chi.Router) {
		r.With(mw.Paginate).
			With(mw.Sort).
			With(mw.CompFilter).
			Get("/", compHandler.GetAllComps)

		r.With(mw.JWT(string(constants.UserRoleCompOrganizer))).
			Post("/", compHandler.CreateComp)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(compHandler.CompCtx)
			r.Get("/results", compHandler.GetResults)
			r.With(mw.JWT(string(constants.UserRoleCompOrganizer))).
				Delete("/", compHandler.DeleteComp)
			r.With(mw.JWT(string(constants.UserRoleSportsman))).
				Post("/sportsman", compHandler.RegSm)
		})
	})

	r.Route("/tcamps", func(r chi.Router) {
		r.With(mw.Paginate).
			With(mw.Sort).
			With(mw.TCampFilter).
			Get("/", tCampHandler.GetAllTCamps)

		r.With(mw.JWT(string(constants.UserRoleTCampOrganizer))).
			Post("/", tCampHandler.CreateTCamp)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(tCampHandler.TCampCtx)
			r.With(mw.JWT(string(constants.UserRoleTCampOrganizer))).
				Delete("/", tCampHandler.DeleteTCamp)
			r.With(mw.JWT(string(constants.UserRoleSportsman))).
				Post("/sportsman", tCampHandler.RegSm)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/signup", userHandler.CreateUser)
		r.Post("/login", userHandler.LoginUser)
		r.Post("/", tCampHandler.CreateTCamp)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(userHandler.UserCtx)
			r.Get("/", userHandler.GetUser)
			r.Delete("/", userHandler.DeleteUser)
			r.Put("/", userHandler.UpdateUser)
			r.With(mw.JWT(string(constants.UserRoleTCampOrganizer))).
				Get("/tcamps", userHandler.GetTCamps)
			r.With(mw.JWT(string(constants.UserRoleCompOrganizer))).
				Get("/competitions", userHandler.GetComps)
		})
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("./swagger/doc.json"),
	))

	http.ListenAndServe(":8000", r)

	/*err = server.StartServer()
	if err != nil {
		log.Fatal(err)
	}*/
	sqlDB, err := db.DB()
	if err != nil {
		logg.Fatal(err)
	}
	err = sqlDB.Close()
	if err != nil {
		logg.Fatal(err)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

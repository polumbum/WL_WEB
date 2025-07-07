package server

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"src/internal/entities"
	"src/internal/service"

	"gorm.io/gorm"

	"github.com/gorilla/sessions"
)

var tmplPath = os.Getenv("TEMPLATES_PATH")
var staticPath = os.Getenv("STATIC_PATH")

var templates = template.Must(template.ParseGlob(tmplPath + "*.html"))

var store = sessions.NewCookieStore([]byte("secret_key"))

type Server struct {
	DB               *gorm.DB
	Logger           *log.Logger
	UserService      service.IUserService
	CoachService     service.ICoachService
	CompService      service.ICompetitionService
	ResultService    service.IResultService
	SportsmanService service.ISportsmanService
	TCampService     service.ITCampService
	AccessService    service.IAccessService
	ADopingService   service.IADopingService
}

func (s *Server) StartServer() error {
	fmt.Println("TEMPLATES_PATH:", tmplPath)
	fmt.Println("STATIC_PATH:", staticPath)

	store.Options.HttpOnly = true
	store.Options.Secure = true
	store.Options.MaxAge = 0
	gob.Register(&entities.User{})

	http.HandleFunc("/", s.handler)
	http.HandleFunc("/error", s.errorPageHandler)

	http.HandleFunc("/login", s.loginHandler)
	http.HandleFunc("/logout", s.logoutHandler)
	http.HandleFunc("/u-not-found", s.unFoundHandler)
	http.HandleFunc("/roles", s.rolesHandler)
	http.HandleFunc("/competitions", s.compsHandler)
	http.HandleFunc("/tcamps", s.tCampsHandler)

	// Registration.
	http.HandleFunc("/sportsman-reg", s.sportsmanRegHandler)
	http.HandleFunc("/coach-reg", s.coachRegHandler)
	http.HandleFunc("/secretary-reg", s.secretaryRegHandler)
	http.HandleFunc("/tcamp-org-reg", s.tCampOrgRegHandler)
	http.HandleFunc("/comp-org-reg", s.compOrgRegHandler)

	// Organizers.
	http.HandleFunc("/tcamp-org", s.tCampOrgHandler)
	http.HandleFunc("/tcamp-create", s.tCampCreateHandler)

	http.HandleFunc("/comp-org", s.compOrgHandler)
	http.HandleFunc("/comp-create", s.compCreateHandler)

	// Sportsman.
	http.HandleFunc("/sportsman", s.sportsmanHandler)
	http.HandleFunc("/profile", s.smProfileHandler)
	http.HandleFunc("/my-results", s.smResultsHandler)
	http.HandleFunc("/sm-comp-reg", s.smCompRegHandler)
	http.HandleFunc("/sm-tcamp-reg", s.smTCampRegHandler)
	http.HandleFunc("/sm-applications", s.smAppHandler)

	// Secretary.
	http.HandleFunc("/secretary", s.secretaryHandler)
	http.HandleFunc("/coach-sportsman", s.coachSmHandler)
	http.HandleFunc("/update-sportsman", s.updateSmHandler)

	// Coach.
	http.HandleFunc("/coach", s.coachHandler)
	http.HandleFunc("/coach-results", s.coachResultsHandler)
	http.HandleFunc("/coach-sm-info", s.coachSmInfoHandler)
	http.HandleFunc("/coach-comp-reg", s.coachCompRegHandler)
	http.HandleFunc("/coach-tcamp-reg", s.coachTCampRegHandler)

	s.Logger.Printf("Server running on http://localhost:8000\n")

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(staticPath))))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return err
	}

	return nil
}

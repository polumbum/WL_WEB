package server

import (
	"net/http"

	"src/internal/constants"
	"src/internal/entities"
	"src/internal/service/dto"

	"time"
)

func (s *Server) tCampOrgHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	if user, ok := val.(*entities.User); !ok || user.Role != constants.UserRoleTCampOrganizer {
		s.Logger.Println("forbidden")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}

	err := templates.ExecuteTemplate(w, "tcamp-org.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) tCampCreateHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	if user, ok := val.(*entities.User); !ok || user.Role != constants.UserRoleTCampOrganizer {
		s.Logger.Println("forbidden")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		req := &dto.CreateTCampReq{}

		req.City = r.FormValue("city")
		req.Address = r.FormValue("address")
		req.BegDate, err = time.Parse("2006-01-02", r.FormValue("begDate"))
		if err != nil {
			s.Logger.Println(err)
			http.Redirect(w, r, "/tcamp-org", http.StatusBadRequest)
			return
		}
		req.EndDate, err = time.Parse("2006-01-02", r.FormValue("endDate"))
		if err != nil {
			s.Logger.Println(err)
			http.Redirect(w, r, "/tcamp-org", http.StatusBadRequest)
			return
		}

		_, err = s.TCampService.Create(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s.Logger.Println("tcamp created")

		http.Redirect(w, r, "/tcamp-org", http.StatusSeeOther)
	}

	err := templates.ExecuteTemplate(w, "tcamp-create.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

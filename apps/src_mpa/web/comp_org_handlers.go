package server

import (
	"net/http"
	"src/internal/constants"
	"src/internal/entities"
	"src/internal/service/dto"
	"strconv"
	"time"
)

func (s *Server) compOrgHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	if user, ok := val.(*entities.User); !ok || user.Role != constants.UserRoleCompOrganizer {
		s.Logger.Println("forbidden")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}

	err := templates.ExecuteTemplate(w, "comp-org.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) compCreateHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	if user, ok := val.(*entities.User); !ok || user.Role != constants.UserRoleCompOrganizer {
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

		req := &dto.CreateCompReq{}

		req.Name = r.FormValue("name")
		req.City = r.FormValue("city")
		req.Address = r.FormValue("address")
		req.BegDate, err = time.Parse("2006-01-02", r.FormValue("begDate"))
		if err != nil {
			s.Logger.Println(err)
			http.Redirect(w, r, "/comp-org", http.StatusBadRequest)
			return
		}
		req.EndDate, err = time.Parse("2006-01-02", r.FormValue("endDate"))
		if err != nil {
			s.Logger.Println(err)
			http.Redirect(w, r, "/comp-org", http.StatusBadRequest)
			return
		}

		req.Age = constants.AgeCategoryT(r.FormValue("age"))
		req.MinSportsCategory = constants.SportsCategoryT(r.FormValue("cat"))
		req.Antidoping, err = strconv.ParseBool(r.FormValue("antidoping"))
		if err != nil {
			s.Logger.Println(err)
			http.Redirect(w, r, "/", http.StatusBadRequest)
			return
		}

		_, err = s.CompService.Create(req)
		if err != nil {
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			s.Logger.Println("error: competition creation")
			return
		}

		s.Logger.Println("competition created")

		http.Redirect(w, r, "/comp-org", http.StatusSeeOther)
	}

	data := struct {
		SportsCat []constants.SportsCategoryT
		AgeCat    []constants.AgeCategoryT
	}{
		SportsCat: constants.GetSportsCat(),
		AgeCat:    constants.GetAgeCat(),
	}

	err := templates.ExecuteTemplate(w, "comp-create.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

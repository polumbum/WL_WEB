package server

import (
	"errors"
	"net/http"
	"net/url"
	"src/internal/constants"
	"src/internal/entities"
	"src/internal/service"
	"src/internal/service/dto"
	"strconv"

	"github.com/google/uuid"
)

func (s *Server) sportsmanHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	if user, ok := val.(*entities.User); !ok || user.Role != constants.UserRoleSportsman {
		s.Logger.Println("forbidden")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}

	err := templates.ExecuteTemplate(w, "sportsman.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) smProfileHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	user, ok := val.(*entities.User)
	if !ok || user.Role != constants.UserRoleSportsman {
		s.Logger.Println("forbidden")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}

	data := struct {
		Sportsman  *entities.Sportsman
		Antidoping *entities.Antidoping
		Access     *entities.CompAccess
		HasAD      bool
		HasCA      bool
	}{}

	sm, err := s.SportsmanService.GetSportsmanByID(user.RoleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data.Sportsman = sm

	ad, err := s.ADopingService.GetADopingByID(user.RoleID)
	if err != nil && !errors.Is(err, service.ErrADopingNotFound) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err == nil {
		data.HasAD = true
		data.Antidoping = ad
	}

	ca, err := s.AccessService.GetAccessByID(user.RoleID)
	if err != nil && !errors.Is(err, service.ErrAccessNotFound) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err == nil {
		data.HasCA = true
		data.Access = ca
	}

	err = templates.ExecuteTemplate(w, "sm-profile.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) smCompRegHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	user, ok := val.(*entities.User)
	if !ok || user.Role != constants.UserRoleSportsman {
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

		req := &dto.RegForCompReq{}
		req.SportsmanID = user.RoleID

		req.CompetitionID, err = uuid.Parse(r.FormValue("compID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		weight, err := strconv.Atoi(r.FormValue("cat"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		req.WeighCategory = constants.WeightCategoryT(weight)

		snatch, err := strconv.Atoi(r.FormValue("snatch"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		req.StartSnatch = snatch

		caj, err := strconv.Atoi(r.FormValue("caj"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		req.StartCleanAndJerk = caj

		_, err = s.CompService.RegisterSportsman(req)
		if err != nil {
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			http.Redirect(w, r, "/error?message="+url.QueryEscape(err.Error()), http.StatusSeeOther)
			return
		}

		s.Logger.Println("sportsman registered for competition")

		http.Redirect(w, r, "/sportsman", http.StatusSeeOther)
	}

	sm, err := s.SportsmanService.GetSportsmanByID(user.RoleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comps, err := s.CompService.ListUpcoming()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Comps        []*entities.Competition
		WeightMale   []constants.WeightCategoryT
		WeightFemale []constants.WeightCategoryT
		Gender       bool
	}{
		Gender:       bool(sm.Gender),
		Comps:        comps,
		WeightMale:   constants.GetWeightMale(),
		WeightFemale: constants.GetWeightFemale(),
	}

	err = templates.ExecuteTemplate(w, "sm-comp-reg.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) smResultsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	user, ok := val.(*entities.User)
	if !ok || user.Role != constants.UserRoleSportsman {
		s.Logger.Println("forbidden")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}

	data := struct {
		Results []*entities.Result
		Comps   []*entities.Competition
	}{}

	results, err := s.SportsmanService.ListResults(user.RoleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data.Results = results
	for _, res := range results {
		comp, err := s.CompService.GetCompetitionByID(res.CompetitionID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data.Comps = append(data.Comps, comp)
	}

	err = templates.ExecuteTemplate(w, "my-results.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) smTCampRegHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	user, ok := val.(*entities.User)
	if !ok || user.Role != constants.UserRoleSportsman {
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

		req := &dto.RegForTCampReq{}
		req.SportsmanID = user.RoleID

		req.TCampID, err = uuid.Parse(r.FormValue("campID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = s.TCampService.RegisterSportsman(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s.Logger.Println("sportsman registered for training camp")

		http.Redirect(w, r, "/sportsman", http.StatusSeeOther)
	}

	tCamps, err := s.TCampService.ListUpcoming()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templates.ExecuteTemplate(w, "sm-tcamp-reg.html", tCamps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) smAppHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	user, ok := val.(*entities.User)
	if !ok || user.Role != constants.UserRoleSportsman {
		s.Logger.Println("forbidden")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}

	data := struct {
		Comps  []*entities.Competition
		TCamps []*entities.TCamp
	}{}

	comps, err := s.CompService.GetUpcoming(user.RoleID)
	if err != nil && !errors.Is(err, service.ErrNotFound) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data.Comps = comps

	tCamps, err := s.TCampService.GetUpcoming(user.RoleID)
	if err != nil && !errors.Is(err, service.ErrNotFound) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data.TCamps = tCamps

	err = templates.ExecuteTemplate(w, "sm-applications.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

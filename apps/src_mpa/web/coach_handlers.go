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

func (s *Server) coachHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	if user, ok := val.(*entities.User); !ok || user.Role != constants.UserRoleCoach {
		s.Logger.Println("forbidden")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}

	err := templates.ExecuteTemplate(w, "coach.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) coachResultsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	user, ok := val.(*entities.User)
	if !ok || user.Role != constants.UserRoleCoach {
		s.Logger.Println("forbidden")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}

	data := struct {
		Sportsmen []struct {
			Name       string
			Surname    string
			Patronymic string
			Results    []*entities.Result
			Comps      []*entities.Competition
		}
	}{}

	sportsmen, err := s.CoachService.ListSportsmen(user.RoleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, sm := range sportsmen {
		record := struct {
			Name       string
			Surname    string
			Patronymic string
			Results    []*entities.Result
			Comps      []*entities.Competition
		}{
			Name:       sm.Name,
			Surname:    sm.Surname,
			Patronymic: sm.Patronymic,
		}
		res, err := s.SportsmanService.ListResults(sm.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		record.Results = res
		for _, r := range res {
			comp, err := s.CompService.GetCompetitionByID(r.CompetitionID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			record.Comps = append(record.Comps, comp)
		}
		data.Sportsmen = append(data.Sportsmen, record)
	}

	err = templates.ExecuteTemplate(w, "coach-results.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) coachSmInfoHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	user, ok := val.(*entities.User)
	if !ok || user.Role != constants.UserRoleCoach {
		s.Logger.Println("forbidden")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}

	data := struct {
		Records []struct {
			Sportsman  *entities.Sportsman
			Antidoping *entities.Antidoping
			Access     *entities.CompAccess
			HasAD      bool
			HasCA      bool
		}
	}{}

	sportsmen, err := s.CoachService.ListSportsmen(user.RoleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, sm := range sportsmen {
		record := struct {
			Sportsman  *entities.Sportsman
			Antidoping *entities.Antidoping
			Access     *entities.CompAccess
			HasAD      bool
			HasCA      bool
		}{
			Sportsman: sm,
		}
		ad, err := s.ADopingService.GetADopingByID(sm.ID)
		if err != nil && !errors.Is(err, service.ErrADopingNotFound) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err == nil {
			record.HasAD = true
			record.Antidoping = ad
		}

		ca, err := s.AccessService.GetAccessByID(sm.ID)
		if err != nil && !errors.Is(err, service.ErrAccessNotFound) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err == nil {
			record.HasCA = true
			record.Access = ca
		}
		data.Records = append(data.Records, record)
	}

	err = templates.ExecuteTemplate(w, "coach-sm-info.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) coachCompRegHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	user, ok := val.(*entities.User)
	if !ok || user.Role != constants.UserRoleCoach {
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
		req.SportsmanID, err = uuid.Parse(r.FormValue("smID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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
			http.Redirect(w, r, "/error?message="+url.QueryEscape(err.Error()), http.StatusSeeOther)
			return
		}

		s.Logger.Println("sportsman registered for competition")

		http.Redirect(w, r, "/coach", http.StatusSeeOther)
	}

	sm, err := s.CoachService.ListSportsmen(user.RoleID)
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
		Sportsmen    []*entities.Sportsman
		Comps        []*entities.Competition
		WeightMale   []constants.WeightCategoryT
		WeightFemale []constants.WeightCategoryT
	}{
		Sportsmen:    sm,
		Comps:        comps,
		WeightMale:   constants.GetWeightMale(),
		WeightFemale: constants.GetWeightFemale(),
	}

	err = templates.ExecuteTemplate(w, "coach-comp-reg.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) coachTCampRegHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	user, ok := val.(*entities.User)
	if !ok || user.Role != constants.UserRoleCoach {
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
		req.SportsmanID, err = uuid.Parse(r.FormValue("smID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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

		s.Logger.Println("sportsman registered for tcamp")

		http.Redirect(w, r, "/coach", http.StatusSeeOther)
	}

	sm, err := s.CoachService.ListSportsmen(user.RoleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tCamps, err := s.TCampService.ListUpcoming()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Sportsmen []*entities.Sportsman
		TCamps    []*entities.TCamp
	}{
		Sportsmen: sm,
		TCamps:    tCamps,
	}

	err = templates.ExecuteTemplate(w, "coach-tcamp-reg.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

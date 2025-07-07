package server

import (
	"net/http"
	"net/url"
	"src/internal/constants"
	"src/internal/entities"
	"src/internal/service/dto"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func (s *Server) secretaryHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	if user, ok := val.(*entities.User); !ok || user.Role != constants.UserRoleChiefSecretary {
		s.Logger.Println("forbidden")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}

	err := templates.ExecuteTemplate(w, "secretary.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) coachSmHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	if user, ok := val.(*entities.User); !ok || user.Role != constants.UserRoleChiefSecretary {
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

		s.Logger.Println("smID", r.FormValue("smID"))
		s.Logger.Println("cID", r.FormValue("cID"))
		smID, err := uuid.Parse(r.FormValue("smID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cID, err := uuid.Parse(r.FormValue("cID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = s.CoachService.AddSportsman(cID, smID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s.Logger.Println("sportsman registered")

		http.Redirect(w, r, "/secretary", http.StatusSeeOther)
	}

	sm, err := s.SportsmanService.ListSportsmen()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	c, err := s.CoachService.ListCoaches()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		Sportsmen []*entities.Sportsman
		Coaches   []*entities.Coach
	}{
		Sportsmen: sm,
		Coaches:   c,
	}

	err = templates.ExecuteTemplate(w, "coach-sportsman.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) updateSmHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	if user, ok := val.(*entities.User); !ok || user.Role != constants.UserRoleChiefSecretary {
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

		req := &dto.UpdateSportsmanReq{}

		s.Logger.Println("smID", r.FormValue("smID"))
		req.ID, err = uuid.Parse(r.FormValue("smID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		req.SportsCategory = constants.SportsCategoryT(r.FormValue("cat"))
		team := r.FormValue("team")
		if team != "" {
			teamVal, err := strconv.ParseBool(team)
			if err != nil {
				s.Logger.Println(err)
				http.Redirect(w, r, "/secretary", http.StatusBadRequest)
				return
			}
			req.MoscowTeam = &teamVal
		}

		reqAD := &dto.UpdateADopingReq{}
		reqAD.SmID = req.ID
		adVal := r.FormValue("antidoping")
		if adVal != "" {
			reqAD.Validity, err = time.Parse("2006-01-02", adVal)
			if err != nil {
				s.Logger.Println(err)
				http.Redirect(w, r, "/secretary", http.StatusBadRequest)
				return
			}

			_, err = s.ADopingService.Update(reqAD)
			if err != nil {
				http.Redirect(w, r, "/error?message="+url.QueryEscape(err.Error()), http.StatusSeeOther)
				return
			}
		}

		reqCA := &dto.UpdateAccessReq{}
		reqCA.SmID = req.ID
		caVal := r.FormValue("accessVal")
		if caVal != "" {
			reqCA.Validity, err = time.Parse("2006-01-02", caVal)
			if err != nil {
				s.Logger.Println(err)
				http.Redirect(w, r, "/secretary", http.StatusBadRequest)
				return
			}
			reqCA.Institution = r.FormValue("accessIns")
			_, err = s.AccessService.Update(reqCA)
			if err != nil {
				http.Redirect(w, r, "/error?message="+url.QueryEscape(err.Error()), http.StatusSeeOther)
				return
			}
		}

		_, err = s.SportsmanService.Update(req)
		if err != nil {
			http.Redirect(w, r, "/error?message="+url.QueryEscape(err.Error()), http.StatusSeeOther)
			return
		}

		s.Logger.Println("sportsman updated")

		http.Redirect(w, r, "/secretary", http.StatusSeeOther)
	}

	sm, err := s.SportsmanService.ListSportsmen()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		SportsCat []constants.SportsCategoryT
		Sportsmen []*entities.Sportsman
	}{
		SportsCat: constants.GetSportsCat(),
		Sportsmen: sm,
	}

	err = templates.ExecuteTemplate(w, "update-sportsman.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

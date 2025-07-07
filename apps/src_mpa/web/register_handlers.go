package server

import (
	"net/http"
	"net/url"
	"src/internal/constants"
	"src/internal/entities"
	"src/internal/service/dto"
	"strconv"
	"time"
)

func (s *Server) sportsmanRegHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	var ok bool
	if _, ok = val.(*entities.User); ok {
		s.Logger.Println("already registered")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}
	if r.Method == http.MethodPost {
		reqUser := &dto.RegisterUserReq{}
		reqUser.Role = constants.UserRoleSportsman
		reqSm := &dto.CreateSportsmanReq{}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		reqSm.Surname = r.FormValue("surname")
		reqSm.Name = r.FormValue("name")
		reqSm.Patronymic = r.FormValue("patronymic")
		reqSm.Birthday, err = time.Parse("2006-01-02", r.FormValue("birthday"))
		if err != nil {
			s.Logger.Println(err)
			http.Redirect(w, r, "/", http.StatusBadRequest)
			return
		}
		reqSm.SportsCategory = constants.SportsCategoryT(r.FormValue("sportsCat"))
		gender, err := strconv.ParseBool(r.FormValue("gender"))
		if err != nil {
			s.Logger.Println(err)
			http.Redirect(w, r, "/", http.StatusBadRequest)
			return
		}
		reqSm.Gender = constants.GenderT(gender)
		reqSm.MoscowTeam, err = strconv.ParseBool(r.FormValue("moscowTeam"))
		if err != nil {
			s.Logger.Println(err)
			http.Redirect(w, r, "/", http.StatusBadRequest)
			return
		}

		sportsman, err := s.SportsmanService.Create(reqSm)
		if err != nil {
			http.Redirect(w, r, "/error?message="+url.QueryEscape(err.Error()), http.StatusSeeOther)
			return
		}
		s.Logger.Println("sportsman", sportsman.ID, "registered")

		reqUser.Email = r.FormValue("email")
		reqUser.Password = r.FormValue("password")
		reqUser.RoleID = sportsman.ID
		user, err := s.UserService.Register(reqUser)
		if err != nil {
			http.Redirect(w, r, "/error?message="+url.QueryEscape(err.Error()), http.StatusSeeOther)
			return
		}
		s.Logger.Println("user", user.ID, "registered")

		session, _ := store.Get(r, "session.id")
		session.Values["user"] = user
		session.Save(r, w)

		http.Redirect(w, r, "/sportsman", http.StatusSeeOther)
	}

	data := struct {
		SportsCat []constants.SportsCategoryT
	}{
		SportsCat: constants.GetSportsCat(),
	}

	err := templates.ExecuteTemplate(w, "sportsman-reg.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) coachRegHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	var ok bool
	if _, ok = val.(*entities.User); ok {
		s.Logger.Println("already registered")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}
	if r.Method == http.MethodPost {
		reqUser := &dto.RegisterUserReq{}
		reqUser.Role = constants.UserRoleCoach
		reqC := &dto.CreateCoachReq{}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		reqC.Surname = r.FormValue("surname")
		reqC.Name = r.FormValue("name")
		reqC.Patronymic = r.FormValue("patronymic")
		reqC.Birthday, err = time.Parse("2006-01-02", r.FormValue("birthday"))
		if err != nil {
			s.Logger.Println(err)
			http.Redirect(w, r, "/", http.StatusBadRequest)
			return
		}
		reqC.Experience, err = strconv.Atoi(r.FormValue("experience"))
		if err != nil {
			s.Logger.Println(err)
			http.Redirect(w, r, "/", http.StatusBadRequest)
			return
		}
		gender, err := strconv.ParseBool(r.FormValue("gender"))
		if err != nil {
			s.Logger.Println(err)
			http.Redirect(w, r, "/", http.StatusBadRequest)
			return
		}
		reqC.Gender = constants.GenderT(gender)

		coach, err := s.CoachService.Create(reqC)
		if err != nil {
			http.Redirect(w, r, "/error?message="+url.QueryEscape(err.Error()), http.StatusSeeOther)
			return
		}
		s.Logger.Println("coach", coach.ID, "registered")

		reqUser.Email = r.FormValue("email")
		reqUser.Password = r.FormValue("password")
		reqUser.RoleID = coach.ID
		user, err := s.UserService.Register(reqUser)
		if err != nil {
			http.Redirect(w, r, "/error?message="+url.QueryEscape(err.Error()), http.StatusSeeOther)
			return
		}
		s.Logger.Println("user", user.ID, "registered")

		session, _ := store.Get(r, "session.id")
		session.Values["user"] = user
		session.Save(r, w)

		http.Redirect(w, r, "/coach", http.StatusSeeOther)
	}

	err := templates.ExecuteTemplate(w, "coach-reg.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) compOrgRegHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	var ok bool
	if _, ok = val.(*entities.User); ok {
		s.Logger.Println("already registered")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}
	if r.Method == http.MethodPost {
		reqUser := &dto.RegisterUserReq{}
		reqUser.Role = constants.UserRoleCompOrganizer

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		reqUser.Email = r.FormValue("email")
		reqUser.Password = r.FormValue("password")
		user, err := s.UserService.Register(reqUser)
		if err != nil {
			http.Redirect(w, r, "/error?message="+url.QueryEscape(err.Error()), http.StatusSeeOther)
			return
		}
		s.Logger.Println("user", user.ID, "registered")

		session, _ := store.Get(r, "session.id")
		session.Values["user"] = user
		session.Save(r, w)

		http.Redirect(w, r, "/comp-org", http.StatusSeeOther)
	}

	err := templates.ExecuteTemplate(w, "comp-org-reg.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) tCampOrgRegHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	var ok bool
	if _, ok = val.(*entities.User); ok {
		s.Logger.Println("already registered")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}
	if r.Method == http.MethodPost {
		reqUser := &dto.RegisterUserReq{}
		reqUser.Role = constants.UserRoleTCampOrganizer

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		reqUser.Email = r.FormValue("email")
		reqUser.Password = r.FormValue("password")
		user, err := s.UserService.Register(reqUser)
		if err != nil {
			http.Redirect(w, r, "/error?message="+url.QueryEscape(err.Error()), http.StatusSeeOther)
			return
		}
		s.Logger.Println("user", user.ID, "registered")

		session, _ := store.Get(r, "session.id")
		session.Values["user"] = user
		session.Save(r, w)

		http.Redirect(w, r, "/tcamp-org", http.StatusSeeOther)
	}

	err := templates.ExecuteTemplate(w, "tcamp-org-reg.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) secretaryRegHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	var ok bool
	if _, ok = val.(*entities.User); ok {
		s.Logger.Println("already registered")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}
	if r.Method == http.MethodPost {
		reqUser := &dto.RegisterUserReq{}
		reqUser.Role = constants.UserRoleChiefSecretary

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		reqUser.Email = r.FormValue("email")
		reqUser.Password = r.FormValue("password")
		user, err := s.UserService.Register(reqUser)
		if err != nil {
			http.Redirect(w, r, "/error?message="+url.QueryEscape(err.Error()), http.StatusSeeOther)
			return
		}
		s.Logger.Println("user", user.ID, "registered")

		session, _ := store.Get(r, "session.id")
		session.Values["user"] = user
		session.Save(r, w)

		http.Redirect(w, r, "/secretary", http.StatusSeeOther)
	}

	err := templates.ExecuteTemplate(w, "secretary-reg.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

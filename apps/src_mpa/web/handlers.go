package server

import (
	"net/http"
	"src/internal/constants"
	"src/internal/entities"
	"src/internal/service/dto"
)

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	if user, ok := val.(*entities.User); ok {
		switch user.Role {
		case constants.UserRoleSportsman:
			http.Redirect(w, r, "/sportsman", http.StatusSeeOther)
			return
		case constants.UserRoleCoach:
			http.Redirect(w, r, "/coach", http.StatusSeeOther)
			return
		case constants.UserRoleCompOrganizer:
			http.Redirect(w, r, "/comp-org", http.StatusSeeOther)
			return
		case constants.UserRoleTCampOrganizer:
			http.Redirect(w, r, "/tcamp-org", http.StatusSeeOther)
			return
		case constants.UserRoleChiefSecretary:
			http.Redirect(w, r, "/secretary", http.StatusSeeOther)
			return
		default:
			http.Error(w, "Invalid role selected", http.StatusBadRequest)
			return
		}
	}

	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		s.Logger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) unFoundHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	if user, ok := val.(*entities.User); ok {
		switch user.Role {
		case constants.UserRoleSportsman:
			http.Redirect(w, r, "/sportsman", http.StatusSeeOther)
			return
		case constants.UserRoleCoach:
			http.Redirect(w, r, "/coach", http.StatusSeeOther)
			return
		case constants.UserRoleCompOrganizer:
			http.Redirect(w, r, "/comp-org", http.StatusSeeOther)
			return
		case constants.UserRoleTCampOrganizer:
			http.Redirect(w, r, "/tcamp-org", http.StatusSeeOther)
			return
		case constants.UserRoleChiefSecretary:
			http.Redirect(w, r, "/secretary", http.StatusSeeOther)
			return
		default:
			http.Error(w, "Invalid role selected", http.StatusBadRequest)
			return
		}
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := s.UserService.Login(&dto.LoginUserReq{
			Email:    email,
			Password: password,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			s.Logger.Println(err)
			http.Redirect(w, r, "/u-not-found", http.StatusSeeOther)
			return
		}

		session, _ := store.Get(r, "session.id")
		session.Values["user"] = user
		session.Save(r, w)

		s.Logger.Println(user.ID, "logged in")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := templates.ExecuteTemplate(w, "u-not-found.html", nil)
	if err != nil {
		s.Logger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	if user, ok := val.(*entities.User); ok {
		switch user.Role {
		case constants.UserRoleSportsman:
			http.Redirect(w, r, "/sportsman", http.StatusSeeOther)
			return
		case constants.UserRoleCoach:
			http.Redirect(w, r, "/coach", http.StatusSeeOther)
			return
		case constants.UserRoleCompOrganizer:
			http.Redirect(w, r, "/comp-org", http.StatusSeeOther)
			return
		case constants.UserRoleTCampOrganizer:
			http.Redirect(w, r, "/tcamp-org", http.StatusSeeOther)
			return
		case constants.UserRoleChiefSecretary:
			http.Redirect(w, r, "/secretary", http.StatusSeeOther)
			return
		default:
			http.Error(w, "Invalid role selected", http.StatusBadRequest)
			return
		}
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := s.UserService.Login(&dto.LoginUserReq{
			Email:    email,
			Password: password,
		})
		if err != nil {
			s.Logger.Println(err)
			http.Redirect(w, r, "/u-not-found", http.StatusSeeOther)
			return
		}

		session, _ := store.Get(r, "session.id")
		session.Values["user"] = user
		session.Save(r, w)

		s.Logger.Println(user.ID, "logged in")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := templates.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	var ok bool
	if _, ok = val.(*entities.User); ok {
		session.Options.MaxAge = -1
		err := session.Save(r, w)
		if err != nil {
			s.Logger.Println("failed to delete session", err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		s.Logger.Println("logout")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	s.Logger.Println("not logged in")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *Server) rolesHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	val := session.Values["user"]
	var ok bool
	if _, ok = val.(*entities.User); ok {
		s.Logger.Println("already registered")
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			s.Logger.Println("Error:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		role := constants.UserRole(r.FormValue("role"))
		s.Logger.Println("Chosen role:", role)

		switch role {
		case constants.UserRoleSportsman:
			http.Redirect(w, r, "/sportsman-reg", http.StatusSeeOther)
			return
		case constants.UserRoleCoach:
			http.Redirect(w, r, "/coach-reg", http.StatusSeeOther)
			return
		case constants.UserRoleCompOrganizer:
			http.Redirect(w, r, "/comp-org-reg", http.StatusSeeOther)
			return
		case constants.UserRoleTCampOrganizer:
			http.Redirect(w, r, "/tcamp-org-reg", http.StatusSeeOther)
			return
		case constants.UserRoleChiefSecretary:
			http.Redirect(w, r, "/secretary-reg", http.StatusSeeOther)
			return
		default:
			http.Error(w, "Invalid role selected", http.StatusBadRequest)
			return
		}
	}

	err := templates.ExecuteTemplate(w, "roles.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) compsHandler(w http.ResponseWriter, r *http.Request) {
	comps, err := s.CompService.ListCompetitions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = templates.ExecuteTemplate(w, "competitions.html", comps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) tCampsHandler(w http.ResponseWriter, r *http.Request) {
	tCamps, err := s.TCampService.ListTCamps()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = templates.ExecuteTemplate(w, "tcamps.html", tCamps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) errorPageHandler(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	templates.ExecuteTemplate(w, "error.html", message)
}

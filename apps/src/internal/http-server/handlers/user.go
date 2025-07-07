package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"src/internal/converters"
	"src/internal/dto"
	http_errors "src/internal/http-server/errors"
	"src/internal/http-server/middleware"
	"src/internal/service"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	userKey KeyStringT = "user"
)

type UserHandler struct {
	serv   service.IUserService
	servTC service.ITCampService
	servC  service.ICompetitionService
}

func NewUserHandler(
	serv service.IUserService,
	servTC service.ITCampService,
	servC service.ICompetitionService,
) *UserHandler {
	return &UserHandler{
		serv:   serv,
		servTC: servTC,
		servC:  servC,
	}
}

func (h *UserHandler) UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		cID := chi.URLParam(r, "id")
		if cID == "" {
			render.Render(w, r, http_errors.ErrInvalidRequest(err))
			return
		}
		id, err := uuid.Parse(cID)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}
		cDomain, err := h.serv.GetUserByID(id)
		if err != nil {
			render.Render(w, r, http_errors.ErrNotFound)
			return
		}
		if cDomain == nil {
			render.Render(w, r, http_errors.ErrNotFound)
			return
		}
		c, err := converters.NewUserConverter().ToDTO(cDomain)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}

		ctx := context.WithValue(r.Context(), userKey, c)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUser godoc
// @Summary Получить пользователя по ID
// @Tags sportsmen
// @Param id path uuid true "ID пользователя"
// @Success 200 {object} UserResp
// @Failure 400 {string} string "invalid ID supplied"
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value(userKey).(*dto.User)
	if err := render.Render(w, r, dto.NewUserResp(u)); err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

// DeleteUser godoc
// @Summary Удалить пользователя
// @Tags users
// @Param id path uuid true "ID пользователя"
// @Success 204 {string} string "ok"
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	c := r.Context().Value(userKey).(*dto.User)

	err := h.serv.Delete(c.ID)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204
}

// CreateUser godoc
// @Summary Создать пользователя
// @Tags users
// @Success 200 {object} UserResp
// @Failure 500 {string} string "internal server error"
// @Router /users/signup [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterUserReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Render(w, r, http_errors.ErrInvalidRequest(err))
		return
	}

	user, err := converters.NewUserConverter().FromRegisterReq(&req)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	user, err = h.serv.Register(user)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	userDTO, err := converters.NewUserConverter().ToDTO(user)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	if err := render.Render(w, r, dto.NewUserResp(userDTO)); err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

// UpdateUser godoc
// @Summary Обновить пароль и почту
// @Tags users
// @Success 200 {object} UserResp
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(userKey).(*dto.User)

	var req dto.UpdateUserReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Render(w, r, http_errors.ErrInvalidRequest(err))
		return
	}

	uDomain, err := converters.NewUserConverter().FromUpdateReq(user, &req)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	uDomain, err = h.serv.Update(uDomain)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	userDTO, err := converters.NewUserConverter().ToDTO(uDomain)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	if err := render.Render(w, r, dto.NewUserResp(userDTO)); err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

// GetTCamps godoc
// @Summary Получить все сборы организатора
// @Tags tcamps
// @Success 200 {object} TCampResp
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /users/{id}/tcamps [get]
func (h *UserHandler) GetTCamps(w http.ResponseWriter, r *http.Request) {
	c := r.Context().Value(userKey).(*dto.User)

	res, err := h.servTC.ListByOrgID(c.ID)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	resList := []*dto.TCamp{}
	for _, item := range res {
		cDTO, err := converters.NewTCampConverter().ToDTO(item)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}
		resList = append(resList, cDTO)
	}

	if err := render.RenderList(w, r, dto.NewTCampListResp(resList)); err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

// GetComps godoc
// @Summary Получить все соревнования организатора
// @Tags competitions
// @Success 200 {object} CompResp
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /users/{id}/competitions [get]
func (h *UserHandler) GetComps(w http.ResponseWriter, r *http.Request) {
	c := r.Context().Value(userKey).(*dto.User)

	res, err := h.servC.ListByOrgID(c.ID)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	resList := []*dto.Competition{}
	for _, item := range res {
		cDTO, err := converters.NewCompConverter().ToDTO(item)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}
		resList = append(resList, cDTO)
	}

	if err := render.RenderList(w, r, dto.NewCompListResp(resList)); err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

// LoginUser godoc
// @Summary Вход в аккаунт
// @Tags users
// @Success 200 {object} LoginResp
// @Failure 500 {string} string "internal server error"
// @Router /users/login [post]
func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginUserReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Render(w, r, http_errors.ErrInvalidRequest(err))
		return
	}

	user, err := converters.NewUserConverter().LoginUserReq(&req)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	user, err = h.serv.Login(user)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	tokenString, err := createToken(user.ID, string(user.Role))
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	userDTO, err := converters.NewUserConverter().ToDTO(user)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	if err := render.Render(w, r, dto.NewLoginResp(userDTO, tokenString)); err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

func createToken(id uuid.UUID, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  id,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString([]byte(middleware.SignJWTStr))
}

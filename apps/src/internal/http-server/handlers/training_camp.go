package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"src/internal/converters"
	"src/internal/dto"
	http_errors "src/internal/http-server/errors"
	"src/internal/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	mw "src/internal/http-server/middleware"
)

const (
	tCampKey KeyStringT = "t_camp"
)

type TCampHandler struct {
	serv   service.ITCampService
	servSm service.ISportsmanService
}

func NewTCampHandler(
	serv service.ITCampService,
	servSm service.ISportsmanService,
) *TCampHandler {
	return &TCampHandler{
		serv:   serv,
		servSm: servSm,
	}
}

// GetAllTCamps godoc
// @Summary Получить все сборы
// @Tags tcamps
// @Param city query string false "Город"
// @Param sort query string false "Сортировка"
// @Param page query string false "Номер страницы"
// @Param batch query string false "Кол-во элементов на странице"
// @Success 200 {object} TCampResp
// @Failure 500 {string} string "internal server error"
// @Router /tcamps [get]
func (h *TCampHandler) GetAllTCamps(w http.ResponseWriter, r *http.Request) {
	pagination := r.Context().Value(mw.PaginationKey).(mw.Pagination)
	sort := r.Context().Value(mw.SortKey).(string)
	filter := r.Context().Value(mw.CityFilterKey).(string)

	log.Printf("Pagination: Page=%d, Batch=%d", pagination.Page, pagination.Batch)
	log.Printf("Sort: %v", sort)
	log.Printf("Filter: %v", filter)

	c, err := h.serv.ListTCamps(pagination.Page,
		pagination.Batch,
		sort,
		filter)

	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	cList := []*dto.TCamp{}

	for _, item := range c {
		cDTO, err := converters.NewTCampConverter().ToDTO(item)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}
		cList = append(cList, cDTO)
	}

	err = render.RenderList(w, r, dto.NewTCampListResp(cList))
	if err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

func (h *TCampHandler) TCampCtx(next http.Handler) http.Handler {
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
		cDomain, err := h.serv.GetTCampByID(id)
		if err != nil {
			render.Render(w, r, http_errors.ErrNotFound)
			return
		}
		if cDomain == nil {
			render.Render(w, r, http_errors.ErrNotFound)
			return
		}
		c, err := converters.NewTCampConverter().ToDTO(cDomain)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}

		ctx := context.WithValue(r.Context(), tCampKey, c)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// DeleteTCamp godoc
// @Summary Удалить сборы
// @Tags tcamps
// @Param id path uuid true "ID сборов"
// @Success 204 {string} string "ok"
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /tcamps/{id} [delete]
func (h *TCampHandler) DeleteTCamp(w http.ResponseWriter, r *http.Request) {
	c := r.Context().Value(tCampKey).(*dto.TCamp)

	err := h.serv.Delete(c.ID)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204
}

// CreateTCamp godoc
// @Summary Создать сборы
// @Tags tcamps
// @Success 200 {object} TCampResp
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /tcamps [post]
func (h *TCampHandler) CreateTCamp(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTCampReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Render(w, r, http_errors.ErrInvalidRequest(err))
		return
	}

	camp, err := converters.NewTCampConverter().FromCreateReq(&req)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	camp, err = h.serv.Create(camp)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	campDTO, err := converters.NewTCampConverter().ToDTO(camp)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	if err := render.Render(w, r, dto.NewTCampResp(campDTO)); err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

// RegSm godoc
// @Summary Запись спортсмена на сборы
// @Tags tcamps
// @Param id path uuid true "ID сборов"
// @Param application body RegForTCampReq true "Заявка"
// @Success 200 {object} RegForTCampResp
// @Failure 400 {string} string "invalid ID supplied"
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /tcamps/{id}/sportsman [post]
func (h *TCampHandler) RegSm(w http.ResponseWriter, r *http.Request) {
	camp := r.Context().Value(tCampKey).(*dto.TCamp)

	var req dto.RegForTCampReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Render(w, r, http_errors.ErrInvalidRequest(err))
		return
	}

	appl, err := converters.NewTCApplConverter().
		FromRegisterReq(camp.ID, &req)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	appl, err = h.serv.RegisterSportsman(appl)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	applDTO, err := converters.NewTCApplConverter().ToDTO(appl)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	if err := render.Render(w, r, dto.NewRegForTCampResp(applDTO)); err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

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
	compKey KeyStringT = "competition"
)

type CompHandler struct {
	serv   service.ICompetitionService
	servSm service.ISportsmanService
}

func NewCompHandler(
	serv service.ICompetitionService,
	servSm service.ISportsmanService,
) *CompHandler {
	return &CompHandler{
		serv:   serv,
		servSm: servSm,
	}
}

// GetAllComps godoc
// @Summary Получить все соревнования
// @Tags competitions
// @Param name query string false "Название"
// @Param city query string false "Город"
// @Param sort query string false "Сортировка"
// @Param page query string false "Номер страницы"
// @Param batch query string false "Кол-во элементов на странице"
// @Success 200 {object} CompResp
// @Failure 500 {string} string "internal server error"
// @Router /competitions [get]
func (h *CompHandler) GetAllComps(w http.ResponseWriter, r *http.Request) {
	pagination := r.Context().Value(mw.PaginationKey).(mw.Pagination)
	sort := r.Context().Value(mw.SortKey).(string)
	filter := r.Context().
		Value(mw.NameFilterKey).(string) + " " + r.Context().
		Value(mw.CityFilterKey).(string)

	log.Printf("Pagination: Page=%d, Batch=%d", pagination.Page, pagination.Batch)
	log.Printf("Sort: %v", sort)
	log.Printf("Filter: %v", filter)

	c, err := h.serv.ListCompetitions(pagination.Page,
		pagination.Batch,
		sort,
		filter)

	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	cList := []*dto.Competition{}

	for _, item := range c {
		cDTO, err := converters.NewCompConverter().ToDTO(item)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}
		cList = append(cList, cDTO)
	}

	err = render.RenderList(w, r, dto.NewCompListResp(cList))
	if err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

func (h *CompHandler) CompCtx(next http.Handler) http.Handler {
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
		cDomain, err := h.serv.GetCompetitionByID(id)
		if err != nil {
			render.Render(w, r, http_errors.ErrNotFound)
			return
		}
		if cDomain == nil {
			render.Render(w, r, http_errors.ErrNotFound)
			return
		}
		c, err := converters.NewCompConverter().ToDTO(cDomain)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}

		ctx := context.WithValue(r.Context(), compKey, c)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetResults godoc
// @Summary Получить все результаты соревнования
// @Tags competitions
// @Param id path uuid true "ID соревнований"
// @Success 200 {object} []ResultResp
// @Failure 400 {string} string "invalid ID supplied"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /competitions/{id}/results [get]
func (h *CompHandler) GetResults(w http.ResponseWriter, r *http.Request) {
	comp := r.Context().Value(compKey).(*dto.Competition)

	res, err := h.serv.ListResults(comp.ID)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	smList := make([]*dto.Sportsman, len(res))
	for i, item := range res {
		sm, err := h.servSm.GetSportsmanByID(item.SportsmanID)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}
		smList[i], err = converters.NewSportsmanConverter().ToDTO(sm)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}
	}

	resList := []*dto.Result{}

	for _, item := range res {
		resDTO, err := converters.NewResultConverter().ToDTO(item)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}
		resList = append(resList, resDTO)
	}

	compsList := make([]*dto.Competition, len(resList))
	for i := range compsList {
		compsList[i] = comp
	}

	err = render.RenderList(w, r,
		dto.NewResultListResp(resList, smList, compsList))
	if err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

// DeleteComp godoc
// @Summary Удалить соревнование
// @Tags competitions
// @Param id path uuid true "ID соревнования"
// @Success 204 {string} string "ok"
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /competitions/{id} [delete]
func (h *CompHandler) DeleteComp(w http.ResponseWriter, r *http.Request) {
	c := r.Context().Value(compKey).(*dto.Competition)

	err := h.serv.Delete(c.ID)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204
}

// CreateComp godoc
// @Summary Создать соревнование
// @Tags competitions
// @Success 200 {object} CompResp
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /competitions [post]
func (h *CompHandler) CreateComp(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCompReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Render(w, r, http_errors.ErrInvalidRequest(err))
		return
	}

	comp, err := converters.NewCompConverter().FromCreateReq(&req)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	comp, err = h.serv.Create(comp)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	compDTO, err := converters.NewCompConverter().ToDTO(comp)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	if err := render.Render(w, r, dto.NewCompResp(compDTO)); err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

// RegSm godoc
// @Summary Запись спортсмена на соревнования
// @Tags competitions
// @Param id path uuid true "ID соревнования"
// @Param application body RegForCompReq true "Заявка"
// @Success 200 {object} RegForCompResp
// @Failure 400 {string} string "invalid ID supplied"
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /competitions/{id}/sportsman [post]
func (h *CompHandler) RegSm(w http.ResponseWriter, r *http.Request) {
	comp := r.Context().Value(compKey).(*dto.Competition)

	var req dto.RegForCompReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Render(w, r, http_errors.ErrInvalidRequest(err))
		return
	}

	appl, err := converters.NewCompApplConverter().FromRegisterReq(comp.ID, &req)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	appl, err = h.serv.RegisterSportsman(appl)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	applDTO, err := converters.NewCompApplConverter().ToDTO(appl)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	if err := render.Render(w, r, dto.NewRegForCompResp(applDTO)); err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

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

type KeyStringT string

const (
	smKey KeyStringT = "sportsman"
)

type SportsmanHandler struct {
	serv      service.ISportsmanService
	servCoach service.ICoachService
	servComp  service.ICompetitionService
}

func NewSportsmanHandler(
	serv service.ISportsmanService,
	servCoach service.ICoachService,
	servComp service.ICompetitionService,
) *SportsmanHandler {
	return &SportsmanHandler{
		serv:      serv,
		servCoach: servCoach,
		servComp:  servComp,
	}
}

// GetAllSportsmen godoc
// @Summary Получить всех спортсменов
// @Tags sportsmen
// @Param fullname query string false "ФИО"
// @Param sort query string false "Сортировка"
// @Param page query string false "Номер страницы"
// @Param batch query string false "Кол-во элементов на странице"
// @Success 200 {object} []SportsmanResp
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /sportsmen [get]
func (h *SportsmanHandler) GetAllSportsmen(w http.ResponseWriter, r *http.Request) {
	pagination := r.Context().Value(mw.PaginationKey).(mw.Pagination)
	sort := r.Context().Value(mw.SortKey).(string)
	filter := r.Context().Value(mw.FNameFilterKey).(string)

	log.Printf("Pagination: Page=%d, Batch=%d", pagination.Page, pagination.Batch)
	log.Printf("Sort: %v", sort)
	log.Printf("Filter: %v", filter)

	sm, err := h.serv.ListSportsmen(pagination.Page,
		pagination.Batch,
		sort,
		filter)

	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	smList := []*dto.Sportsman{}

	for _, item := range sm {
		smDTO, err := converters.NewSportsmanConverter().ToDTO(item)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}
		smList = append(smList, smDTO)
	}

	err = render.RenderList(w, r, dto.NewSportsmanListResp(smList))
	if err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

func (h *SportsmanHandler) SportsmanCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// SportsmanCtx middleware is used to load a Sportsman object from
		// the URL parameters passed through as the request. In case
		// the Sportsman could not be found, we stop here and return a 404.
		var err error
		smID := chi.URLParam(r, "id")
		if smID == "" {
			render.Render(w, r, http_errors.ErrInvalidRequest(err))
			return
		}
		id, err := uuid.Parse(smID)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}
		smDomain, err := h.serv.GetSportsmanByID(id)
		if err != nil {
			render.Render(w, r, http_errors.ErrNotFound)
			return
		}
		if smDomain == nil {
			render.Render(w, r, http_errors.ErrNotFound)
			return
		}
		sm, err := converters.NewSportsmanConverter().ToDTO(smDomain)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}

		ctx := context.WithValue(r.Context(), smKey, sm)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetSportsman godoc
// @Summary Получить спортсмена по ID
// @Tags sportsmen
// @Param id path uuid true "ID спортсмена"
// @Success 200 {object} SportsmanResp
// @Failure 400 {string} string "invalid ID supplied"
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /sportsmen/{id} [get]
func (h *SportsmanHandler) GetSportsman(w http.ResponseWriter, r *http.Request) {
	sm := r.Context().Value(smKey).(*dto.Sportsman)
	if err := render.Render(w, r, dto.NewSportsmanResp(sm)); err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

// UpdateSportsman godoc
// @Summary Обновление секретарем информации о спортсмене
// @Description Обновление информации о сертификате и допуске на соревнования
// @Tags sportsmen
// @Param id path uuid true "ID спортсмена"
// @Param sportsman body UpdateSportsmanReq true "Спортсмен"
// @Success 200 {object} SportsmanResp
// @Failure 400 {string} string "invalid ID supplied"
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /sportsmen/{id} [patch]
func (h *SportsmanHandler) UpdateSportsman(w http.ResponseWriter, r *http.Request) {
	sm := r.Context().Value(smKey).(*dto.Sportsman)
	var req dto.UpdateSportsmanReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Render(w, r, http_errors.ErrInvalidRequest(err))
		return
	}

	smDomain, err := converters.NewSportsmanConverter().
		FromUpdateReq(sm, &req)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	smDomain, err = h.serv.Update(smDomain) // domain
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}
	if smDomain == nil {
		render.Render(w, r, http_errors.ErrNotFound)
		return
	}

	sm, err = converters.NewSportsmanConverter().ToDTO(smDomain)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	if err := render.Render(w, r, dto.NewSportsmanResp(sm)); err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

func (h *SportsmanHandler) DeleteSportsman(w http.ResponseWriter, r *http.Request) {
	sm := r.Context().Value(smKey).(*dto.Sportsman)

	err := h.serv.Delete(sm.ID)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204
}

// RegForCoach godoc
// @Summary Запись спортсмена к тренеру
// @Tags sportsmen
// @Param id path uuid true "ID спортсмена"
// @Param coach body CoachIDReq true "ID тренера"
// @Success 200 {object} SmCoachResp
// @Failure 400 {string} string "invalid ID supplied"
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /sportsmen/{id}/coach [post]
func (h *SportsmanHandler) RegForCoach(w http.ResponseWriter, r *http.Request) {
	sm := r.Context().Value(smKey).(*dto.Sportsman)

	var req dto.CoachIDReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Render(w, r, http_errors.ErrInvalidRequest(err))
		return
	}

	res, err := h.servCoach.AddSportsman(req.CID, sm.ID)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	if err := render.Render(w, r, dto.NewSmCoachResp(res.SportsmanID, res.CoachID)); err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

// GetResults godoc
// @Summary Получить все результаты спортсмена
// @Tags sportsmen
// @Param id path uuid true "ID спортсмена"
// @Success 200 {object} []ResultResp
// @Failure 400 {string} string "invalid ID supplied"
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /sportsmen/{id}/results [get]
func (h *SportsmanHandler) GetResults(w http.ResponseWriter, r *http.Request) {
	sm := r.Context().Value(smKey).(*dto.Sportsman)

	res, err := h.serv.ListResults(sm.ID)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	comps, err := h.servComp.ListCompsByRes(res)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	compsList := []*dto.Competition{}
	for _, item := range comps {
		compDTO, err := converters.NewCompConverter().ToDTO(item)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}
		compsList = append(compsList, compDTO)
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

	smList := make([]*dto.Sportsman, len(resList))
	for i := range smList {
		smList[i] = sm
	}

	err = render.RenderList(w, r,
		dto.NewResultListResp(resList, smList, compsList))
	if err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

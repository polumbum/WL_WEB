package handlers

import (
	"context"
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
	coachKey KeyStringT = "coach"
)

type CoachHandler struct {
	serv     service.ICoachService
	servComp service.ICompetitionService
	servSm   service.ISportsmanService
}

func NewCoachHandler(
	serv service.ICoachService,
	servComp service.ICompetitionService,
	servSm service.ISportsmanService,

) *CoachHandler {
	return &CoachHandler{
		serv:     serv,
		servComp: servComp,
		servSm:   servSm,
	}
}

// GetAllCoaches godoc
// @Summary Получить всех тренеров
// @Tags сoaches
// @Param fullname query string false "ФИО"
// @Param sort query string false "Сортировка"
// @Param page query string false "Номер страницы"
// @Param batch query string false "Кол-во элементов на странице"
// @Success 200 {object} CoachResp
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Router /coaches [get]
func (h *CoachHandler) GetAllCoaches(w http.ResponseWriter, r *http.Request) {
	pagination := r.Context().Value(mw.PaginationKey).(mw.Pagination)
	sort := r.Context().Value(mw.SortKey).(string)
	filter := r.Context().Value(mw.FNameFilterKey).(string)

	c, err := h.serv.ListCoaches(pagination.Page,
		pagination.Batch,
		sort,
		filter)

	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}

	cList := []*dto.Coach{}

	for _, item := range c {
		cDTO, err := converters.NewCoachConverter().ToDTO(item)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}
		cList = append(cList, cDTO)
	}

	err = render.RenderList(w, r, dto.NewCoachListResp(cList))
	if err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

func (h *CoachHandler) CoachCtx(next http.Handler) http.Handler {
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
		cEntity, err := h.serv.GetCoachByID(id)
		if err != nil {
			render.Render(w, r, http_errors.ErrNotFound)
			return
		}
		if cEntity == nil {
			render.Render(w, r, http_errors.ErrNotFound)
			return
		}
		c, err := converters.NewCoachConverter().ToDTO(cEntity)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}

		ctx := context.WithValue(r.Context(), coachKey, c)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *CoachHandler) SportsmanCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		smID := chi.URLParam(r, "sm_id")
		if smID == "" {
			render.Render(w, r, http_errors.ErrInvalidRequest(err))
			return
		}
		id, err := uuid.Parse(smID)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}
		smDomain, err := h.servSm.GetSportsmanByID(id)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}
		if smDomain == nil {
			render.Render(w, r, http_errors.ErrNotFound)
			return
		}
		c, err := converters.NewSportsmanConverter().ToDTO(smDomain)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}

		ctx := context.WithValue(r.Context(), smKey, c)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetCoach godoc
// @Summary Получить тренера по ID
// @Tags coaches
// @Param id path uuid true "ID тренера"
// @Success 200 {object} CoachResp
// @Failure 400 {string} string "invalid ID supplied"
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /coaches/{id} [get]
func (h *CoachHandler) GetCoach(w http.ResponseWriter, r *http.Request) {
	c := r.Context().Value(coachKey).(*dto.Coach)
	if err := render.Render(w, r, dto.NewCoachResp(c)); err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

// GetSmResults godoc
// @Summary Получить все результаты спортсменов тренера
// @Tags coaches
// @Param id path uuid true "ID тренера"
// @Success 200 {object} CoachResp
// @Failure 400 {string} string "invalid ID supplied"
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /coaches/{id}/sportsmen/results [get]
func (h *CoachHandler) GetResults(w http.ResponseWriter, r *http.Request) {
	c := r.Context().Value(coachKey).(*dto.Coach)

	//sm := r.Context().Value(smKey).(*dto.Sportsman)

	sm, res, err := h.serv.ListResults(c.ID)
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

	smList := []*dto.Sportsman{}
	for _, item := range sm {
		smDTO, err := converters.NewSportsmanConverter().ToDTO(item)
		if err != nil {
			render.Render(w, r, http_errors.ErrServer)
			return
		}
		smList = append(smList, smDTO)
	}

	err = render.RenderList(w, r,
		dto.NewResultListResp(resList, smList, compsList))
	if err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

// GetSportsmen godoc
// @Summary Получить всех спортсменов тренера
// @Tags coaches
// @Param id path uuid true "ID тренера"
// @Success 200 {object} CoachResp
// @Failure 400 {string} string "invalid ID supplied"
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /coaches/{id}/sportsmen [get]
func (h *CoachHandler) GetSportsmen(w http.ResponseWriter, r *http.Request) {
	pagination := r.Context().Value(mw.PaginationKey).(mw.Pagination)
	sort := r.Context().Value(mw.SortKey).(string)
	filter := r.Context().Value(mw.FNameFilterKey).(string)

	c := r.Context().Value(coachKey).(*dto.Coach)

	sm, err := h.serv.ListSportsmen(c.ID,
		pagination.Page,
		pagination.Batch,
		sort,
		filter,
	)
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

	err = render.RenderList(w, r,
		dto.NewSportsmanListResp(smList))
	if err != nil {
		render.Render(w, r, http_errors.ErrRender(err))
		return
	}
}

// RemoveSportsman godoc
// @Summary Отменить запись спортсмена к тренеру
// @Tags coaches
// @Param id path uuid true "ID тренера"
// @Param sm_id path uuid true "ID спортсмена"
// @Success 204 {string} string "ok"
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /coaches/{id}/sportsmen/{sm_id} [delete]
func (h *CoachHandler) RemoveSportsman(w http.ResponseWriter, r *http.Request) {
	c := r.Context().Value(coachKey).(*dto.Coach)
	sm := r.Context().Value(smKey).(*dto.Sportsman)

	err := h.serv.RemoveSportsman(c.ID, sm.ID)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}
}

// DeleteCoach godoc
// @Summary Удалить тренера
// @Tags coaches
// @Param id path uuid true "ID тренера"
// @Success 204 {string} string "ok"
// @Failure 401 {string} string "unauthorized"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /coaches/{id} [delete]
func (h *CoachHandler) DeleteCoach(w http.ResponseWriter, r *http.Request) {
	c := r.Context().Value(coachKey).(*dto.Coach)

	err := h.serv.Delete(c.ID)
	if err != nil {
		render.Render(w, r, http_errors.ErrServer)
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204
}

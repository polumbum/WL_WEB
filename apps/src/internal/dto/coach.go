package dto

import (
	"net/http"
	errors "src/internal/http-server/errors"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// Coach model info
// @Description Тренер
type Coach struct {
	ID         uuid.UUID `json:"id"`
	Surname    string    `json:"surname"`
	Name       string    `json:"name"`
	Patronymic string    `json:"patronymic"`
	Birthday   time.Time `json:"birthday"`
	Experience int       `json:"experience"`
	Gender     bool      `json:"gender"`
}

// Coach model info
// @Description Тренер
type CoachReq struct {
	*Coach
} // @name CoachReq

type UpdateCoachReq struct {
	Surname    string    `json:"surname"`
	Name       string    `json:"name"`
	Patronymic string    `json:"patronymic"`
	Experience int       `json:"experience"`
	Birthday   time.Time `json:"birthday"`
	Gender     bool      `json:"gender"`
} // @name UpdateCoachReq

type CreateCoachReq struct {
	Surname    string    `json:"surname"`
	Name       string    `json:"name"`
	Patronymic string    `json:"patronymic"`
	Experience int       `json:"experience"`
	Birthday   time.Time `json:"birthday"`
	Gender     bool      `json:"gender"`
} // @name CreateCoachReq

func (s *CoachReq) Bind(r *http.Request) error {
	if s.Coach == nil {
		return errors.ErrSmFields
	}
	return nil
}

// Coach
// @Description Тренер
type CoachResp struct {
	*Coach
	//additional fields
} // @name CoachResp

func NewCoachResp(c *Coach) *CoachResp {
	resp := &CoachResp{Coach: c}

	return resp
}

func (rd *CoachResp) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil

}

func NewCoachListResp(c []*Coach) []render.Renderer {
	list := []render.Renderer{}
	for _, item := range c {
		list = append(list, NewCoachResp(item))
	}

	return list
}

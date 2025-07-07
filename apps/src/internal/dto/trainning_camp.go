package dto

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// TCamp model info
// @Description Спортивные сборы
type TCamp struct {
	ID      uuid.UUID `json:"id"`
	City    string    `json:"city"`
	Address string    `json:"address"`
	BegDate time.Time `json:"beg_date"`
	EndDate time.Time `json:"end_date"`
}

// TCamp model info
// @Description Спортивные сборы
type TCampReq struct {
	*TCamp
} // @name TCampReq

// Create TCamp
// @Description Создать спортивные сборы
type CreateTCampReq struct {
	City    string    `json:"city"`
	Address string    `json:"address"`
	BegDate time.Time `json:"beg_date"`
	EndDate time.Time `json:"end_date"`
} // @name CreateTCampReq

// Register for TCamp
// @Description Зарегистрировать спортсмена на сборы
type RegForTCampReq struct {
	SmID uuid.UUID `json:"sm_id"`
} // @name RegForTCampReq

// TCamp model info
// @Description Спортивные сборы
type TCampResp struct {
	*TCamp
} // @name TCampResp

func NewTCampResp(c *TCamp) *TCampResp {
	resp := &TCampResp{TCamp: c}

	return resp
}

func (rd *TCampResp) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil

}

func NewTCampListResp(c []*TCamp) []render.Renderer {
	list := []render.Renderer{}
	for _, item := range c {
		list = append(list, NewTCampResp(item))
	}

	return list
}

// Register for tcamp
// @Description Зарегистрировать спортсмена на сборы
type TCampApplication struct {
	SmID    uuid.UUID `json:"sm_id"`
	TCampID uuid.UUID `json:"tcamp_id"`
} // @name TCampApplication

// Register for competiton
// @Description Зарегистрировать спортсмена на сборы
type RegForTCampResp struct {
	*TCampApplication
} // @name RegForTCampResp

func NewRegForTCampResp(c *TCampApplication) *RegForTCampResp {
	resp := &RegForTCampResp{TCampApplication: c}

	return resp
}

func (rd *RegForTCampResp) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil

}

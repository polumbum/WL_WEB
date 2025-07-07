package dto

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// Competiton model info
// @Description Соревнование
type Competition struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	City              string    `json:"city"`
	Address           string    `json:"address"`
	BegDate           time.Time `json:"beg_date"`
	EndDate           time.Time `json:"end_date"`
	Age               string    `json:"age"`
	MinSportsCategory string    `json:"min_sports_cat"`
	Antidoping        bool      `json:"antidoping"`
}

// Competiton model info
// @Description Соревнование
type CompReq struct {
	*Competition
} // @name CompReq

// Create competiton
// @Description Создать соревнование
type CreateCompReq struct {
	Name              string    `json:"name"`
	City              string    `json:"city"`
	Address           string    `json:"address"`
	BegDate           time.Time `json:"beg_date"`
	EndDate           time.Time `json:"end_date"`
	Age               string    `json:"age"`
	MinSportsCategory string    `json:"min_sports_cat"`
	Antidoping        bool      `json:"antidoping"`
	OrgID             uuid.UUID `json:"org_id"`
} // @name CreateCompReq

// Register for competiton
// @Description Зарегистрировать спортсмена на соревнования
type RegForCompReq struct {
	SmID              uuid.UUID `json:"sm_id"`
	WeighCategory     int       `json:"weight_category"`
	StartSnatch       int       `json:"start_snatch"`
	StartCleanAndJerk int       `json:"start_clean_and_jerk"`
} // @name RegForCompReq

// Competiton model info
// @Description Соревнование
type CompResp struct {
	*Competition
	//additional fields
} // @name CompResp

func NewCompResp(c *Competition) *CompResp {
	resp := &CompResp{Competition: c}

	return resp
}

func (rd *CompResp) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil

}

func NewCompListResp(c []*Competition) []render.Renderer {
	list := []render.Renderer{}
	for _, item := range c {
		list = append(list, NewCompResp(item))
	}

	return list
}

// Register for competiton
// @Description Зарегистрировать спортсмена на соревнования
type CompApplication struct {
	CompID            uuid.UUID `json:"comp_id"`
	SmID              uuid.UUID `json:"sm_id"`
	WeightCategory    int       `json:"weight_category"`
	StartSnatch       int       `json:"start_snatch"`
	StartCleanAndJerk int       `json:"start_clean_and_jerk"`
} // @name CompApplication

// Register for competiton
// @Description Зарегистрировать спортсмена на соревнования
type RegForCompResp struct {
	*CompApplication
} // @name RegForCompResp

func NewRegForCompResp(c *CompApplication) *RegForCompResp {
	resp := &RegForCompResp{CompApplication: c}

	return resp
}

func (rd *RegForCompResp) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil

}

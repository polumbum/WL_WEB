package payload

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type Result struct {
	SportsmanID    uuid.UUID `json:"sm_id"`
	CompetitionID  uuid.UUID `json:"comp_id"`
	WeightCategory string    `json:"weight_category"`
	Snatch         int       `json:"snatch"`
	CleanAndJerk   int       `json:"clean_and_jerk"`
	Place          int       `json:"place"`
}

type ResultResp struct {
	CompID         uuid.UUID `json:"comp_id"`
	Name           string    `json:"comp_name"`
	SmFullname     string    `json:"sm_fullname"`
	BegDate        time.Time `json:"beg_date"`
	EndDate        time.Time `json:"end_date"`
	City           string    `json:"city"`
	WeightCategory string    `json:"weight_category"`
	Snatch         int       `json:"snatch"`
	CleanAndJerk   int       `json:"clean_and_jerk"`
	Place          int       `json:"place"`
}

func NewResultResp(r *Result, sm *Sportsman, comp *Competition) *ResultResp {
	resp := &ResultResp{
		CompID:         r.CompetitionID,
		Name:           comp.Name,
		SmFullname:     sm.Surname + " " + sm.Name + " " + sm.Patronymic,
		BegDate:        comp.BegDate,
		EndDate:        comp.EndDate,
		City:           comp.City,
		WeightCategory: r.WeightCategory,
		Snatch:         r.Snatch,
		CleanAndJerk:   r.CleanAndJerk,
		Place:          r.Place,
	}

	return resp
}

func (rd *ResultResp) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil

}

func NewResultListResp(r []*Result,
	sm []*Sportsman,
	comp []*Competition) []render.Renderer {
	list := []render.Renderer{}

	smMap := make(map[uuid.UUID]*Sportsman)
	for _, item := range sm {
		smMap[item.ID] = item
	}

	compMap := make(map[uuid.UUID]*Competition)
	for _, item := range comp {
		compMap[item.ID] = item
	}

	for _, res := range r {
		if sm, ok := smMap[res.SportsmanID]; ok {
			if comp, ok := compMap[res.CompetitionID]; ok {
				list = append(list, NewResultResp(res, sm, comp))
			}
		}
	}

	return list
}

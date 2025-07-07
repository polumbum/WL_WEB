package payload

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

/*
type UpdateCoachReq struct {
	SportsCategory     string    `json:"sports_category"`
	MoscowTeam         bool      `json:"moscow_team"`
	Adoping_validity   time.Time `json:"adoping_validity"`
	Access_validity    time.Time `json:"access_validity"`
	Access_institution string    `json:"access_institution"`
} // @name UpdateSportsmanReq
*/

/*
// Coach ID
// @Description ID тренера
type CoachIDReq struct {
	CID uuid.UUID `json:"c_id"`
} // @name CoachIDReq
*/

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
	//Adoping_validity   time.Time `json:"adoping_validity"`
	//Access_validity    time.Time `json:"access_validity"`
	//Access_institution string    `json:"access_institution"`
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

// SmCoachResp
// @Description Запись спортсмена к тренеру
/*type SmCoachResp struct {
	SmID uuid.UUID `json:"sm_id"`
	CID  uuid.UUID `json:"c_id"`
} // @name SmCoachResp*/

/*
func NewSmCoachResp(smID, cID uuid.UUID) *SmCoachResp {
	return &SmCoachResp{
		SmID: smID,
		CID:  cID,
	}
}

func (rd *SmCoachResp) Render(w http.ResponseWriter, r *http.Request) error {
	return nil

}*/

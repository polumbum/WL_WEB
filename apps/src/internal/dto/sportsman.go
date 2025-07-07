package dto

import (
	"net/http"
	"time"

	errors "src/internal/http-server/errors"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// Sportsman model info
// @Description Спортсмен
type Sportsman struct {
	ID             uuid.UUID `json:"id"`
	Surname        string    `json:"surname"`
	Name           string    `json:"name"`
	Patronymic     string    `json:"patronymic"`
	Birthday       time.Time `json:"birthday"`
	SportsCategory string    `json:"sports_category"`
	Gender         bool      `json:"gender"`
	MoscowTeam     bool      `json:"moscow_team"`
}

// Sportsman model info
// @Description Спортсмен
type SportsmanReq struct {
	*Sportsman
} // @name SportsmanReq

// Update sportsman
// @Description Обновление информации о спортсмене
type UpdateSportsmanReq struct {
	SportsCategory     string    `json:"sports_category"`
	MoscowTeam         bool      `json:"moscow_team"`
	Adoping_validity   time.Time `json:"adoping_validity"`
	Access_validity    time.Time `json:"access_validity"`
	Access_institution string    `json:"access_institution"`
} // @name UpdateSportsmanReq

// Coach ID
// @Description ID тренера
type CoachIDReq struct {
	CID uuid.UUID `json:"c_id"`
} // @name CoachIDReq

func (s *SportsmanReq) Bind(r *http.Request) error {
	if s.Sportsman == nil {
		return errors.ErrSmFields
	}

	return nil
}

// Sportsman - extended info
// @Description Расширенная информация о спортсмене
type SportsmanResp struct {
	*Sportsman
	//additional fields
	Adoping_validity   time.Time `json:"adoping_validity"`
	Access_validity    time.Time `json:"access_validity"`
	Access_institution string    `json:"access_institution"`
} // @name SportsmanResp

func NewSportsmanResp(sm *Sportsman) *SportsmanResp {
	resp := &SportsmanResp{Sportsman: sm}

	return resp
}

func (rd *SportsmanResp) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil

}

func NewSportsmanListResp(sm []*Sportsman) []render.Renderer {
	list := []render.Renderer{}
	for _, item := range sm {
		list = append(list, NewSportsmanResp(item))
	}

	return list
}

// SmCoachResp
// @Description Запись спортсмена к тренеру
type SmCoachResp struct {
	SmID uuid.UUID `json:"sm_id"`
	CID  uuid.UUID `json:"c_id"`
} // @name SmCoachResp

func NewSmCoachResp(smID, cID uuid.UUID) *SmCoachResp {
	return &SmCoachResp{
		SmID: smID,
		CID:  cID,
	}
}

func (rd *SmCoachResp) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Сreate sportsman
// @Description Создание спортсмена
type CreateSportsmanReq struct {
	Surname        string    `json:"surname"`
	Name           string    `json:"name"`
	Patronymic     string    `json:"patronymic"`
	Birthday       time.Time `json:"birthday"`
	Gender         bool      `json:"gender"`
	SportsCategory string    `json:"sports_category"`
	MoscowTeam     bool      `json:"moscow_team"`
} // @name CreateSportsmanReq

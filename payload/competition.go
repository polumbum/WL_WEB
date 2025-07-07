package payload

import (
	"time"

	"github.com/google/uuid"
)

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

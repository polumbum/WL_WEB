package ui

import (
	"fmt"
	"src/internal/entities"
	"src/internal/service/dto"

	"github.com/google/uuid"
)

func (ui *UI) PrintTCamp(comp *entities.TCamp) error {
	if comp == nil {
		return ErrNilPointer
	}
	fmt.Printf("%s\t\t%s\t\t%s\t\t%s\n",
		comp.City,
		comp.Address,
		comp.BegDate.Format("02-01-2006"),
		comp.EndDate.Format("02-01-2006"),
	)

	return nil
}

func (ui *UI) ListTCamps() ([]*entities.TCamp, error) {
	tCamps, err := ui.TCampService.ListTCamps()
	if err != nil {
		return nil, err
	}

	for i, tc := range tCamps {
		fmt.Printf("%d\t", i+1)
		err = ui.PrintTCamp(tc)
		if err != nil {
			return nil, err
		}
	}
	return tCamps, nil
}

func (ui *UI) RegForTCamp(smID uuid.UUID) (*entities.TCampApplication, error) {
	sportsman, err := ui.SportsmanService.GetSportsmanByID(smID)
	if err != nil {
		return nil, err
	}

	tCamps, err := ui.ListTCamps()
	if err != nil {
		return nil, err
	}
	if len(tCamps) == 0 {
		return nil, ErrNoTCamps
	}

	fmt.Println("Введите номер сборов:")
	var choice int
	_, err = fmt.Scanf("%d", &choice)
	for err != nil || choice <= 0 || choice > len(tCamps) {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}

	req := &dto.RegForTCampReq{
		TCampID:     tCamps[choice-1].ID,
		SportsmanID: sportsman.ID,
	}

	appl, err := ui.TCampService.RegisterSportsman(req)
	if err != nil {
		return nil, err
	}

	return appl, nil
}

func (ui *UI) CreateTCamp() (*entities.TCamp, error) {
	city := ui.InputString("Введите город:")
	address := ui.InputString("Введите адрес:")
	begDate := ui.InputDate("Введите дату начала")
	endDate := ui.InputDate("Введите дату окончания")

	req := &dto.CreateTCampReq{
		City:    city,
		Address: address,
		BegDate: begDate,
		EndDate: endDate,
	}

	tCamp, err := ui.TCampService.Create(req)
	if err != nil {
		return nil, err
	}
	return tCamp, nil
}

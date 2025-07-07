package ui

import (
	"fmt"
	"src/internal/entities"
	"src/internal/service/dto"

	"github.com/google/uuid"
)

func (ui *UI) PrintCompetition(comp *entities.Competition) error {
	if comp == nil {
		return ErrNilPointer
	}

	fmt.Printf("%s\t\t%s\t\t%s\t\t%s\t\t%s\t\t%s\t\t%s\t\t%t\n",
		comp.Name,
		comp.City,
		comp.Address,
		comp.BegDate.Format("02-01-2006"),
		comp.EndDate.Format("02-01-2006"),
		comp.Age,
		comp.MinSportsCategory,
		comp.Antidoping,
	)

	return nil
}

func (ui *UI) ListCompetitions() ([]*entities.Competition, error) {
	comps, err := ui.CompService.ListCompetitions()
	if err != nil {
		return nil, err
	}

	for i, c := range comps {
		fmt.Printf("%d\t", i+1)
		err = ui.PrintCompetition(c)
		if err != nil {
			return nil, err
		}
	}
	return comps, nil
}

func (ui *UI) RegForComp(smID uuid.UUID) (*entities.CompApplication, error) {
	sportsman, err := ui.SportsmanService.GetSportsmanByID(smID)
	if err != nil {
		return nil, err
	}

	comps, err := ui.ListCompetitions()
	if err != nil {
		return nil, err
	}
	if len(comps) == 0 {
		return nil, ErrNoComps
	}

	fmt.Println("Введите номер соревнования:")
	var choice int
	_, err = fmt.Scanf("%d", &choice)
	for err != nil || choice <= 0 || choice > len(comps) {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}

	req := &dto.RegForCompReq{
		CompetitionID: comps[choice-1].ID,
		SportsmanID:   sportsman.ID,
	}

	fmt.Println("Выберите весовую категорию:")
	req.WeighCategory = ui.InputWeightCat(sportsman.Gender)

	fmt.Println("Введите начальный вес в упражнении 'рывок':")
	_, err = fmt.Scanf("%d", &req.StartSnatch)
	for err != nil {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &req.StartSnatch)
	}

	fmt.Println("Введите начальный вес в упражнении 'толчок':")
	_, err = fmt.Scanf("%d", &req.StartCleanAndJerk)
	for err != nil {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &req.StartCleanAndJerk)
	}

	appl, err := ui.CompService.RegisterSportsman(req)
	if err != nil {
		return nil, err
	}

	return appl, nil
}

func (ui *UI) CreateComp() (*entities.Competition, error) {
	name := ui.InputString("Введите название соревнований:")
	city := ui.InputString("Введите город:")
	address := ui.InputString("Введите адрес:")
	begDate := ui.InputDate("Введите дату начала")
	endDate := ui.InputDate("Введите дату окончания")
	age := ui.InputAgeCat("Введите возрастную категорию:")
	minSCat := ui.InputSportsCat("Введите минимальный разряд:")

	req := &dto.CreateCompReq{
		Name:              name,
		City:              city,
		Address:           address,
		BegDate:           begDate,
		EndDate:           endDate,
		Age:               age,
		MinSportsCategory: minSCat,
	}

	comp, err := ui.CompService.Create(req)
	if err != nil {
		return nil, err
	}
	return comp, nil
}

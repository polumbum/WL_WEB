package ui

import (
	"fmt"
	"src/internal/entities"
	"src/internal/service/dto"

	"github.com/google/uuid"
)

func (ui *UI) CreateCoach() (*entities.Coach, error) {
	req := &dto.CreateCoachReq{}

	ui.InputFullName(&req.Surname, &req.Name, &req.Patronymic)
	req.Birthday = ui.InputDate("Введите дату рождения")
	req.Gender = ui.InputGender()
	req.Experience = ui.InputExperience()

	coach, err := ui.CoachService.Create(req)
	if err != nil {
		return nil, err
	}

	return coach, nil
}

func (ui *UI) ListCSportsmen(cID uuid.UUID) ([]*entities.Sportsman, error) {
	sportsmen, err := ui.CoachService.ListSportsmen(cID)
	if err != nil {
		return nil, err
	}

	for i, sm := range sportsmen {
		fmt.Printf("%d\t", i+1)
		err = ui.PrintSportsman(sm)
		if err != nil {
			return nil, err
		}
	}

	return sportsmen, nil
}

func (ui *UI) PrintCoach(c *entities.Coach) error {
	if c == nil {
		return ErrNilPointer
	}
	fmt.Printf("%s\t\t%s\t\t%s\t\t%s\t\t%d\t\t%t\n",
		c.Surname,
		c.Name,
		c.Patronymic,
		c.Birthday.Format("02-01-2006"),
		c.Experience,
		c.Gender,
	)
	return nil
}

func (ui *UI) ListCoaches() ([]*entities.Coach, error) {
	coaches, err := ui.CoachService.ListCoaches()
	if err != nil {
		return nil, err
	}

	for i, c := range coaches {
		fmt.Printf("%d\t", i+1)
		err = ui.PrintCoach(c)
		if err != nil {
			return nil, err
		}
	}
	return coaches, nil
}

func (ui *UI) PrintCSmInfo(cID uuid.UUID) error {
	sportsmen, err := ui.CoachService.ListSportsmen(cID)
	if err != nil {
		return err
	}

	for _, sm := range sportsmen {
		err = ui.PrintSmInfo(sm.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ui *UI) ListCResults(cID uuid.UUID) error {
	sportsmen, err := ui.CoachService.ListSportsmen(cID)
	if err != nil {
		return err
	}

	for _, sm := range sportsmen {
		_, err = ui.ListResults(sm.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ui *UI) CRegForComp(cID uuid.UUID) (*entities.CompApplication, error) {
	sportsmen, err := ui.ListCSportsmen(cID)
	if err != nil {
		return nil, err
	}
	if len(sportsmen) == 0 {
		return nil, ErrNoSm
	}

	fmt.Println("Введите номер спортсмена:")
	var choice int
	_, err = fmt.Scanf("%d", &choice)
	for err != nil || choice <= 0 || choice > len(sportsmen) {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}

	appl, err := ui.RegForComp(sportsmen[choice-1].ID)
	if err != nil {
		return nil, err
	}

	return appl, nil
}

func (ui *UI) CRegForTCamp(cID uuid.UUID) (*entities.TCampApplication, error) {
	sportsmen, err := ui.ListCSportsmen(cID)
	if err != nil {
		return nil, err
	}
	if len(sportsmen) == 0 {
		return nil, ErrNoSm
	}

	fmt.Println("Введите номер спортсмена:")
	var choice int
	_, err = fmt.Scanf("%d", &choice)
	for err != nil || choice <= 0 || choice > len(sportsmen) {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}

	appl, err := ui.RegForTCamp(sportsmen[choice-1].ID)
	if err != nil {
		return nil, err
	}

	return appl, nil
}

func (ui *UI) AddSmToCoach() (*entities.SportsmenCoach, error) {
	sportsmen, err := ui.ListSportsmen()
	if err != nil {
		return nil, err
	}
	if len(sportsmen) == 0 {
		return nil, ErrNoSm
	}

	fmt.Println("Введите номер спортсмена:")
	var sm int
	_, err = fmt.Scanf("%d", &sm)
	for err != nil || sm <= 0 || sm > len(sportsmen) {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &sm)
	}

	coaches, err := ui.ListCoaches()
	if err != nil {
		return nil, err
	}
	if len(coaches) == 0 {
		return nil, ErrNoCoaches
	}

	fmt.Println("Введите номер тренера:")
	var c int
	_, err = fmt.Scanf("%d", &c)
	for err != nil || c <= 0 || c > len(coaches) {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &c)
	}

	res, err := ui.CoachService.AddSportsman(coaches[c-1].ID, sportsmen[sm-1].ID)

	return res, err
}

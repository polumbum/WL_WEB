package ui

import (
	"errors"
	"fmt"
	"src/internal/entities"
	"src/internal/service"
	"src/internal/service/dto"

	"github.com/google/uuid"
)

func (ui *UI) CreateSportsman() (*entities.Sportsman, error) {
	req := &dto.CreateSportsmanReq{}

	ui.InputFullName(&req.Surname, &req.Name, &req.Patronymic)
	req.Birthday = ui.InputDate("Введите дату рождения")
	req.Gender = ui.InputGender()
	req.SportsCategory = ui.InputSportsCat("Введите разряд/звание:")
	req.MoscowTeam = ui.InputMoscowTeam()

	sportsman, err := ui.SportsmanService.Create(req)
	if err != nil {
		return nil, err
	}

	return sportsman, nil
}

func (ui *UI) ListSportsmen() ([]*entities.Sportsman, error) {
	sportsmen, err := ui.SportsmanService.ListSportsmen()
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

func (ui *UI) PrintResult(comp *entities.Competition, res *entities.Result) error {
	if res == nil || comp == nil {
		return ErrNilPointer
	}

	err := ui.PrintCompetition(comp)
	if err != nil {
		return err
	}

	fmt.Printf("Весовая категория: %d\tРывок: %d\tТолчок: %d\tМесто: %d\n",
		res.WeightCategory,
		res.Snatch,
		res.CleanAndJerk,
		res.Place,
	)

	return nil
}

func (ui *UI) ListResults(smID uuid.UUID) ([]*entities.Result, error) {
	results, err := ui.SportsmanService.ListResults(smID)
	if err != nil {
		return nil, err
	}

	sm, err := ui.SportsmanService.GetSportsmanByID(smID)
	if err != nil {
		return nil, err
	}

	fmt.Print("Спортсмен: ")
	err = ui.PrintSportsman(sm)
	if err != nil {
		return nil, err
	}

	for i, res := range results {
		comp, err := ui.CompService.GetCompetitionByID(res.CompetitionID)
		if err != nil {
			return nil, err
		}

		fmt.Printf("%d\t", i+1)
		err = ui.PrintResult(comp, res)
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}

func (ui *UI) UpdateSportsman() (*entities.Sportsman, error) {
	sportsmen, err := ui.ListSportsmen()
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

	req := &dto.UpdateSportsmanReq{
		ID: sportsmen[choice-1].ID,
		//Surname:        sportsmen[choice-1].Surname,
		//Name:           sportsmen[choice-1].Name,
		//Patronymic:     sportsmen[choice-1].Patronymic,
		//Birthday:       sportsmen[choice-1].Birthday,
		SportsCategory: sportsmen[choice-1].SportsCategory,
		//Gender:         sportsmen[choice-1].Gender,
		MoscowTeam: &sportsmen[choice-1].MoscowTeam,
	}

	fmt.Printf("Обновить разряд/звание?\n1. Да\n2. Нет'n")
	_, err = fmt.Scanf("%d", &choice)
	for err != nil || (choice != 1 && choice != 2) {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}
	if choice == 1 {
		req.SportsCategory = ui.InputSportsCat("Введите разряд/звание:")
	}
	fmt.Printf("Обновить статус членства в сборной Москвы?\n1. Да\n2. Нет\n")
	_, err = fmt.Scanf("%d", &choice)
	for err != nil || (choice != 1 && choice != 2) {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}
	if choice == 1 {
		team := ui.InputMoscowTeam()
		req.MoscowTeam = &team
	}

	sportsman, err := ui.SportsmanService.Update(req)
	if err != nil {
		return nil, err
	}

	fmt.Println("Обновить статус сертификата о знании антидопинговых правил?\n1. Да\n2. Нет")
	_, err = fmt.Scanf("%d", &choice)
	for err != nil || (choice != 1 && choice != 2) {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}
	if choice == 1 {
		_, err = ui.UpdateADoping(sportsman.ID)
		if err != nil {
			return sportsman, err
		}
	}

	fmt.Println("Обновить статус допуска к соревнованиям?\n1. Да\n2. Нет")
	_, err = fmt.Scanf("%d", &choice)
	for err != nil || (choice != 1 && choice != 2) {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}
	if choice == 1 {
		_, err = ui.UpdateAccess(sportsman.ID)
		if err != nil {
			return sportsman, err
		}
	}

	return sportsman, nil
}

func (ui *UI) UpdateADoping(smID uuid.UUID) (*entities.Antidoping, error) {
	ad, err := ui.ADopingService.GetADopingByID(smID)
	if errors.Is(err, service.ErrADopingNotFound) && err != nil {
		return nil, err
	}

	validity := ui.InputDate("Введите дату окончания действия сертификата")

	if ad == nil {
		req := &dto.CreateADopingReq{
			SmID:     smID,
			Validity: validity,
		}
		ad, err = ui.ADopingService.Create(req)
		if err != nil {
			return nil, err
		}
	} else {
		req := &dto.UpdateADopingReq{
			SmID:     smID,
			Validity: validity,
		}
		ad, err = ui.ADopingService.Update(req)
		if err != nil {
			return nil, err
		}
	}

	return ad, nil
}

func (ui *UI) UpdateAccess(smID uuid.UUID) (*entities.CompAccess, error) {
	ca, err := ui.AccessService.GetAccessByID(smID)
	if errors.Is(err, service.ErrAccessNotFound) && err != nil {
		return nil, err
	}

	validity := ui.InputDate("Введите дату окончания действия допуска")
	institution := ui.InputString("Введите название мед. учреждения:")

	if ca == nil {
		req := &dto.CreateAccessReq{
			SmID:        smID,
			Validity:    validity,
			Institution: institution,
		}
		ca, err = ui.AccessService.Create(req)
		if err != nil {
			return nil, err
		}
	} else {
		req := &dto.UpdateAccessReq{
			SmID:        smID,
			Validity:    validity,
			Institution: institution,
		}
		ca, err = ui.AccessService.Update(req)
		if err != nil {
			return nil, err
		}
	}

	return ca, nil
}

func (ui *UI) PrintSportsman(sm *entities.Sportsman) error {
	if sm == nil {
		return ErrNilPointer
	}
	fmt.Printf("%s\t\t%s\t\t%s\t\t%s\t\t%s\t\t%t\t\t%t\n",
		sm.Surname,
		sm.Name,
		sm.Patronymic,
		sm.Birthday.Format("02-01-2006"),
		sm.SportsCategory,
		sm.Gender,
		sm.MoscowTeam,
	)
	return nil
}

func (ui *UI) PrintAntidoping(ad *entities.Antidoping) error {
	fmt.Print("Антидопинг: ")
	if ad == nil {
		fmt.Println("нет сертификата")
	} else {
		fmt.Println("до", ad.Validity.Format("02-01-2006"))
	}
	return nil
}

func (ui *UI) PrintAccess(ca *entities.CompAccess) error {
	fmt.Print("Допуск: ")
	if ca == nil {
		fmt.Println("нет допуска")
	} else {
		fmt.Println(ca.Institution, "до", ca.Validity.Format("02-01-2006"))
	}
	return nil
}

func (ui *UI) PrintSmInfo(smID uuid.UUID) error {
	sportsman, err := ui.SportsmanService.GetSportsmanByID(smID)
	if err != nil {
		return err
	}

	ad, _ := ui.ADopingService.GetADopingByID(smID)
	ca, _ := ui.AccessService.GetAccessByID(smID)

	err = ui.PrintSportsman(sportsman)
	if err != nil {
		return err
	}

	err = ui.PrintAntidoping(ad)
	if err != nil {
		return err
	}

	err = ui.PrintAccess(ca)
	if err != nil {
		return err
	}

	return nil
}

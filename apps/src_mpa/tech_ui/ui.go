package ui

import (
	"fmt"
	"log"
	"src/internal/constants"
	"src/internal/entities"
	"src/internal/service"
	"src/internal/service/dto"

	"github.com/google/uuid"
)

type UI struct {
	UserService      service.IUserService
	CoachService     service.ICoachService
	CompService      service.ICompetitionService
	ResultService    service.IResultService
	SportsmanService service.ISportsmanService
	TCampService     service.ITCampService
	AccessService    service.IAccessService
	ADopingService   service.IADopingService
}

/*func NewUI(repo repository.ICoachRepository) *CoachService {
	return &CoachService{repo: repo}
}*/

func (ui *UI) Authorization() (*entities.User, error) {
	fmt.Println("1. Войти\n2. Зарегистрироваться\n3. Продолжить без входа\n0. Выход")
	var choice int
	var err error
	var user *entities.User
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		user, err = ui.Login()
	case 2:
		user, err = ui.Register()
	case 3:
		user = &entities.User{Role: constants.UserRoleGuest}
		err = nil
	case 0:
		user = nil
		err = nil
	}
	return user, err
}

func (ui *UI) GuestFunc(choice int) error {
	switch choice {
	case 1:
		_, err := ui.ListCompetitions()
		if err != nil {
			return err
		}
	case 2:
		_, err := ui.ListTCamps()
		if err != nil {
			return err
		}
	}
	return nil
}

func (ui *UI) SportsmanFunc(id uuid.UUID, choice int) error {
	switch choice {
	case 3:
		_, err := ui.RegForComp(id)
		if err != nil {
			fmt.Println("Ошибка. Отмена регистрации.")
			return err
		}
		fmt.Println("Регистрация прошла успешно.")
	case 4:
		_, err := ui.RegForTCamp(id)
		if err != nil {
			fmt.Println("Ошибка. Отмена регистрации.")
			return err
		}
		fmt.Println("Регистрация прошла успешно.")
	case 5:
		err := ui.PrintSmInfo(id)
		if err != nil {
			return err
		}
	case 6:
		_, err := ui.ListResults(id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ui *UI) SportsmanUI(id uuid.UUID) (int, error) {
	fmt.Println(constants.UserRoleSportsman)
	fmt.Println("\n1. Посмотреть список соревнований")
	fmt.Println("2. Посмотреть список сборов")
	fmt.Println("3. Записаться на соревнования")
	fmt.Println("4. Записаться на сборы")
	fmt.Println("5. Информация о спортсмене")
	fmt.Println("6. Статистика")
	fmt.Println("0. Выход")

	var choice int
	_, err := fmt.Scanf("%d", &choice)
	for err != nil || choice < 0 || choice > 6 {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}
	switch {
	case choice == 1 || choice == 2:
		err = ui.GuestFunc(choice)
	case choice > 2 && choice <= 6:
		err = ui.SportsmanFunc(id, choice)
	}
	return choice, err
}

func (ui *UI) CoachFunc(id uuid.UUID, choice int) error {
	switch choice {
	case 3:
		_, err := ui.CRegForComp(id)
		if err != nil {
			fmt.Println("Ошибка. Отмена регистрации.")
			return err
		}
		fmt.Println("Регистрация прошла успешно.")

	case 4:
		_, err := ui.CRegForTCamp(id)
		if err != nil {
			fmt.Println("Ошибка. Отмена регистрации.")
			return err
		}
		fmt.Println("Регистрация прошла успешно.")
	case 5:
		err := ui.PrintCSmInfo(id)
		if err != nil {
			return err
		}

	case 6:
		err := ui.ListCResults(id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ui *UI) CoachUI(id uuid.UUID) (int, error) {
	fmt.Println(constants.UserRoleCoach)
	fmt.Println("\n1. Посмотреть список соревнований")
	fmt.Println("2. Посмотреть список сборов")
	fmt.Println("3. Записать спортсмена на соревнования")
	fmt.Println("4. Записать спортсмена на сборы")
	fmt.Println("5. Информация о подопечных")
	fmt.Println("6. Статистика подопечных")
	fmt.Println("0. Выход")

	var choice int
	_, err := fmt.Scanf("%d", &choice)
	for err != nil || choice < 0 || choice > 6 {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}
	switch {
	case choice == 1 || choice == 2:
		err = ui.GuestFunc(choice)
	case choice > 2 && choice <= 6:
		err = ui.CoachFunc(id, choice)
	}
	return choice, err
}

func (ui *UI) GuestUI() (int, error) {
	fmt.Println(constants.UserRoleGuest)
	fmt.Println("\n1. Посмотреть список соревнований")
	fmt.Println("2. Посмотреть список сборов")
	fmt.Println("0. Выход")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	for err != nil || choice < 0 || choice > 2 {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}
	if choice == 1 || choice == 2 {
		err = ui.GuestFunc(choice)
	}
	return choice, err
}

func (ui *UI) SecretaryFunc(choice int) error {
	switch choice {
	case 3:
		_, err := ui.AddSmToCoach()
		if err != nil {
			fmt.Println("Отмена записи.")
			return err
		}
		fmt.Println("Запись выполнена.")
	case 4:
		sportsman, err := ui.UpdateSportsman()
		if err != nil {
			return err
		}
		if sportsman != nil {
			fmt.Println("Запись обновлена.")
		} else {
			fmt.Println("Отмена операции.")
		}
	}

	return nil
}

func (ui *UI) SecretaryUI() (int, error) {
	fmt.Println(constants.UserRoleChiefSecretary)
	fmt.Println("\n1. Посмотреть список соревнований")
	fmt.Println("2. Посмотреть список сборов")
	fmt.Println("3. Записать спортсмена к тренеру")
	fmt.Println("4. Обновить информацию о спортсмене")
	fmt.Println("0. Выход")

	var choice int
	_, err := fmt.Scanf("%d", &choice)
	for err != nil || choice < 0 || choice > 4 {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}
	switch {
	case choice == 1 || choice == 2:
		err = ui.GuestFunc(choice)
	case choice > 2:
		err = ui.SecretaryFunc(choice)
	}
	return choice, err
}

func (ui *UI) CompOrgFunc(choice int) error {
	if choice == 3 {
		_, err := ui.CreateComp()
		if err != nil {
			fmt.Println("Ошибка. Отмена организации.")
			return err
		}
		fmt.Println("Соревнование организовано.")
	}
	return nil
}

func (ui *UI) CompOrgUI() (int, error) {
	fmt.Println(constants.UserRoleCompOrganizer)
	fmt.Println("\n1. Посмотреть список соревнований")
	fmt.Println("2. Посмотреть список сборов")
	fmt.Println("3. Организовать соревнование")
	fmt.Println("0. Выход")

	var choice int
	_, err := fmt.Scanf("%d", &choice)
	for err != nil || choice < 0 || choice > 4 {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}
	switch {
	case choice == 1 || choice == 2:
		err = ui.GuestFunc(choice)
	case choice > 2:
		err = ui.CompOrgFunc(choice)
	}
	return choice, err
}

func (ui *UI) TCampOrgFunc(choice int) error {
	if choice == 3 {
		_, err := ui.CreateTCamp()
		if err != nil {
			fmt.Println("Ошибка. Отмена организации.")
			return err
		}
		fmt.Println("Сборы организованы.")
	}
	return nil
}

func (ui *UI) TCampOrgUI() (int, error) {
	fmt.Println("\n1. Посмотреть список соревнований")
	fmt.Println("2. Посмотреть список сборов")
	fmt.Println("3. Организовать сборы")
	fmt.Println("0. Выход")

	var choice int
	_, err := fmt.Scanf("%d", &choice)
	for err != nil || choice < 0 || choice > 4 {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}
	switch {
	case choice == 1 || choice == 2:
		err = ui.GuestFunc(choice)
	case choice > 2:
		err = ui.TCampOrgFunc(choice)
	}
	return choice, err
}

func (ui *UI) RunUI(logger *log.Logger) error {
	user, err := ui.Authorization()
	for err != nil {
		logger.Println(err)
		fmt.Println(err)
		user, err = ui.Authorization()
	}
	for user != nil {
		choice := 1
		role := user.Role
		switch role {
		case constants.UserRoleGuest:
			for choice != 0 {
				choice, err = ui.GuestUI()
				if err != nil {
					logger.Println(err)
					fmt.Println(err)
				}
			}
		case constants.UserRoleChiefSecretary:
			for choice != 0 {
				choice, err = ui.SecretaryUI()
				if err != nil {
					logger.Println(err)
					fmt.Println(err)
				}
			}
		case constants.UserRoleCoach:
			for choice != 0 {
				choice, err = ui.CoachUI(user.RoleID)
				if err != nil {
					logger.Println(err)
					fmt.Println(err)
				}
			}
		case constants.UserRoleSportsman:
			for choice != 0 {
				choice, err = ui.SportsmanUI(user.RoleID)
				if err != nil {
					logger.Println(err)
					fmt.Println(err)
				}
			}
		case constants.UserRoleCompOrganizer:
			for choice != 0 {
				choice, err = ui.CompOrgUI()
				if err != nil {
					logger.Println(err)
					fmt.Println(err)
				}
			}
		case constants.UserRoleTCampOrganizer:
			for choice != 0 {
				choice, err = ui.TCampOrgUI()
				if err != nil {
					logger.Println(err)
					fmt.Println(err)
				}
			}
		}
		user, err = ui.Authorization()
		for err != nil {
			logger.Println(err)
			fmt.Println(err)
			user, err = ui.Authorization()
		}
	}

	return nil
}

func (ui *UI) Login() (*entities.User, error) {
	var email string
	var password string
	fmt.Println("Введите email:")
	fmt.Scanf("%s", &email)
	fmt.Println("Введите пароль:")
	fmt.Scanf("%s", &password)

	req := &dto.LoginUserReq{
		Email:    email,
		Password: password,
	}

	user, err := ui.UserService.Login(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ui *UI) Register() (*entities.User, error) {
	var email string
	var password string
	fmt.Println("Введите email:")
	fmt.Scanf("%s", &email)
	fmt.Println("Введите пароль:")
	fmt.Scanf("%s", &password)

	req := &dto.RegisterUserReq{
		Email:    email,
		Password: password,
	}

	fmt.Println("Кто вы?")
	fmt.Println("1. Спортсмен\n2. Тренер")
	fmt.Println("3. Организатор соревнований")
	fmt.Println("4. Организатор сборов\n5. Главный секретарь ФТАМ")
	var roleChoice int
	fmt.Scanf("%d", &roleChoice)
	switch roleChoice {
	case 1:
		req.Role = constants.UserRoleSportsman
		sportsman, err := ui.CreateSportsman()
		if err != nil {
			return nil, err
		}
		req.RoleID = sportsman.ID
	case 2:
		req.Role = constants.UserRoleCoach
		coach, err := ui.CreateCoach()
		if err != nil {
			return nil, err
		}
		req.RoleID = coach.ID
	case 3:
		req.Role = constants.UserRoleCompOrganizer
		req.RoleID = uuid.Nil
	case 4:
		req.Role = constants.UserRoleTCampOrganizer
		req.RoleID = uuid.Nil
	case 5:
		req.Role = constants.UserRoleChiefSecretary
		req.RoleID = uuid.Nil
	default:
		return nil, ErrRole
	}

	user, err := ui.UserService.Register(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

package ui

import (
	"bufio"
	"fmt"
	"os"
	"src/internal/constants"
	"strings"
	"time"
)

func (ui *UI) InputWeightCat(gender constants.GenderT) constants.WeightCategoryT {
	var cat map[int]constants.WeightCategoryT
	if gender == constants.Female {
		cat = map[int]constants.WeightCategoryT{
			1:  constants.WC29,
			2:  constants.WC31,
			3:  constants.WC33,
			4:  constants.WC35,
			5:  constants.WC37,
			6:  constants.WC40,
			7:  constants.WC45,
			8:  constants.WC49,
			9:  constants.WC55,
			10: constants.WC59,
			11: constants.WC64,
			12: constants.WC71,
			13: constants.WC76,
			14: constants.WC81,
			15: constants.WC87,
			16: constants.WC87plus,
		}
	} else {
		cat = map[int]constants.WeightCategoryT{
			1:  constants.WC29,
			2:  constants.WC31,
			3:  constants.WC33,
			4:  constants.WC35,
			5:  constants.WC37,
			6:  constants.WC41,
			7:  constants.WC45,
			8:  constants.WC49,
			9:  constants.WC55,
			10: constants.WC61,
			11: constants.WC67,
			12: constants.WC73,
			13: constants.WC81,
			14: constants.WC89,
			15: constants.WC96,
			16: constants.WC102,
			17: constants.WC109,
			18: constants.WC109plus,
		}
	}

	for i := 1; i <= len(cat); i++ {
		fmt.Printf("%d. %d\n", i, cat[i])
	}

	fmt.Println("Введите весовую категорию:")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	for err != nil || choice < 1 || choice > len(cat) {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}

	return cat[choice]
}

func (ui *UI) InputAgeCat(text string) constants.AgeCategoryT {
	cat := map[int]constants.AgeCategoryT{
		1: constants.AgeCategoryY15_23,
		2: constants.AgeCategoryY19_20,
		3: constants.AgeCategoryY15_18,
		4: constants.AgeCategoryBG13_17,
		5: constants.AgeCategoryBG13_15,
		6: constants.AgeCategoryBG10_12,
		7: constants.AgeCategoryY21_23,
		8: constants.AgeCategoryY17_25,
		9: constants.AgeCategoryMW,
	}
	for i := 1; i <= len(cat); i++ {
		fmt.Printf("%d. %s\n", i, cat[i])
	}

	fmt.Println(text)
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	for err != nil || choice < 1 || choice > len(cat) {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}

	return cat[choice]
}

func (ui *UI) InputSportsCat(text string) constants.SportsCategoryT {
	cat := map[int]constants.SportsCategoryT{
		1:  constants.SportsCategory3youth,
		2:  constants.SportsCategory2youth,
		3:  constants.SportsCategory1youth,
		4:  constants.SportsCategory3,
		5:  constants.SportsCategory2,
		6:  constants.SportsCategory1,
		7:  constants.SportsCategoryCMS,
		8:  constants.SportsCategoryMS,
		9:  constants.SportsCategoryMSIC,
		10: constants.SportsCategoryT(""),
	}
	for i := 1; i < len(cat); i++ {
		fmt.Printf("%d. %s\n", i, cat[i])
	}
	fmt.Println(len(cat), ". Нет разряда")

	fmt.Println(text)
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	for err != nil || choice > 10 || choice < 1 {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &choice)
	}

	return cat[choice]
}

func (ui *UI) InputGender() constants.GenderT {
	fmt.Println("Введите пол:\n1. Женщина\n2. Мужчина")
	var gender int
	_, err := fmt.Scanf("%d", &gender)
	for err != nil || (gender != 1 && gender != 2) {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &gender)
	}
	if gender == 1 {
		return constants.Female
	}
	return constants.Male
}

func (ui *UI) InputMoscowTeam() bool {
	fmt.Println("Являетесь ли членом сборной Москвы?\n1. Да\n2. Нет")
	var moscowTeam int
	_, err := fmt.Scanf("%d", &moscowTeam)
	for err != nil || (moscowTeam != 1 && moscowTeam != 2) {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &moscowTeam)
	}

	return moscowTeam == 1
}

func (ui *UI) InputDate(text string) time.Time {
	fmt.Println(text, "(в формате дд.мм.гггг):")
	var day, year int
	var month time.Month
	_, err := fmt.Scanf("%d.%d.%d", &day, &month, &year)
	for err != nil {
		fmt.Println("Неверный формат даты. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d.%d.%d", &day, &month, &year)
	}
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func (ui *UI) InputString(text string) string {
	fmt.Println(text)
	reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
	res = strings.TrimRight(res, "\n")
	return res
}

func (ui *UI) InputFullName(surname, name, patronymic *string) {
	fmt.Println("Введите Фамилию Имя Отчество (через пробел, отчество - если есть):")
	count, _ := fmt.Scanf("%s %s %s", surname, name, patronymic)
	for count < 2 {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		count, _ = fmt.Scanf("%s %s %s", surname, name, patronymic)
	}
}

func (ui *UI) InputExperience() int {
	fmt.Println("Введите опыт работы (в годах):")
	var years int
	_, err := fmt.Scanf("%d", &years)
	for err != nil {
		fmt.Println("Неверный ввод. Попробуйте ещё раз.")
		_, err = fmt.Scanf("%d", &years)
	}

	return years
}

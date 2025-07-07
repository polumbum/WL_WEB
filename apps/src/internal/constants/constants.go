package constants

const (
	MinWeight int = 15
)

type SportsCategoryT string

const (
	SportsCategory3youth SportsCategoryT = "III(юн)"
	SportsCategory2youth SportsCategoryT = "II(юн)"
	SportsCategory1youth SportsCategoryT = "I(юн)"
	SportsCategory3      SportsCategoryT = "III"
	SportsCategory2      SportsCategoryT = "II"
	SportsCategory1      SportsCategoryT = "I"
	SportsCategoryCMS    SportsCategoryT = "КМС"
	SportsCategoryMS     SportsCategoryT = "МС"
	SportsCategoryMSIC   SportsCategoryT = "МСМК"
)

func GetSportsCat() []SportsCategoryT {
	sportsCat := []SportsCategoryT{
		SportsCategory3youth,
		SportsCategory2youth,
		SportsCategory1youth,
		SportsCategory3,
		SportsCategory2,
		SportsCategory1,
		SportsCategoryCMS,
		SportsCategoryMS,
		SportsCategoryMSIC,
	}
	return sportsCat
}

type GenderT bool

const (
	Male   GenderT = false
	Female GenderT = true
)

type WeightCategoryT int

const (
	WC29      WeightCategoryT = 29  // F
	WC31      WeightCategoryT = 31  // F
	WC33      WeightCategoryT = 33  // F M
	WC35      WeightCategoryT = 35  // F M
	WC37      WeightCategoryT = 37  // F M
	WC40      WeightCategoryT = 40  // F
	WC41      WeightCategoryT = 41  // M
	WC45      WeightCategoryT = 45  // F M
	WC49      WeightCategoryT = 49  // F M
	WC55      WeightCategoryT = 55  // F M
	WC59      WeightCategoryT = 59  // F
	WC61      WeightCategoryT = 61  // M
	WC64      WeightCategoryT = 64  // F
	WC67      WeightCategoryT = 67  // M
	WC71      WeightCategoryT = 71  // F
	WC73      WeightCategoryT = 73  // M
	WC76      WeightCategoryT = 76  // F
	WC81      WeightCategoryT = 81  // F M
	WC87      WeightCategoryT = 87  // F
	WC87plus  WeightCategoryT = 88  // F
	WC89      WeightCategoryT = 89  // M
	WC96      WeightCategoryT = 96  // M
	WC102     WeightCategoryT = 102 // M
	WC109     WeightCategoryT = 109 // M
	WC109plus WeightCategoryT = 110 // M
)

func GetWeightMale() []WeightCategoryT {
	weightMale := []WeightCategoryT{
		WC33,
		WC35,
		WC37,
		WC41,
		WC45,
		WC49,
		WC55,
		WC61,
		WC67,
		WC73,
		WC81,
		WC89,
		WC96,
		WC102,
		WC109,
		WC109plus,
	}
	return weightMale
}

func GetWeightFemale() []WeightCategoryT {
	weightFemale := []WeightCategoryT{
		WC29,
		WC31,
		WC33,
		WC35,
		WC37,
		WC40,
		WC45,
		WC49,
		WC55,
		WC59,
		WC64,
		WC71,
		WC76,
		WC81,
		WC87,
		WC87plus,
	}
	return weightFemale
}

type AgeCategoryT string

const (
	AgeCategoryY15_23  AgeCategoryT = "юниоры, юниорки (15-23 года)" //
	AgeCategoryY19_20  AgeCategoryT = "юниоры, юниорки (19-20 лет)"  //
	AgeCategoryY15_18  AgeCategoryT = "юниоры, юниорки (15-18 лет)"  //
	AgeCategoryBG13_17 AgeCategoryT = "юноши, девушки (13-17 лет)"   //
	AgeCategoryBG13_15 AgeCategoryT = "юноши, девушки (13-15 лет)"   //
	AgeCategoryBG10_12 AgeCategoryT = "юноши, девушки (10-12 лет)"   //
	AgeCategoryY21_23  AgeCategoryT = "юниоры, юниорки (21-23 лет)"  //
	AgeCategoryY17_25  AgeCategoryT = "юниоры, юниорки (17-25 лет)"  //
	AgeCategoryMW      AgeCategoryT = "мужчины, женщины"
)

func GetAgeCat() []AgeCategoryT {
	ageCat := []AgeCategoryT{
		AgeCategoryY15_23,
		AgeCategoryY19_20,
		AgeCategoryY15_18,
		AgeCategoryBG13_17,
		AgeCategoryBG13_15,
		AgeCategoryBG10_12,
		AgeCategoryY21_23,
		AgeCategoryY17_25,
		AgeCategoryMW,
	}
	return ageCat
}

type UserRole string

const (
	UserRoleGuest          UserRole = "guest"
	UserRoleChiefSecretary UserRole = "secretary"
	UserRoleCoach          UserRole = "coach"
	UserRoleSportsman      UserRole = "sportsman"
	UserRoleCompOrganizer  UserRole = "competition organizer"
	UserRoleTCampOrganizer UserRole = "training camp organizer"
)

func CompareSportsCategory(cat1, cat2 *SportsCategoryT) int {
	categories := map[SportsCategoryT]int{
		"":                   0, // undefined category
		SportsCategory3youth: 1,
		SportsCategory2youth: 2,
		SportsCategory1youth: 3,
		SportsCategory3:      4,
		SportsCategory2:      5,
		SportsCategory1:      6,
		SportsCategoryCMS:    7,
		SportsCategoryMS:     8,
		SportsCategoryMSIC:   9,
	}

	switch {
	case categories[*cat1] < categories[*cat2]:
		return -1
	case categories[*cat1] > categories[*cat2]:
		return 1
	default:
		return 0
	}
}

func ValidateAgeCategory(age int, cat *AgeCategoryT) bool {
	validateAge := func(minAge, maxAge int) bool {
		return age >= minAge && age <= maxAge
	}

	ageRanges := map[AgeCategoryT][2]int{
		AgeCategoryBG10_12: {10, 12},
		AgeCategoryBG13_17: {13, 17},
		AgeCategoryBG13_15: {13, 15},
		AgeCategoryY15_18:  {15, 18},
		AgeCategoryY15_23:  {15, 23},
		AgeCategoryY17_25:  {17, 25},
		AgeCategoryY19_20:  {19, 20},
		AgeCategoryY21_23:  {21, 23},
		AgeCategoryMW:      {18, 100},
	}

	ageRange, ok := ageRanges[*cat]
	if !ok {
		return false
	}

	return validateAge(ageRange[0], ageRange[1])
}

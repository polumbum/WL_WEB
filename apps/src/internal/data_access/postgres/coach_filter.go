package dataaccess

import (
	"strings"

	"gorm.io/gorm"
)

// условно: имя, фамилия и отчество
type CoachFilter struct {
	first  string
	second string
	third  string
}

func NewCoachFilter(fullname string) *CoachFilter {
	nameParts := strings.Fields(fullname)
	filter := &CoachFilter{}

	fields := []*string{&filter.first, &filter.second, &filter.third}

	for i, part := range nameParts {
		if i < len(fields) {
			*fields[i] = part
		}
	}

	return filter
}

func (f *CoachFilter) Apply(query *gorm.DB) *gorm.DB {
	if f.first != "" {
		query = query.Where("LOWER(surname) LIKE LOWER(?) OR LOWER(name) LIKE LOWER(?) OR LOWER(patronymic) LIKE LOWER(?)", "%"+f.first+"%", "%"+f.first+"%", "%"+f.first+"%")
	}
	if f.second != "" {
		query = query.Where("LOWER(surname) LIKE LOWER(?) OR LOWER(name) LIKE LOWER(?) OR LOWER(patronymic) LIKE LOWER(?)", "%"+f.second+"%", "%"+f.second+"%", "%"+f.second+"%")
	}
	if f.third != "" {
		query = query.Where("LOWER(surname) LIKE LOWER(?) OR LOWER(name) LIKE LOWER(?) OR LOWER(patronymic) LIKE LOWER(?)", "%"+f.third+"%", "%"+f.third+"%", "%"+f.third+"%")
	}
	return query
}

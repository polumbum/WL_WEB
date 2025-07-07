package dataaccess

import (
	"strings"

	"gorm.io/gorm"
)

// условно: имя, фамилия и отчество
type SportsmanFilter struct {
	first  string
	second string
	third  string
}

func NewSportsmanFilter(filterStr string) *SportsmanFilter {
	nameParts := strings.Fields(filterStr)
	filter := &SportsmanFilter{}

	fields := []*string{&filter.first, &filter.second, &filter.third}

	for i, part := range nameParts {
		if i < len(fields) {
			*fields[i] = part
		}
	}

	return filter
}

func (f *SportsmanFilter) Apply(query *gorm.DB) *gorm.DB {
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

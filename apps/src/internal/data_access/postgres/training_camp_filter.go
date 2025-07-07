package dataaccess

import (
	"strings"

	"gorm.io/gorm"
)

type TCampFilter struct {
	city string
}

func NewTCampFilter(filterStr string) *TCampFilter {
	parts := strings.Fields(filterStr)
	filter := &TCampFilter{}

	fields := []*string{&filter.city}

	for i, part := range parts {
		if i < len(fields) {
			*fields[i] = part
		}
	}

	return filter
}

func (f *TCampFilter) Apply(query *gorm.DB) *gorm.DB {
	if f.city != "" {
		query = query.Where("LOWER(city) LIKE LOWER(?)", "%"+f.city+"%")
	}

	return query
}

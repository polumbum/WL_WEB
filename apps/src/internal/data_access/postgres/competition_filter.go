package dataaccess

import (
	"strings"

	"gorm.io/gorm"
)

type CompFilter struct {
	name string
	city string
}

func NewCompFilter(filterStr string) *CompFilter {
	parts := strings.Fields(filterStr)
	filter := &CompFilter{}

	fields := []*string{&filter.name, &filter.city}

	for i, part := range parts {
		if i < len(fields) {
			*fields[i] = part
		}
	}

	return filter
}

func (f *CompFilter) Apply(query *gorm.DB) *gorm.DB {
	if f.name != "" {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+f.name+"%")
	}
	if f.city != "" {
		query = query.Where("LOWER(city) LIKE LOWER(?)", "%"+f.city+"%")
	}

	return query
}

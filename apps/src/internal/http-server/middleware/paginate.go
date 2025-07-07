package middleware

import (
	"context"
	"net/http"
	"strconv"
)

type KeyStringT string

const (
	PaginationKey KeyStringT = "pagination"
)

type Pagination struct {
	Page  int
	Batch int
}

func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		batchStr := r.URL.Query().Get("batch")

		page := 1
		batch := 10

		if pageStr != "" {
			if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
				page = p
			}
		}
		if batchStr != "" {
			if l, err := strconv.Atoi(batchStr); err == nil && l > 0 {
				batch = l
			}
		}

		pagination := Pagination{
			Page:  page,
			Batch: batch,
		}

		ctx := context.WithValue(r.Context(), PaginationKey, pagination)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

const (
	SortKey KeyStringT = "sort"
)

func Sort(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sortStr := r.URL.Query().Get("sort")
		sortParam := strings.Replace(sortStr, ".", " ", -1)

		fmt.Println("sort")
		fmt.Println(sortParam)

		ctx := context.WithValue(r.Context(), SortKey, sortParam)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

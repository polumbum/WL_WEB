package middleware

import (
	"context"
	"log"
	"net/http"
)

const (
	FNameFilterKey KeyStringT = "fullname"
	NameFilterKey  KeyStringT = "name"
	CityFilterKey  KeyStringT = "city"
)

func FNameFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filterStr := r.URL.Query().Get(string(FNameFilterKey))
		//sortParam := strings.Replace(sortStr, ".", " ", -1)

		log.Println(FNameFilterKey, filterStr)

		ctx := context.WithValue(r.Context(), FNameFilterKey, filterStr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CompFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get(string(NameFilterKey))
		city := r.URL.Query().Get(string(CityFilterKey))

		//log.Println(FNameFilterKey, filterStr)

		ctx := context.WithValue(r.Context(), NameFilterKey, name)
		ctx = context.WithValue(ctx, CityFilterKey, city)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TCampFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		city := r.URL.Query().Get(string(CityFilterKey))

		//log.Println(FNameFilterKey, filterStr)

		ctx := context.WithValue(r.Context(), CityFilterKey, city)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

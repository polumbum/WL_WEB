package middleware

import (
	"fmt"
	"net/http"
	http_errors "src/internal/http-server/errors"

	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
)

const SignJWTStr = "secret"

func JWT(requiredRole string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				render.Render(w, r, http_errors.ErrAuth)
				return
			}

			tokenString = tokenString[len("Bearer "):]

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(SignJWTStr), nil
			})

			if err != nil || !token.Valid {
				render.Render(w, r, http_errors.ErrUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				userRole := claims["role"].(string)

				if userRole != requiredRole && !(requiredRole == "" && userRole != "") {
					render.Render(w, r, http_errors.ErrUnauthorized)
					return
				}
				next.ServeHTTP(w, r)
			} else {
				render.Render(w, r, http_errors.ErrUnauthorized)
				//return
			}
		})
	}
}

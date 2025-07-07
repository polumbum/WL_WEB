package http_errors

import (
	"errors"
)

var (
	ErrSmFields = errors.New("missing required Sportsman fields")
)

var (
	ErrServer = &ErrResponse{HTTPStatusCode: 500, StatusText: "Internal Server Error"}

	ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found"}

	ErrAuth = &ErrResponse{HTTPStatusCode: 401, StatusText: "Authorization header missing"}

	ErrUnauthorized = &ErrResponse{HTTPStatusCode: 401, StatusText: "Invalid token"}
)

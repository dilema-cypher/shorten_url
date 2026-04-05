package utils

import (
	"errors"
	"net/http"

	customErrors "github.com/dilema-cypher/shorten_url/internal/errors"
)

func StatusCodeByError(err error) int {
	var invalidURLError customErrors.InvalidURLError
	var databaseError customErrors.DatabaseError
	var redisError customErrors.RedisError

	switch {
	case errors.As(err, &invalidURLError):
		return http.StatusBadRequest
	case errors.As(err, &databaseError):
		return http.StatusInternalServerError
	case errors.As(err, &redisError):
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}
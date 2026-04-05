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
	var urlLongError customErrors.URLLong

	switch {
	case errors.As(err, &invalidURLError):
		return http.StatusBadRequest
	case errors.As(err, &databaseError):
		return http.StatusInternalServerError
	case errors.As(err, &redisError):
		return http.StatusServiceUnavailable
	case errors.As(err, &urlLongError):
		return http.StatusRequestURITooLong
	default:
		return http.StatusInternalServerError
	}
}
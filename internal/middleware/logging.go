package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request:= fmt.Sprintf("%s %s", r.Method, r.URL.Path)
		slog.Info("message", "Request:", request)
		next.ServeHTTP(w, r)
	})
}
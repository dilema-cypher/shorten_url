package router

import (
	"net/http"

	"github.com/dilema-cypher/shorten_url/internal/handler"
	"github.com/dilema-cypher/shorten_url/internal/middleware"
	"github.com/dilema-cypher/shorten_url/pkg/config"
)

func V1Handler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("POST /shorten", http.HandlerFunc(handler.ShortenHandler))

	return mux
}

func SetupRouter(cfg config.Config) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", V1Handler()))

	limiter := middleware.NewRateLimiterFromConfig(cfg.RateLimitRequests, cfg.RateLimitWindow)

	mw := middleware.RecoveryMiddleware(mux)
	mw = middleware.LoggingMiddleware(mw)
	mw = middleware.RateLimitMiddleware(limiter)(mw)

	return mw
}
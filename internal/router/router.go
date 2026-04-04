package router

import (
	"encoding/json"
	"net/http"

	"github.com/dilema-cypher/shorten_url/internal/handler"
	"github.com/dilema-cypher/shorten_url/internal/middleware"
	"github.com/dilema-cypher/shorten_url/internal/service"
	"github.com/dilema-cypher/shorten_url/pkg/config"
)

func V1Handler() http.Handler {
	mux := http.NewServeMux()

	svc := service.NewShortenerService()
	h := handler.NewHandler(svc)

	mux.Handle("POST /shorten", http.HandlerFunc(h.Shorten))

	return mux
}

func SetupRouter(cfg config.Config) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "OK",
		})
	}))
	
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", V1Handler()))

	limiter := middleware.NewRateLimiterFromConfig(cfg.RateLimitRequests, cfg.RateLimitWindow)

	mw := middleware.RecoveryMiddleware(mux)
	mw = middleware.LoggingMiddleware(mw)
	mw = middleware.RateLimitMiddleware(limiter)(mw)
	mw = middleware.JSONAplicationMiddleware(mw)

	return mw
}
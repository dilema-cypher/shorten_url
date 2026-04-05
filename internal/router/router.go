package router

import (
	"encoding/json"
	"net/http"

	"github.com/dilema-cypher/shorten_url/internal/handler"
	"github.com/dilema-cypher/shorten_url/internal/middleware"
	"github.com/dilema-cypher/shorten_url/internal/redis"
	"github.com/dilema-cypher/shorten_url/internal/service"
	"github.com/dilema-cypher/shorten_url/pkg/config"
	"github.com/dilema-cypher/shorten_url/pkg/query/shorten_urls"
	"github.com/gocql/gocql"
)

func SetupRouter(cfg config.Config, redisClient *redis.RedisClient, session *gocql.Session) http.Handler {
	mux := http.NewServeMux()

	queries := shorten_urls.LoadQueries(cfg)
	svc := service.NewShortenerService(redisClient, session, queries, cfg.RedisSaltKey)
	h := handler.NewHandler(svc)

	limiter := middleware.NewRateLimiterFromConfig(cfg.RateLimitRequests, cfg.RateLimitWindow)

	mux.HandleFunc("POST /api/v1/shorten", h.Shorten)
	mux.HandleFunc("GET /api/v1/{id}", h.RedirectByUrl)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "OK",
		})
	})

	mw := middleware.RecoveryMiddleware(mux)
	mw = middleware.LoggingMiddleware(mw)
	mw = middleware.RateLimitMiddleware(limiter)(mw)
	mw = middleware.JSONAplicationMiddleware(mw)

	return mw
}
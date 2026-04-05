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

func V1Handler(redisClient *redis.RedisClient, session *gocql.Session, queries shorten_urls.Queries, saltKey string) http.Handler {
	mux := http.NewServeMux()

	svc := service.NewShortenerService(redisClient, session, queries, saltKey)
	h := handler.NewHandler(svc)

	mux.Handle("POST /shorten", http.HandlerFunc(h.Shorten))

	return mux
}

func SetupRouter(cfg config.Config, redisClient *redis.RedisClient, session *gocql.Session) http.Handler {
	mux := http.NewServeMux()

	queries := shorten_urls.LoadQueries(cfg)

	mux.Handle("GET /health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "OK",
		})
	}))

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", V1Handler(redisClient, session, queries, cfg.RedisSaltKey)))

	limiter := middleware.NewRateLimiterFromConfig(cfg.RateLimitRequests, cfg.RateLimitWindow)

	mw := middleware.RecoveryMiddleware(mux)
	mw = middleware.LoggingMiddleware(mw)
	mw = middleware.RateLimitMiddleware(limiter)(mw)
	mw = middleware.JSONAplicationMiddleware(mw)

	return mw
}
package router

import (
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
	svc := service.NewShortenerService(redisClient, session, queries, cfg.RedisSaltKey, cfg.URLApiRedirect)
	h := handler.NewHandler(svc)

	limiter := middleware.NewRateLimiterFromConfig(cfg.RateLimitRequests, cfg.RateLimitWindow)

	middleware.InitAuthMiddleware(cfg.UrlAuth)

	mux.HandleFunc("POST /api/v1/shorten", h.Shorten)
	mux.HandleFunc("GET /{id}", h.RedirectByUrl)
	mux.HandleFunc("GET /health", h.Health)

	mw := middleware.RecoveryMiddleware(mux)
	mw = middleware.LoggingMiddleware(mw)
	mw = middleware.RateLimitMiddleware(limiter)(mw)
	mw = middleware.JSONApplicationMiddleware(mw)
	mw = middleware.AuthMiddleware(mw)

	return mw
}
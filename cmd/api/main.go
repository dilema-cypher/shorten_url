package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dilema-cypher/shorten_url/internal/logger"
	"github.com/dilema-cypher/shorten_url/internal/router"
	"github.com/dilema-cypher/shorten_url/internal/redis"
	"github.com/dilema-cypher/shorten_url/pkg/config"
	"github.com/dilema-cypher/shorten_url/pkg/db"
)

func main() {
	cfg := config.LoadEnv()

	logger.Configure()

	session, err := db.ConnCassandra(cfg)
	if err != nil {
		slog.Error("Failed to connect to Cassandra:", "error", err)
		return
	}
	defer db.CloseSession(session)

	redisConfig := redis.RedisConfig{
		Host:         cfg.RedisHost,
		DB:           cfg.RedisDB,
		PoolSize:     cfg.RedisPoolSize,
		MinIdleConns: cfg.RedisMinIdleConns,
		ReadTimeout:  cfg.RedisReadTimeout,
		WriteTimeout: cfg.RedisWriteTimeout,
		DialTimeout:  cfg.RedisDialTimeout,
	}

	redisClient, err := redis.NewRedisClient(redisConfig)
	if err != nil {
		slog.Error("Failed to connect to Redis:", "error", err)
		return
	}
	defer redisClient.Close()

	r := router.SetupRouter(cfg, redisClient, session)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	slog.Info("Server started", "port", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		slog.Error("Failed to start server:", "error", err)
	}

}
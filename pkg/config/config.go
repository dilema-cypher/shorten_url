package config

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	CassandraHost     string
	CassandraUser     string
	CassandraPass     string
	CassandraKeyspace string
	RateLimitRequests int
	RateLimitWindow   int // em segundos
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

func LoadEnv() Config{
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file")
	}
	return Config{
		CassandraHost:     GetEnv("CASSANDRA_HOST", "localhost:9042"),
		CassandraUser:     GetEnv("CASSANDRA_USER", "cassandra"),
		CassandraPass:     GetEnv("CASSANDRA_PASSWORD", ""),
		CassandraKeyspace: GetEnv("CASSANDRA_KEYSPACE", "cassandra"),
		RateLimitRequests: GetEnvInt("RATE_LIMIT_REQUESTS", 60),
		RateLimitWindow:   GetEnvInt("RATE_LIMIT_WINDOW_SECONDS", 60),
	}
}
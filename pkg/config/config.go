package config

import (
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	CassandraHost     string
	CassandraUser     string
	CassandraPass     string
	CassandraKeyspace string
	RateLimitRequests int
	RateLimitWindow   int // em segundos
	RedisHost         string
	RedisDB           int
	RedisSaltKey      string
	RedisPoolSize     int
	RedisMinIdleConns int
	RedisReadTimeout  time.Duration
	RedisWriteTimeout time.Duration
	RedisDialTimeout  time.Duration
	URLApiRedirect    string
	ServerPort        string
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

func GetEnvDuration(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		return defaultValue
	}
	return duration
}

func LoadEnv() Config {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file")
	}
	return Config{
		CassandraHost:      GetEnv("CASSANDRA_HOST", "localhost:9042"),
		CassandraUser:      GetEnv("CASSANDRA_USER", "cassandra"),
		CassandraPass:      GetEnv("CASSANDRA_PASSWORD", ""),
		CassandraKeyspace:  GetEnv("CASSANDRA_KEYSPACE", "cassandra"),
		RateLimitRequests:  GetEnvInt("RATE_LIMIT_REQUESTS", 60),
		RateLimitWindow:    GetEnvInt("RATE_LIMIT_WINDOW_SECONDS", 60),
		RedisHost:          GetEnv("REDIS_HOST", "localhost:6379"),
		RedisDB:            GetEnvInt("REDIS_DB", 1),
		RedisSaltKey:       GetEnv("REDIS_SALT_KEY", "salt_shorten_url"),
		RedisPoolSize:      GetEnvInt("REDIS_POOL_SIZE", 10),
		RedisMinIdleConns:  GetEnvInt("REDIS_MIN_IDLE_CONNS", 2),
		RedisReadTimeout:   GetEnvDuration("REDIS_READ_TIMEOUT", 3*time.Second),
		RedisWriteTimeout:  GetEnvDuration("REDIS_WRITE_TIMEOUT", 3*time.Second),
		RedisDialTimeout:   GetEnvDuration("REDIS_DIAL_TIMEOUT", 5*time.Second),
		URLApiRedirect:     GetEnv("URL_API_REDIRECT", "http://localhost:8080"),
		ServerPort:         GetEnv("SERVER_PORT", "8080"),
	}
}
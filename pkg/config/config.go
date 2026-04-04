package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	CassandraHost    string
	CassandraUser   string
	CassandraPass   string
	CassandraKeyspace string
}

func GetEnv(key string, defaultValue string) string{
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func LoadEnv() Config{
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file")
	}
	return Config{
		CassandraHost:    GetEnv("CASSANDRA_HOST", "localhost:9042"),
		CassandraUser:   GetEnv("CASSANDRA_USER", "cassandra"),
		CassandraPass:   GetEnv("CASSANDRA_PASSWORD", ""),
		CassandraKeyspace: GetEnv("CASSANDRA_KEYSPACE", "cassandra"),
	}
}
package repository

import (
	"context"
	"log/slog"

	"github.com/dilema-cypher/shorten_url/internal/redis"
	"github.com/dilema-cypher/shorten_url/internal/utils"
	"github.com/dilema-cypher/shorten_url/pkg/query/shorten_urls"
	"github.com/gocql/gocql"
)

type ShortenerRepository interface {
	IncrRedis(ctx context.Context, key string) (int64, error)
	CreateShortenURL(ctx context.Context, id string, url string) error
}

type shortenerRepository struct {
	cliRedis *redis.RedisClient
	session  *gocql.Session
	queries  shorten_urls.Queries
}

func NewShortenerRepository(client *redis.RedisClient, session *gocql.Session, queries shorten_urls.Queries) ShortenerRepository {
	return &shortenerRepository{
		cliRedis: client,
		session:  session,
		queries:  queries,
	}
}

func (s *shortenerRepository) IncrRedis(ctx context.Context, key string) (int64, error) {
	id, err := s.cliRedis.Incr(ctx, key)
	if err != nil {
		slog.Error("IncrRedis error", "error", err)
		return 0, err
	}

	return id, nil
}

func (s *shortenerRepository) CreateShortenURL(ctx context.Context, id string, url string) error {
	err := s.session.Query(s.queries.Insert).Bind(id, url, utils.TimeNow()).Exec()
	if err != nil {
		slog.Error("CreateShortenURL error", "error", err)
		return err
	}

	return nil
}
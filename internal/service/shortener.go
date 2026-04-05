package service

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"strings"

	"github.com/dilema-cypher/shorten_url/internal/errors"
	"github.com/dilema-cypher/shorten_url/internal/redis"
	"github.com/dilema-cypher/shorten_url/internal/repository"
	"github.com/dilema-cypher/shorten_url/internal/utils"
	"github.com/dilema-cypher/shorten_url/pkg/query/shorten_urls"
	"github.com/gocql/gocql"
)

type ShortenerService interface {
	Shorten(ctx context.Context, url string) (string, error)
	Get(ctx context.Context, id string) (string, error)
}

type shortenerService struct {
	repo    repository.ShortenerRepository
	saltKey string
}

func NewShortenerService(redisClient *redis.RedisClient, session *gocql.Session, queries shorten_urls.Queries, saltKey string) ShortenerService {
	repo := repository.NewShortenerRepository(redisClient, session, queries)
	return &shortenerService{repo: repo, saltKey: saltKey}
}

func (s *shortenerService) Shorten(ctx context.Context, urlStr string) (string, error) {
	if strings.TrimSpace(urlStr) == "" {
		return "", errors.InvalidURLError{Message: "url is required"}
	}
	if _, err := url.ParseRequestURI(urlStr); err != nil {
		return "", errors.InvalidURLError{Message: "invalid url format"}
	}

	id, err := s.repo.IncrRedis(ctx, s.saltKey)
	if err != nil {
		slog.Error("Error incrementing Redis", "error", err)
		return "", errors.RedisError{Message: "failed to increment counter", Err: err}
	}

	urlWithID := fmt.Sprintf("%s%d", urlStr, id)
	idBase62 := utils.StringToBase62(urlWithID)

	err = s.repo.CreateShortenURL(ctx, idBase62, urlStr)
	if err != nil {
		slog.Error("Error creating short URL in Cassandra", "error", err)
		return "", errors.DatabaseError{Message: "failed to save URL", Err: err}
	}

	shortURL := fmt.Sprintf("%s/%s", "http://localhost:8080/api/v1/", idBase62)

	return shortURL, nil
}

func (s *shortenerService) Get(ctx context.Context, id string) (string, error) {
	if strings.TrimSpace(id) == "" {
		return "", errors.InvalidURLError{Message: "id is required"}
	}

	url, err := s.repo.GetURL(ctx, id)
	if err != nil {
		slog.Error("Error getting URL from Cassandra", "error", err)
		return "", errors.DatabaseError{Message: "failed to get URL", Err: err}
	}

	return url, nil
}
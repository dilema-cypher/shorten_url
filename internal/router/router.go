package router

import (
	"net/http"

	"github.com/dilema-cypher/shorten_url/internal/handler"
)

func V1Handler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("POST /shorten", http.HandlerFunc(handler.ShortenHandler))

	return mux
}

func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/api/v1", http.StripPrefix("/api/v1", V1Handler()))

	return mux
}
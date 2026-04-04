package main

import (
	"log/slog"
	"net/http"

	"github.com/dilema-cypher/shorten_url/internal/router"
	"github.com/dilema-cypher/shorten_url/pkg/config"
	"github.com/dilema-cypher/shorten_url/pkg/db"
)

func main() {
	cfg := config.LoadEnv()

	session, err := db.ConnCassandra(cfg)
	if err != nil {
		slog.Error("Failed to connect to Cassandra:", "error", err)
		return
	}
	
	defer db.CloseSession(session)

	r := router.SetupRouter(cfg)

	slog.Info("Server started on :8080")
	
	if err := http.ListenAndServe(":8080", r); err != nil {
		slog.Error("Failed to start server:", "error", err)
	}
	
}
package shorten_urls

import (
	"embed"
	"strings"

	"github.com/dilema-cypher/shorten_url/pkg/config"
)

//go:embed *.cql
var cqlFiles embed.FS

type Queries struct {
	Insert string
	Get    string
}

func LoadQueries(cfg config.Config) Queries {
	keyspace := cfg.CassandraKeyspace

	queries := Queries{}

	if data, err := cqlFiles.ReadFile("insert.cql"); err == nil {
		queries.Insert = strings.ReplaceAll(string(data), "%%KEYSPACE%%", keyspace)
	}

	if data, err := cqlFiles.ReadFile("get.cql"); err == nil {
		queries.Get = strings.ReplaceAll(string(data), "%%KEYSPACE%%", keyspace)
	}

	return queries
}
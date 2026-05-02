package db

import (
	"log/slog"
	"time"

	"github.com/dilema-cypher/shorten_url/pkg/config"
	"github.com/gocql/gocql"
)

func ConnCassandra(cfg config.Config) (*gocql.Session, error) {

	cluster := gocql.NewCluster(cfg.CassandraHost)
	cluster.Keyspace = cfg.CassandraKeyspace
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = 5 * time.Second
	cluster.ConnectTimeout = 5 * time.Second
	cluster.NumConns = 3
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cfg.CassandraUser,
		Password: cfg.CassandraPass,
	}

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	if err := PingCassandra(session); err != nil {
		return nil, err
	}

	slog.Info("Cassandra connected successfully")

	return session, nil
}

func CloseSession(session *gocql.Session) {
	session.Close()
	slog.Info("cassandra session closed")
}

func PingCassandra(session *gocql.Session) error {
	if err := session.Query("SELECT now() FROM system.local;").Exec(); err != nil {
		return err
	} else {
		return nil
	}
}
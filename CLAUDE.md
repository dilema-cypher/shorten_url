# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go-based URL shortener service using Cassandra as the database.

## Common Commands

- `go run cmd/api/main.go` - Run the API server
- `go build ./...` - Build all packages
- `go test ./...` - Run all tests
- `go test ./... -run TestName` - Run a specific test
- `go mod tidy` - Clean up go.mod dependencies

## Architecture

```
cmd/api/main.go     - Application entry point
pkg/config/config.go - Configuration from .env file
pkg/db/connect.go   - Database connection (Cassandra)
.env                - Environment variables (contains CASSANDRA_URL)
```

The project uses a standard Go workspace layout with `cmd/` for binaries and `pkg/` for reusable packages.

## Configuration

All configuration is loaded from `.env` file using `godotenv`. The `Config` struct in `pkg/config/config.go` provides:
- `CassandraURL` - Cassandra database connection string (default: `localhost:9042`)

To change the database, update the `CASSANDRA_URL` in `.env`.
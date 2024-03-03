package store

import (
	"GO_Redis/internal/config"
	"database/sql"
)

type PostgresStore struct {
	DB *sql.DB
}

func NewStore(cfg config.Postgres) (*PostgresStore, error) {
	// connect to store

	return &PostgresStore{}, nil
}

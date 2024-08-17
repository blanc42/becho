package db

import (
	"github.com/jackc/pgx/v5"
)

type DbStore struct {
	*Queries
	db *pgx.Conn
}

func NewDbStore(db *pgx.Conn) *DbStore {
	return &DbStore{
		db:      db,
		Queries: New(db),
	}
}

package store

import (
	"encore.dev/storage/sqldb"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db  *sqldb.Database
	dbx *sqlx.DB
}

func NewStore(db *sqldb.Database, dbx *sqlx.DB) *Store {
	if db == nil {
		panic("database is nil")
	}
	if dbx == nil {
		panic("dbx is nil")
	}
	return &Store{db: db, dbx: dbx}
}

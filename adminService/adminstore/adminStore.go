package adminstore

import "github.com/jmoiron/sqlx"

type AdminStore struct {
	db *sqlx.DB
}

func NewAdminStore(db *sqlx.DB) *AdminStore {
	return &AdminStore{db: db}
}

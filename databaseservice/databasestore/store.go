package databasestore

import "encore.dev/storage/sqldb"

type DatabaseStore struct {
	db *sqldb.Database
}

func NewDatabaseStore(db *sqldb.Database) *DatabaseStore {
	if db == nil {
		panic("database is nil")
	}
	return &DatabaseStore{db: db}
}

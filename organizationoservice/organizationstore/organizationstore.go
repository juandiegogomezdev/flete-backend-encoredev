package organizationstore

import (
	"encore.dev/storage/sqldb"
	"github.com/jmoiron/sqlx"
)

type OrganizationStore struct {
	db  *sqldb.Database
	dbx *sqlx.DB
}

func NewOrganizationStore(db *sqldb.Database, dbx *sqlx.DB) *OrganizationStore {
	return &OrganizationStore{db: db, dbx: dbx}
}

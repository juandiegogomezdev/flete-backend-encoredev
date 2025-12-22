package organizationstore

import "encore.dev/storage/sqldb"

type OrganizationStore struct {
	db *sqldb.Database
}

func NewOrganizationStore(db *sqldb.Database) *OrganizationStore {
	return &OrganizationStore{db: db}
}

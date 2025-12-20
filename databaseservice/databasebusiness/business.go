package databasebusiness

import (
	"encore.app/databaseservice/databasestore"
)

type BusinessDatabase struct {
	store *databasestore.DatabaseStore
}

func NewDatabaseBusiness(store *databasestore.DatabaseStore) *BusinessDatabase {
	return &BusinessDatabase{store: store}
}

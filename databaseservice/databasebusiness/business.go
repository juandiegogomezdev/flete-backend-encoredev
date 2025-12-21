package databasebusiness

import (
	"encore.app/databaseservice/databasestore"
)

type DatabaseBusiness struct {
	store *databasestore.DatabaseStore
}

func NewDatabaseBusiness(store *databasestore.DatabaseStore) *DatabaseBusiness {
	return &DatabaseBusiness{store: store}
}

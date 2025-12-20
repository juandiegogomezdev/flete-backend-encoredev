package databaseservice

import (
	"encore.app/databaseservice/databasebusiness"
	"encore.app/databaseservice/databasestore"
	"encore.dev/storage/sqldb"
)

var primaryDB = sqldb.NewDatabase("primary_db", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

//encore:service
type DatabaseService struct {
	b *databasebusiness.BusinessDatabase
}

func initDatabaseService() (*DatabaseService, error) {
	s := databasestore.NewDatabaseStore(primaryDB)
	b := databasebusiness.NewDatabaseBusiness(s)
	return &DatabaseService{b: b}, nil
}

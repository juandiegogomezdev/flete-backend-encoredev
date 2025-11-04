package adminservice

import (
	"context"

	"encore.app/adminService/adminBusiness"
	"encore.app/adminService/adminstore"
	"encore.app/adminService/seeders"
	"encore.dev/storage/sqldb"
	"github.com/jmoiron/sqlx"
)

var fleteDB = sqldb.Named("db_app")
var fleteDBX = sqlx.NewDb(fleteDB.Stdlib(), "postgres")

//encore:service
type adminService struct {
	b *adminBusiness.AdminBusiness
}

func initAdminService() (*adminService, error) {
	fleteDBX := sqlx.NewDb(fleteDB.Stdlib(), "postgres")

	store := adminstore.NewAdminStore(fleteDBX)
	business := adminBusiness.NewAdminBusiness(store)

	return &adminService{
		b: business,
	}, nil
}

//encore:api private method=POST path=/admin/seed
func (s *adminService) SeedDatabase(ctx context.Context) (ResponseSeedDatabase, error) {
	return seeders.RunSeed(s.db)
}

type ResponseSeedDatabase struct {
	Message string `json:"message"`
}

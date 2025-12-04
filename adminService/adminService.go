package adminservice

import (
	"context"

	"encore.app/adminService/adminBusiness"
	"encore.app/adminService/adminstore"
	"encore.dev/storage/sqldb"
	"github.com/jmoiron/sqlx"
)

var fleteDB = sqldb.Named("db_app")
var fleteDBX = sqlx.NewDb(fleteDB.Stdlib(), "postgres")

//encore:service
type AdminService struct {
	b *adminBusiness.AdminBusiness
}

func initAdminService() (*AdminService, error) {

	store := adminstore.NewAdminStore(fleteDBX)
	business := adminBusiness.NewAdminBusiness(store)

	return &AdminService{
		b: business,
	}, nil
}

//encore:api private method=POST path=/admin/seed
func (s *AdminService) SeedDatabase(ctx context.Context) error {
	return s.b.SeedNotificationTemplates(ctx)
}

// type ResponseSeedDatabase struct {
// 	Message string `json:"message"`
// }

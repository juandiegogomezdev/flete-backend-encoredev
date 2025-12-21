package adminservice

import (
	"context"

	"encore.app/adminService/adminBusiness"
	"encore.app/adminService/adminstore"
	"encore.dev/storage/sqldb"
	"github.com/jmoiron/sqlx"
)

var primaryDB = sqldb.Named("primary_db")
var primaryDBX = sqlx.NewDb(primaryDB.Stdlib(), "postgres")

//encore:service
type AdminService struct {
	b *adminBusiness.AdminBusiness
}

func initAdminService() (*AdminService, error) {

	store := adminstore.NewAdminStore(primaryDBX)
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

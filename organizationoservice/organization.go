package organizationoservice

import (
	"encore.app/organizationoservice/organizationbusiness"
	"encore.app/organizationoservice/organizationstore"
	"encore.dev/storage/sqldb"
	"github.com/jmoiron/sqlx"
)

var primaryDB = sqldb.Named("primary_db")

var primaryDBX = sqlx.NewDb(primaryDB.Stdlib(), "postgres")

//encore:service
type OrganizationService struct {
	b *organizationbusiness.OrganizationBusiness
}

func initOrganizationService() (*OrganizationService, error) {
	store := organizationstore.NewOrganizationStore(primaryDB, primaryDBX)
	business := organizationbusiness.NewOrganizationBusiness(store)
	return &OrganizationService{b: business}, nil
}

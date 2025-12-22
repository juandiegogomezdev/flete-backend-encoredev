package organizationoservice

import (
	"context"
	"fmt"

	"encore.app/organizationoservice/organizationbusiness"
	"encore.app/organizationoservice/organizationstore"
	"encore.dev/beta/auth"
	"encore.dev/storage/sqldb"
)

var primaryDB = sqldb.Named("primary_db")

//encore:service
type OrganizationService struct {
	b *organizationbusiness.OrganizationBusiness
}

func initOrganizationService() *OrganizationService {
	store := organizationstore.NewOrganizationStore(primaryDB)
	business := organizationbusiness.NewOrganizationBusiness(store)
	return &OrganizationService{b: business}
}

//encore:api public method=POST path=/create-organization
func (s *OrganizationService) CreateOrganization(ctx context.Context) error {
	userId := auth.UserID()
	fmt.Println("userId", userId)
	return nil
}

package organizationbusiness

import (
	"encore.app/organizationoservice/organizationstore"
)

type OrganizationBusiness struct {
	s *organizationstore.OrganizationStore
}

func NewOrganizationBusiness(store *organizationstore.OrganizationStore) *OrganizationBusiness {
	return &OrganizationBusiness{s: store}
}

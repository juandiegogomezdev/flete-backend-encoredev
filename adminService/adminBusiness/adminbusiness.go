package adminBusiness

import (
	"context"

	"encore.app/adminService/adminseed"
	"encore.app/adminService/adminstore"
	"encore.dev/beta/errs"
)

type AdminBusiness struct {
	s *adminstore.AdminStore
}

func NewAdminBusiness(s *adminstore.AdminStore) *AdminBusiness {
	if s == nil {
		panic("adminStore cannot be nil in NewAdminBusiness")
	}
	return &AdminBusiness{s: s}
}

func (b *AdminBusiness) SeedNotificationTemplates(ctx context.Context) error {
	// import data
	notificationTemplates := adminseed.SeedDataNotificationTemplates()

	err := b.s.SeedNotificationTemplates(ctx, notificationTemplates)
	if err != nil {
		return &errs.Error{
			Code:    errs.Internal,
			Message: `failed to seed notification templates: ` + err.Error(),
		}
	}

	return nil

}

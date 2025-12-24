package organizationbusiness

import (
	"context"

	"encore.dev/beta/errs"
	"encore.dev/types/uuid"
)

func (b *OrganizationBusiness) SaveOrganizationLogoKey(ctx context.Context, orgId uuid.UUID, logoKey string) error {
	err := b.s.SaveLogoKeyOrganization(ctx, orgId, logoKey)
	if err != nil {
		return &errs.Error{
			Message: "Error al guardar el logo de la organización",
			Code:    errs.Internal,
		}
	}
	return nil
}

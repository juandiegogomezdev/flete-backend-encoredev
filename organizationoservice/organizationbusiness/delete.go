package organizationbusiness

import (
	"context"

	"encore.dev/beta/errs"
	"encore.dev/types/uuid"
)

func (b *OrganizationBusiness) DeleteOrganizationLogo(ctx context.Context, orgID uuid.UUID) error {
	err := b.s.DeleteOrganizationLogo(ctx, orgID)
	if err != nil {
		return &errs.Error{
			Message: "Error al eliminar el logo de la organización",
			Code:    errs.Internal,
		}
	}
	return nil

}

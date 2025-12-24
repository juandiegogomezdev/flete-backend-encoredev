package organizationbusiness

import (
	"context"

	"encore.app/databaseservice/models"
	"encore.dev/beta/errs"
	"encore.dev/types/uuid"
)

func (b *OrganizationBusiness) GetOrganizationByMembershipId(ctx context.Context, memID uuid.UUID) (models.Organization, error) {
	org, err := b.s.GetOrganizationByMembershipId(ctx, memID)
	if err != nil {
		return models.Organization{}, &errs.Error{
			Message: "Error al encontrar la organización.",
			Code:    errs.Internal,
		}
	}
	return org, nil
}

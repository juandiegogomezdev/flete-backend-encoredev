package organizationbusiness

import (
	"context"

	"encore.app/databaseservice/models"
	"encore.dev/beta/errs"
	"encore.dev/types/uuid"
)

func (b *OrganizationBusiness) GetOrganizationById(ctx context.Context, orgID uuid.UUID) (*models.Organization, error) {
	org, err := b.s.GetOrganizationByID(ctx, orgID)
	if err != nil {
		return nil, &errs.Error{
			Message: "Error al encontrar la organización.",
			Code:    errs.Internal,
		}
	}
	return &org, nil
}

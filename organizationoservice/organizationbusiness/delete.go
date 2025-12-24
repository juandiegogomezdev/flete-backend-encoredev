package organizationbusiness

import (
	"context"

	"encore.app/databaseservice/models"
	"encore.dev/beta/errs"
	"encore.dev/types/uuid"
)

func (b *OrganizationBusiness) DeleteOrganization(ctx context.Context, orgID uuid.UUID) error {

	// Delete organization
	err := b.s.DeleteOrganization(ctx, orgID)
	if err != nil {
		return &errs.Error{
			Message: "Error al eliminar la organización",
			Code:    errs.Internal,
		}
	}

	return nil
}

func (b *OrganizationBusiness) DeleteOrganizationLogoInDatabase(ctx context.Context, org *models.Organization) error {

	// Delete logo
	err := b.s.DeleteOrganizationLogo(ctx, org.ID)
	if err != nil {
		return &errs.Error{
			Message: "Error al eliminar el logo de la organización",
			Code:    errs.Internal,
		}
	}
	return nil

}

func (b *OrganizationBusiness) DeleteOrganizationManifestFiles(ctx context.Context, orgID uuid.UUID) error {
	err := b.s.DeleteOrganizationManifestFiles(ctx, orgID)
	if err != nil {
		return &errs.Error{
			Message: "Error al eliminar los archivos de manifiesto de la organización",
			Code:    errs.Internal,
		}
	}
	return nil
}

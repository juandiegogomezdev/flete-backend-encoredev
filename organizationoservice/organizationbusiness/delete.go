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

func (b *OrganizationBusiness) DeleteOrganizationManifestFiles(ctx context.Context, orgID uuid.UUID) ([]*models.File, error) {
	// Get all manifest files
	files, err := b.s.GetOrganizationFiles(ctx, orgID, "manifest")
	if err != nil {
		return nil, &errs.Error{
			Message: "Error al buscar los manifiestos de carga",
			Code:    errs.Internal,
		}
	}

	// Extract IDs
	var ids []uuid.UUID
	for _, file := range *files {
		ids = append(ids, file.ID)
	}

	// Delete files
	err = b.s.DeleteFilesByIDs(ctx, ids)
	if err != nil {
		return nil, &errs.Error{
			Message: "Error al borrar los manifiestos de carga",
			Code:    errs.Internal,
		}
	}

	return files, nil
}

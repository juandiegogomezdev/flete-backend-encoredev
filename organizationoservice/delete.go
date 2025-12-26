package organizationoservice

import (
	"context"
	"fmt"

	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
	"encore.dev/types/uuid"
)

//encore:api auth method=DELETE path=/delete-organization/{orgID}
func (s *OrganizationService) DeleteOrganization(ctx context.Context, p *DeleteOrganizationRequest) error {
	// Get user ID and OrgID
	userID, _ := auth.UserID()

	// Only owner can delete organization
	org, err := s.b.GetOrganizationById(ctx, p.OrgID)
	if err != nil {
		return err
	}
	if org.OwnerID != string(userID) {
		return &errs.Error{
			Message: "Solo el dueño de la empresa puede eliminar la organización",
			Code:    errs.PermissionDenied,
		}
	}

	// Delete organization
	err = s.b.DeleteOrganization(ctx, org.ID)
	if err != nil {
		return err
	}

	return nil

}

//encore:api auth method=DELETE path=/delete-organization-logo
func (s *OrganizationService) DeleteOrganizationLogo(ctx context.Context, p *DeleteOrganizationLogoRequest) error {
	userID, _ := auth.UserID()

	// Get organization

	org, err := s.b.GetOrganizationById(ctx, p.OrgID)
	if err != nil {
		return err
	}

	// Check if the member is the owner of the organization
	if org.OwnerID != string(userID) {
		return &errs.Error{
			Message: "Solo el dueño de la empresa puede eliminar el logo",
			Code:    errs.PermissionDenied,
		}
	}

	err = s.b.DeleteOrganizationLogoInDatabase(ctx, p.OrgID)
	if err != nil {
		return err
	}

	//TODO:Delete with bucket

	return nil
}

type DeleteOrganizationLogoRequest struct {
	OrgID uuid.UUID `query:"org_id"`
}

type DeleteOrganizationRequest struct {
	OrgID uuid.UUID `query:"org_id"`
}

//encore:api auth method=DELETE path=/delete-organization-manifest-files
func (s *OrganizationService) DeleteOrganizationManifestFiles(ctx context.Context, p *DeleteOrganizationMovementFilesRequest) error {
	userID, _ := auth.UserID()

	// Get organization
	org, err := s.b.GetOrganizationById(ctx, p.OrgID)

	if org.OwnerID != string(userID) {
		return &errs.Error{
			Message: "Solo el dueño de la empresa puede eliminar la organización",
			Code:    errs.PermissionDenied,
		}
	}

	// Delete manifests files
	files, err := s.b.DeleteOrganizationManifestFiles(ctx, p.OrgID)
	if err != nil {
		return err
	}

	fmt.Println("files", files)

	return nil
}

type DeleteOrganizationManifestFilesRequest struct {
	OrgID uuid.UUID `query:"org_id"`
}

//encore:api auth method=DELETE path=/delete-organization-movement-files
func (s *OrganizationService) DeleteOrganizationMovementFiles(ctx context.Context, p *DeleteOrganizationMovementFilesRequest) error {
	userID, _ := auth.UserID()

	// Get organization
	org, err := s.b.GetOrganizationById(ctx, p.OrgID)

	if org.OwnerID != string(userID) {
		return &errs.Error{
			Message: "Solo el dueño de la empresa puede eliminar la organización",
			Code:    errs.PermissionDenied,
		}
	}

	// Delete movements files
	deletedFiles, err := s.b.DeleteOrganizationMovementFiles(ctx, org.ID)
	if err != nil {
		return err
	}

	fmt.Println("Deleted movements files", deletedFiles)
	return nil
}

type DeleteOrganizationMovementFilesRequest struct {
	OrgID uuid.UUID `query:"org_id"`
}

//encore:api auth method=DELETE path=/delete-organization-documents
func (s *OrganizationService) DeleteOrganizationDocuments(ctx context.Context, p *DeleteOrganizationDocumentsRequest) error {
	userID, _ := auth.UserID()

	// Get organization
	org, err := s.b.GetOrganizationById(ctx, p.OrgID)

	if org.OwnerID != string(userID) {
		return &errs.Error{
			Message: "Solo el dueño de la empresa puede eliminar la organización",
			Code:    errs.PermissionDenied,
		}
	}

	// Delete documents files
	deletedFiles, err := s.b.DeleteOrganizationDocumentFiles(ctx, p.OrgID)
	if err != nil {
		return err
	}
	fmt.Println("Delted documents", deletedFiles)
	return err
}

type DeleteOrganizationDocumentsRequest struct {
	OrgID uuid.UUID `query:"org_id"`
}

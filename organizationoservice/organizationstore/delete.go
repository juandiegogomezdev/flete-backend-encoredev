package organizationstore

import (
	"context"
	"log"

	"encore.dev/types/uuid"
)

func (s *OrganizationStore) DeleteOrganization(ctx context.Context, orgId uuid.UUID) error {
	// Start transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		log.Println("Error starting transaction: ", err)
		return err
	}
	defer tx.Rollback()

	// Suspend all memberships
	q := `UPDATE memberships SET status = 'suspended' WHERE org_id = $1`
	_, err = tx.Exec(ctx, q, orgId)

	if err != nil {
		log.Println("Error suspending memberships: ", err)
		return err
	}

	// Suspend all invitations
	q = `UPDATE job_invitations SET status = 'revoked' WHERE org_id = $1`
	_, err = tx.Exec(ctx, q, orgId)
	if err != nil {
		log.Println("Error suspending invitations: ", err)
		return err
	}

	// Suspend organization
	q = `UPDATE organizations SET active = FALSE WHERE id = $1`
	_, err = tx.Exec(ctx, q, orgId)
	if err != nil {
		log.Println("Error suspending organization: ", err)
		return err
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		log.Println("Error committing transaction: ", err)
		return err
	}

	// Check if organization
	return nil
}

func (s *OrganizationStore) DeleteOrganizationLogo(ctx context.Context, orgID uuid.UUID) error {
	q := `UPDATE organizations SET image_key = NULL WHERE id = $1`
	_, err := s.db.Exec(ctx, q, orgID)
	if err != nil {
		log.Println("Error deleting organization: ", err)
		return err
	}
	return nil
}

// func (s *OrganizationStore) DeleteOrganizationManifestFiles(ctx context.Context, orgID uuid.UUID) error {
// 	q := `DELETE FROM organization_manifest_files WHERE org_id = $1`
// 	_, err := s.db.Exec(ctx, q, orgID)
// 	if err != nil {
// 		log.Println("Error deleting organization manifest files: ", err)
// 		return err
// 	}
// 	return nil
// }

func (s *OrganizationStore) DeleteFilesByIDs(ctx context.Context, ids []uuid.UUID) error {
	q := `DELETE FROM files WHERE ID = ANY($1)`
	_, err := s.db.Exec(ctx, q, ids)
	if err != nil {
		log.Println("Error deleting files by IDs: ", err)
		return err
	}
	return nil
}

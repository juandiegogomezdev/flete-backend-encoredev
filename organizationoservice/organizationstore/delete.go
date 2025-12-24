package organizationstore

import (
	"context"
	"log"

	"encore.dev/types/uuid"
)

func (s *OrganizationStore) DeleteOrganizationLogo(ctx context.Context, orgID uuid.UUID) error {
	q := `UPDATE organizations SET image_key = NULL WHERE id = $1`
	_, err := s.db.Exec(ctx, q, orgID)
	if err != nil {
		log.Println("Error deleting organization: ", err)
		return err
	}
	return nil
}

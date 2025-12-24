package organizationstore

import (
	"context"
	"fmt"
	"log"

	"encore.dev/types/uuid"
)

func (s *OrganizationStore) SaveLogoKeyOrganization(ctx context.Context, orgId uuid.UUID, logoKey string) error {
	q := `UPDATE organizations SET image_key = $1 WHERE id = $2`
	_, err := s.db.Exec(ctx, q, logoKey, orgId)
	if err != nil {
		log.Println("Error updating organization logo key: ", err)
		return fmt.Errorf("error updating organization logo key: %w", err)
	}
	return err
}

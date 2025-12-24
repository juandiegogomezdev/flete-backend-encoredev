package organizationstore

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"encore.app/databaseservice/models"
	"encore.dev/beta/auth"
	"encore.dev/storage/sqldb"
	"encore.dev/types/uuid"
)

func (s *OrganizationStore) IsAvailableOrganizationName(ctx context.Context, name string, userId auth.UID) (bool, error) {
	q := `SELECT COUNT(1) FROM organizations WHERE name = $1, owner_id = $2`
	var count int
	err := s.db.QueryRow(ctx, q, name, userId).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking organization name availability: %w", err)
	}
	return count == 0, nil
}

func (s *OrganizationStore) GetOrganizationByMembershipId(ctx context.Context, memID uuid.UUID) (models.Organization, error) {
	// Query to get organization by membership ID
	q := `SELECT o.* FROM organizations o
		  JOIN memberships m ON o.id = m.organization_id
		  WHERE m.id = $1 LIMIT 1`
	// q := `SELECT * FROM organization WHERE id = $1 LIMIT 1`
	var org models.Organization
	err := s.dbx.GetContext(ctx, &org, q, memID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("error: organization not found for membership id", memID)
			return models.Organization{}, fmt.Errorf("membership %s not found", memID)
		}
		log.Println("error querying organization id for membership id", memID, ":", err)
		return models.Organization{}, fmt.Errorf("error querying organization id: %w", err)
	}
	return org, nil
}

func (s *OrganizationStore) GetRoleIdByName(ctx context.Context, roleName string) (uuid.UUID, error) {
	// Query to get role ID by name
	q := `SELECT id FROM roles WHERE name = $1 LIMIT 1`
	var roleId uuid.UUID
	err := s.db.QueryRow(ctx, q, roleName).Scan(&roleId)
	if err != nil {
		if err == sqldb.ErrNoRows {
			return uuid.Nil, fmt.Errorf("role %s not found", roleName)
		}
		return uuid.Nil, fmt.Errorf("error querying role id: %w", err)
	}
	return roleId, nil

}

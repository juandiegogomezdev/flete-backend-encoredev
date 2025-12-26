package organizationstore

import (
	"context"
	"fmt"
	"time"

	"encore.app/organizationoservice/shared/types"
	"encore.app/pkg/utils"
	"encore.dev/beta/auth"
	"encore.dev/types/uuid"
)

func (s *OrganizationStore) CreateOrganizationAndMembership(ctx context.Context, name string, ownerId auth.UID, roleID uuid.UUID) (types.Membership, error) {
	// Begin transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return types.Membership{}, err
	}
	defer tx.Rollback()

	// Insert organization
	orgID, err := utils.GenerateUUID()
	if err != nil {
		return types.Membership{}, fmt.Errorf("error generating organization ID: %w", err)
	}

	q := `INSERT INTO organizations (id, name, owner_id) VALUES ($1, $2, $3) RETURNING id`
	var orgId uuid.UUID
	err = tx.QueryRow(ctx, q, orgID, name, ownerId).Scan(&orgId)
	if err != nil {
		return types.Membership{}, fmt.Errorf("error inserting organization: %w", err)
	}

	// Insert membership
	memID, err := utils.GenerateUUID()
	if err != nil {
		return types.Membership{}, fmt.Errorf("error generating membership ID: %w", err)
	}

	q = `INSERT INTO memberships (id, user_id, organization_id, role_id) VALUES ($1, $2, $3, $4) RETURNING created_at`
	var createdAt time.Time
	err = tx.QueryRow(ctx, q, memID, ownerId, orgId, roleID).Scan(&createdAt)
	if err != nil {
		return types.Membership{}, fmt.Errorf("error inserting membership: %w", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return types.Membership{}, fmt.Errorf("error committing transaction: %w", err)
	}
	return types.Membership{
		ID:        memID,
		OrgName:   name,
		CreatedAt: createdAt,
	}, nil
}

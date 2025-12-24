package organizationstore

import (
	"context"
	"fmt"
	"time"

	"encore.app/organizationoservice/shared/types"
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
	q := `INSERT INTO organizations (name, owner_id) VALUES ($1, $2) RETURNING id`
	var orgId uuid.UUID
	err = tx.QueryRow(ctx, q, name, ownerId).Scan(&orgId)
	if err != nil {
		return types.Membership{}, fmt.Errorf("error inserting organization: %w", err)
	}

	// Insert membership
	q = `INSERT INTO memberships (user_id, organization_id, role_id) VALUES ($1, $2, $3) RETURNING id, created_at`
	var membershipId uuid.UUID
	var createdAt time.Time
	err = tx.QueryRow(ctx, q, ownerId, orgId, roleID).Scan(&membershipId, &createdAt)
	if err != nil {
		return types.Membership{}, fmt.Errorf("error inserting membership: %w", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return types.Membership{}, fmt.Errorf("error committing transaction: %w", err)
	}
	return types.Membership{
		Id:        membershipId,
		OrgName:   name,
		CreatedAt: createdAt,
	}, nil
}

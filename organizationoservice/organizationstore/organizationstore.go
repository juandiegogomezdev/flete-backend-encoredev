package organizationstore

import (
	"context"
	"fmt"
	"log"
	"time"

	"encore.app/organizationoservice/shared/types"
	"encore.dev/beta/auth"
	"encore.dev/storage/sqldb"
	"encore.dev/types/uuid"
	"github.com/jmoiron/sqlx"
)

type OrganizationStore struct {
	db  *sqldb.Database
	dbx *sqlx.DB
}

func NewOrganizationStore(db *sqldb.Database, dbx *sqlx.DB) *OrganizationStore {
	return &OrganizationStore{db: db, dbx: dbx}
}

func (s *OrganizationStore) GetOrganizationIdByMembershipId(ctx context.Context, memID uuid.UUID) (uuid.UUID, error) {
	// Query to get organization ID by membership ID
	q := `SELECT organization_id FROM memberships WHERE id = $1 LIMIT 1`
	var orgId uuid.UUID
	err := s.db.QueryRow(ctx, q, memID).Scan(&orgId)
	if err != nil {
		if err == sqldb.ErrNoRows {
			log.Println("error: organization not found for membership id", memID)
			return uuid.Nil, fmt.Errorf("membership %s not found", memID)
		}
		log.Println("error querying organization id for membership id", memID, ":", err)
		return uuid.Nil, fmt.Errorf("error querying organization id: %w", err)
	}
	return orgId, nil
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

func (s *OrganizationStore) IsAvailableOrganizationName(ctx context.Context, name string, userId auth.UID) (bool, error) {
	q := `SELECT COUNT(1) FROM organizations WHERE name = $1, owner_id = $2`
	var count int
	err := s.db.QueryRow(ctx, q, name, userId).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking organization name availability: %w", err)
	}
	return count == 0, nil
}

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

func (s *OrganizationStore) SaveLogoKeyOrganization(ctx context.Context, orgId uuid.UUID, logoKey string) error {
	q := `UPDATE organizations SET image_key = $1 WHERE id = $2`
	_, err := s.db.Exec(ctx, q, logoKey, orgId)
	if err != nil {
		log.Println("Error updating organization logo key: ", err)
		return fmt.Errorf("error updating organization logo key: %w", err)
	}
	return err
}

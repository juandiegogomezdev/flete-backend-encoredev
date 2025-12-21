package authstore

import (
	"context"
	"fmt"

	"encore.dev/storage/sqldb"
)

type AuthStore struct {
	db *sqldb.Database
}

func NewAuthStore(db *sqldb.Database) *AuthStore {
	if db == nil {
		panic("database is nil")
	}
	return &AuthStore{db: db}
}

func (s *AuthStore) CheckUserExists(ctx context.Context, userID string) (bool, error) {
	var q = `SELECT COUNT(1) FROM users WHERE id = $1`
	var count int
	err := s.db.QueryRow(ctx, q, userID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("Error query user exists %w", err)
	}
	return count > 0, nil
}

func (s *AuthStore) CreateUser(ctx context.Context, usr *UserCreateParams) error {
	var q = `INSERT INTO users (id, email, first_name, last_name, phone) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.Exec(ctx, q, usr.ID, usr.Email, usr.FirstName, usr.LastName, usr.Phone)
	if err != nil {
		return fmt.Errorf("Error creating user: %w", err)
	}
	return nil
}

type UserCreateParams struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Phone     string
}

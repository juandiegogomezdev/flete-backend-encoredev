package authstore

import (
	"context"

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
		return false, err
	}
	return count > 0, nil
}

func (s *AuthStore) CreateUser(ctx context.Context, usr *UserCreateParams) error {
	var q = `INSERT INTO users (id, email, first_name, last_name) VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(ctx, q, usr.ID, usr.Email, usr.FirstName, usr.LastName)
	if err != nil {
		return err
	}
	return nil
}

type UserCreateParams struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
}

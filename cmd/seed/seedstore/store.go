package seedstore

import (
	"database/sql"

	"encore.app/databaseservice/models"
)

type SeedStore struct {
	db *sql.DB
}

func NewSeedStore(db *sql.DB) *SeedStore {
	return &SeedStore{db: db}
}

func (s *SeedStore) InsertDepartments(departments []models.Department) error {
	q := `INSERT INTO departments (id, name, code) VALUES ($1 , $2, $3)`

	// Create a transaction
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, dept := range departments {
		_, err := tx.Exec(q, dept.ID, dept.Name, dept.Code)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *SeedStore) InsertCities(cities []models.City) error {
	q := `INSERT INTO cities (id, code, name, department_id, latitude, longitude) VALUES ($1 , $2, $3, $4, $5, $6)`

	// Create a transaction
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, city := range cities {
		_, err := tx.Exec(q, city.ID, city.Code, city.Name, city.DepartmentID, city.Latitude, city.Longitude)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *SeedStore) InsertUnitCategories(unitCategories []models.UnitCategory) error {
	return nil
}
func (s *SeedStore) InsertUnits(units []models.Unit) error {
	return nil
}

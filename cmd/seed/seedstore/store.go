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
	return nil
}

func (s *SeedStore) InsertCities(cities []models.City) error {
	return nil
}

func (s *SeedStore) InsertUnitCategories(unitCategories []models.UnitCategory) error {
	return nil
}
func (s *SeedStore) InsertUnits(units []models.Unit) error {
	return nil
}

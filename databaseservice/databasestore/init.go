package databasestore

import (
	"context"
	"fmt"

	"encore.app/databaseservice/models"
)

func (s *DatabaseStore) GetDepartments(ctx context.Context) ([]models.Department, error) {
	q := `SELECT id, name, code FROM departments`

	// Execute the query using s.db and handle the result
	rows, err := s.db.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("query departments: %w", err)
	}

	defer rows.Close()

	var departments []models.Department
	for rows.Next() {
		var dept models.Department
		if err := rows.Scan(&dept.ID, &dept.Name, &dept.Code); err != nil {
			return nil, fmt.Errorf("scan department: %w", err)
		}
		departments = append(departments, dept)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return departments, nil

}

func (s *DatabaseStore) GetCities(ctx context.Context) ([]models.City, error) {
	q := `SELECT id, code, name, department_id, latitude, longitude FROM cities`

	// Execute the query using s.db and handle the result
	rows, err := s.db.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("query cities: %w", err)
	}

	defer rows.Close()

	var cities []models.City
	for rows.Next() {
		var city models.City
		if err := rows.Scan(&city.ID, &city.Code, &city.Name, &city.DepartmentID, &city.Latitude, &city.Longitude); err != nil {
			return nil, fmt.Errorf("scan city: %w", err)
		}
		cities = append(cities, city)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return cities, nil
}

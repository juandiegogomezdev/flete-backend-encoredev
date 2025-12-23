package databaseservice

import (
	"context"
	"fmt"

	"encore.app/databaseservice/databasebusiness"
	"encore.app/databaseservice/databasestore"
	"encore.app/databaseservice/models"
	"encore.dev/storage/objects"
	"encore.dev/storage/sqldb"
)

var primaryDB = sqldb.NewDatabase("primary_db", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})
var PrimaryBucketPublic = objects.NewBucket("primary-bucket-public", objects.BucketConfig{
	Public: true,
})
var PrimaryBucketPrivate = objects.NewBucket("primary-bucket-private", objects.BucketConfig{
	Public: false,
})

//encore:service
type DatabaseService struct {
	b *databasebusiness.DatabaseBusiness
}

func initDatabaseService() (*DatabaseService, error) {
	s := databasestore.NewDatabaseStore(primaryDB)
	b := databasebusiness.NewDatabaseBusiness(s)
	return &DatabaseService{b: b}, nil
}

//encore:api auth method=GET path=/departments
func (s *DatabaseService) GetDepartments(ctx context.Context) (GetDepartmentsResponse, error) {
	departments, err := s.b.GetDepartments(ctx)
	fmt.Println("Departments", departments)
	if err != nil {
		return GetDepartmentsResponse{}, err
	}
	return GetDepartmentsResponse{Departments: departments}, nil
}

type GetDepartmentsResponse struct {
	Departments []models.Department `json:"departments"`
}

//encore:api auth method=GET path=/cities
func (s *DatabaseService) GetCities(ctx context.Context) (GetCitiesResponse, error) {
	cities, err := s.b.GetCities(ctx)
	if err != nil {
		return GetCitiesResponse{}, err
	}
	return GetCitiesResponse{Cities: cities}, nil
}

type GetCitiesResponse struct {
	Cities []models.City `json:"cities"`
}

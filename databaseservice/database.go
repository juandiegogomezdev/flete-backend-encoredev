package databaseservice

import (
	"context"

	"encore.app/databaseservice/databasebusiness"
	"encore.app/databaseservice/databasestore"
	"encore.app/databaseservice/models"
	"encore.dev/storage/sqldb"
)

var primaryDB = sqldb.NewDatabase("primary_db", sqldb.DatabaseConfig{
	Migrations: "./migrations",
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
	if err != nil {
		return GetDepartmentsResponse{}, err
	}
	return GetDepartmentsResponse{Departments: departments}, nil
}

type GetDepartmentsResponse struct {
	Departments []models.Department `json:"departments"`
}

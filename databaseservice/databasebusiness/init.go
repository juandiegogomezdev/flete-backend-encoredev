package databasebusiness

import (
	"context"
	"fmt"

	"encore.app/databaseservice/models"
	"encore.dev/beta/errs"
)

func (b *DatabaseBusiness) GetDepartments(ctx context.Context) ([]models.Department, error) {
	departments, err := b.store.GetDepartments(ctx)
	if err != nil {
		fmt.Println("Error in GetDepartments: ", err)
		return nil, &errs.Error{
			Message: "Error al obtener los departamentos",
			Code:    errs.Internal,
		}
	}
	return departments, nil

}

func (b *DatabaseBusiness) GetCities(ctx context.Context) ([]models.City, error) {
	cities, err := b.store.GetCities(ctx)
	if err != nil {
		fmt.Println("Error in GetCities: ", err)
		return nil, &errs.Error{
			Message: "Error al obtener las ciudades",
			Code:    errs.Internal,
		}
	}
	return cities, nil
}

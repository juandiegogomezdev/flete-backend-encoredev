package seedbusiness

import (
	"fmt"
	"strconv"

	"encore.app/cmd/seed/seedstore"
	"encore.app/databaseservice/models"
	"encore.app/pkg/utils"
)

type SeedBusiness struct {
	store *seedstore.SeedStore
}

func NewSeedBusiness(store *seedstore.SeedStore) *SeedBusiness {
	return &SeedBusiness{store: store}
}

type DepartmentsAndCities struct {
	DepartmentCode string `json:"Código Departamento"`
	DepartmentName string `json:"Departamento"`
	CityCode       string `json:"Código Municipio"`
	CityName       string `json:"Municipio"`
	CityType       string `json:"Tipo"`
	Longitude      string `json:"Longitud"`
	Latitude       string `json:"Latitud"`
}

func (b *SeedBusiness) SeedDepartmentsAndCities() error {

	// Read data from JSON files and seed the database
	citiesAndDepartments, err := utils.ReadJsonFile[DepartmentsAndCities]("cmd/seed/data/1_cities&departments.json")
	if err != nil {
		fmt.Println("departments & cities => Error: ", err)
		return err

	}

	// Get unique departments
	setDepartments := make(map[string]models.Department)
	for _, item := range citiesAndDepartments {
		_, exists := setDepartments[item.DepartmentCode]
		if exists {
			continue
		}
		departmentId, err := utils.MustNewUUID()
		if err != nil {
			return err
		}
		setDepartments[item.DepartmentCode] = models.Department{
			ID:   departmentId,
			Code: item.DepartmentCode,
			Name: item.DepartmentName,
		}
	}

	// Convert map to slice
	var departments []models.Department
	for _, dept := range setDepartments {
		dept.ID, err = utils.MustNewUUID()
		departments = append(departments, dept)
	}

	// Cities
	var cities []models.City
	for _, item := range citiesAndDepartments {
		departmentId := setDepartments[item.DepartmentCode].ID
		longitude, err := strconv.ParseFloat(item.Longitude, 64)
		if err != nil {
			return err
		}
		latitude, err := strconv.ParseFloat(item.Latitude, 64)
		if err != nil {
			return err
		}
		city := models.City{
			Name:         item.CityName,
			DepartmentID: departmentId,
			Longitude:    longitude,
			Latitude:     latitude,
		}
		cities = append(cities, city)
	}

	err = b.store.InsertDepartments(departments)
	if err != nil {
		fmt.Println("1_departments => Error: ", err)
		return err
	}

	err = b.store.InsertCities(cities)
	if err != nil {
		fmt.Println("1_cities => Error: ", err)
		return err
	}

	fmt.Println("1_departments => Ok")
	return nil
}

func (b *SeedBusiness) SeedUnitCategories() error {

	unitCategories, err := utils.ReadJsonFile[models.UnitCategory]("cmd/seed/data/3_unit_categories.json")
	if err != nil {
		fmt.Println("3_unit_categories => Error: ", err)
		return err
	}
	err = b.store.InsertUnitCategories(unitCategories)
	if err != nil {
		fmt.Println("3_unit_categories => Error: ", err)
		return err
	}

	fmt.Println("3_unit_categories => Ok")
	return nil
}

func (b *SeedBusiness) SeedUnits() error {
	units, err := utils.ReadJsonFile[models.Unit]("cmd/seed/data/4_units.json")

	if err != nil {
		fmt.Println("4_units => Error: ", err)
		return err
	}
	err = b.store.InsertUnits(units)
	if err != nil {
		fmt.Println("4_units => Error: ", err)
		return err
	}

	fmt.Println("4_units => Ok")
	return nil
}

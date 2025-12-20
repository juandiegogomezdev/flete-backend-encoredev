package seedbusiness

import (
	"fmt"

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

func (b *SeedBusiness) SeedDepartments() error {

	// Read data from JSON files and seed the database
	departments, err := utils.ReadJsonFile[models.Department]("cmd/seed/data/1_departments.json")
	if err != nil {
		fmt.Println("1_departments => Error: ", err)
		return err

	}
	err = b.store.InsertDepartments(departments)
	if err != nil {
		fmt.Println("1_departments => Error: ", err)
		return err
	}

	return nil
}

func (b *SeedBusiness) SeedCities() error {

	cities, err := utils.ReadJsonFile[models.City]("cmd/seed/data/2_cities.json")
	if err != nil {
		fmt.Println("2_cities => Error: ", err)
		return nil
	}
	err = b.store.InsertCities(cities)
	if err != nil {
		fmt.Println("2_cities => Error: ", err)
		return nil
	}

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

	return nil
}

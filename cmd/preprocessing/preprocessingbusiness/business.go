package preprocessingbusiness

import (
	"fmt"

	"encore.app/pkg/utils"
	"encore.dev/types/uuid"
)

type DataPreprocessingBusiness struct{}

func NewDataPreprocessingBusiness() *DataPreprocessingBusiness {
	return &DataPreprocessingBusiness{}
}

type unit struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

type processedUnit struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Abbreviation string    `json:"abbreviation"`
}

func (b *DataPreprocessingBusiness) PreprocessUnit() {
	// Read data from JSON files and seed the database
	units, err := utils.ReadJsonFile[unit]("cmd/data/rawData/3_units.json")
	if err != nil {
		fmt.Println("3_units => Error reading units: ", err)
		return
	}

	var processedUnits []processedUnit
	for _, item := range units {
		unitID, err := utils.GenerateUUID()
		if err != nil {
			return
		}

		processedUnit := processedUnit{
			ID:           unitID,
			Name:         item.Name,
			Abbreviation: item.Abbreviation,
		}
		processedUnits = append(processedUnits, processedUnit)
	}

	// Store processed units for later use
	err = utils.SaveJsonFile("cmd/data/processedData/3_units.json", processedUnits)
	if err != nil {
		fmt.Println("3_units => Error saving processed units: ", err)
		return
	}

}

type roles struct {
	Name string `json:"name"`
}
type processedRoles struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (b *DataPreprocessingBusiness) PreprocessRoles() {
	// Read data from JSON files and seed the database
	rolesData, err := utils.ReadJsonFile[roles]("cmd/data/rawData/5_roles.json")
	if err != nil {
		fmt.Println("5_roles => Error reading roles: ", err)
		return
	}
	var processedRolesData []processedRoles
	for _, item := range rolesData {
		roleID, err := utils.GenerateUUID()
		if err != nil {
			return
		}

		processedRole := processedRoles{
			ID:   roleID,
			Name: item.Name,
		}
		processedRolesData = append(processedRolesData, processedRole)
	}

	// Store processed roles for later use
	err = utils.SaveJsonFile("cmd/data/processedData/5_roles.json", processedRolesData)
	if err != nil {
		fmt.Println("5_roles => Error saving processed roles: ", err)
		return
	}

	fmt.Println("05_Roles => Ok")
}

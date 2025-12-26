package preprocessingbusiness

import (
	"fmt"
	"strconv"
	"strings"

	"encore.app/databaseservice/models"
	"encore.app/pkg/utils"
	"encore.dev/types/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type DepartmentsAndCities struct {
	DepartmentCode interface{} `json:"department_code"`
	DepartmentName string      `json:"department_name"`
	CityCode       interface{} `json:"city_code"`
	CityName       string      `json:"city_name"`
	CityType       string      `json:"type"`
	Longitude      string      `json:"longitude"`
	Latitude       string      `json:"latitude"`
}

type DepartmentsAndCities2 struct {
	DepartmentCode string `json:"department_code"`
	DepartmentName string `json:"department_name"`
	CityCode       string `json:"city_code"`
	CityName       string `json:"city_name"`
	CityType       string `json:"type"`
	Longitude      string `json:"longitude"`
	Latitude       string `json:"latitude"`
}

type Department struct {
	ID   uuid.UUID
	Code string
	Name string
}

func (b *DataPreprocessingBusiness) ExtractCitiesAndDepartments() {

	// Read data from JSON files and seed the database
	citiesAndDepartments, err := utils.ReadJsonFile[DepartmentsAndCities]("cmd/data/rawData/1_cities&departments.json")
	if err != nil {
		fmt.Println("departments & cities => Error: ", err)
		return
	}

	// Transform data
	// 1- Normalize department codes and city codes to string.
	// 2- ToLowercase department and city names and Capitalize first letter.
	var normalized = make([]DepartmentsAndCities2, len(citiesAndDepartments))
	for i, item := range citiesAndDepartments {
		departmentCodeStr := fmt.Sprintf("%v", item.DepartmentCode)
		cityCodeStr := fmt.Sprintf("%v", item.CityCode)

		capitalizer := cases.Title(language.Spanish)

		cityName := capitalizer.String(strings.ToLower(strings.TrimSpace(item.CityName)))
		departmentName := capitalizer.String(strings.ToLower(strings.TrimSpace(item.DepartmentName)))

		normalized[i] = DepartmentsAndCities2{
			DepartmentCode: departmentCodeStr,
			DepartmentName: departmentName,
			CityCode:       cityCodeStr,
			CityName:       cityName,
			CityType:       item.CityType,
			Longitude:      item.Longitude,
			Latitude:       item.Latitude,
		}
	}

	// Extract unique departments
	var departmentMap = make(map[string]Department) // code -> name
	for _, item := range normalized {
		if _, exists := departmentMap[item.DepartmentCode]; exists {
			continue
		}

		departmentId, err := utils.GenerateUUID()
		if err != nil {
			return
		}

		departmentMap[item.DepartmentCode] = Department{
			ID:   departmentId,
			Code: item.DepartmentCode,
			Name: item.DepartmentName,
		}
	}

	// Convert map to slice
	var departments []models.Department
	for _, dept := range departmentMap {
		departments = append(departments, models.Department{
			ID:   dept.ID,
			Code: dept.Code,
			Name: dept.Name,
		})
	}

	// Extract cities and departments
	var cities []models.City

	for _, item := range normalized {
		departmentId := departmentMap[item.DepartmentCode].ID
		cityId, err := utils.GenerateUUID()
		if err != nil {
			return
		}
		lon, err := parseFloatLocale(item.Longitude)
		if err != nil {
			fmt.Println("invalid longitude ", item.Longitude)
			return
		}
		lat, err := parseFloatLocale(item.Latitude)
		if err != nil {
			fmt.Println("invalid latitude ", item.Latitude)
			return
		}
		cities = append(cities, models.City{
			ID:           cityId,
			Name:         item.CityName,
			Code:         item.CityCode,
			DepartmentID: departmentId,
			Longitude:    lon,
			Latitude:     lat,
		})
	}

	// Save cities and departments to JSON files
	err = utils.SaveJsonFile("cmd/data/processedData/1_departments.json", departments)
	if err != nil {
		fmt.Println("01_departments => Error: ", err)
		return
	}
	fmt.Println("01_departments => Ok")

	err = utils.SaveJsonFile("cmd/data/processedData/2_cities.json", cities)
	if err != nil {
		fmt.Println("02_cities => Error: ", err)
		return
	}
	fmt.Println("02_cities => Ok")

}

func parseFloatLocale(s string) (float64, error) {
	s = strings.TrimSpace(s)
	// convierte coma decimal a punto
	s = strings.ReplaceAll(s, ",", ".")
	// opcional: elimina separadores de miles si llegaran en el dataset (e.g. "1.234,56")
	// si contiene más de un '.', asumimos que todos excepto el último son miles:
	if strings.Count(s, ".") > 1 {
		last := strings.LastIndex(s, ".")
		s = strings.ReplaceAll(s[:last], ".", "") + s[last:]
	}
	return strconv.ParseFloat(s, 64)
}

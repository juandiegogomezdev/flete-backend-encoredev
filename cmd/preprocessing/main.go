package main

import "encore.app/cmd/preprocessing/preprocessingbusiness"

func main() {

	b := preprocessingbusiness.NewDataPreprocessingBusiness()
	b.ExtractCitiesAndDepartments()
	b.PreprocessUnit()
	b.PreprocessRoles()

}

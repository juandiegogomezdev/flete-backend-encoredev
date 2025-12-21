package main

import "encore.app/cmd/preprocessing/preprocessingbusiness"

func main() {

	b := preprocessingbusiness.NewDataPreprocessingBusiness()
	err := b.ExtractCitiesAndDepartments()
	if err != nil {
		panic(err)
	}

}

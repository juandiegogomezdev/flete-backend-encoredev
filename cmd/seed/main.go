package main

import (
	"database/sql"
	"fmt"

	"encore.app/cmd/seed/seedbusiness"
	"encore.app/cmd/seed/seedstore"
)

func main() {
	connStr := "user=youruser dbname=yourdb sslmode=disable" // Update with your DB connection string
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Seed data
	store := seedstore.NewSeedStore(db)
	b := seedbusiness.NewSeedBusiness(store)

	b.SeedDepartments()
	b.SeedCities()
	b.SeedUnitCategories()
	b.SeedUnits()

	fmt.Println("Seeding completed.")
}

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"encore.app/cmd/seed/seedbusiness"
	"encore.app/cmd/seed/seedstore"
)

func main() {
	connStr := "postgresql://encoretruck-zpoi:local@127.0.0.1:9500/primary_db?sslmode=disable" // Update with your DB connection string
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Seed data
	store := seedstore.NewSeedStore(db)
	b := seedbusiness.NewSeedBusiness(store)

	fmt.Print(b)

	b.SeedDepartments()
	b.SeedCities()
	b.SeedUnits()
	b.SeedRoles()

	fmt.Println("Seeding completed.")
}

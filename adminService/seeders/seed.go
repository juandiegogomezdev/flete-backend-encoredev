package seeders

import (
	"fmt"

	"encore.dev/beta/errs"
	"github.com/jmoiron/sqlx"
)

func RunSeed(db *sqlx.DB) error {
	// postgresql://encoretruck-zpoi:local@127.0.0.1:9500/db_app?sslmode=disable
	fmt.Println("-------------------------")
	fmt.Println("Executing database seed...")

	// Create transaction fot seeding
	tx, err := db.Beginx()
	if err != nil {
		return &errs.Error{
			Code:    errs.Internal,
			Message: fmt.Sprintf("Error starting transaction for seeding: %v", err),
		}
	}

	InsertNotificationTemplate(tx)

	err = tx.Commit()
	if err != nil {
		return &errs.Error{
			Code:    errs.Internal,
			Message: fmt.Sprintf("Error committing transaction for seeding: %v", err),
		}
	}
	fmt.Println("Database seeded successfully")
	return nil
}

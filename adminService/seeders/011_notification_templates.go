package seeders

import (
	"fmt"

	"encore.dev/types/uuid"
	"github.com/jmoiron/sqlx"
)

type notificationTemplate struct {
	ID              uuid.UUID `json:"id" db:"id"`
	Category        string    `json:"category" db:"category"`
	Version         int       `json:"version" db:"version"`
	SubjectTemplate string    `json:"subject_template" db:"subject_template"`
	BodyTemplate    string    `json:"body_template" db:"body_template"`
	DefaultActions  string    `json:"default_actions" db:"default_actions"`
	IsCurrent       bool
}

func InsertNotificationTemplate(tx *sqlx.Tx) error {
	for i := 0; i < len(templates); i++ {

		fmt.Println("iteration, ", i)
		ID, err := uuid.NewV4()
		if err != nil {
			return err
		}
		templates[i].ID = ID

		query := `
			INSERT INTO notification_templates (id, category, version, subject_template, body_template, default_actions, is_current)
			VALUES (:id, :category, :version, :subject_template, :body_template, :default_actions, :is_current)
		`

		_, err = tx.NamedExec(query, templates[i])
		fmt.Println("Inserted notification template: ", templates[i].Category)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting notification template: %v", err)

		}
	}
	return nil

}

var templates = []notificationTemplate{
	{
		ID:              uuid.Nil,
		Category:        "welcome",
		Version:         1,
		SubjectTemplate: "Bienvenido a Flete, {{.UserName}}!",
		BodyTemplate:    "Nos alegra tenerte como parte de nuestra plataforma de transporte.",
		DefaultActions:  "",
		IsCurrent:       true,
	},
}

// CREATE TABLE IF NOT EXISTS notification_templates (
//     id UUID PRIMARY KEY,
//     category category_notification NOT NULL,
//     version SMALLINT NOT NULL DEFAULT 1,
//     subject_template VARCHAR(150) NOT NULL,
//     body_template TEXT NOT NULL,
//     default_actions JSONB,
//     is_current BOOLEAN NOT NULL DEFAULT TRUE,
//     created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
// )

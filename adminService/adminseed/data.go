package adminseed

import "encore.dev/types/uuid"

type NotificationTemplate struct {
	ID              uuid.UUID `json:"id" db:"id"`
	Category        string    `json:"category" db:"category"`
	Version         int       `json:"version" db:"version"`
	SubjectTemplate string    `json:"subject_template" db:"subject_template"`
	BodyTemplate    string    `json:"body_template" db:"body_template"`
	DefaultActions  string    `json:"default_actions" db:"default_actions"`
	IsCurrent       bool      `json:"is_current" db:"is_current"`
}

func SeedDataNotificationTemplates() []NotificationTemplate {
	return []NotificationTemplate{
		{
			ID:              uuid.Must(uuid.NewV4()),
			Category:        "document_expiry",
			Version:         1,
			SubjectTemplate: "Bienvenido a Flete, {{.UserName}}!",
			BodyTemplate:    "Nos alegra tenerte como parte de nuestra plataforma de transporte.",
			DefaultActions:  "{}",
			IsCurrent:       true,
		},
	}
}

package adminstore

import (
	"context"

	"encore.app/adminService/adminseed"
)

func (s *AdminStore) SeedNotificationTemplates(ctx context.Context, notificationTemplates []adminseed.NotificationTemplate) error {
	// start transaction
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, template := range notificationTemplates {
		_, err := tx.Exec(`
			INSERT INTO notification_templates (id, category, version, subject_template, body_template, default_actions, is_current)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`, template.ID, template.Category, template.Version, template.SubjectTemplate, template.BodyTemplate, template.DefaultActions, template.IsCurrent)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

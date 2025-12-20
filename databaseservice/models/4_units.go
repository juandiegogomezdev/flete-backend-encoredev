package models

import "encore.dev/types/uuid"

type Unit struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	CategoryID uuid.UUID `json:"category_id"`
}

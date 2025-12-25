package models

import "encore.dev/types/uuid"

type Role struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

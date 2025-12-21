package models

import "encore.dev/types/uuid"

type Department struct {
	ID   uuid.UUID `json:"id"`
	Code string    `json:"code"`
	Name string    `json:"name"`
}

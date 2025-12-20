package models

import "encore.dev/types/uuid"

type City struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	DepartmentID uuid.UUID `json:"department_id"`
}

package models

import "encore.dev/types/uuid"

type City struct {
	ID           uuid.UUID `json:"id"`
	Code         string    `json:"code"`
	Name         string    `json:"name"`
	DepartmentID uuid.UUID `json:"department_id"`
	Longitude    float64   `json:"longitude"`
	Latitude     float64   `json:"latitude"`
}

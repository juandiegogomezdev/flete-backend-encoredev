package models

import (
	"time"

	"encore.dev/types/uuid"
)

type Organization struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	OwnerID   string    `db:"owner_id"`
	ImageKey  *string   `db:"image_key"`
	CreatedAt time.Time `db:"created_at"`
	Active    bool      `db:"active"`
}

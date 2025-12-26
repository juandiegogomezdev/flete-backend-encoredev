package types

import (
	"time"

	"encore.dev/types/uuid"
)

// This is the type returned when a new organization is created along with the membership.
type Membership struct {
	ID        uuid.UUID `json:"id"`
	OrgName   string    `json:"orgName"`
	ImageUrl  string    `json:"imageUrl"`
	RoleName  string    `json:"roleName"`
	CreatedAt time.Time `json:"createdAt"`
}

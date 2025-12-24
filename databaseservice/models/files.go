type file struct {
	Id         int64     `db:"id"`
	PublicId   string    `db:"public_id"`
	EntityID   uuid.UUID `db:"entity_id"`
	EntityType string    `db:"entity_type"`

	OwnerUserID string `db:"owner_user_id"`
	OrgID       string `db:"org_id"`

	SizeBytes int64  `db:"size_bytes"`
	MimeType  string `db:"mime_type"`

	Bucket    string `db:"bucket"`
	BucketKey string `db:"bucket_key"`

	CreatedBy string     `db:"created_by"`
	CreatedAt time.Time  `db:"created_at"`
	DeletedBy *string    `db:"deleted_by"`
	DeletedAt *time.Time `db:"deleted_at"`
}
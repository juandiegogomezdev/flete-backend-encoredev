CREATE TABLE files_entity {
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    public_id UUID UNIQUE NOT NULL DEFAULT gen_random_uuid(),

-- Ownership and Organization
    owner_user_id VARCHAR(255) REFERENCES users(id),
    org_id UUID REFERENCES organizations(id),

-- Metadata
    size_bytes BIGINT NOT NULL,
    mime_type VARCHAR(100) NOT NULL,

-- Storage Info
    bucket TEXT NOT NULL,
    key TEXT NOT NULL,

-- State
    created_by VARCHAR(255) REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_by VARCHAR(255) REFERENCES users(id),
    deleted_at TIMESTAMP WITH TIME ZONE
}
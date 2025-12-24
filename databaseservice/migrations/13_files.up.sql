CREATE TYPE files_entity_types AS ENUM ('movement', 'travel', 'document');

CREATE TABLE files (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    public_id UUID UNIQUE NOT NULL DEFAULT gen_random_uuid(),
    entity_id UUID NOT NULL,
    entity_type files_entity_types NOT NULL,

-- Ownership and Organization
    owner_user_id VARCHAR(255) NOT NULL REFERENCES users(id),
    org_id UUID NOT NULL REFERENCES organizations(id),

-- Metadata
    size_bytes BIGINT NOT NULL,
    mime_type VARCHAR(100) NOT NULL,

-- Storage Info
    bucket TEXT NOT NULL,
    bucket_key TEXT NOT NULL,

-- State
    created_by VARCHAR(255) NOT NULL REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_by VARCHAR(255) REFERENCES users(id),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_files_entity ON files (entity_type, entity_id);
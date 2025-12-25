CREATE TYPE membership_status AS ENUM ('active', 'suspended');

CREATE TABLE IF NOT EXISTS memberships (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    org_id UUID REFERENCES organizations(id),
    user_id VARCHAR(255) REFERENCES users(id),
    role_id UUID REFERENCES roles(id),
    status membership_status NOT NULL DEFAULT 'active',
    created_by VARCHAR(255) REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
)
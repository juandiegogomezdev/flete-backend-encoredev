CREATE type type_organizations AS ENUM ('company', 'personal');

CREATE TABLE IF NOT EXISTS organizations (
    id UUID PRIMARY KEY,
    owner_user_id UUID REFERENCES users(id) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    CONSTRAINT unique_owner_user_id_name UNIQUE (owner_user_id, name)
)
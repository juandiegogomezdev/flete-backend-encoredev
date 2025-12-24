CREATE TYPE job_invitation_status AS ENUM ('pending', 'accepted', 'rejected', 'revoked');

CREATE TABLE job_invitations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    org_id UUID REFERENCES organizations(id),
    email VARCHAR(255) NOT NULL,
    role_id UUID REFERENCES roles(id),
    status job_invitation_status NOT NULL DEFAULT 'pending',
    created_by VARCHAR(255) REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
)
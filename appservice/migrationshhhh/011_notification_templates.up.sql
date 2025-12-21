CREATE TYPE category_notification AS ENUM (
    'document_expiry',
    'maintenance',
    'information',
    'invitation',
    'global_announcement',
    'welcome'
);

CREATE TABLE IF NOT EXISTS notification_templates (
    id UUID PRIMARY KEY,
    category category_notification NOT NULL,
    version SMALLINT NOT NULL DEFAULT 1,
    subject_template VARCHAR(150) NOT NULL,
    body_template TEXT NOT NULL,
    default_actions JSONB,
    is_current BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- Add contraints
-- 1. CONSTRAINT unique_template_notification_version UNIQUE (name, version)
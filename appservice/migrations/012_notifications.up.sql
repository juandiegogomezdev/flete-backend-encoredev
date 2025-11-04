CREATE TABLE IF NOT EXISTS notifications (
    id UUID PRIMARY KEY,
    template_id UUID REFERENCES notification_templates(id) NOT NULL,
    template_data jsonb NOT NULL DEFAULT '{}',
    sender_id UUID REFERENCES users(id) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE type notification_user_status AS ENUM (
    'unread',
    'read',
    'action_taken'
);


CREATE TABLE IF NOT EXISTS notification_users (
    id UUID PRIMARY KEY,
    notification_id UUID REFERENCES notifications(id) NOT NULL,
    user_id UUID REFERENCES users(id) NOT NULL,
    status notification_user_status NOT NULL DEFAULT 'unread',
    action_taken VARCHAR(50),
    action_taken_at TIMESTAMP WITH TIME ZONE,
    read_at TIMESTAMP WITH TIME ZONE
);
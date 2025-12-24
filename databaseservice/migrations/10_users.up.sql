CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    phone VARCHAR(20) ,
    email VARCHAR(255) UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    storage_limit_bytes BIGINT NOT NULL DEFAULT 52428800,
    storege_used_bytes BIGINT NOT NULL DEFAULT 0,
)
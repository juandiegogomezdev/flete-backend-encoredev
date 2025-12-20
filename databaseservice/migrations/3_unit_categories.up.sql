CREATE TABLE IF NOT EXISTS unit_categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT UNIQUE NOT NULL 
)
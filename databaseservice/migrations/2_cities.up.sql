CREATE TABLE IF NOT EXISTS cities (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    department_id UUID REFERENCES departments(id),
    code TEXT UNIQUE NOT NULL,
    longitude DECIMAL(3, 6), 
    latitude DECIMAL(2, 6),
    UNIQUE(name, department_id)


)
-- Lontitude range -180 to 180
-- Latitude range -90 to 90
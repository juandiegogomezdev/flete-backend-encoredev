CREATE TABLE IF NOT EXISTS cities (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    name TEXT NOT NULL,
    department_id UUID REFERENCES departments(id),
    code TEXT UNIQUE NOT NULL,
    longitude DECIMAL(9, 6), 
    latitude DECIMAL(8, 6),
    UNIQUE(name, department_id)


)
CREATE TYPE entity_types AS ENUM ('movements', 'travels', 'documents');


CREATE TABLE entity_files (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    file_id BIGINT REFERENCES files(id),
    entity_type entity_types NOT NULL,
    entity_id UUID NOT NULL,
    
    INDEX idx_file_id (file_id)
)
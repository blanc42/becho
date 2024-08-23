ALTER TABLE categories ADD COLUMN unique_identifier VARCHAR(255) UNIQUE NOT NULL;

CREATE INDEX idx_categories_unique_identifier ON categories(unique_identifier);

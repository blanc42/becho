-- Drop foreign key constraints referencing images.id
ALTER TABLE variant_options DROP CONSTRAINT IF EXISTS fk_variant_option_image;
ALTER TABLE product_variant_images DROP CONSTRAINT IF EXISTS fk_product_variant_image_image;

-- Alter images table
ALTER TABLE images
    DROP COLUMN IF EXISTS product_variant_id,
    DROP COLUMN IF EXISTS display_order,
    DROP COLUMN IF EXISTS user_id,
    DROP COLUMN IF EXISTS store_id;

-- Create a new id column
ALTER TABLE images ADD COLUMN new_id SERIAL;

-- Copy data from old id to new_id
UPDATE images SET new_id = CAST(id AS INTEGER);

-- Drop the old id column and rename new_id to id
ALTER TABLE images DROP COLUMN id;
ALTER TABLE images RENAME COLUMN new_id TO id;
ALTER TABLE images ADD PRIMARY KEY (id);

-- Rename image_url to image_id and change its type to UUID
ALTER TABLE images RENAME COLUMN image_url TO image_id;
ALTER TABLE images ALTER COLUMN image_id SET NOT NULL;

ALTER TABLE images
    ALTER COLUMN title DROP NOT NULL;

-- Update product_variant_images table
ALTER TABLE product_variant_images
    DROP COLUMN id,
    ADD COLUMN id SERIAL PRIMARY KEY,
    DROP COLUMN image_id,
    ADD COLUMN image_id INTEGER,
    DROP COLUMN IF EXISTS store_id,
    ADD COLUMN display_order INTEGER NOT NULL DEFAULT 0;

-- Update variant_options table
ALTER TABLE variant_options
    DROP COLUMN image_id,
    ADD COLUMN image_id INTEGER;

-- Add image_id column to users table
ALTER TABLE users
    ADD COLUMN image_id INTEGER;

-- Add image_id column to stores table
ALTER TABLE stores
    ADD COLUMN image_id INTEGER;

-- Add image_id column to categories table
ALTER TABLE categories
    ADD COLUMN image_id INTEGER;

-- Re-add foreign key constraints
ALTER TABLE variant_options
    ADD CONSTRAINT fk_variant_option_image
    FOREIGN KEY (image_id) 
    REFERENCES images(id);

ALTER TABLE product_variant_images
    ADD CONSTRAINT fk_product_variant_image_image
    FOREIGN KEY (image_id) 
    REFERENCES images(id);

ALTER TABLE users
    ADD CONSTRAINT fk_user_image
    FOREIGN KEY (image_id) 
    REFERENCES images(id);

ALTER TABLE stores
    ADD CONSTRAINT fk_store_image
    FOREIGN KEY (image_id) 
    REFERENCES images(id);

ALTER TABLE categories
    ADD CONSTRAINT fk_category_image
    FOREIGN KEY (image_id) 
    REFERENCES images(id);

-- Create indexes for the image_id columns
CREATE INDEX idx_product_variant_images_image_id ON product_variant_images(image_id);
CREATE INDEX idx_variant_option_image_id ON variant_options(image_id);
CREATE INDEX idx_user_image_id ON users(image_id);
CREATE INDEX idx_store_image_id ON stores(image_id);
CREATE INDEX idx_category_image_id ON categories(image_id);
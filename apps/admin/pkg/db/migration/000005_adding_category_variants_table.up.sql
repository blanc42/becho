-- Create CategoryVariant table
CREATE TABLE IF NOT EXISTS category_variants (
    id CHAR(11) PRIMARY KEY,
    category_id CHAR(11) NOT NULL,
    variant_id CHAR(11) NOT NULL,
    CONSTRAINT fk_category_variant_category
        FOREIGN KEY (category_id) 
        REFERENCES categories(id),
    CONSTRAINT fk_category_variant_variant
        FOREIGN KEY (variant_id) 
        REFERENCES variants(id)
);

CREATE INDEX idx_category_variant_category_id ON category_variants(category_id);
CREATE INDEX idx_category_variant_variant_id ON category_variants(variant_id);

-- Remove variants column from categories table
ALTER TABLE categories DROP COLUMN variants;

-- Add logo to stores table
ALTER TABLE stores ADD COLUMN logo VARCHAR(255);

-- Add brand and number_of_ratings to products table
ALTER TABLE products 
    ADD COLUMN brand VARCHAR(255),
    ADD COLUMN number_of_ratings INTEGER DEFAULT 0;

-- Create product_variant_images table
CREATE TABLE IF NOT EXISTS product_variant_images (
    id CHAR(11) PRIMARY KEY,
    image_id CHAR(11) NOT NULL,
    product_variant_id CHAR(11) NOT NULL,
    store_id CHAR(11) NOT NULL,
    CONSTRAINT fk_product_variant_image_image
        FOREIGN KEY (image_id) 
        REFERENCES images(id),
    CONSTRAINT fk_product_variant_image_product_variant
        FOREIGN KEY (product_variant_id) 
        REFERENCES product_variants(id),
    CONSTRAINT fk_product_variant_image_store
        FOREIGN KEY (store_id) 
        REFERENCES stores(id)
);

CREATE INDEX idx_product_variant_images_image_id ON product_variant_images(image_id);
CREATE INDEX idx_product_variant_images_product_variant_id ON product_variant_images(product_variant_id);
CREATE INDEX idx_product_variant_images_store_id ON product_variant_images(store_id);

-- Add store_id and user_id to images table
ALTER TABLE images 
    ADD COLUMN store_id CHAR(11),
    ADD COLUMN user_id CHAR(11),
    ADD CONSTRAINT fk_image_store
        FOREIGN KEY (store_id) 
        REFERENCES stores(id),
    ADD CONSTRAINT fk_image_user
        FOREIGN KEY (user_id) 
        REFERENCES users(id);

CREATE INDEX idx_image_store_id ON images(store_id);
CREATE INDEX idx_image_user_id ON images(user_id);

-- Add title to product_variants table
ALTER TABLE product_variants ADD COLUMN title VARCHAR(255);

-- Create wishlist table
CREATE TABLE IF NOT EXISTS wishlists (
    id CHAR(11) PRIMARY KEY,
    user_id CHAR(11) NOT NULL,
    product_variant_id CHAR(11) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_wishlist_user
        FOREIGN KEY (user_id) 
        REFERENCES users(id),
    CONSTRAINT fk_wishlist_product_variant
        FOREIGN KEY (product_variant_id) 
        REFERENCES product_variants(id)
);

CREATE INDEX idx_wishlist_user_id ON wishlists(user_id);
CREATE INDEX idx_wishlist_product_variant_id ON wishlists(product_variant_id);

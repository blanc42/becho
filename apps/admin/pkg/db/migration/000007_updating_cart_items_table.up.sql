-- Drop existing constraints and indexes
ALTER TABLE cart_items DROP CONSTRAINT fk_cart_item_cart;
DROP INDEX idx_cart_item_cart_id;
ALTER TABLE cart_items DROP CONSTRAINT fk_cart_item_product_item;
DROP INDEX idx_cart_item_product_item_id;

-- Rename product_item_id to product_variant_id
ALTER TABLE cart_items RENAME COLUMN product_item_id TO product_variant_id;

-- Add user_id and store_id columns
ALTER TABLE cart_items ADD COLUMN user_id CHAR(11) NOT NULL;
ALTER TABLE cart_items ADD COLUMN store_id CHAR(11) NOT NULL;

-- Remove cart_id column
ALTER TABLE cart_items DROP COLUMN cart_id;

-- Add new foreign key constraints
ALTER TABLE cart_items
    ADD CONSTRAINT fk_cart_item_product_variant
    FOREIGN KEY (product_variant_id) 
    REFERENCES product_variants(id);

ALTER TABLE cart_items
    ADD CONSTRAINT fk_cart_item_user
    FOREIGN KEY (user_id) 
    REFERENCES users(id);

ALTER TABLE cart_items
    ADD CONSTRAINT fk_cart_item_store
    FOREIGN KEY (store_id) 
    REFERENCES stores(id);

-- Create new indexes
CREATE INDEX idx_cart_item_product_variant_id ON cart_items(product_variant_id);
CREATE INDEX idx_cart_item_user_id ON cart_items(user_id);
CREATE INDEX idx_cart_item_store_id ON cart_items(store_id);

-- Create composite index on store_id and user_id
CREATE INDEX idx_cart_item_store_user ON cart_items(store_id, user_id);

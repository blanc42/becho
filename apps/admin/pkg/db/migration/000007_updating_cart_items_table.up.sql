-- Drop existing constraints and indexes
ALTER TABLE cart_items DROP CONSTRAINT fk_cart_item_product_item;
DROP INDEX idx_cart_item_product_item_id;

-- Rename product_item_id to product_variant_id
ALTER TABLE cart_items RENAME COLUMN product_item_id TO product_variant_id;

ALTER TABLE cart_items ADD COLUMN store_id CHAR(11) NOT NULL;


ALTER TABLE cart_items
    ADD CONSTRAINT fk_cart_item_cart
    FOREIGN KEY (cart_id) 
    REFERENCES carts(id);

ALTER TABLE cart_items
    ADD CONSTRAINT fk_cart_item_store
    FOREIGN KEY (store_id) 
    REFERENCES stores(id);

CREATE INDEX idx_cart_item_cart_id ON cart_items(cart_id);
CREATE INDEX idx_cart_item_store_id ON cart_items(store_id);

-- Create composite index on store_id and user_id
CREATE INDEX idx_cart_item_store_user ON cart_items(store_id, cart_id);

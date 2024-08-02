-- Drop all tables in reverse order of creation to avoid foreign key constraints

DROP TABLE IF EXISTS addresses;
DROP TABLE IF EXISTS countries;
DROP TABLE IF EXISTS cart_items;
DROP TABLE IF EXISTS carts;
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS product_images;
DROP TABLE IF EXISTS product_items;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS variants;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS stores;
DROP TABLE IF EXISTS users;

-- Drop all indexes (although dropping tables will automatically drop their indexes)

DROP INDEX IF EXISTS idx_address_country_id;
DROP INDEX IF EXISTS idx_address_pincode;
DROP INDEX IF EXISTS idx_address_city;
DROP INDEX IF EXISTS idx_country_name;
DROP INDEX IF EXISTS idx_cart_item_cart_id;
DROP INDEX IF EXISTS idx_cart_item_product_item_id;
DROP INDEX IF EXISTS idx_cart_customer_id;
DROP INDEX IF EXISTS idx_order_item_order_id;
DROP INDEX IF EXISTS idx_order_item_product_item_id;
DROP INDEX IF EXISTS idx_order_customer_id;
DROP INDEX IF EXISTS idx_order_store_id;
DROP INDEX IF EXISTS idx_order_order_status;
DROP INDEX IF EXISTS idx_order_payment_status;
DROP INDEX IF EXISTS idx_order_order_number;
DROP INDEX IF EXISTS idx_product_image_product_item_id;
DROP INDEX IF EXISTS idx_product_item_price;
DROP INDEX IF EXISTS idx_product_item_product_id;
DROP INDEX IF EXISTS idx_product_item_sku;
DROP INDEX IF EXISTS idx_product_store_id;
DROP INDEX IF EXISTS idx_product_category_id;
DROP INDEX IF EXISTS idx_product_is_archived;
DROP INDEX IF EXISTS idx_product_is_featured;
DROP INDEX IF EXISTS idx_product_name;
DROP INDEX IF EXISTS idx_variant_name;
DROP INDEX IF EXISTS idx_category_parent_id;
DROP INDEX IF EXISTS idx_category_store_id;
DROP INDEX IF EXISTS idx_category_name;
DROP INDEX IF EXISTS idx_store_user_id;
DROP INDEX IF EXISTS idx_store_name;
DROP INDEX IF EXISTS idx_user_store_id;
DROP INDEX IF EXISTS idx_user_role;
DROP INDEX IF EXISTS idx_user_email;
DROP INDEX IF EXISTS idx_user_username;
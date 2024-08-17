-- Create an enum type for Role
CREATE TYPE user_role AS ENUM ('customer', 'owner', 'admin');

-- Add a comment to explain the enum
COMMENT ON TYPE user_role IS 'Enum for user roles: customer, owner, or admin';


CREATE TABLE IF NOT EXISTS users (
    id CHAR(11) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role user_role NOT NULL,
    store_id CHAR(11) NULL
);

CREATE UNIQUE INDEX idx_user_username ON users(username);
CREATE UNIQUE INDEX idx_user_email ON users(email);
CREATE INDEX idx_user_role ON users(role);
CREATE INDEX idx_user_store_id ON users(store_id);

CREATE TABLE IF NOT EXISTS stores (
    id CHAR(11) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    user_id CHAR(11) NOT NULL,
    CONSTRAINT fk_store_user
        FOREIGN KEY (user_id) 
        REFERENCES users(id)
);

CREATE INDEX idx_store_name ON stores(name);
CREATE INDEX idx_store_user_id ON stores(user_id);

CREATE TABLE IF NOT EXISTS categories (
    id CHAR(11) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    store_id CHAR(11) NOT NULL,
    parent_id CHAR(11),
    variants JSONB, -- in order
    CONSTRAINT fk_category_store
        FOREIGN KEY (store_id) 
        REFERENCES stores(id),
    CONSTRAINT fk_category_parent
        FOREIGN KEY (parent_id) 
        REFERENCES categories(id)
);

CREATE INDEX idx_category_name ON categories(name);
CREATE INDEX idx_category_parent_id ON categories(parent_id);

CREATE TABLE IF NOT EXISTS variants (
    id CHAR(11) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    store_id CHAR(11) NOT NULL,
    CONSTRAINT fk_variant_store
        FOREIGN KEY (store_id) 
        REFERENCES stores(id)
);

CREATE INDEX idx_variant_name ON variants(name);
CREATE INDEX idx_variant_store_id ON variants(store_id);

CREATE TABLE IF NOT EXISTS variant_options (
    id CHAR(11) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    variant_id CHAR(11) NOT NULL,
    name VARCHAR(255) NOT NULL,
    display_order INT NOT NULL,
    CONSTRAINT fk_variant_option_variant
        FOREIGN KEY (variant_id) 
        REFERENCES variants(id)
);

CREATE INDEX idx_variant_option_variant_id ON variant_options(variant_id);
CREATE INDEX idx_variant_option_display_order ON variant_options(display_order);

CREATE TABLE IF NOT EXISTS products (
    id CHAR(11) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    rating FLOAT CHECK (rating > 0 AND rating < 6),
    is_featured BOOLEAN,
    is_archived BOOLEAN,
    has_variants BOOLEAN,
    category_id CHAR(11) NOT NULL,
    store_id CHAR(11) NOT NULL,
    category_name VARCHAR(255) NOT NULL,
    variants JSONB, -- Array of variant IDs in order
    CONSTRAINT fk_product_category
        FOREIGN KEY (category_id) 
        REFERENCES categories(id),
    CONSTRAINT fk_product_store
        FOREIGN KEY (store_id) 
        REFERENCES stores(id)
);

CREATE INDEX idx_product_name ON products(name);
CREATE INDEX idx_product_category_id ON products(category_id);
CREATE INDEX idx_product_is_featured ON products(is_featured);
    CREATE INDEX idx_product_is_archived ON products(is_archived);
CREATE INDEX idx_product_store_id ON products(store_id);

CREATE TABLE IF NOT EXISTS product_variants (
    id CHAR(11) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    product_id CHAR(11) NOT NULL,
    sku VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    price FLOAT NOT NULL,
    discounted_price FLOAT, -- should be less than price
    cost_price FLOAT,
    CONSTRAINT fk_product_item_product
        FOREIGN KEY (product_id) 
        REFERENCES products(id)
);

CREATE INDEX idx_product_variants_product_id ON product_variants(product_id);
CREATE INDEX idx_product_variants_sku ON product_variants(sku);
CREATE INDEX idx_product_variants_price ON product_variants(price);

CREATE TABLE product_variant_options (
    product_variant_id CHAR(11) NOT NULL,
    variant_option_id CHAR(11) NOT NULL,
    PRIMARY KEY (product_variant_id, variant_option_id),
    FOREIGN KEY (product_variant_id) REFERENCES product_variants(id),
    FOREIGN KEY (variant_option_id) REFERENCES variant_options(id)
);

CREATE TABLE IF NOT EXISTS images (
    id CHAR(11) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title VARCHAR(255) NOT NULL,
    product_variant_id CHAR(11),
    display_order INT NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    CONSTRAINT fk_product_image_product_variant
        FOREIGN KEY (product_variant_id) 
        REFERENCES product_variants(id)
);

CREATE INDEX idx_image_product_variant_id ON images(product_variant_id);
CREATE INDEX idx_image_display_order ON images(display_order);

CREATE TABLE IF NOT EXISTS orders (
    id CHAR(11) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    order_number VARCHAR(255) NOT NULL,
    payment_status VARCHAR(50) NOT NULL,
    order_status VARCHAR(50) NOT NULL,
    store_id CHAR(11) NOT NULL,
    customer_id CHAR(11) NOT NULL,
    CONSTRAINT fk_order_store
        FOREIGN KEY (store_id) 
        REFERENCES stores(id),
    CONSTRAINT fk_order_customer
        FOREIGN KEY (customer_id) 
        REFERENCES users(id)
);

CREATE UNIQUE INDEX idx_order_order_number ON orders(order_number);
CREATE INDEX idx_order_payment_status ON orders(payment_status);
CREATE INDEX idx_order_order_status ON orders(order_status);
CREATE INDEX idx_order_store_id ON orders(store_id);
CREATE INDEX idx_order_customer_id ON orders(customer_id);

CREATE TABLE IF NOT EXISTS order_items (
    id CHAR(11) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    product_item_id CHAR(11) NOT NULL,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    order_id CHAR(11) NOT NULL,
    CONSTRAINT fk_order_item_product_item
        FOREIGN KEY (product_item_id) 
        REFERENCES product_variants(id),
    CONSTRAINT fk_order_item_order
        FOREIGN KEY (order_id) 
        REFERENCES orders(id)
);

CREATE INDEX idx_order_item_product_item_id ON order_items(product_item_id);
CREATE INDEX idx_order_item_order_id ON order_items(order_id);

CREATE TABLE IF NOT EXISTS carts (
    id CHAR(11) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    customer_id CHAR(11) NOT NULL,
    total_price FLOAT NOT NULL,
    total_discounted_price FLOAT,
    total_quantity INTEGER NOT NULL,
    CONSTRAINT fk_cart_customer
        FOREIGN KEY (customer_id) 
        REFERENCES users(id)
);

CREATE INDEX idx_cart_customer_id ON carts(customer_id);

CREATE TABLE IF NOT EXISTS cart_items (
    id CHAR(11) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    product_item_id CHAR(11) NOT NULL,
    quantity INTEGER NOT NULL,
    cart_id CHAR(11) NOT NULL,
    CONSTRAINT fk_cart_item_product_item
        FOREIGN KEY (product_item_id) 
        REFERENCES product_variants(id),
    CONSTRAINT fk_cart_item_cart
        FOREIGN KEY (cart_id) 
        REFERENCES carts(id)
);

CREATE INDEX idx_cart_item_product_item_id ON cart_items(product_item_id);
CREATE INDEX idx_cart_item_cart_id ON cart_items(cart_id);

CREATE TABLE IF NOT EXISTS countries (
    id CHAR(11) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    country VARCHAR(255) NOT NULL
);

CREATE UNIQUE INDEX idx_country_name ON countries(country);

CREATE TABLE IF NOT EXISTS addresses (
    id CHAR(11) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    address_line_1 VARCHAR(255) NOT NULL,
    address_line_2 VARCHAR(255),
    city VARCHAR(255) NOT NULL,
    pincode VARCHAR(255) NOT NULL,
    country_id CHAR(11) NOT NULL,
    CONSTRAINT fk_address_country
        FOREIGN KEY (country_id) 
        REFERENCES countries(id)
);

CREATE INDEX idx_address_city ON addresses(city);
CREATE INDEX idx_address_pincode ON addresses(pincode);
CREATE INDEX idx_address_country_id ON addresses(country_id);
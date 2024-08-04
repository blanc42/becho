-- Users
-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, username, email, password, role, store_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at
LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET username = $2, email = $3, password = $4, role = $5, store_id = $6, updated_at = $7
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- Stores
-- name: CreateStore :one
INSERT INTO stores (id, created_at, updated_at, name, description, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetStore :one
SELECT * FROM stores
WHERE id = $1 LIMIT 1;

-- name: ListStores :many
SELECT * FROM stores
ORDER BY created_at
LIMIT $1 OFFSET $2;

-- name: UpdateStore :one
UPDATE stores
SET name = $2, description = $3, updated_at = $4
WHERE id = $1
RETURNING *;

-- name: DeleteStore :exec
DELETE FROM stores
WHERE id = $1;

-- Categories
-- name: CreateCategory :one
INSERT INTO categories (id, created_at, updated_at, name, description, store_id, parent_category_id, variants)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1  LIMIT 1;

-- name: ListCategories :many
SELECT * FROM categories
WHERE store_id = $1
ORDER BY created_at;

-- name: UpdateCategory :one
UPDATE categories
SET name = $2, description = $3, parent_category_id = $4, variants = $5, updated_at = $6
WHERE id = $1 AND store_id = $7
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1 AND store_id = $2;

-- Variants
-- name: CreateVariant :one
INSERT INTO variants (id, created_at, updated_at, name, description, options, store_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetVariant :one
SELECT * FROM variants
WHERE id = $1 LIMIT 1;

-- name: ListVariants :many
SELECT * FROM variants
WHERE store_id = $1
ORDER BY created_at;

-- name: UpdateVariant :one
UPDATE variants
SET name = $2, description = $3, options = $4, updated_at = $5
WHERE id = $1 AND store_id = $6
RETURNING *;

-- name: DeleteVariant :exec
DELETE FROM variants
WHERE id = $1 AND store_id = $2;

-- Products
-- name: CreateProduct :one
INSERT INTO products (id, created_at, updated_at, name, description, rating, is_featured, is_archived, has_variants, category_id, store_id, category_name, variants)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
RETURNING *;

-- name: GetProduct :one
SELECT p.*,
       COALESCE(
       json_agg(json_build_object(
           'id', pi.id,
           'sku', pi.sku,
           'quantity', pi.quantity,
           'price', pi.price,
           'discounted_price', pi.discounted_price,
           'cost_price', pi.cost_price,
           'options', pi.options
       )) ,'[]'::json
       ) AS product_items
FROM products p
LEFT JOIN product_items pi ON p.id = pi.product_id
WHERE p.id = $1
GROUP BY p.id;


-- name: GetProducts :many
WITH product_price_info AS (
    SELECT p.id,
           MIN(pi.price) AS min_price
    FROM products p
    LEFT JOIN product_items pi ON p.id = pi.product_id
    GROUP BY p.id
),
variant_filters AS (
    SELECT jsonb_object_keys(sqlc.narg('variants')::jsonb) AS variant_name,
           jsonb_array_elements_text(sqlc.narg('variants')::jsonb->jsonb_object_keys(sqlc.narg('variants')::jsonb)) AS variant_value
)
SELECT p.*,
       COALESCE(json_agg(json_build_object(
           'id', pi.id,
           'sku', pi.sku,
           'quantity', pi.quantity,
           'price', pi.price,
           'discounted_price', pi.discounted_price,
           'cost_price', pi.cost_price,
           'options', pi.options
       )) FILTER (WHERE pi.id IS NOT NULL), '[]'::json) AS product_items,
       ppi.min_price
FROM products p
LEFT JOIN product_items pi ON p.id = pi.product_id
JOIN product_price_info ppi ON p.id = ppi.id
WHERE p.store_id = sqlc.narg('store_id')
    AND (sqlc.narg('category_id')::text IS NULL OR p.category_id = sqlc.narg('category_id'))
    AND (sqlc.narg('is_featured')::boolean IS NULL OR p.is_featured = sqlc.narg('is_featured'))
    AND (sqlc.narg('is_archived')::boolean IS NULL OR p.is_archived = sqlc.narg('is_archived'))
    AND (sqlc.narg('min_price')::decimal IS NULL OR ppi.min_price >= sqlc.narg('min_price'))
    AND (sqlc.narg('max_price')::decimal IS NULL OR ppi.min_price <= sqlc.narg('max_price'))
    AND (sqlc.narg('search')::text IS NULL OR p.name ILIKE '%' || sqlc.narg('search') || '%' OR p.description ILIKE '%' || sqlc.narg('search') || '%')
    AND (
        sqlc.narg('variants')::jsonb IS NULL
        OR EXISTS (
            SELECT 1
            FROM product_items sub_pi
            WHERE sub_pi.product_id = p.id
            AND (
                SELECT bool_and(
                    sub_pi.options ->> vf.variant_name IN (
                        SELECT jsonb_array_elements_text(sqlc.narg('variants')::jsonb->vf.variant_name)
                    )
                )
                FROM variant_filters vf
            )
        )
    )
GROUP BY p.id, ppi.min_price
LIMIT COALESCE(sqlc.narg('limit')::integer, 10)
OFFSET COALESCE(sqlc.narg('offset')::integer, 0);


-- name: UpdateProduct :one
UPDATE products
SET name = $2, description = $3, rating = $4, is_featured = $5, is_archived = $6, has_variants = $7, category_id = $8, category_name = $9, variants = $10, updated_at = $11
WHERE id = $1 AND store_id = $12
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1 AND store_id = $2;

-- Product Items
-- name: CreateProductItem :one
INSERT INTO product_items (id, created_at, updated_at, product_id, sku, quantity, price, discounted_price, cost_price, options)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: GetProductItem :one
SELECT * FROM product_items
WHERE id = $1 LIMIT 1;

-- name: ListProductItems :many
SELECT * FROM product_items
WHERE product_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3;

-- name: UpdateProductItem :one
UPDATE product_items
SET sku = $2, quantity = $3, price = $4, discounted_price = $5, cost_price = $6, options = $7, updated_at = $8
WHERE id = $1
RETURNING *;

-- name: DeleteProductItem :exec
DELETE FROM product_items
WHERE id = $1;

-- Product Images
-- name: CreateProductImage :one
INSERT INTO product_images (id, created_at, updated_at, product_item_id, image_url)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetProductImage :one
SELECT * FROM product_images
WHERE id = $1 LIMIT 1;

-- name: ListProductImages :many
SELECT * FROM product_images
WHERE product_item_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3;

-- name: UpdateProductImage :one
UPDATE product_images
SET image_url = $2, updated_at = $3
WHERE id = $1
RETURNING *;

-- name: DeleteProductImage :exec
DELETE FROM product_images
WHERE id = $1;

-- Orders
-- name: CreateOrder :one
INSERT INTO orders (id, created_at, updated_at, order_number, payment_status, order_status, store_id, customer_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: ListOrders :many
SELECT * FROM orders
WHERE store_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3;

-- name: UpdateOrder :one
UPDATE orders
SET payment_status = $2, order_status = $3, updated_at = $4
WHERE id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1;

-- Order Items
-- name: CreateOrderItem :one
INSERT INTO order_items (id, created_at, updated_at, product_item_id, quantity, order_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetOrderItem :one
SELECT * FROM order_items
WHERE id = $1 LIMIT 1;

-- name: ListOrderItems :many
SELECT * FROM order_items
WHERE order_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3;

-- name: UpdateOrderItem :one
UPDATE order_items
SET quantity = $2, updated_at = $3
WHERE id = $1
RETURNING *;

-- name: DeleteOrderItem :exec
DELETE FROM order_items
WHERE id = $1;

-- Carts
-- name: CreateCart :one
INSERT INTO carts (id, created_at, updated_at, customer_id, total_price, total_discounted_price, total_quantity)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetCart :one
SELECT * FROM carts
WHERE id = $1 LIMIT 1;

-- name: GetCartByCustomer :one
SELECT * FROM carts
WHERE customer_id = $1 LIMIT 1;

-- name: UpdateCart :one
UPDATE carts
SET total_price = $2, total_discounted_price = $3, total_quantity = $4, updated_at = $5
WHERE id = $1
RETURNING *;

-- name: DeleteCart :exec
DELETE FROM carts
WHERE id = $1;

-- Cart Items
-- name: CreateCartItem :one
INSERT INTO cart_items (id, created_at, updated_at, product_item_id, quantity, cart_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetCartItem :one
SELECT * FROM cart_items
WHERE id = $1 LIMIT 1;

-- name: ListCartItems :many
SELECT * FROM cart_items
WHERE cart_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3;

-- name: UpdateCartItem :one
UPDATE cart_items
SET quantity = $2, updated_at = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCartItem :exec
DELETE FROM cart_items
WHERE id = $1;

-- Countries
-- name: CreateCountry :one
INSERT INTO countries (id, created_at, updated_at, country)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCountry :one
SELECT * FROM countries
WHERE id = $1 LIMIT 1;

-- name: ListCountries :many
SELECT * FROM countries
ORDER BY country
LIMIT $1 OFFSET $2;

-- name: UpdateCountry :one
UPDATE countries
SET country = $2, updated_at = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCountry :exec
DELETE FROM countries
WHERE id = $1;

-- Addresses
-- name: CreateAddress :one
INSERT INTO addresses (id, created_at, updated_at, address_line_1, address_line_2, city, pincode, country_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetAddress :one
SELECT * FROM addresses
WHERE id = $1 LIMIT 1;

-- name: ListAddresses :many
SELECT * FROM addresses
WHERE country_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3;

-- name: UpdateAddress :one
UPDATE addresses
SET address_line_1 = $2, address_line_2 = $3, city = $4, pincode = $5, country_id = $6, updated_at = $7
WHERE id = $1
RETURNING *;

-- name: DeleteAddress :exec
DELETE FROM addresses
WHERE id = $1;
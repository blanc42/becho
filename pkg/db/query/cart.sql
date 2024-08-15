
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

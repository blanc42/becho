
-- Cart Items
-- name: CreateCartItem :one
INSERT INTO cart_items (id, created_at, updated_at, product_variant_id, quantity, user_id, store_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetCartItem :one
SELECT ci.*, pv.sku, pv.price, pv.discounted_price, pv.title
FROM cart_items ci
JOIN product_variants pv ON ci.product_variant_id = pv.id
WHERE ci.id = $1 LIMIT 1;

-- name: ListCartItems :many
SELECT ci.*, pv.sku, pv.price, pv.discounted_price, pv.title
FROM cart_items ci
JOIN product_variants pv ON ci.product_variant_id = pv.id
WHERE ci.user_id = $1 AND ci.store_id = $2
ORDER BY ci.created_at
LIMIT $3 OFFSET $4;

-- name: UpdateCartItem :one
UPDATE cart_items
SET quantity = $2, updated_at = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCartItem :exec
DELETE FROM cart_items
WHERE id = $1;

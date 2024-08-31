-- name: CreateCart :one
INSERT INTO carts (id, user_id )
VALUES ($1, $2)
RETURNING *;

-- name: GetCart :one
SELECT * FROM carts
WHERE id = $1 LIMIT 1;

-- name: DeleteCart :exec
DELETE FROM carts
WHERE id = $1;

-- name: ListCarts :many
SELECT * FROM carts
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CreateCartItem :one
INSERT INTO cart_items (id, cart_id, product_variant_id, quantity)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCartItem :one
SELECT ci.*
FROM cart_items ci
WHERE ci.id = $1 LIMIT 1;

-- name: UpdateCartItem :one
UPDATE cart_items
SET quantity = $2, updated_at = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCartItem :exec
DELETE FROM cart_items
WHERE id = $1;

-- name: ListCartItems :many
SELECT ci.*
FROM cart_items ci
WHERE ci.cart_id = $1
ORDER BY ci.created_at DESC
LIMIT $2 OFFSET $3;
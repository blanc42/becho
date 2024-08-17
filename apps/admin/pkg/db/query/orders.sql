

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
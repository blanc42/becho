
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
WHERE user_id = $1
ORDER BY created_at;

-- name: UpdateStore :one
UPDATE stores
SET name = $2, description = $3, updated_at = $4
WHERE id = $1
RETURNING *;

-- name: DeleteStore :exec
DELETE FROM stores
WHERE id = $1;

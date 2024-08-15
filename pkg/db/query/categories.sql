
-- Categories
-- name: CreateCategory :one
INSERT INTO categories (id, created_at, updated_at, name, description, store_id, parent_id, variants)
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
SET name = $2, description = $3, parent_id = $4, variants = $5, updated_at = $6
WHERE id = $1 AND store_id = $7
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1 AND store_id = $2;
-- name: CreateImage :one
INSERT INTO images (id, created_at, updated_at, title, product_variant_id, display_order, image_url)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateImage :one
UPDATE images
SET title = $2, image_url = $3, display_order = $4, updated_at = $5
WHERE id = $1
RETURNING *;
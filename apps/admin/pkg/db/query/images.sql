-- name: CreateImage :one
INSERT INTO images (created_at, updated_at, title, image_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateImage :one
UPDATE images
SET title = $2, image_id = $3, updated_at = $4
WHERE id = $1
RETURNING *;

-- name: DeleteImage :exec
DELETE FROM images
WHERE id = $1;

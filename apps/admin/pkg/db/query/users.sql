-- Users
-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, username, email, password, role, store_id, image_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsersForStore :many
SELECT * FROM users
WHERE store_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3;

-- name: UpdateUser :one
UPDATE users
SET username = $2, email = $3, password = $4, updated_at = $5, image_id = $6
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


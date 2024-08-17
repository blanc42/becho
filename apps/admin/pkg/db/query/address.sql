

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
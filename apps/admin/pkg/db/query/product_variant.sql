
-- Product Items
-- name: CreateProductItem :one
INSERT INTO product_variants (id, created_at, updated_at, product_id, sku, quantity, price, discounted_price, cost_price)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetProductVariant :one
SELECT * FROM product_variants
WHERE id = $1 LIMIT 1;

-- name: ListProductVariants :many
SELECT * FROM product_variants
WHERE product_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3;

-- name: UpdateProductVariant :one
UPDATE product_variants
SET sku = $2, quantity = $3, price = $4, discounted_price = $5, cost_price = $6, updated_at = $7
WHERE id = $1
RETURNING *;

-- name: DeleteProductItem :exec
DELETE FROM product_variants
WHERE id = $1;


-- name: CreateProductVariantImage :one
INSERT INTO product_variant_images (product_variant_id, image_id, display_order)
VALUES ($1, $2, $3)
RETURNING *;
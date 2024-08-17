-- name: CreateProductVariantOption :one
INSERT INTO product_variant_options (
    product_variant_id, variant_option_id
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetProductVariantOption :one
SELECT * FROM product_variant_options
WHERE product_variant_id = $1 AND variant_option_id = $2;

-- name: ListProductVariantOptionsByProductVariantID :many
SELECT * FROM product_variant_options
WHERE product_variant_id = $1;

-- name: ListProductVariantOptionsByVariantOptionID :many
SELECT * FROM product_variant_options
WHERE variant_option_id = $1;

-- name: DeleteProductVariantOption :exec
DELETE FROM product_variant_options
WHERE product_variant_id = $1 AND variant_option_id = $2;

-- name: DeleteProductVariantOptionsByProductVariantID :exec
DELETE FROM product_variant_options
WHERE product_variant_id = $1;
-- name: CreateVariant :one
INSERT INTO variants (id, created_at, updated_at, name, label, description, store_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetVariant :one
SELECT 
    v.*,
    COALESCE(
        json_agg(
            json_build_object(
                'id', vo.id,
                'created_at', vo.created_at,
                'updated_at', vo.updated_at,
                'variant_id', vo.variant_id,
                'value', vo.value,
                'data', vo.data,
                'image_id', vo.image_id,
                'display_order', vo.display_order
            ) ORDER BY vo.display_order
        ) FILTER (WHERE vo.id IS NOT NULL),
        '[]'::json
    ) AS options
FROM variants v
LEFT JOIN variant_options vo ON v.id = vo.variant_id
WHERE v.id = $1 AND v.store_id = $2
GROUP BY v.id
LIMIT 1;

-- name: ListVariants :many
SELECT 
    v.*,
    COALESCE(
        json_agg(
            json_build_object(
                'id', vo.id,
                'created_at', vo.created_at,
                'updated_at', vo.updated_at,
                'variant_id', vo.variant_id,
                'value', vo.value,
                'data', vo.data,
                'image_id', vo.image_id,
                'display_order', vo.display_order
            ) ORDER BY vo.display_order
        ) FILTER (WHERE vo.id IS NOT NULL),
        '[]'::json
    ) AS options
FROM variants v
LEFT JOIN variant_options vo ON v.id = vo.variant_id
WHERE v.store_id = $1
GROUP BY v.id
ORDER BY v.created_at
LIMIT COALESCE($2, 10) OFFSET COALESCE($3, 0);

-- name: ListVariantsByIds :many
SELECT 
    v.*,
    COALESCE(
        json_agg(
            json_build_object(
                'id', vo.id,
                'created_at', vo.created_at,
                'updated_at', vo.updated_at,
                'variant_id', vo.variant_id,
                'value', vo.value,
                'data', vo.data,
                'image_id', vo.image_id,
                'display_order', vo.display_order
            ) ORDER BY vo.display_order
        ) FILTER (WHERE vo.id IS NOT NULL),
        '[]'::json
    )::json AS options
FROM variants v
LEFT JOIN variant_options vo ON v.id = vo.variant_id
WHERE v.id = ANY($1::char(11)[]) and v.store_id = $2
GROUP BY v.id
ORDER BY v.created_at;

-- name: UpdateVariant :one
UPDATE variants
SET name = $2, description = $3, updated_at = $4, label = $5
WHERE id = $1
RETURNING *;

-- name: DeleteVariant :exec
DELETE FROM variants
WHERE id = $1;

-- name: CreateVariantOption :one
INSERT INTO variant_options (id, created_at, updated_at, variant_id, value, display_order, data, image_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetVariantOption :one
SELECT * FROM variant_options
WHERE id = $1 LIMIT 1;

-- name: ListVariantOptions :many
SELECT * FROM variant_options
WHERE variant_id = $1
ORDER BY display_order
LIMIT $2 OFFSET $3;

-- name: UpdateVariantOption :one
UPDATE variant_options
SET value = $2, display_order = $3, updated_at = $4, image_id = $5, data = $6
WHERE id = $1
RETURNING *;

-- name: DeleteVariantOption :exec
DELETE FROM variant_options
WHERE id = $1;

-- name: GetVariantAndOptionsArrayForVariantIds :many
SELECT
    v.id AS variant_id,
    COALESCE(
        json_agg(vo.id) FILTER (WHERE vo.id IS NOT NULL),
        '[]'::json
    )::json AS variant_options
FROM
    variants v
LEFT JOIN
    variant_options vo ON v.id = vo.variant_id
WHERE
    v.id = ANY($1::char(11)[])
    AND v.store_id = $2
GROUP BY
    v.id
ORDER BY
    v.created_at;


-- name: GetVariantsWithOptionIds :many
SELECT
    v.id AS variant_id,
    COALESCE(
        json_agg(
            json_build_object(
                'id', vo.id,
                'value', vo.value,
                'data', vo.data,
                'image_id', vo.image_id,
                'display_order', vo.display_order
            )
            ORDER BY vo.display_order
        ) FILTER (WHERE vo.id IS NOT NULL),
        '[]'::json
    )::json AS option_ids
FROM
    variants v
LEFT JOIN
    variant_options vo ON v.id = vo.variant_id
WHERE
    v.store_id = $1
GROUP BY
    v.id
ORDER BY
    v.created_at;

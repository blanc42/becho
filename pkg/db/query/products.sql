-- Products
-- name: CreateProduct :one
INSERT INTO products (id, created_at, updated_at, name, description, rating, is_featured, is_archived, has_variants, category_id, store_id, category_name, variants)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
RETURNING *;



---- ----
WITH product_price_info AS (
    SELECT p.id,
           MIN(pv.price) AS min_price
    FROM products p
    LEFT JOIN product_variants pv ON p.id = pv.product_id
    GROUP BY p.id
),
variant_filters AS (
    SELECT jsonb_object_keys(sqlc.narg('variants')::jsonb) AS variant_name,
           jsonb_array_elements_text(sqlc.narg('variants')::jsonb->jsonb_object_keys(sqlc.narg('variants')::jsonb)) AS variant_value
)
SELECT p.*,
       COALESCE(json_agg(json_build_object(
           'id', pv.id,
           'sku', pv.sku,
           'quantity', pv.quantity,
           'price', pv.price,
           'discounted_price', pv.discounted_price,
           'cost_price', pv.cost_price,
           'options', pv.options
       )) FILTER (WHERE pv.id IS NOT NULL), '[]'::json) AS product_variants,
       ppi.min_price
FROM products p
LEFT JOIN product_variants pv ON p.id = pv.product_id
JOIN product_price_info ppi ON p.id = ppi.id
WHERE p.store_id = sqlc.narg('store_id')
    AND (sqlc.narg('category_id')::text IS NULL OR p.category_id = sqlc.narg('category_id'))
    AND (sqlc.narg('is_featured')::boolean IS NULL OR p.is_featured = sqlc.narg('is_featured'))
    AND (sqlc.narg('is_archived')::boolean IS NULL OR p.is_archived = sqlc.narg('is_archived'))
    AND (sqlc.narg('min_price')::decimal IS NULL OR ppi.min_price >= sqlc.narg('min_price'))
    AND (sqlc.narg('max_price')::decimal IS NULL OR ppi.min_price <= sqlc.narg('max_price'))
    AND (sqlc.narg('search')::text IS NULL OR p.name ILIKE '%' || sqlc.narg('search') || '%' OR p.description ILIKE '%' || sqlc.narg('search') || '%')
    AND (
        sqlc.narg('variants')::jsonb IS NULL
        OR EXISTS (
            SELECT 1
            FROM product_variants sub_pv
            WHERE sub_pv.product_id = p.id
            AND (
                SELECT bool_and(
                    sub_pv.options ->> vf.variant_name IN (
                        SELECT jsonb_array_elements_text(sqlc.narg('variants')::jsonb->vf.variant_name)
                    )
                )
                FROM variant_filters vf
            )
        )
    )
GROUP BY p.id, ppi.min_price
LIMIT COALESCE(sqlc.narg('limit')::integer, 10)
OFFSET COALESCE(sqlc.narg('offset')::integer, 0);


-- name: UpdateProduct :one
UPDATE products
SET name = $2, description = $3, rating = $4, is_featured = $5, is_archived = $6, has_variants = $7, category_id = $8, category_name = $9, variants = $10, updated_at = $11
WHERE id = $1 AND store_id = $12
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1 AND store_id = $2;

-- name: GetFilteredProducts :many
WITH filter_data AS (
  SELECT *
  FROM jsonb_each(sqlc.narg('variants')::jsonb) AS f(variant_id, option_ids)
),
filter_options AS (
  SELECT 
    f.variant_id,
    jsonb_array_elements_text(f.option_ids) AS option_id
  FROM filter_data f
)
SELECT p.id,
	COALESCE(
    JSON_AGG(
          JSON_BUILD_OBJECT(
            'id', pv.id,
            'sku', pv.sku,
            'price', pv.price
          )
        ) FILTER (WHERE pv.id IS NOT NULL),
        '[]'::JSON
      ) AS product_variants
FROM products p
	LEFT JOIN product_variants pv ON pv.product_id = p.id
  LEFT JOIN product_variant_options pvo ON pv.id = pvo.product_variant_id
  LEFT JOIN variant_options vo ON vo.id = pvo.variant_option_id
	LEFT JOIN filter_options fo ON vo.id = fo.option_id AND fo.variant_id = vo.variant_id 
	LEFT JOIN variants v ON vo.variant_id = v.id
WHERE 
	p.store_id = sqlc.narg('store_id')
  AND (sqlc.narg('category_id')::text IS NULL OR p.category_id = sqlc.narg('category_id'))
  AND (sqlc.narg('is_featured')::boolean IS NULL OR p.is_featured = sqlc.narg('is_featured'))
  AND (sqlc.narg('is_archived')::boolean IS NULL OR p.is_archived = sqlc.narg('is_archived'))
  AND (sqlc.narg('search')::text IS NULL OR p.name ILIKE '%' || sqlc.narg('search') || '%' OR p.description ILIKE '%' || sqlc.narg('search') || '%')
GROUP BY p.id
LIMIT COALESCE(sqlc.narg('limit')::integer, 10)
OFFSET COALESCE(sqlc.narg('offset')::integer, 0);
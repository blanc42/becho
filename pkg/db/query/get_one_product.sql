-- name: GetProduct :one
SELECT
    p.id AS product_id,
    p.name AS product_name,
    p.description AS product_description,
    p.rating,
    p.is_featured,
    p.is_archived,
    p.has_variants,
    p.category_id,
    p.store_id,
    p.category_name,
    p.variants AS variants_order,
    COALESCE(
        JSON_AGG(
            JSON_BUILD_OBJECT(
                'variant_id', pv.id,
                'sku', pv.sku,
                'quantity', pv.quantity,
                'price', pv.price,
                'discounted_price', pv.discounted_price,
                'cost_price', pv.cost_price,
                'created_at', pv.created_at,
                'updated_at', pv.updated_at,
                'options', (
                    SELECT JSON_OBJECT_AGG(v.name, vo.value)
                    FROM variant_options vo
                    LEFT JOIN product_variant_options pvo ON vo.id = pvo.variant_option_id
                    LEFT JOIN variants v ON vo.variant_id = v.id
                    WHERE pvo.product_variant_id = pv.id
                )
            )
        ) FILTER (WHERE pv.id IS NOT NULL), '[]'
    ) AS product_variants,
    COALESCE(
        JSON_AGG(
            jsonb_build_object('product_variant_id', pv.id) || 
            COALESCE(
                (SELECT jsonb_object_agg(v.id, vo.id)
                 FROM variant_options vo
                 LEFT JOIN product_variant_options pvo ON vo.id = pvo.variant_option_id
                 LEFT JOIN variants v ON vo.variant_id = v.id
                 WHERE pvo.product_variant_id = pv.id),
                '{}'::jsonb
            )
        ) FILTER (WHERE pv.id IS NOT NULL),
        '[]'
    ) AS available_combinations,
    COALESCE(
        (SELECT 
            JSON_AGG(
                JSON_BUILD_OBJECT(
                    'variant_id', gv.variant_id,
                    'name', gv.variant_name,
                    'options', gv.options
                )
                ORDER BY gv.ord
            )
         FROM (
             SELECT
                 v.id AS variant_id,
                 v.name AS variant_name,
                 JSON_AGG(
                     JSON_BUILD_OBJECT(
                         'id', vo.id,
                         'value', vo.value,
                         'data', vo.data
                     )
                 ) AS options,
                 idx.ord
             FROM
                 variants v
             LEFT JOIN variant_options vo ON v.id = vo.variant_id
             LEFT JOIN product_variant_options pvo ON vo.id = pvo.variant_option_id
             LEFT JOIN product_variants pv_inner ON pvo.product_variant_id = pv_inner.id
             LEFT JOIN LATERAL (
                 SELECT ordinality AS ord
                 FROM jsonb_array_elements_text(p.variants) WITH ORDINALITY
                 WHERE value = v.id::text
             ) idx ON true
             WHERE
                 pv_inner.product_id = p.id
             GROUP BY
                 v.id, v.name, idx.ord
         ) gv
        ),
        '[]'::JSON
    ) AS variants
FROM
    products p
LEFT JOIN
    product_variants pv ON p.id = pv.product_id
WHERE
    p.id = $1 AND p.store_id = $2
GROUP BY
    p.id;
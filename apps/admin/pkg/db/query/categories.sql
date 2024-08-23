
-- Categories
-- name: CreateCategory :one
INSERT INTO 
    categories (id, created_at, updated_at, name, description, store_id, parent_id, level, unique_identifier)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetCategory :one
SELECT *, ARRAY_AGG(cv.variant_id) AS variants
FROM categories c
LEFT JOIN category_variants cv ON c.id = cv.category_id
WHERE c.id = $1 AND c.store_id = $2 
GROUP BY c.id
LIMIT 1;

-- name: GetAllCategoriesRecursive :many
WITH RECURSIVE category_tree AS (
    -- Base case: Select root categories (those without a parent)
    SELECT
        c.id,
        c.name,
        c.description,
        c.parent_id,
        0 AS level,
        CAST(c.id AS TEXT) AS path
    FROM 
        categories c
    WHERE 
        c.store_id = $1 AND c.parent_id IS NULL

    UNION ALL

    -- Recursive case: Select child categories
    SELECT 
        c.id,
        c.name,
        c.description,
        c.parent_id,
        ct.level + 1,
        ct.path || '.' || c.id
    FROM 
        categories c
    JOIN 
        category_tree ct ON c.parent_id = ct.id
    WHERE 
        c.store_id = $1
)
SELECT 
    id,
    name,
    description,
    parent_id,
    level,
    path
FROM 
    category_tree
ORDER BY 
    path;


-- name: GetAllCategories :many
SELECT 
    c.id,
    c.name,
    c.description,
    c.parent_id,
    c.level,
    c.unique_identifier,
    ARRAY_AGG(cv.variant_id) AS variants
FROM categories c
LEFT JOIN category_variants cv ON c.id = cv.category_id
WHERE store_id = $1
GROUP BY c.id
ORDER BY created_at;



-- name: ListCategories :many
SELECT * FROM categories
WHERE store_id = $1
ORDER BY created_at;

-- name: UpdateCategory :one
UPDATE categories
SET name = $2, description = $3, parent_id = $4, updated_at = $5
WHERE id = $1 AND store_id = $6
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1 AND store_id = $2;


-- name: CreateCategoryVariant :exec
INSERT INTO category_variants (id, category_id, variant_id)
VALUES ($1, $2, $3);
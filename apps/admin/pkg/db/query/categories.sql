
-- Categories
-- name: CreateCategory :one
INSERT INTO categories (id, created_at, updated_at, name, description, store_id, parent_id, variants)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1  LIMIT 1;

-- name: GetAllCategoriesRecursive :many
WITH RECURSIVE category_tree AS (
    -- Base case: Select root categories (those without a parent)
    SELECT 
        c.id,
        c.name,
        c.description,
        c.parent_id,
        c.variants,
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
        c.variants,
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
    variants,
    level,
    path
FROM 
    category_tree
ORDER BY 
    path;

-- name: ListCategories :many
SELECT * FROM categories
WHERE store_id = $1
ORDER BY created_at;

-- name: UpdateCategory :one
UPDATE categories
SET name = $2, description = $3, parent_id = $4, variants = $5, updated_at = $6
WHERE id = $1 AND store_id = $7
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1 AND store_id = $2;
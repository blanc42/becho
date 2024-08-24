// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: categories.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO 
    categories (id, created_at, updated_at, name, description, store_id, parent_id, level, unique_identifier)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, created_at, updated_at, name, description, store_id, parent_id, level, unique_identifier
`

type CreateCategoryParams struct {
	ID               string           `json:"id"`
	CreatedAt        pgtype.Timestamp `json:"created_at"`
	UpdatedAt        pgtype.Timestamp `json:"updated_at"`
	Name             string           `json:"name"`
	Description      pgtype.Text      `json:"description"`
	StoreID          string           `json:"store_id"`
	ParentID         pgtype.Text      `json:"parent_id"`
	Level            int32            `json:"level"`
	UniqueIdentifier string           `json:"unique_identifier"`
}

// Categories
func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRow(ctx, createCategory,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Description,
		arg.StoreID,
		arg.ParentID,
		arg.Level,
		arg.UniqueIdentifier,
	)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.StoreID,
		&i.ParentID,
		&i.Level,
		&i.UniqueIdentifier,
	)
	return i, err
}

const createCategoryVariant = `-- name: CreateCategoryVariant :exec
INSERT INTO category_variants (id, category_id, variant_id)
VALUES ($1, $2, $3)
`

type CreateCategoryVariantParams struct {
	ID         string `json:"id"`
	CategoryID string `json:"category_id"`
	VariantID  string `json:"variant_id"`
}

func (q *Queries) CreateCategoryVariant(ctx context.Context, arg CreateCategoryVariantParams) error {
	_, err := q.db.Exec(ctx, createCategoryVariant, arg.ID, arg.CategoryID, arg.VariantID)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1 AND store_id = $2
`

type DeleteCategoryParams struct {
	ID      string `json:"id"`
	StoreID string `json:"store_id"`
}

func (q *Queries) DeleteCategory(ctx context.Context, arg DeleteCategoryParams) error {
	_, err := q.db.Exec(ctx, deleteCategory, arg.ID, arg.StoreID)
	return err
}

const getAllCategories = `-- name: GetAllCategories :many
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
ORDER BY created_at
`

type GetAllCategoriesRow struct {
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	Description      pgtype.Text `json:"description"`
	ParentID         pgtype.Text `json:"parent_id"`
	Level            int32       `json:"level"`
	UniqueIdentifier string      `json:"unique_identifier"`
	Variants         interface{} `json:"variants"`
}

func (q *Queries) GetAllCategories(ctx context.Context, storeID string) ([]GetAllCategoriesRow, error) {
	rows, err := q.db.Query(ctx, getAllCategories, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllCategoriesRow
	for rows.Next() {
		var i GetAllCategoriesRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.ParentID,
			&i.Level,
			&i.UniqueIdentifier,
			&i.Variants,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllCategoriesRecursive = `-- name: GetAllCategoriesRecursive :many
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
    path
`

type GetAllCategoriesRecursiveRow struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description pgtype.Text `json:"description"`
	ParentID    pgtype.Text `json:"parent_id"`
	Level       int32       `json:"level"`
	Path        string      `json:"path"`
}

func (q *Queries) GetAllCategoriesRecursive(ctx context.Context, storeID string) ([]GetAllCategoriesRecursiveRow, error) {
	rows, err := q.db.Query(ctx, getAllCategoriesRecursive, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllCategoriesRecursiveRow
	for rows.Next() {
		var i GetAllCategoriesRecursiveRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.ParentID,
			&i.Level,
			&i.Path,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategory = `-- name: GetCategory :one
SELECT c.id, created_at, updated_at, name, description, store_id, parent_id, level, unique_identifier, cv.id, category_id, variant_id, ARRAY_AGG(cv.variant_id) AS variants
FROM categories c
LEFT JOIN category_variants cv ON c.id = cv.category_id
WHERE c.id = $1 AND c.store_id = $2 
GROUP BY c.id
LIMIT 1
`

type GetCategoryParams struct {
	ID      string `json:"id"`
	StoreID string `json:"store_id"`
}

type GetCategoryRow struct {
	ID               string           `json:"id"`
	CreatedAt        pgtype.Timestamp `json:"created_at"`
	UpdatedAt        pgtype.Timestamp `json:"updated_at"`
	Name             string           `json:"name"`
	Description      pgtype.Text      `json:"description"`
	StoreID          string           `json:"store_id"`
	ParentID         pgtype.Text      `json:"parent_id"`
	Level            int32            `json:"level"`
	UniqueIdentifier string           `json:"unique_identifier"`
	ID_2             pgtype.Text      `json:"id_2"`
	CategoryID       pgtype.Text      `json:"category_id"`
	VariantID        pgtype.Text      `json:"variant_id"`
	Variants         interface{}      `json:"variants"`
}

func (q *Queries) GetCategory(ctx context.Context, arg GetCategoryParams) (GetCategoryRow, error) {
	row := q.db.QueryRow(ctx, getCategory, arg.ID, arg.StoreID)
	var i GetCategoryRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.StoreID,
		&i.ParentID,
		&i.Level,
		&i.UniqueIdentifier,
		&i.ID_2,
		&i.CategoryID,
		&i.VariantID,
		&i.Variants,
	)
	return i, err
}

const listCategories = `-- name: ListCategories :many
SELECT id, created_at, updated_at, name, description, store_id, parent_id, level, unique_identifier FROM categories
WHERE store_id = $1
ORDER BY created_at
`

func (q *Queries) ListCategories(ctx context.Context, storeID string) ([]Category, error) {
	rows, err := q.db.Query(ctx, listCategories, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Description,
			&i.StoreID,
			&i.ParentID,
			&i.Level,
			&i.UniqueIdentifier,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :one
UPDATE categories
SET name = $2, description = $3, parent_id = $4, updated_at = $5
WHERE id = $1 AND store_id = $6
RETURNING id, created_at, updated_at, name, description, store_id, parent_id, level, unique_identifier
`

type UpdateCategoryParams struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description pgtype.Text      `json:"description"`
	ParentID    pgtype.Text      `json:"parent_id"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	StoreID     string           `json:"store_id"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.db.QueryRow(ctx, updateCategory,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.ParentID,
		arg.UpdatedAt,
		arg.StoreID,
	)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.StoreID,
		&i.ParentID,
		&i.Level,
		&i.UniqueIdentifier,
	)
	return i, err
}

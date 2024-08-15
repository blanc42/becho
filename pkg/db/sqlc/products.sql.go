// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: products.sql

package db

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (id, created_at, updated_at, name, description, rating, is_featured, is_archived, has_variants, category_id, store_id, category_name, variants)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
RETURNING id, created_at, updated_at, name, description, rating, is_featured, is_archived, has_variants, category_id, store_id, category_name, variants
`

type CreateProductParams struct {
	ID           string           `json:"id"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	Name         string           `json:"name"`
	Description  pgtype.Text      `json:"description"`
	Rating       pgtype.Float8    `json:"rating"`
	IsFeatured   pgtype.Bool      `json:"is_featured"`
	IsArchived   pgtype.Bool      `json:"is_archived"`
	HasVariants  pgtype.Bool      `json:"has_variants"`
	CategoryID   string           `json:"category_id"`
	StoreID      string           `json:"store_id"`
	CategoryName string           `json:"category_name"`
	Variants     json.RawMessage  `json:"variants"`
}

// Products
func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Description,
		arg.Rating,
		arg.IsFeatured,
		arg.IsArchived,
		arg.HasVariants,
		arg.CategoryID,
		arg.StoreID,
		arg.CategoryName,
		arg.Variants,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.Rating,
		&i.IsFeatured,
		&i.IsArchived,
		&i.HasVariants,
		&i.CategoryID,
		&i.StoreID,
		&i.CategoryName,
		&i.Variants,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1 AND store_id = $2
`

type DeleteProductParams struct {
	ID      string `json:"id"`
	StoreID string `json:"store_id"`
}

func (q *Queries) DeleteProduct(ctx context.Context, arg DeleteProductParams) error {
	_, err := q.db.Exec(ctx, deleteProduct, arg.ID, arg.StoreID)
	return err
}

const getFilteredProducts = `-- name: GetFilteredProducts :many
WITH filter_data AS (
  SELECT variant_id, option_ids
  FROM jsonb_each($8::jsonb) AS f(variant_id, option_ids)
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
	p.store_id = $1
  AND ($2::text IS NULL OR p.category_id = $2)
  AND ($3::boolean IS NULL OR p.is_featured = $3)
  AND ($4::boolean IS NULL OR p.is_archived = $4)
  AND ($5::text IS NULL OR p.name ILIKE '%' || $5 || '%' OR p.description ILIKE '%' || $5 || '%')
GROUP BY p.id
LIMIT COALESCE($7::integer, 10)
OFFSET COALESCE($6::integer, 0)
`

type GetFilteredProductsParams struct {
	StoreID    pgtype.Text     `json:"store_id"`
	CategoryID pgtype.Text     `json:"category_id"`
	IsFeatured pgtype.Bool     `json:"is_featured"`
	IsArchived pgtype.Bool     `json:"is_archived"`
	Search     pgtype.Text     `json:"search"`
	Offset     pgtype.Int4     `json:"offset"`
	Limit      pgtype.Int4     `json:"limit"`
	Variants   json.RawMessage `json:"variants"`
}

type GetFilteredProductsRow struct {
	ID              string      `json:"id"`
	ProductVariants interface{} `json:"product_variants"`
}

func (q *Queries) GetFilteredProducts(ctx context.Context, arg GetFilteredProductsParams) ([]GetFilteredProductsRow, error) {
	rows, err := q.db.Query(ctx, getFilteredProducts,
		arg.StoreID,
		arg.CategoryID,
		arg.IsFeatured,
		arg.IsArchived,
		arg.Search,
		arg.Offset,
		arg.Limit,
		arg.Variants,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFilteredProductsRow
	for rows.Next() {
		var i GetFilteredProductsRow
		if err := rows.Scan(&i.ID, &i.ProductVariants); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE products
SET name = $2, description = $3, rating = $4, is_featured = $5, is_archived = $6, has_variants = $7, category_id = $8, category_name = $9, variants = $10, updated_at = $11
WHERE id = $1 AND store_id = $12
RETURNING id, created_at, updated_at, name, description, rating, is_featured, is_archived, has_variants, category_id, store_id, category_name, variants
`

type UpdateProductParams struct {
	ID           string           `json:"id"`
	Name         string           `json:"name"`
	Description  pgtype.Text      `json:"description"`
	Rating       pgtype.Float8    `json:"rating"`
	IsFeatured   pgtype.Bool      `json:"is_featured"`
	IsArchived   pgtype.Bool      `json:"is_archived"`
	HasVariants  pgtype.Bool      `json:"has_variants"`
	CategoryID   string           `json:"category_id"`
	CategoryName string           `json:"category_name"`
	Variants     json.RawMessage  `json:"variants"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	StoreID      string           `json:"store_id"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, updateProduct,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Rating,
		arg.IsFeatured,
		arg.IsArchived,
		arg.HasVariants,
		arg.CategoryID,
		arg.CategoryName,
		arg.Variants,
		arg.UpdatedAt,
		arg.StoreID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.Rating,
		&i.IsFeatured,
		&i.IsArchived,
		&i.HasVariants,
		&i.CategoryID,
		&i.StoreID,
		&i.CategoryName,
		&i.Variants,
	)
	return i, err
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: product_variant.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProductItem = `-- name: CreateProductItem :one
INSERT INTO product_variants (id, created_at, updated_at, product_id, sku, quantity, price, discounted_price, cost_price)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, created_at, updated_at, product_id, sku, quantity, price, discounted_price, cost_price, title
`

type CreateProductItemParams struct {
	ID              string           `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
	ProductID       string           `json:"product_id"`
	Sku             string           `json:"sku"`
	Quantity        int32            `json:"quantity"`
	Price           float64          `json:"price"`
	DiscountedPrice pgtype.Float8    `json:"discounted_price"`
	CostPrice       pgtype.Float8    `json:"cost_price"`
}

// Product Items
func (q *Queries) CreateProductItem(ctx context.Context, arg CreateProductItemParams) (ProductVariant, error) {
	row := q.db.QueryRow(ctx, createProductItem,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ProductID,
		arg.Sku,
		arg.Quantity,
		arg.Price,
		arg.DiscountedPrice,
		arg.CostPrice,
	)
	var i ProductVariant
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductID,
		&i.Sku,
		&i.Quantity,
		&i.Price,
		&i.DiscountedPrice,
		&i.CostPrice,
		&i.Title,
	)
	return i, err
}

const deleteProductItem = `-- name: DeleteProductItem :exec
DELETE FROM product_variants
WHERE id = $1
`

func (q *Queries) DeleteProductItem(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteProductItem, id)
	return err
}

const getProductVariant = `-- name: GetProductVariant :one
SELECT id, created_at, updated_at, product_id, sku, quantity, price, discounted_price, cost_price, title FROM product_variants
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProductVariant(ctx context.Context, id string) (ProductVariant, error) {
	row := q.db.QueryRow(ctx, getProductVariant, id)
	var i ProductVariant
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductID,
		&i.Sku,
		&i.Quantity,
		&i.Price,
		&i.DiscountedPrice,
		&i.CostPrice,
		&i.Title,
	)
	return i, err
}

const listProductVariants = `-- name: ListProductVariants :many
SELECT id, created_at, updated_at, product_id, sku, quantity, price, discounted_price, cost_price, title FROM product_variants
WHERE product_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3
`

type ListProductVariantsParams struct {
	ProductID string `json:"product_id"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
}

func (q *Queries) ListProductVariants(ctx context.Context, arg ListProductVariantsParams) ([]ProductVariant, error) {
	rows, err := q.db.Query(ctx, listProductVariants, arg.ProductID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductVariant
	for rows.Next() {
		var i ProductVariant
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProductID,
			&i.Sku,
			&i.Quantity,
			&i.Price,
			&i.DiscountedPrice,
			&i.CostPrice,
			&i.Title,
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

const updateProductVariant = `-- name: UpdateProductVariant :one
UPDATE product_variants
SET sku = $2, quantity = $3, price = $4, discounted_price = $5, cost_price = $6, updated_at = $7
WHERE id = $1
RETURNING id, created_at, updated_at, product_id, sku, quantity, price, discounted_price, cost_price, title
`

type UpdateProductVariantParams struct {
	ID              string           `json:"id"`
	Sku             string           `json:"sku"`
	Quantity        int32            `json:"quantity"`
	Price           float64          `json:"price"`
	DiscountedPrice pgtype.Float8    `json:"discounted_price"`
	CostPrice       pgtype.Float8    `json:"cost_price"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateProductVariant(ctx context.Context, arg UpdateProductVariantParams) (ProductVariant, error) {
	row := q.db.QueryRow(ctx, updateProductVariant,
		arg.ID,
		arg.Sku,
		arg.Quantity,
		arg.Price,
		arg.DiscountedPrice,
		arg.CostPrice,
		arg.UpdatedAt,
	)
	var i ProductVariant
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductID,
		&i.Sku,
		&i.Quantity,
		&i.Price,
		&i.DiscountedPrice,
		&i.CostPrice,
		&i.Title,
	)
	return i, err
}

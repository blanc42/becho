// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: cart.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCart = `-- name: CreateCart :one
INSERT INTO carts (id, user_id )
VALUES ($1, $2)
RETURNING id, created_at, updated_at, user_id
`

type CreateCartParams struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
}

func (q *Queries) CreateCart(ctx context.Context, arg CreateCartParams) (Cart, error) {
	row := q.db.QueryRow(ctx, createCart, arg.ID, arg.UserID)
	var i Cart
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const createCartItem = `-- name: CreateCartItem :one
INSERT INTO cart_items (id, cart_id, product_variant_id, quantity)
VALUES ($1, $2, $3, $4)
RETURNING id, created_at, updated_at, product_variant_id, quantity, cart_id, store_id
`

type CreateCartItemParams struct {
	ID               string `json:"id"`
	CartID           string `json:"cart_id"`
	ProductVariantID string `json:"product_variant_id"`
	Quantity         int32  `json:"quantity"`
}

func (q *Queries) CreateCartItem(ctx context.Context, arg CreateCartItemParams) (CartItem, error) {
	row := q.db.QueryRow(ctx, createCartItem,
		arg.ID,
		arg.CartID,
		arg.ProductVariantID,
		arg.Quantity,
	)
	var i CartItem
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductVariantID,
		&i.Quantity,
		&i.CartID,
		&i.StoreID,
	)
	return i, err
}

const deleteCart = `-- name: DeleteCart :exec
DELETE FROM carts
WHERE id = $1
`

func (q *Queries) DeleteCart(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteCart, id)
	return err
}

const deleteCartItem = `-- name: DeleteCartItem :exec
DELETE FROM cart_items
WHERE id = $1
`

func (q *Queries) DeleteCartItem(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteCartItem, id)
	return err
}

const getCart = `-- name: GetCart :one
SELECT id, created_at, updated_at, user_id FROM carts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCart(ctx context.Context, id string) (Cart, error) {
	row := q.db.QueryRow(ctx, getCart, id)
	var i Cart
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const getCartItem = `-- name: GetCartItem :one
SELECT ci.id, ci.created_at, ci.updated_at, ci.product_variant_id, ci.quantity, ci.cart_id, ci.store_id
FROM cart_items ci
WHERE ci.id = $1 LIMIT 1
`

func (q *Queries) GetCartItem(ctx context.Context, id string) (CartItem, error) {
	row := q.db.QueryRow(ctx, getCartItem, id)
	var i CartItem
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductVariantID,
		&i.Quantity,
		&i.CartID,
		&i.StoreID,
	)
	return i, err
}

const listCartItems = `-- name: ListCartItems :many
SELECT ci.id, ci.created_at, ci.updated_at, ci.product_variant_id, ci.quantity, ci.cart_id, ci.store_id
FROM cart_items ci
WHERE ci.cart_id = $1
ORDER BY ci.created_at DESC
LIMIT $2 OFFSET $3
`

type ListCartItemsParams struct {
	CartID string `json:"cart_id"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListCartItems(ctx context.Context, arg ListCartItemsParams) ([]CartItem, error) {
	rows, err := q.db.Query(ctx, listCartItems, arg.CartID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CartItem
	for rows.Next() {
		var i CartItem
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProductVariantID,
			&i.Quantity,
			&i.CartID,
			&i.StoreID,
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

const listCarts = `-- name: ListCarts :many
SELECT id, created_at, updated_at, user_id FROM carts
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type ListCartsParams struct {
	UserID string `json:"user_id"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListCarts(ctx context.Context, arg ListCartsParams) ([]Cart, error) {
	rows, err := q.db.Query(ctx, listCarts, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Cart
	for rows.Next() {
		var i Cart
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
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

const updateCartItem = `-- name: UpdateCartItem :one
UPDATE cart_items
SET quantity = $2, updated_at = $3
WHERE id = $1
RETURNING id, created_at, updated_at, product_variant_id, quantity, cart_id, store_id
`

type UpdateCartItemParams struct {
	ID        string           `json:"id"`
	Quantity  int32            `json:"quantity"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) UpdateCartItem(ctx context.Context, arg UpdateCartItemParams) (CartItem, error) {
	row := q.db.QueryRow(ctx, updateCartItem, arg.ID, arg.Quantity, arg.UpdatedAt)
	var i CartItem
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductVariantID,
		&i.Quantity,
		&i.CartID,
		&i.StoreID,
	)
	return i, err
}

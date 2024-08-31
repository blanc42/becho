package domain

import (
	"context"

	db "github.com/blanc42/becho/apps/admin/pkg/db/sqlc"
)

type CartRepository interface {
	CreateCart(ctx context.Context, arg db.CreateCartParams) (db.Cart, error)
	GetCart(ctx context.Context, id string) (db.Cart, error)
	DeleteCart(ctx context.Context, id string) error
	ListCarts(ctx context.Context, customerID string) ([]db.Cart, error)
}

type CartItemRepository interface {
	CreateCartItem(ctx context.Context, arg db.CreateCartItemParams) (db.CartItem, error)
	GetCartItem(ctx context.Context, id string) (db.CartItem, error)
	UpdateCartItem(ctx context.Context, arg db.UpdateCartItemParams) (db.CartItem, error)
	DeleteCartItem(ctx context.Context, id string) error
	ListCartItems(ctx context.Context, arg db.ListCartItemsParams) ([]db.CartItem, error)
}

type cartRepository struct {
	db *db.DbStore
}

func NewCartRepository(db *db.DbStore) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) CreateCart(ctx context.Context, arg db.CreateCartParams) (db.Cart, error) {
	return r.db.CreateCart(ctx, arg)
}

func (r *cartRepository) GetCart(ctx context.Context, id string) (db.Cart, error) {
	return r.db.GetCart(ctx, id)
}

func (r *cartRepository) DeleteCart(ctx context.Context, id string) error {
	return r.db.DeleteCart(ctx, id)
}

func (r *cartRepository) ListCarts(ctx context.Context, customerID string) ([]db.Cart, error) {
	return r.db.ListCarts(ctx, db.ListCartsParams{
		UserID: customerID,
		Limit:  100, // You might want to make this configurable
		Offset: 0,
	})
}

type cartItemRepository struct {
	db *db.DbStore
}

func NewCartItemRepository(db *db.DbStore) CartItemRepository {
	return &cartItemRepository{db: db}
}

func (r *cartItemRepository) CreateCartItem(ctx context.Context, arg db.CreateCartItemParams) (db.CartItem, error) {
	return r.db.CreateCartItem(ctx, arg)
}

func (r *cartItemRepository) GetCartItem(ctx context.Context, id string) (db.CartItem, error) {
	return r.db.GetCartItem(ctx, id)
}

func (r *cartItemRepository) UpdateCartItem(ctx context.Context, arg db.UpdateCartItemParams) (db.CartItem, error) {
	return r.db.UpdateCartItem(ctx, arg)
}

func (r *cartItemRepository) DeleteCartItem(ctx context.Context, id string) error {
	return r.db.DeleteCartItem(ctx, id)
}

func (r *cartItemRepository) ListCartItems(ctx context.Context, arg db.ListCartItemsParams) ([]db.CartItem, error) {
	return r.db.ListCartItems(ctx, arg)
}

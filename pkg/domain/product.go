package domain

import (
	"context"

	db "github.com/blanc42/becho/pkg/db/sqlc"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, arg db.CreateProductParams) (db.Product, error)
	GetProduct(ctx context.Context, product_id string) (db.GetProductRow, error)
	UpdateProduct(ctx context.Context, arg db.UpdateProductParams) (db.Product, error)
	// DeleteProduct(ctx context.Context, id string) error
	GetProducts(ctx context.Context, arg db.GetProductsParams) ([]db.GetProductsRow, error)
}

type productRepository struct {
	db *db.DbStore
}

func NewProductRepository(db *db.DbStore) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(ctx context.Context, arg db.CreateProductParams) (db.Product, error) {
	return r.db.CreateProduct(ctx, arg)
}

func (r *productRepository) GetProduct(ctx context.Context, product_id string) (db.GetProductRow, error) {
	return r.db.GetProduct(ctx, product_id)
}

func (r *productRepository) UpdateProduct(ctx context.Context, arg db.UpdateProductParams) (db.Product, error) {
	return r.db.UpdateProduct(ctx, arg)
}

func (r *productRepository) GetProducts(ctx context.Context, arg db.GetProductsParams) ([]db.GetProductsRow, error) {
	return r.db.GetProducts(ctx, arg)
}

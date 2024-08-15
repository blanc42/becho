package domain

import (
	"context"

	db "github.com/blanc42/becho/pkg/db/sqlc"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, arg db.CreateProductParams) (db.Product, error)
	GetProduct(ctx context.Context, arg db.GetProductParams) (db.GetProductRow, error)
	UpdateProduct(ctx context.Context, arg db.UpdateProductParams) (db.Product, error)
	// DeleteProduct(ctx context.Context, id string) error
	GetProducts(ctx context.Context, arg db.GetFilteredProductsParams) ([]db.GetFilteredProductsRow, error)
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

func (r *productRepository) GetProduct(ctx context.Context, arg db.GetProductParams) (db.GetProductRow, error) {
	return r.db.GetProduct(ctx, arg)
}

func (r *productRepository) UpdateProduct(ctx context.Context, arg db.UpdateProductParams) (db.Product, error) {
	return r.db.UpdateProduct(ctx, arg)
}

func (r *productRepository) GetProducts(ctx context.Context, arg db.GetFilteredProductsParams) ([]db.GetFilteredProductsRow, error) {
	return r.db.GetFilteredProducts(ctx, arg)
}

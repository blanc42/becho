package domain

import (
	"context"

	db "github.com/blanc42/becho/pkg/db/sqlc"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, arg db.CreateCategoryParams) (db.Category, error)
	GetCategory(ctx context.Context, id string) (db.Category, error)
	UpdateCategory(ctx context.Context, arg db.UpdateCategoryParams) (db.Category, error)
	DeleteCategory(ctx context.Context, id string, storeID string) error
	ListCategories(ctx context.Context) ([]db.Category, error)
}

type categoryRepository struct {
	db *db.DbStore
}

func NewCategoryRepository(db *db.DbStore) CategoryRepository {
	return &categoryRepository{db: db}
}

func (cr *categoryRepository) CreateCategory(ctx context.Context, arg db.CreateCategoryParams) (db.Category, error) {
	return cr.db.CreateCategory(ctx, arg)
}

func (cr *categoryRepository) GetCategory(ctx context.Context, id string) (db.Category, error) {
	return cr.db.GetCategory(ctx, id)
}

func (cr *categoryRepository) UpdateCategory(ctx context.Context, arg db.UpdateCategoryParams) (db.Category, error) {
	return cr.db.UpdateCategory(ctx, arg)
}

func (cr *categoryRepository) DeleteCategory(ctx context.Context, id string, storeID string) error {
	return cr.db.DeleteCategory(ctx, db.DeleteCategoryParams{
		ID:      id,
		StoreID: storeID,
	})
}

func (cr *categoryRepository) ListCategories(ctx context.Context) ([]db.Category, error) {
	return cr.db.ListCategories(ctx, "")
}

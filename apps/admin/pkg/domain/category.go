package domain

import (
	"context"

	db "github.com/blanc42/becho/apps/admin/pkg/db/sqlc"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, arg db.CreateCategoryParams) (db.Category, error)
	GetCategory(ctx context.Context, id string, storeID string) (db.GetCategoryRow, error)
	UpdateCategory(ctx context.Context, arg db.UpdateCategoryParams) (db.Category, error)
	DeleteCategory(ctx context.Context, id string, storeID string) error
	ListCategories(ctx context.Context) ([]db.Category, error)
	GetAllCategoriesRecursive(ctx context.Context, storeID string) ([]db.GetAllCategoriesRecursiveRow, error)
	GetAllCategories(ctx context.Context, storeID string) ([]db.GetAllCategoriesRow, error)
	CreateCategoryVariant(ctx context.Context, arg db.CreateCategoryVariantParams) error
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

func (cr *categoryRepository) GetCategory(ctx context.Context, id string, storeID string) (db.GetCategoryRow, error) {
	return cr.db.GetCategory(ctx, db.GetCategoryParams{
		ID:      id,
		StoreID: storeID,
	})
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

func (cr *categoryRepository) GetAllCategoriesRecursive(ctx context.Context, storeID string) ([]db.GetAllCategoriesRecursiveRow, error) {
	return cr.db.GetAllCategoriesRecursive(ctx, storeID)
}

func (cr *categoryRepository) GetAllCategories(ctx context.Context, storeID string) ([]db.GetAllCategoriesRow, error) {
	return cr.db.GetAllCategories(ctx, storeID)
}

func (cr *categoryRepository) CreateCategoryVariant(ctx context.Context, arg db.CreateCategoryVariantParams) error {
	return cr.db.CreateCategoryVariant(ctx, arg)
}

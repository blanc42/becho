package domain

import (
	"context"

	db "github.com/blanc42/becho/pkg/db/sqlc"
)

type VariantRepository interface {
	CreateVariant(ctx context.Context, arg db.CreateVariantParams) (db.Variant, error)
	GetVariant(ctx context.Context, id string) (db.Variant, error)
	UpdateVariant(ctx context.Context, arg db.UpdateVariantParams) (db.Variant, error)
	DeleteVariant(ctx context.Context, id string, storeID string) error
	ListVariants(ctx context.Context) ([]db.Variant, error)
}

type variantRepository struct {
	db *db.DbStore
}

func NewVariantRepository(db *db.DbStore) VariantRepository {
	return &variantRepository{db: db}
}

func (vr *variantRepository) CreateVariant(ctx context.Context, arg db.CreateVariantParams) (db.Variant, error) {
	return vr.db.CreateVariant(ctx, arg)
}

func (vr *variantRepository) GetVariant(ctx context.Context, id string) (db.Variant, error) {
	return vr.db.GetVariant(ctx, id)
}

func (vr *variantRepository) UpdateVariant(ctx context.Context, arg db.UpdateVariantParams) (db.Variant, error) {
	return vr.db.UpdateVariant(ctx, arg)
}

func (vr *variantRepository) DeleteVariant(ctx context.Context, id string, storeID string) error {
	return vr.db.DeleteVariant(ctx, db.DeleteVariantParams{
		ID:      id,
		StoreID: storeID,
	})
}

func (vr *variantRepository) ListVariants(ctx context.Context) ([]db.Variant, error) {
	return vr.db.ListVariants(ctx, "")
}

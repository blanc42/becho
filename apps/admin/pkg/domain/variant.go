package domain

import (
	"context"

	db "github.com/blanc42/becho/pkg/db/sqlc"
)

type VariantRepository interface {
	CreateVariant(ctx context.Context, arg db.CreateVariantParams) (db.Variant, error)
	GetVariant(ctx context.Context, id string, storeID string) (db.GetVariantRow, error)
	UpdateVariant(ctx context.Context, arg db.UpdateVariantParams) (db.Variant, error)
	DeleteVariant(ctx context.Context, id string, storeID string) error
	ListVariants(ctx context.Context, storeID string) ([]db.ListVariantsRow, error)
	CreateVariantOption(ctx context.Context, arg db.CreateVariantOptionParams) (db.VariantOption, error)
	UpdateVariantOption(ctx context.Context, arg db.UpdateVariantOptionParams) (db.VariantOption, error)
	DeleteVariantOption(ctx context.Context, id string, variantID string) error
	ListVariantOptions(ctx context.Context, variantID string) ([]db.VariantOption, error)
	GetVariantAndOptionsArrayForVariantIds(ctx context.Context, variantIDs []string, storeID string) ([]db.GetVariantAndOptionsArrayForVariantIdsRow, error)
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

func (vr *variantRepository) GetVariant(ctx context.Context, id string, storeID string) (db.GetVariantRow, error) {
	return vr.db.GetVariant(ctx, db.GetVariantParams{
		ID:      id,
		StoreID: storeID,
	})
}

func (vr *variantRepository) UpdateVariant(ctx context.Context, arg db.UpdateVariantParams) (db.Variant, error) {
	return vr.db.UpdateVariant(ctx, arg)
}

func (vr *variantRepository) DeleteVariant(ctx context.Context, id string, storeID string) error {
	return vr.db.DeleteVariant(ctx, id)
}

func (vr *variantRepository) ListVariants(ctx context.Context, storeID string) ([]db.ListVariantsRow, error) {
	return vr.db.ListVariants(ctx, db.ListVariantsParams{
		StoreID: storeID,
	})
}

func (vr *variantRepository) CreateVariantOption(ctx context.Context, arg db.CreateVariantOptionParams) (db.VariantOption, error) {
	return vr.db.CreateVariantOption(ctx, arg)
}

func (vr *variantRepository) UpdateVariantOption(ctx context.Context, arg db.UpdateVariantOptionParams) (db.VariantOption, error) {
	return vr.db.UpdateVariantOption(ctx, arg)
}

func (vr *variantRepository) DeleteVariantOption(ctx context.Context, id string, variantID string) error {
	return vr.db.DeleteVariantOption(ctx, id)
}

func (vr *variantRepository) ListVariantOptions(ctx context.Context, variantID string) ([]db.VariantOption, error) {
	return vr.db.ListVariantOptions(ctx, db.ListVariantOptionsParams{
		VariantID: variantID,
	})
}

func (vr *variantRepository) GetVariantAndOptionsArrayForVariantIds(ctx context.Context, variantIDs []string, storeID string) ([]db.GetVariantAndOptionsArrayForVariantIdsRow, error) {
	return vr.db.GetVariantAndOptionsArrayForVariantIds(ctx, db.GetVariantAndOptionsArrayForVariantIdsParams{
		Column1: variantIDs,
		StoreID: storeID,
	})
}

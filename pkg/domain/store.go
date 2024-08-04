package domain

import (
	"context"

	db "github.com/blanc42/becho/pkg/db/sqlc"
)

type StoreRepository interface {
	CreateStore(ctx context.Context, arg db.CreateStoreParams) (db.Store, error)
	GetStore(ctx context.Context, store_id string) (db.Store, error)
	UpdateStore(ctx context.Context, arg db.UpdateStoreParams) (db.Store, error)
	DeleteStore(ctx context.Context, id string) error
	ListStores(ctx context.Context, arg db.ListStoresParams) ([]db.Store, error)
}

type storeRepository struct {
	db *db.DbStore
}

func NewStoreRepository(db *db.DbStore) StoreRepository {
	return &storeRepository{db: db}
}

func (sr *storeRepository) CreateStore(ctx context.Context, arg db.CreateStoreParams) (db.Store, error) {
	return sr.db.CreateStore(ctx, arg)
}

func (sr *storeRepository) GetStore(ctx context.Context, id string) (db.Store, error) {
	return sr.db.GetStore(ctx, id)
}

func (sr *storeRepository) UpdateStore(ctx context.Context, arg db.UpdateStoreParams) (db.Store, error) {
	return sr.db.UpdateStore(ctx, arg)
}

func (sr *storeRepository) DeleteStore(ctx context.Context, id string) error {
	return sr.db.DeleteStore(ctx, id)
}

func (sr *storeRepository) ListStores(ctx context.Context, arg db.ListStoresParams) ([]db.Store, error) {
	return sr.db.ListStores(ctx, arg)
}

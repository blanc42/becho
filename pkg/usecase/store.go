package usecase

import (
	"context"
	"fmt"
	"time"

	db "github.com/blanc42/becho/pkg/db/sqlc"
	"github.com/blanc42/becho/pkg/domain"
	"github.com/blanc42/becho/pkg/handlers/request"
	"github.com/blanc42/becho/pkg/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type StoreUsecase interface {
	CreateStore(ctx context.Context, userID string, storeRequest request.CreateStoreRequest) (string, error)
	GetStore(ctx context.Context, id string) (db.Store, error)
	UpdateStore(ctx context.Context, id string, storeRequest request.UpdateStoreRequest) (db.Store, error)
	DeleteStore(ctx context.Context, id string) error
	ListStores(ctx context.Context, limit, offset int32) ([]db.Store, error)
}

type storeUseCase struct {
	storeRepo domain.StoreRepository
	userRepo  domain.UserRepository
}

func NewStoreUseCase(storeRepo domain.StoreRepository, userRepo domain.UserRepository) StoreUsecase {
	return &storeUseCase{
		storeRepo: storeRepo,
		userRepo:  userRepo,
	}
}

func (s *storeUseCase) CreateStore(ctx context.Context, userID string, storeRequest request.CreateStoreRequest) (string, error) {
	// Check if the user is an admin
	user, err := s.userRepo.GetUser(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %w", err)
	}

	if user.Role != "admin" {
		return "", fmt.Errorf("only admins can create stores")
	}

	// Generate a new ID for the store
	id, err := utils.GenerateShortID()
	if err != nil {
		return "", fmt.Errorf("failed to generate store ID: %w", err)
	}

	// Create the store
	newStore := db.CreateStoreParams{
		ID:          id,
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
		Name:        storeRequest.Name,
		Description: pgtype.Text{String: storeRequest.Description, Valid: true},
		UserID:      userID,
	}

	createdStore, err := s.storeRepo.CreateStore(ctx, newStore)
	if err != nil {
		return "", fmt.Errorf("failed to create store: %w", err)
	}

	return createdStore.ID, nil
}

// TODO : return storeResponse and not db.Store
func (s *storeUseCase) GetStore(ctx context.Context, store_id string) (db.Store, error) {
	store, err := s.storeRepo.GetStore(ctx, store_id)
	if err != nil {
		return db.Store{}, fmt.Errorf("failed to get store: %w", err)
	}
	return store, nil
}

func (s *storeUseCase) UpdateStore(ctx context.Context, id string, storeRequest request.UpdateStoreRequest) (db.Store, error) {
	existingStore, err := s.storeRepo.GetStore(ctx, id)
	if err != nil {
		return db.Store{}, fmt.Errorf("failed to get store: %w", err)
	}

	updateParams := db.UpdateStoreParams{
		ID:        id,
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if storeRequest.Name != nil {
		updateParams.Name = *storeRequest.Name
	} else {
		updateParams.Name = existingStore.Name
	}

	if storeRequest.Description != nil {
		updateParams.Description = pgtype.Text{String: *storeRequest.Description, Valid: true}
	} else {
		updateParams.Description = existingStore.Description
	}

	updatedStore, err := s.storeRepo.UpdateStore(ctx, updateParams)
	if err != nil {
		return db.Store{}, fmt.Errorf("failed to update store: %w", err)
	}

	return updatedStore, nil
}

func (s *storeUseCase) DeleteStore(ctx context.Context, id string) error {
	err := s.storeRepo.DeleteStore(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete store: %w", err)
	}
	return nil
}

func (s *storeUseCase) ListStores(ctx context.Context, limit, offset int32) ([]db.Store, error) {
	stores, err := s.storeRepo.ListStores(ctx, db.ListStoresParams{Limit: limit, Offset: offset})
	if err != nil {
		return nil, fmt.Errorf("failed to list stores: %w", err)
	}
	return stores, nil
}

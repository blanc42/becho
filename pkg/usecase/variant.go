package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	db "github.com/blanc42/becho/pkg/db/sqlc"
	"github.com/blanc42/becho/pkg/domain"
	"github.com/blanc42/becho/pkg/handlers/request"
	"github.com/blanc42/becho/pkg/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type VariantUsecase interface {
	CreateVariant(ctx context.Context, variantRequest request.CreateVariantRequest) (string, error)
	GetVariant(ctx context.Context, id string) (db.Variant, error)
	UpdateVariant(ctx context.Context, id string, variantRequest request.UpdateVariantRequest) (db.Variant, error)
	DeleteVariant(ctx context.Context, id string, storeID string) error
	ListVariants(ctx context.Context) ([]db.Variant, error)
}

type variantUseCase struct {
	variantRepo domain.VariantRepository
	storeRepo   domain.StoreRepository
}

func NewVariantUseCase(variantRepo domain.VariantRepository, storeRepo domain.StoreRepository) VariantUsecase {
	return &variantUseCase{
		variantRepo: variantRepo,
		storeRepo:   storeRepo,
	}
}

func (v *variantUseCase) CreateVariant(ctx context.Context, variantRequest request.CreateVariantRequest) (string, error) {
	// Check if the store exists
	_, err := v.storeRepo.GetStore(ctx, variantRequest.StoreID)
	if err != nil {
		return "", fmt.Errorf("failed to get store: %w", err)
	}

	id, err := utils.GenerateShortID()
	if err != nil {
		return "", fmt.Errorf("failed to generate variant ID: %w", err)
	}

	optionsJSON, err := json.Marshal(variantRequest.Options)
	if err != nil {
		return "", fmt.Errorf("failed to marshal options: %w", err)
	}

	newVariant := db.CreateVariantParams{
		ID:          id,
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
		Name:        variantRequest.Name,
		Description: pgtype.Text{String: variantRequest.Description, Valid: true},
		Options:     optionsJSON,
		StoreID:     variantRequest.StoreID,
	}

	createdVariant, err := v.variantRepo.CreateVariant(ctx, newVariant)
	if err != nil {
		return "", fmt.Errorf("failed to create variant: %w", err)
	}

	return createdVariant.ID, nil
}

func (v *variantUseCase) GetVariant(ctx context.Context, id string) (db.Variant, error) {
	variant, err := v.variantRepo.GetVariant(ctx, id)
	if err != nil {
		return db.Variant{}, fmt.Errorf("failed to get variant: %w", err)
	}
	return variant, nil
}

func (v *variantUseCase) UpdateVariant(ctx context.Context, id string, variantRequest request.UpdateVariantRequest) (db.Variant, error) {
	existingVariant, err := v.variantRepo.GetVariant(ctx, id)
	if err != nil {
		return db.Variant{}, fmt.Errorf("failed to get variant: %w", err)
	}

	updateParams := db.UpdateVariantParams{
		ID:        id,
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if variantRequest.Name != nil {
		updateParams.Name = *variantRequest.Name
	} else {
		updateParams.Name = existingVariant.Name
	}

	if variantRequest.Description != nil {
		updateParams.Description = pgtype.Text{String: *variantRequest.Description, Valid: true}
	} else {
		updateParams.Description = existingVariant.Description
	}

	if variantRequest.Options != nil {
		optionsJSON, err := json.Marshal(variantRequest.Options)
		if err != nil {
			return db.Variant{}, fmt.Errorf("failed to marshal options: %w", err)
		}
		updateParams.Options = optionsJSON
	} else {
		updateParams.Options = existingVariant.Options
	}

	updatedVariant, err := v.variantRepo.UpdateVariant(ctx, updateParams)
	if err != nil {
		return db.Variant{}, fmt.Errorf("failed to update variant: %w", err)
	}

	return updatedVariant, nil
}

func (v *variantUseCase) DeleteVariant(ctx context.Context, id string, storeID string) error {
	err := v.variantRepo.DeleteVariant(ctx, id, storeID)
	if err != nil {
		return fmt.Errorf("failed to delete variant: %w", err)
	}
	return nil
}

func (v *variantUseCase) ListVariants(ctx context.Context) ([]db.Variant, error) {
	variants, err := v.variantRepo.ListVariants(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list variants: %w", err)
	}
	return variants, nil
}

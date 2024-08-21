package usecase

import (
	"context"
	"fmt"
	"time"

	db "github.com/blanc42/becho/apps/admin/pkg/db/sqlc"
	"github.com/blanc42/becho/apps/admin/pkg/domain"
	"github.com/blanc42/becho/apps/admin/pkg/handlers/request"
	"github.com/blanc42/becho/apps/admin/pkg/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type VariantUsecase interface {
	CreateVariant(ctx context.Context, variantRequest request.CreateVariantRequest) (string, error)
	GetVariant(ctx context.Context, id string, storeID string) (db.GetVariantRow, error)
	UpdateVariant(ctx context.Context, id string, variantRequest request.UpdateVariantRequest) (db.Variant, error)
	DeleteVariant(ctx context.Context, id string, storeID string) error
	ListVariants(ctx context.Context, storeID string) ([]db.ListVariantsRow, error)
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

func (v *variantUseCase) CreateVariant(ctx context.Context, req request.CreateVariantRequest) (string, error) {
	// Check if the store exists
	_, err := v.storeRepo.GetStore(ctx, req.StoreID)
	if err != nil {
		return "", fmt.Errorf("failed to get store: %w", err)
	}

	id, err := utils.GenerateShortID()
	if err != nil {
		return "", fmt.Errorf("failed to generate variant ID: %w", err)
	}

	newVariant := db.CreateVariantParams{
		ID:          id,
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
		Name:        req.Name,
		Description: pgtype.Text{String: *req.Description, Valid: req.Description != nil},
		StoreID:     req.StoreID,
		Label:       req.Label,
	}

	createdVariant, err := v.variantRepo.CreateVariant(ctx, newVariant)
	if err != nil {
		return "", fmt.Errorf("failed to create variant: %w", err)
	}

	// Create variant options
	for _, option := range req.Options {
		optionID, _ := utils.GenerateShortID()
		newOption := db.CreateVariantOptionParams{
			ID:           optionID,
			CreatedAt:    pgtype.Timestamp{Time: time.Now(), Valid: true},
			UpdatedAt:    pgtype.Timestamp{Time: time.Now(), Valid: true},
			VariantID:    createdVariant.ID,
			Value:        option.Value,
			DisplayOrder: option.DisplayOrder,
		}

		if option.ImageId != nil {
			newOption.ImageID = pgtype.Text{String: *option.ImageId, Valid: true}
		}
		if option.Data != nil {
			newOption.Data = pgtype.Text{String: *option.Data, Valid: true}
		}

		_, err := v.variantRepo.CreateVariantOption(ctx, newOption)
		if err != nil {
			return "", fmt.Errorf("failed to create variant option: %w", err)
		}
	}

	return createdVariant.ID, nil
}

func (v *variantUseCase) UpdateVariant(ctx context.Context, id string, req request.UpdateVariantRequest) (db.Variant, error) {
	existingVariant, err := v.variantRepo.GetVariant(ctx, id, req.StoreID)
	if err != nil {
		return db.Variant{}, fmt.Errorf("failed to get variant: %w", err)
	}

	updateParams := db.UpdateVariantParams{
		ID:          id,
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
		Name:        existingVariant.Name,
		Description: existingVariant.Description,
	}

	if req.Name != nil {
		updateParams.Name = *req.Name
	}
	if req.Description != nil {
		updateParams.Description = pgtype.Text{String: *req.Description, Valid: true}
	}

	updatedVariant, err := v.variantRepo.UpdateVariant(ctx, updateParams)
	if err != nil {
		return db.Variant{}, fmt.Errorf("failed to update variant: %w", err)
	}

	// Update or create variant options
	for _, option := range req.Options {
		if option.ID != nil {
			// Update existing option
			updateOptionParams := db.UpdateVariantOptionParams{
				ID:           *option.ID,
				UpdatedAt:    pgtype.Timestamp{Time: time.Now(), Valid: true},
				Value:        *option.Value,
				DisplayOrder: *option.DisplayOrder,
				ImageID:      pgtype.Text{String: *option.ImageId, Valid: option.ImageId != nil},
				Data:         pgtype.Text{String: *option.Data, Valid: option.Data != nil},
			}
			_, err := v.variantRepo.UpdateVariantOption(ctx, updateOptionParams)
			if err != nil {
				return db.Variant{}, fmt.Errorf("failed to update variant option: %w", err)
			}
		} else {
			// Create new option
			newOptionID, _ := utils.GenerateShortID()
			createOptionParams := db.CreateVariantOptionParams{
				ID:           newOptionID,
				CreatedAt:    pgtype.Timestamp{Time: time.Now(), Valid: true},
				UpdatedAt:    pgtype.Timestamp{Time: time.Now(), Valid: true},
				VariantID:    id,
				Value:        *option.Value,
				DisplayOrder: *option.DisplayOrder,
				ImageID:      pgtype.Text{String: *option.ImageId, Valid: option.ImageId != nil},
				Data:         pgtype.Text{String: *option.Data, Valid: option.Data != nil},
			}
			_, err := v.variantRepo.CreateVariantOption(ctx, createOptionParams)
			if err != nil {
				return db.Variant{}, fmt.Errorf("failed to create new variant option: %w", err)
			}
		}
	}

	return updatedVariant, nil
}

func (v *variantUseCase) GetVariant(ctx context.Context, id string, storeID string) (db.GetVariantRow, error) {
	variant, err := v.variantRepo.GetVariant(ctx, id, storeID)
	if err != nil {
		return db.GetVariantRow{}, fmt.Errorf("failed to get variant: %w", err)
	}
	return variant, nil
}

func (v *variantUseCase) DeleteVariant(ctx context.Context, id string, storeID string) error {
	err := v.variantRepo.DeleteVariant(ctx, id, storeID)
	if err != nil {
		return fmt.Errorf("failed to delete variant: %w", err)
	}
	return nil
}

func (v *variantUseCase) ListVariants(ctx context.Context, storeID string) ([]db.ListVariantsRow, error) {
	variants, err := v.variantRepo.ListVariants(ctx, storeID)
	if err != nil {
		return nil, fmt.Errorf("failed to list variants: %w", err)
	}
	if len(variants) == 0 {
		return []db.ListVariantsRow{}, nil
	}

	return variants, nil
}

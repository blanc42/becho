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

type CategoryUsecase interface {
	CreateCategory(ctx context.Context, categoryRequest request.CreateCategoryRequest) (string, error)
	GetCategory(ctx context.Context, id string) (db.Category, error)
	UpdateCategory(ctx context.Context, id string, categoryRequest request.UpdateCategoryRequest) (db.Category, error)
	DeleteCategory(ctx context.Context, id string, storeID string) error
	ListCategories(ctx context.Context) ([]db.Category, error)
}

type categoryUseCase struct {
	categoryRepo domain.CategoryRepository
	storeRepo    domain.StoreRepository
}

func NewCategoryUseCase(categoryRepo domain.CategoryRepository, storeRepo domain.StoreRepository) CategoryUsecase {
	return &categoryUseCase{
		categoryRepo: categoryRepo,
		storeRepo:    storeRepo,
	}
}

func (c *categoryUseCase) CreateCategory(ctx context.Context, categoryRequest request.CreateCategoryRequest) (string, error) {
	// Check if the store exists
	_, err := c.storeRepo.GetStore(ctx, categoryRequest.StoreID)
	if err != nil {
		return "", fmt.Errorf("failed to get store: %w", err)
	}

	id, err := utils.GenerateShortID()
	if err != nil {
		return "", fmt.Errorf("failed to generate category ID: %w", err)
	}

	newCategory := db.CreateCategoryParams{
		ID:               id,
		CreatedAt:        pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt:        pgtype.Timestamp{Time: time.Now(), Valid: true},
		Name:             categoryRequest.Name,
		Description:      pgtype.Text{String: categoryRequest.Description, Valid: true},
		ParentCategoryID: pgtype.Text{String: categoryRequest.ParentID, Valid: categoryRequest.ParentID != ""},
		StoreID:          categoryRequest.StoreID,
	}

	createdCategory, err := c.categoryRepo.CreateCategory(ctx, newCategory)
	if err != nil {
		return "", fmt.Errorf("failed to create category: %w", err)
	}

	return createdCategory.ID, nil
}

func (c *categoryUseCase) GetCategory(ctx context.Context, id string) (db.Category, error) {
	category, err := c.categoryRepo.GetCategory(ctx, id)
	if err != nil {
		return db.Category{}, fmt.Errorf("failed to get category: %w", err)
	}
	return category, nil
}

func (c *categoryUseCase) UpdateCategory(ctx context.Context, id string, categoryRequest request.UpdateCategoryRequest) (db.Category, error) {
	existingCategory, err := c.categoryRepo.GetCategory(ctx, id)
	if err != nil {
		return db.Category{}, fmt.Errorf("failed to get category: %w", err)
	}

	updateParams := db.UpdateCategoryParams{
		ID:        id,
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if categoryRequest.Name != nil {
		updateParams.Name = *categoryRequest.Name
	} else {
		updateParams.Name = existingCategory.Name
	}

	if categoryRequest.Description != nil {
		updateParams.Description = pgtype.Text{String: *categoryRequest.Description, Valid: true}
	} else {
		updateParams.Description = existingCategory.Description
	}

	if categoryRequest.ParentID != nil {
		updateParams.ParentCategoryID = pgtype.Text{String: *categoryRequest.ParentID, Valid: true}
	} else {
		updateParams.ParentCategoryID = existingCategory.ParentCategoryID
	}

	updatedCategory, err := c.categoryRepo.UpdateCategory(ctx, updateParams)
	if err != nil {
		return db.Category{}, fmt.Errorf("failed to update category: %w", err)
	}

	return updatedCategory, nil
}

func (c *categoryUseCase) DeleteCategory(ctx context.Context, id string, storeID string) error {
	err := c.categoryRepo.DeleteCategory(ctx, id, storeID)
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}
	return nil
}

func (c *categoryUseCase) ListCategories(ctx context.Context) ([]db.Category, error) {
	categories, err := c.categoryRepo.ListCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list categories: %w", err)
	}
	return categories, nil
}

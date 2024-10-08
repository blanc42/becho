package usecase

import (
	"context"
	"fmt"
	"time"

	db "github.com/blanc42/becho/apps/admin/pkg/db/sqlc"
	"github.com/blanc42/becho/apps/admin/pkg/domain"
	"github.com/blanc42/becho/apps/admin/pkg/handlers/request"
	"github.com/blanc42/becho/apps/admin/pkg/handlers/response"
	"github.com/blanc42/becho/apps/admin/pkg/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type CategoryUsecase interface {
	CreateCategory(ctx context.Context, categoryRequest request.CreateCategoryRequest) (string, error)
	GetCategory(ctx context.Context, id string, storeID string) (db.GetCategoryRow, error)
	UpdateCategory(ctx context.Context, id string, storeID string, categoryRequest request.UpdateCategoryRequest) (db.Category, error)
	DeleteCategory(ctx context.Context, id string, storeID string) error
	ListCategories(ctx context.Context) ([]db.Category, error)
	GetAllCategoriesRecursive(ctx context.Context, storeID string) ([]*response.CategoryTreeNode, error)
}

type categoryUseCase struct {
	categoryRepo domain.CategoryRepository
	storeRepo    domain.StoreRepository
	variantRepo  domain.VariantRepository
}

func NewCategoryUseCase(categoryRepo domain.CategoryRepository, storeRepo domain.StoreRepository, variantRepo domain.VariantRepository) CategoryUsecase {
	return &categoryUseCase{
		categoryRepo: categoryRepo,
		storeRepo:    storeRepo,
		variantRepo:  variantRepo,
	}
}

func (c *categoryUseCase) CreateCategory(ctx context.Context, categoryRequest request.CreateCategoryRequest) (string, error) {

	// TODO: check if the parent category exists in the store and if it does, check if the new category has all the variants of the parent category
	_, err := c.storeRepo.GetStore(ctx, categoryRequest.StoreID)
	if err != nil {
		return "", fmt.Errorf("failed to get store: %w", err)
	}

	id, err := utils.GenerateShortID()
	if err != nil {
		return "", fmt.Errorf("failed to generate category ID: %w", err)
	}

	variants, err := c.variantRepo.GetVariantAndOptionsArrayForVariantIds(ctx, categoryRequest.Variants, categoryRequest.StoreID)
	if err != nil {
		return "", fmt.Errorf("failed to get variants: %w", err)
	}

	newCategory := db.CreateCategoryParams{
		ID:               id,
		CreatedAt:        pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt:        pgtype.Timestamp{Time: time.Now(), Valid: true},
		Name:             categoryRequest.Name,
		StoreID:          categoryRequest.StoreID,
		UniqueIdentifier: categoryRequest.UniqueIdentifier,
	}

	if categoryRequest.ParentID != nil {
		newCategory.ParentID = pgtype.Text{String: *categoryRequest.ParentID, Valid: true}
	}

	if categoryRequest.Description != nil {
		newCategory.Description = pgtype.Text{String: *categoryRequest.Description, Valid: true}
	}

	createdCategory, err := c.categoryRepo.CreateCategory(ctx, newCategory)
	if err != nil {
		return "", fmt.Errorf("failed to create category: %w", err)
	}

	for _, v := range variants {
		id, _ := utils.GenerateShortID()
		err := c.categoryRepo.CreateCategoryVariant(ctx, db.CreateCategoryVariantParams{
			ID:         id,
			CategoryID: createdCategory.ID,
			VariantID:  v.VariantID,
		})
		if err != nil {
			return "", fmt.Errorf("failed to create category variant: %w", err)
		}
	}

	return createdCategory.ID, nil
}

func (c *categoryUseCase) GetCategory(ctx context.Context, id string, storeID string) (db.GetCategoryRow, error) {
	category, err := c.categoryRepo.GetCategory(ctx, id, storeID)
	if err != nil {
		return db.GetCategoryRow{}, fmt.Errorf("failed to get category: %w", err)
	}
	return category, nil
}

func (c *categoryUseCase) UpdateCategory(ctx context.Context, id string, storeID string, categoryRequest request.UpdateCategoryRequest) (db.Category, error) {
	existingCategory, err := c.categoryRepo.GetCategory(ctx, id, storeID)
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

	// if categoryRequest.ParentID != nil {
	// 	updateParams.ParentCategoryID = pgtype.Text{String: *categoryRequest.ParentID, Valid: true}
	// } else {
	// 	updateParams.ParentCategoryID = existingCategory.ParentCategoryID
	// }

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

func (c *categoryUseCase) GetAllCategoriesRecursive(ctx context.Context, storeID string) ([]*response.CategoryTreeNode, error) {
	categories, err := c.categoryRepo.GetAllCategories(ctx, storeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get all categories recursively: %w", err)
	}
	if len(categories) == 0 {
		return []*response.CategoryTreeNode{}, nil
	}

	categoryMap := make(map[string][]*response.CategoryTreeNode)
	var rootCategories []*response.CategoryTreeNode

	for _, cat := range categories {
		variants := make([]response.VariantInsideCategoryResponse, 0)
		if cat.Variants != nil {
			if variantSlice, ok := cat.Variants.([]interface{}); ok {
				for _, v := range variantSlice {
					if variantMap, ok := v.(map[string]interface{}); ok {
						variant := response.VariantInsideCategoryResponse{
							ID:   variantMap["id"].(string),
							Name: variantMap["name"].(string),
						}
						variants = append(variants, variant)
					}
				}
			}
		}

		node := &response.CategoryTreeNode{
			ID:               cat.ID,
			Name:             cat.Name,
			Description:      cat.Description.String,
			Level:            cat.Level,
			UniqueIdentifier: cat.UniqueIdentifier,
			Variants:         variants,
		}
		if cat.ParentID.Valid {
			parentID := cat.ParentID.String
			node.ParentID = &parentID
		}

		if cat.ParentID.String == "" {
			rootCategories = append(rootCategories, node)
		} else {
			categoryMap[cat.ParentID.String] = append(categoryMap[cat.ParentID.String], node)
		}
	}

	var orderedCategories []*response.CategoryTreeNode
	for _, root := range rootCategories {
		appendCategoryAndChildren(root, &orderedCategories, categoryMap)
	}

	return orderedCategories, nil
}

func appendCategoryAndChildren(cat *response.CategoryTreeNode, orderedCategories *[]*response.CategoryTreeNode, categoryMap map[string][]*response.CategoryTreeNode) {
	*orderedCategories = append(*orderedCategories, cat)
	if children, exists := categoryMap[cat.ID]; exists {
		for _, child := range children {
			appendCategoryAndChildren(child, orderedCategories, categoryMap)
		}
	}
}

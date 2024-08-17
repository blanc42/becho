package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	db "github.com/blanc42/becho/apps/admin/pkg/db/sqlc"
	"github.com/blanc42/becho/apps/admin/pkg/domain"
	"github.com/blanc42/becho/apps/admin/pkg/handlers/request"
	"github.com/blanc42/becho/apps/admin/pkg/handlers/response"
	"github.com/blanc42/becho/apps/admin/pkg/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, product request.CreateProductRequest) (db.GetProductRow, error)
	GetProduct(ctx context.Context, id string, store_id string) (db.GetProductRow, error)
	UpdateProduct(ctx context.Context, id string, req request.UpdateProductRequest) (response.ProductResponse, error)
	GetProducts(ctx context.Context, arg db.GetFilteredProductsParams) ([]db.GetFilteredProductsRow, error)
}

type productUseCase struct {
	productRepo domain.ProductRepository
	storeRepo   domain.StoreRepository
	variantRepo domain.VariantRepository
}

func NewProductUseCase(productRepo domain.ProductRepository, storeRepo domain.StoreRepository, variantRepo domain.VariantRepository) ProductUsecase {
	return &productUseCase{
		productRepo: productRepo,
		storeRepo:   storeRepo,
		variantRepo: variantRepo,
	}
}

func (u *productUseCase) CreateProduct(ctx context.Context, product request.CreateProductRequest) (db.GetProductRow, error) {
	// // Check if the store exists
	// _, err := u.storeRepo.GetStore(ctx, req.StoreID)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to get store: %w", err)
	// }

	// id, err := utils.GenerateShortID()
	// if err != nil {
	// 	return "", fmt.Errorf("failed to generate product ID: %w", err)
	// }

	// variants, err := json.Marshal(req.VariantIDs)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to marshal variants: %w", err)
	// }

	// product, err := u.productRepo.CreateProduct(ctx, db.CreateProductParams{
	// 	ID:          id,
	// 	CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	// 	UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	// 	Name:        req.Name,
	// 	Description: pgtype.Text{String: req.Description, Valid: true},
	// 	CategoryID:  req.CategoryID,
	// 	StoreID:     req.StoreID,
	// 	Variants:    variants,
	// })

	// if err != nil {
	// 	return "", fmt.Errorf("failed to create product: %w", err)
	// }

	// return product.ID, nil

	// product := request.CreateProductRequest{}
	// err := ctx.ShouldBindJSON(&product)
	// if err != nil {
	// 	return db.GetProductRow{}, fmt.Errorf("failed to bind product: %w", err)
	// }

	// store_id, ok := ctx.Value("store_id").(string)
	// if !ok {
	// 	return db.GetProductRow{}, fmt.Errorf("failed to get store ID")
	// }

	variantOptionsArray, err := u.variantRepo.GetVariantAndOptionsArrayForVariantIds(ctx, product.Variants, product.StoreID)
	if err != nil {
		return db.GetProductRow{}, fmt.Errorf("failed to fetch variants and options: %w", err)
	}
	// utils.PrettyPrintJSON(variantOptionsArray)

	// Validate variant options
	variantOptionsMap := make(map[string]map[string]struct{})
	for _, variantOption := range variantOptionsArray {
		optionMap := make(map[string]struct{})
		var options []string
		err := json.Unmarshal(variantOption.VariantOptions, &options)
		if err != nil {
			return db.GetProductRow{}, fmt.Errorf("failed to unmarshal variant options: %w", err)
		}
		for _, option := range options {
			optionMap[option] = struct{}{}
		}
		variantOptionsMap[variantOption.VariantID] = optionMap
	}

	// Validate product items
	for _, item := range product.Items {
		for variantID, optionID := range item.VariantOptions {
			if _, ok := variantOptionsMap[variantID]; !ok {
				return db.GetProductRow{}, fmt.Errorf("invalid variant ID: %s", variantID)
			}
			if _, ok := variantOptionsMap[variantID][optionID]; !ok {
				return db.GetProductRow{}, fmt.Errorf("invalid variant option combination: %s - %s", variantID, optionID)
			}
		}
	}

	// Create product
	id, _ := utils.GenerateShortID()
	createdAt := pgtype.Timestamp{Time: time.Now(), Valid: true}
	updatedAt := pgtype.Timestamp{Time: time.Now(), Valid: true}

	variantsJSON, err := json.Marshal(product.Variants)
	if err != nil {
		return db.GetProductRow{}, fmt.Errorf("failed to marshal variants: %w", err)
	}

	newProduct, err := u.productRepo.CreateProduct(ctx, db.CreateProductParams{
		ID:          id,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Name:        product.Name,
		IsFeatured:  pgtype.Bool{Bool: product.IsFeatured, Valid: true},
		IsArchived:  pgtype.Bool{Bool: false, Valid: true},
		HasVariants: pgtype.Bool{Bool: product.HasVariants, Valid: true},
		CategoryID:  product.CategoryID,
		Variants:    variantsJSON,
		StoreID:     product.StoreID,
	})
	if err != nil {
		return db.GetProductRow{}, fmt.Errorf("failed to create product: %w", err)
	}

	if product.Description != nil {
		newProduct.Description = pgtype.Text{String: *product.Description, Valid: true}
	}

	// Create product variants
	for _, item := range product.Items {
		variantID, _ := utils.GenerateShortID()
		productVariant, err := u.productRepo.CreateProductItem(ctx, db.CreateProductItemParams{
			ID:              variantID,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
			ProductID:       newProduct.ID,
			Sku:             item.SKU,
			Quantity:        int32(item.Quantity),
			Price:           float64(item.Price),
			DiscountedPrice: pgtype.Float8{Float64: float64(item.DiscountedPrice), Valid: item.DiscountedPrice != 0},
			CostPrice:       pgtype.Float8{Float64: float64(item.CostPrice), Valid: item.CostPrice != 0},
		})
		if err != nil {
			return db.GetProductRow{}, fmt.Errorf("failed to create product variant: %w", err)
		}

		// Populate product_variant_options join table
		for _, optionID := range item.VariantOptions {
			_, err := u.productRepo.CreateProductVariantOption(ctx, db.CreateProductVariantOptionParams{
				ProductVariantID: productVariant.ID,
				VariantOptionID:  optionID,
			})
			if err != nil {
				return db.GetProductRow{}, fmt.Errorf("failed to create product variant option: %w", err)
			}
		}
	}

	// Fetch the created product
	createdProduct, err := u.productRepo.GetProduct(ctx, db.GetProductParams{
		ID:      newProduct.ID,
		StoreID: product.StoreID,
	})
	if err != nil {
		return db.GetProductRow{}, fmt.Errorf("failed to fetch created product: %w", err)
	}

	return createdProduct, nil

}

func (u *productUseCase) GetProduct(ctx context.Context, id string, store_id string) (db.GetProductRow, error) {
	product, err := u.productRepo.GetProduct(ctx, db.GetProductParams{
		ID:      id,
		StoreID: store_id,
	})
	if err != nil {
		return db.GetProductRow{}, fmt.Errorf("failed to get product: %w", err)
	}

	return product, nil
	// var variants []response.Variant
	// err = json.Unmarshal(product.Variants, &variants)
	// if err != nil {
	// 	return response.ProductResponse{}, fmt.Errorf("failed to unmarshal variants: %w", err)
	// }

	// return response.ProductResponse{
	// 	ID:          product.ID,
	// 	Name:        product.Name,
	// 	Description: product.Description.String,
	// 	CategoryID:  product.CategoryID,
	// 	StoreID:     product.StoreID,
	// 	Variants:    variants,
	// }, nil
}

func (u *productUseCase) UpdateProduct(ctx context.Context, id string, req request.UpdateProductRequest) (response.ProductResponse, error) {
	store_id, ok := ctx.Value("store_id").(string)
	if !ok {
		return response.ProductResponse{}, fmt.Errorf("failed to get store ID")
	}
	existingProduct, err := u.productRepo.GetProduct(ctx, db.GetProductParams{
		ID:      id,
		StoreID: store_id,
	})
	if err != nil {
		return response.ProductResponse{}, fmt.Errorf("failed to get product: %w", err)
	}

	updateParams := db.UpdateProductParams{
		ID:        id,
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		StoreID:   existingProduct.StoreID,
	}

	if req.Name != nil {
		updateParams.Name = *req.Name
	} else {
		updateParams.Name = existingProduct.ProductName
	}

	if req.Description != nil {
		updateParams.Description = pgtype.Text{String: *req.Description, Valid: true}
	} else {
		updateParams.Description = existingProduct.ProductDescription
	}

	if req.CategoryID != nil {
		updateParams.CategoryID = *req.CategoryID
	} else {
		updateParams.CategoryID = existingProduct.CategoryID
	}

	if req.VariantIDs != nil {
		variants, err := json.Marshal(req.VariantIDs)
		if err != nil {
			return response.ProductResponse{}, fmt.Errorf("failed to marshal variants: %w", err)
		}
		updateParams.Variants = variants
	} else {
		updateParams.Variants = existingProduct.VariantsOrder
	}

	updatedProduct, err := u.productRepo.UpdateProduct(ctx, updateParams)
	if err != nil {
		return response.ProductResponse{}, fmt.Errorf("failed to update product: %w", err)
	}

	var variants []response.Variant
	err = json.Unmarshal(updatedProduct.Variants, &variants)
	if err != nil {
		return response.ProductResponse{}, fmt.Errorf("failed to unmarshal variants: %w", err)
	}

	return response.ProductResponse{
		ID:          updatedProduct.ID,
		Name:        updatedProduct.Name,
		Description: updatedProduct.Description.String,
		CategoryID:  updatedProduct.CategoryID,
		StoreID:     updatedProduct.StoreID,
		Variants:    variants,
	}, nil
}

func (u *productUseCase) GetProducts(ctx context.Context, arg db.GetFilteredProductsParams) ([]db.GetFilteredProductsRow, error) {
	// products, err := u.productRepo.GetProducts(ctx, db.GetFilteredProductsParams{
	// 	StoreID: pgtype.Text{String: storeID, Valid: true},
	// 	Limit:   pgtype.Int4{Int32: limit, Valid: true},
	// 	Offset:  pgtype.Int4{Int32: offset, Valid: true},
	// })
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get products: %w", err)
	// }

	// var productResponses []response.ProductResponse
	// for _, product := range products {
	// 	var variants []response.Variant
	// 	err = json.Unmarshal(product.Variants, &variants)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to unmarshal variants: %w", err)
	// 	}

	// 	productResponses = append(productResponses, response.ProductResponse{
	// 		ID:          product.ID,
	// 		Name:        product.Name,
	// 		Description: product.Description.String,
	// 		CategoryID:  product.CategoryID,
	// 		StoreID:     product.StoreID,
	// 		Variants:    variants,
	// 	})
	// }
	products, err := u.productRepo.GetProducts(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	if len(products) == 0 {
		return []db.GetFilteredProductsRow{}, nil
	}

	return products, nil
}

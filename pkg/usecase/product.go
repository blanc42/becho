package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	db "github.com/blanc42/becho/pkg/db/sqlc"
	"github.com/blanc42/becho/pkg/domain"
	"github.com/blanc42/becho/pkg/handlers/request"
	"github.com/blanc42/becho/pkg/handlers/response"
	"github.com/blanc42/becho/pkg/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, req request.CreateProductRequest) (string, error)
	GetProduct(ctx context.Context, id string, store_id string) (db.GetProductRow, error)
	UpdateProduct(ctx context.Context, id string, req request.UpdateProductRequest) (response.ProductResponse, error)
	GetProducts(ctx context.Context, arg db.GetFilteredProductsParams) ([]db.GetFilteredProductsRow, error)
}

type productUseCase struct {
	productRepo domain.ProductRepository
	storeRepo   domain.StoreRepository
}

func NewProductUseCase(productRepo domain.ProductRepository, storeRepo domain.StoreRepository) ProductUsecase {
	return &productUseCase{
		productRepo: productRepo,
		storeRepo:   storeRepo,
	}
}

func (u *productUseCase) CreateProduct(ctx context.Context, req request.CreateProductRequest) (string, error) {
	// Check if the store exists
	_, err := u.storeRepo.GetStore(ctx, req.StoreID)
	if err != nil {
		return "", fmt.Errorf("failed to get store: %w", err)
	}

	id, err := utils.GenerateShortID()
	if err != nil {
		return "", fmt.Errorf("failed to generate product ID: %w", err)
	}

	variants, err := json.Marshal(req.VariantIDs)
	if err != nil {
		return "", fmt.Errorf("failed to marshal variants: %w", err)
	}

	product, err := u.productRepo.CreateProduct(ctx, db.CreateProductParams{
		ID:          id,
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
		Name:        req.Name,
		Description: pgtype.Text{String: req.Description, Valid: true},
		CategoryID:  req.CategoryID,
		StoreID:     req.StoreID,
		Variants:    variants,
	})

	if err != nil {
		return "", fmt.Errorf("failed to create product: %w", err)
	}

	return product.ID, nil
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

	return products, nil
}

package usecase

import (
	"context"
	"errors"

	db "github.com/blanc42/becho/apps/admin/pkg/db/sqlc"
	"github.com/blanc42/becho/apps/admin/pkg/domain"
	"github.com/blanc42/becho/apps/admin/pkg/handlers/request"
	"github.com/blanc42/becho/apps/admin/pkg/handlers/response"
	"github.com/blanc42/becho/apps/admin/pkg/utils"
)

type CartUsecase interface {
	AddToCart(ctx context.Context, req request.AddToCartRequest) (response.CartResponse, error)
	RemoveFromCart(ctx context.Context, cartItemID string) error
	UpdateCartItem(ctx context.Context, req request.UpdateCartItemRequest) (response.CartResponse, error)
	GetCart(ctx context.Context, userID string) (response.CartResponse, error)
}

type cartUseCase struct {
	cartRepo     domain.CartRepository
	cartItemRepo domain.CartItemRepository
	productRepo  domain.ProductRepository
}

func NewCartUseCase(cartRepo domain.CartRepository, cartItemRepo domain.CartItemRepository, productRepo domain.ProductRepository) CartUsecase {
	return &cartUseCase{
		cartRepo:     cartRepo,
		cartItemRepo: cartItemRepo,
		productRepo:  productRepo,
	}
}

func (c *cartUseCase) AddToCart(ctx context.Context, req request.AddToCartRequest) (response.CartResponse, error) {
	cart, err := c.cartRepo.GetCart(ctx, req.UserID)
	if err != nil {
		// If cart doesn't exist, create a new one
		cart, err = c.cartRepo.CreateCart(ctx, db.CreateCartParams{
			ID:     utils.GenerateShortID(),
			UserID: req.UserID,
		})
		if err != nil {
			return response.CartResponse{}, err
		}
	}

	// Check if the product variant exists
	_, err = c.productRepo.GetProductVariant(ctx, req.ProductVariantID)
	if err != nil {
		return response.CartResponse{}, errors.New("product variant not found")
	}

	id, _ := utils.GenerateShortID()

	_, err = c.cartItemRepo.CreateCartItem(ctx, db.CreateCartItemParams{
		ID:               id,
		CartID:           cart.ID,
		ProductVariantID: req.ProductVariantID,
		Quantity:         req.Quantity,
	})
	if err != nil {
		return response.CartResponse{}, err
	}

	return c.getCartResponse(ctx, cart.ID)
}

func (c *cartUseCase) RemoveFromCart(ctx context.Context, cartItemID string) error {
	return c.cartItemRepo.DeleteCartItem(ctx, cartItemID)
}

func (c *cartUseCase) UpdateCartItem(ctx context.Context, req request.UpdateCartItemRequest) (response.CartResponse, error) {
	_, err := c.cartItemRepo.GetCartItem(ctx, req.CartItemID)
	if err != nil {
		return response.CartResponse{}, errors.New("cart item not found")
	}

	updatedCartItem, err := c.cartItemRepo.UpdateCartItem(ctx, db.UpdateCartItemParams{
		ID:       req.CartItemID,
		Quantity: req.Quantity,
	})
	if err != nil {
		return response.CartResponse{}, err
	}

	return c.getCartResponse(ctx, updatedCartItem.CartID)
}

func (c *cartUseCase) GetCart(ctx context.Context, userID string) (response.CartResponse, error) {
	cart, err := c.cartRepo.GetCart(ctx, userID)
	if err != nil {
		return response.CartResponse{}, err
	}

	return c.getCartResponse(ctx, cart.ID)
}

func (c *cartUseCase) getCartResponse(ctx context.Context, cartID string) (response.CartResponse, error) {
	cartItems, err := c.cartItemRepo.ListCartItems(ctx, db.ListCartItemsParams{
		CartID: cartID,
	})
	if err != nil {
		return response.CartResponse{}, err
	}

	var totalPrice float64
	var items []response.CartItemResponse

	for _, item := range cartItems {
		variant, err := c.productRepo.GetProductVariant(ctx, item.ProductVariantID)
		if err != nil {
			return response.CartResponse{}, err
		}

		itemPrice := float64(item.Quantity) * variant.Price
		totalPrice += itemPrice

		items = append(items, response.CartItemResponse{
			ID:               item.ID,
			ProductVariantID: item.ProductVariantID,
			Quantity:         item.Quantity,
			Price:            variant.Price,
			TotalPrice:       itemPrice,
		})
	}

	return response.CartResponse{
		ID:         cartID,
		Items:      items,
		TotalPrice: totalPrice,
	}, nil
}

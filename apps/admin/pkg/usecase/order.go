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

type OrderUsecase interface {
	CreateOrder(ctx context.Context, req request.CreateOrderRequest) (response.OrderResponse, error)
	GetOrder(ctx context.Context, orderID string) (response.OrderResponse, error)
	UpdateOrderStatus(ctx context.Context, req request.UpdateOrderStatusRequest) (response.OrderResponse, error)
	ListOrders(ctx context.Context, userID string, limit, offset int32) ([]response.OrderResponse, error)
}

type orderUseCase struct {
	orderRepo     domain.OrderRepository
	orderItemRepo domain.OrderItemRepository
	cartRepo      domain.CartRepository
	productRepo   domain.ProductRepository
}

func NewOrderUseCase(orderRepo domain.OrderRepository, orderItemRepo domain.OrderItemRepository, cartRepo domain.CartRepository, productRepo domain.ProductRepository) OrderUsecase {
	return &orderUseCase{
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
		cartRepo:      cartRepo,
		productRepo:   productRepo,
	}
}

func (o *orderUseCase) CreateOrder(ctx context.Context, req request.CreateOrderRequest) (response.OrderResponse, error) {
	// Get user's cart
	cart, err := o.cartRepo.GetCart(ctx, req.UserID)
	if err != nil {
		return response.OrderResponse{}, errors.New("cart not found")
	}
	id, _ := utils.GenerateShortID()
	orderNumber, _ := utils.GenerateOrderNumber()
	// Create order
	order, err := o.orderRepo.CreateOrder(ctx, db.CreateOrderParams{
		ID:            id,
		CustomerID:    req.UserID,
		OrderNumber:   orderNumber,
		PaymentStatus: "pending",
		OrderStatus:   "created",
	})
	if err != nil {
		return response.OrderResponse{}, err
	}

	// Get cart items and create order items
	cartItems, err := o.cartRepo.ListCartItems(ctx, db.ListCartItemsParams{CartID: cart.ID})
	if err != nil {
		return response.OrderResponse{}, err
	}

	var totalAmount float64
	for _, cartItem := range cartItems {
		variant, err := o.productRepo.GetProductVariant(ctx, cartItem.ProductVariantID)
		if err != nil {
			return response.OrderResponse{}, err
		}

		_, err = o.orderItemRepo.CreateOrderItem(ctx, db.CreateOrderItemParams{
			ID:            utils.GenerateUUID(),
			OrderID:       order.ID,
			ProductItemID: cartItem.ProductVariantID,
			Quantity:      cartItem.Quantity,
		})
		if err != nil {
			return response.OrderResponse{}, err
		}

		totalAmount += float64(cartItem.Quantity) * variant.Price
	}

	// Clear the cart
	err = o.cartRepo.DeleteCart(ctx, cart.ID)
	if err != nil {
		return response.OrderResponse{}, err
	}

	return o.getOrderResponse(ctx, order.ID)
}

func (o *orderUseCase) GetOrder(ctx context.Context, orderID string) (response.OrderResponse, error) {
	return o.getOrderResponse(ctx, orderID)
}

func (o *orderUseCase) UpdateOrderStatus(ctx context.Context, req request.UpdateOrderStatusRequest) (response.OrderResponse, error) {
	order, err := o.orderRepo.GetOrder(ctx, req.OrderID)
	if err != nil {
		return response.OrderResponse{}, errors.New("order not found")
	}

	updatedOrder, err := o.orderRepo.UpdateOrder(ctx, db.UpdateOrderParams{
		ID:          req.OrderID,
		OrderStatus: req.Status,
	})
	if err != nil {
		return response.OrderResponse{}, err
	}

	return o.getOrderResponse(ctx, updatedOrder.ID)
}

func (o *orderUseCase) ListOrders(ctx context.Context, userID string, limit, offset int32) ([]response.OrderResponse, error) {
	orders, err := o.orderRepo.ListOrders(ctx, db.ListOrdersParams{
		CustomerID: userID,
		Limit:      limit,
		Offset:     offset,
	})
	if err != nil {
		return nil, err
	}

	var orderResponses []response.OrderResponse
	for _, order := range orders {
		orderResponse, err := o.getOrderResponse(ctx, order.ID)
		if err != nil {
			return nil, err
		}
		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses, nil
}

func (o *orderUseCase) getOrderResponse(ctx context.Context, orderID string) (response.OrderResponse, error) {
	order, err := o.orderRepo.GetOrder(ctx, orderID)
	if err != nil {
		return response.OrderResponse{}, err
	}

	orderItems, err := o.orderItemRepo.ListOrderItems(ctx, db.ListOrderItemsParams{
		OrderID: orderID,
	})
	if err != nil {
		return response.OrderResponse{}, err
	}

	var totalAmount float64
	var items []response.OrderItemResponse

	for _, item := range orderItems {
		variant, err := o.productRepo.GetProductVariant(ctx, item.ProductItemID)
		if err != nil {
			return response.OrderResponse{}, err
		}

		itemPrice := float64(item.Quantity) * variant.Price
		totalAmount += itemPrice

		items = append(items, response.OrderItemResponse{
			ID:               item.ID,
			ProductVariantID: item.ProductItemID,
			Quantity:         item.Quantity,
			Price:            variant.Price,
			TotalPrice:       itemPrice,
		})
	}

	return response.OrderResponse{
		ID:            order.ID,
		OrderNumber:   order.OrderNumber,
		CustomerID:    order.CustomerID,
		Items:         items,
		TotalAmount:   totalAmount,
		PaymentStatus: order.PaymentStatus,
		OrderStatus:   order.OrderStatus,
		CreatedAt:     order.CreatedAt.Time,
		UpdatedAt:     order.UpdatedAt.Time,
	}, nil
}

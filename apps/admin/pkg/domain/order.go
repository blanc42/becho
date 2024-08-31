package domain

import (
	"context"

	db "github.com/blanc42/becho/apps/admin/pkg/db/sqlc"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, arg db.CreateOrderParams) (db.Order, error)
	GetOrder(ctx context.Context, id string) (db.Order, error)
	UpdateOrder(ctx context.Context, arg db.UpdateOrderParams) (db.Order, error)
	DeleteOrder(ctx context.Context, id string) error
	ListOrders(ctx context.Context, arg db.ListOrdersParams) ([]db.Order, error)
}

type OrderItemRepository interface {
	CreateOrderItem(ctx context.Context, arg db.CreateOrderItemParams) (db.OrderItem, error)
	GetOrderItem(ctx context.Context, id string) (db.OrderItem, error)
	UpdateOrderItem(ctx context.Context, arg db.UpdateOrderItemParams) (db.OrderItem, error)
	DeleteOrderItem(ctx context.Context, id string) error
	ListOrderItems(ctx context.Context, arg db.ListOrderItemsParams) ([]db.OrderItem, error)
}

type orderRepository struct {
	db *db.DbStore
}

func NewOrderRepository(db *db.DbStore) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(ctx context.Context, arg db.CreateOrderParams) (db.Order, error) {
	return r.db.CreateOrder(ctx, arg)
}

func (r *orderRepository) GetOrder(ctx context.Context, id string) (db.Order, error) {
	return r.db.GetOrder(ctx, id)
}

func (r *orderRepository) UpdateOrder(ctx context.Context, arg db.UpdateOrderParams) (db.Order, error) {
	return r.db.UpdateOrder(ctx, arg)
}

func (r *orderRepository) DeleteOrder(ctx context.Context, id string) error {
	return r.db.DeleteOrder(ctx, id)
}

func (r *orderRepository) ListOrders(ctx context.Context, arg db.ListOrdersParams) ([]db.Order, error) {
	return r.db.ListOrders(ctx, arg)
}

type orderItemRepository struct {
	db *db.DbStore
}

func NewOrderItemRepository(db *db.DbStore) OrderItemRepository {
	return &orderItemRepository{db: db}
}

func (r *orderItemRepository) CreateOrderItem(ctx context.Context, arg db.CreateOrderItemParams) (db.OrderItem, error) {
	return r.db.CreateOrderItem(ctx, arg)
}

func (r *orderItemRepository) GetOrderItem(ctx context.Context, id string) (db.OrderItem, error) {
	return r.db.GetOrderItem(ctx, id)
}

func (r *orderItemRepository) UpdateOrderItem(ctx context.Context, arg db.UpdateOrderItemParams) (db.OrderItem, error) {
	return r.db.UpdateOrderItem(ctx, arg)
}

func (r *orderItemRepository) DeleteOrderItem(ctx context.Context, id string) error {
	return r.db.DeleteOrderItem(ctx, id)
}

func (r *orderItemRepository) ListOrderItems(ctx context.Context, arg db.ListOrderItemsParams) ([]db.OrderItem, error) {
	return r.db.ListOrderItems(ctx, arg)
}

package request

type CreateOrderRequest struct {
	UserID    string `json:"user_id" binding:"required"`
	AddressID string `json:"address_id" binding:"required"`
}

type UpdateOrderStatusRequest struct {
	OrderID string `json:"order_id" binding:"required"`
	Status  string `json:"status" binding:"required"`
}

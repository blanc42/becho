package response

import "time"

type OrderItemResponse struct {
	ID               string  `json:"id"`
	ProductVariantID string  `json:"product_variant_id"`
	ProductName      string  `json:"product_name"`
	Quantity         int32   `json:"quantity"`
	Price            float64 `json:"price"`
	TotalPrice       float64 `json:"total_price"`
}

type OrderResponse struct {
	ID            string              `json:"id"`
	UserID        string              `json:"user_id"`
	OrderNumber   string              `json:"order_number"`
	Items         []OrderItemResponse `json:"items"`
	TotalPrice    float64             `json:"total_price"`
	Status        string              `json:"status"`
	PaymentStatus string              `json:"payment_status"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
}

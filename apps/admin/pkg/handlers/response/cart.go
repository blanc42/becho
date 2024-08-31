package response

type CartItemResponse struct {
	ID               string  `json:"id"`
	ProductVariantID string  `json:"product_variant_id"`
	ProductName      string  `json:"product_name"`
	Quantity         int32   `json:"quantity"`
	Price            float64 `json:"price"`
	TotalPrice       float64 `json:"total_price"`
}

type CartResponse struct {
	ID         string             `json:"id"`
	UserID     string             `json:"user_id"`
	Items      []CartItemResponse `json:"items"`
	TotalPrice float64            `json:"total_price"`
}

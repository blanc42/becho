package request

type AddToCartRequest struct {
	UserID           string `json:"user_id" binding:"required"`
	ProductVariantID string `json:"product_variant_id" binding:"required"`
	Quantity         int32  `json:"quantity" binding:"required,min=1"`
}

type UpdateCartItemRequest struct {
	CartItemID string `json:"cart_item_id" binding:"required"`
	Quantity   int32  `json:"quantity" binding:"required,min=1"`
}

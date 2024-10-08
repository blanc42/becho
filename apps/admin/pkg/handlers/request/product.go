package request

type CreateProductRequest struct {
	Name        string                        `json:"name" binding:"required"`
	Description *string                       `json:"description"`
	IsFeatured  bool                          `json:"is_featured"`                    // has to remove the binding since gin binding is causing troubles
	IsArchived  bool                          `json:"is_archived"`                    // has to remove the binding since gin binding is causing troubles
	CategoryID  string                        `json:"category_id" binding:"required"` // TODO: make the category optional
	StoreID     string                        `json:"store_id"`
	Variants    []string                      `json:"variants"`
	Items       []CreateProductVariantRequest `json:"items"`
}

type CreateProductVariantRequest struct {
	SKU             string            `json:"sku" binding:"required"`
	Quantity        int64             `json:"quantity" binding:"required"`
	Price           int64             `json:"price" binding:"required"`
	CostPrice       int64             `json:"cost_price" binding:"required"`
	DiscountedPrice int64             `json:"discounted_price" binding:"required"`
	VariantOptions  map[string]string `json:"variant_options"`
	Images          []string          `json:"images"`
}

type UpdateProductRequest struct {
	Name        *string   `json:"name,omitempty" binding:"omitempty,min=3,max=100"`
	Description *string   `json:"description,omitempty"`
	Price       *float64  `json:"price,omitempty" binding:"omitempty,min=0"`
	CategoryID  *string   `json:"category_id,omitempty"`
	VariantIDs  *[]string `json:"variant_ids,omitempty"`
}

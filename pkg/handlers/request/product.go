package request

type CreateProductRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Price       float64  `json:"price" binding:"required,min=0"`
	CategoryID  string   `json:"category_id" binding:"required"`
	StoreID     string   `json:"store_id" binding:"required"`
	VariantIDs  []string `json:"variant_ids"`
}

type UpdateProductRequest struct {
	Name        *string   `json:"name,omitempty" binding:"omitempty,min=3,max=100"`
	Description *string   `json:"description,omitempty"`
	Price       *float64  `json:"price,omitempty" binding:"omitempty,min=0"`
	CategoryID  *string   `json:"category_id,omitempty"`
	VariantIDs  *[]string `json:"variant_ids,omitempty"`
}

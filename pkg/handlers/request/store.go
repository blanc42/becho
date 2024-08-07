package request

type CreateStoreRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateStoreRequest struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,min=3,max=100"`
	Description *string `json:"description,omitempty"`
}

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ParentID    string `json:"parent_id,omitempty"`
	StoreID     string `json:"store_id" binding:"required"`
}

type UpdateCategoryRequest struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,min=3,max=100"`
	Description *string `json:"description,omitempty"`
	ParentID    *string `json:"parent_id,omitempty"`
}

type CreateVariantRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Options     []string `json:"options" binding:"required"`
	StoreID     string   `json:"store_id" binding:"required"`
}

type UpdateVariantRequest struct {
	Name        *string   `json:"name,omitempty" binding:"omitempty,min=3,max=100"`
	Description *string   `json:"description,omitempty"`
	Options     *[]string `json:"options,omitempty"`
}

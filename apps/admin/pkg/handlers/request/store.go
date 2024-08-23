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
	Name             string   `json:"name" binding:"required"`
	Description      *string  `json:"description"`
	ParentID         *string  `json:"parent_id,omitempty"`
	StoreID          string   `json:"store_id"`
	UniqueIdentifier string   `json:"unique_identifier" binding:"required"`
	Variants         []string `json:"variants,omitempty"`
}

type UpdateCategoryRequest struct {
	Name        *string `json:"name,omitempty" binding:"omitempty,min=3,max=100"`
	Description *string `json:"description,omitempty"`
	ParentID    *string `json:"parent_id,omitempty"`
}

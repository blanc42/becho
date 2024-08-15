package request

type CreateVariantRequest struct {
	Name        string                       `json:"name"`
	Description *string                      `json:"description"`
	StoreID     string                       `json:"store_id"`
	Options     []CreateVariantOptionRequest `json:"options"`
}

type CreateVariantOptionRequest struct {
	Value        string  `json:"value" binding:"required"`
	DisplayOrder int32   `json:"display_order" binding:"required"`
	ImageId      *string `json:"image_id" binding:"len=11"`
	Data         *string `json:"data"`
}

type UpdateVariantRequest struct {
	ID          string                       `json:"id" binding:"required,len=11"`
	Name        *string                      `json:"name"`
	Description *string                      `json:"description"`
	Options     []UpdateVariantOptionRequest `json:"options"`
	StoreID     string                       `json:"store_id" binding:"required,len=11"`
}

type UpdateVariantOptionRequest struct {
	ID           *string `json:"id"`
	Value        *string `json:"value"`
	DisplayOrder *int32  `json:"display_order"`
	ImageId      *string `json:"image_id" binding:"omitempty,len=11"`
	Data         *string `json:"data"`
}

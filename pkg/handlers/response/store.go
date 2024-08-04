package response

import db "github.com/blanc42/becho/pkg/db/sqlc"

type StoreResponse struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Categories  []CategoryResponse `json:"categories"`
}

type CategoryResponse struct {
	ID            string              `json:"id"`
	Name          string              `json:"name"`
	Description   string              `json:"description"`
	SubCategories []*CategoryResponse `json:"sub_categories"`
	Variants      []*db.Variant       `json:"variants"`
}

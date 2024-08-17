package response

import (
	"time"

	db "github.com/blanc42/becho/apps/admin/pkg/db/sqlc"
)

type CreateVariantResponse struct {
	ID          string             `json:"id"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	Name        string             `json:"name"`
	StoreID     string             `json:"store_id"`
	Description *string            `json:"description"`
	Options     []db.VariantOption `json:"options"`
}

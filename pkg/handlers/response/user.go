package response

import db "github.com/blanc42/becho/pkg/db/sqlc"

type AdminLoginResponse struct {
	ID       string     `json:"id"`       // User ID
	Username string     `json:"username"` // User name
	Email    string     `json:"email"`    // User email
	Role     string     `json:"role"`     // User role
	Stores   []db.Store `json:"stores"`   // User stores
}

type CustomerLoginResponse struct {
	ID       string `json:"id"`       // User ID
	Username string `json:"username"` // User name
	Email    string `json:"email"`    // User email
	Role     string `json:"role"`     // User role
}

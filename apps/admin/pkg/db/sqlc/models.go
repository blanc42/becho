// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

// Enum for user roles: customer, owner, or admin
type UserRole string

const (
	UserRoleCustomer UserRole = "customer"
	UserRoleOwner    UserRole = "owner"
	UserRoleAdmin    UserRole = "admin"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole `json:"user_role"`
	Valid    bool     `json:"valid"` // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type Address struct {
	ID           string           `json:"id"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	AddressLine1 string           `json:"address_line_1"`
	AddressLine2 pgtype.Text      `json:"address_line_2"`
	City         string           `json:"city"`
	Pincode      string           `json:"pincode"`
	CountryID    string           `json:"country_id"`
}

type Cart struct {
	ID        string           `json:"id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	UserID    string           `json:"user_id"`
}

type CartItem struct {
	ID               string           `json:"id"`
	CreatedAt        pgtype.Timestamp `json:"created_at"`
	UpdatedAt        pgtype.Timestamp `json:"updated_at"`
	ProductVariantID string           `json:"product_variant_id"`
	Quantity         int32            `json:"quantity"`
	CartID           string           `json:"cart_id"`
	StoreID          string           `json:"store_id"`
}

type Category struct {
	ID               string           `json:"id"`
	CreatedAt        pgtype.Timestamp `json:"created_at"`
	UpdatedAt        pgtype.Timestamp `json:"updated_at"`
	Name             string           `json:"name"`
	Description      pgtype.Text      `json:"description"`
	StoreID          string           `json:"store_id"`
	ParentID         pgtype.Text      `json:"parent_id"`
	Level            int32            `json:"level"`
	UniqueIdentifier string           `json:"unique_identifier"`
	ImageID          pgtype.Int4      `json:"image_id"`
}

type CategoryVariant struct {
	ID         string `json:"id"`
	CategoryID string `json:"category_id"`
	VariantID  string `json:"variant_id"`
}

type Country struct {
	ID        string           `json:"id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	Country   string           `json:"country"`
}

type Image struct {
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	Title     pgtype.Text      `json:"title"`
	ImageID   string           `json:"image_id"`
	ID        pgtype.Int4      `json:"id"`
}

type Order struct {
	ID            string           `json:"id"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
	OrderNumber   string           `json:"order_number"`
	PaymentStatus string           `json:"payment_status"`
	OrderStatus   string           `json:"order_status"`
	StoreID       string           `json:"store_id"`
	CustomerID    string           `json:"customer_id"`
}

type OrderItem struct {
	ID            string           `json:"id"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
	ProductItemID string           `json:"product_item_id"`
	Quantity      int32            `json:"quantity"`
	OrderID       string           `json:"order_id"`
}

type Product struct {
	ID              string           `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
	Name            string           `json:"name"`
	Description     pgtype.Text      `json:"description"`
	Rating          pgtype.Float8    `json:"rating"`
	IsFeatured      pgtype.Bool      `json:"is_featured"`
	IsArchived      pgtype.Bool      `json:"is_archived"`
	HasVariants     pgtype.Bool      `json:"has_variants"`
	CategoryID      string           `json:"category_id"`
	StoreID         string           `json:"store_id"`
	Variants        json.RawMessage  `json:"variants"`
	Brand           pgtype.Text      `json:"brand"`
	NumberOfRatings pgtype.Int4      `json:"number_of_ratings"`
}

type ProductVariant struct {
	ID              string           `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
	ProductID       string           `json:"product_id"`
	Sku             string           `json:"sku"`
	Quantity        int32            `json:"quantity"`
	Price           float64          `json:"price"`
	DiscountedPrice pgtype.Float8    `json:"discounted_price"`
	CostPrice       pgtype.Float8    `json:"cost_price"`
	Title           pgtype.Text      `json:"title"`
}

type ProductVariantImage struct {
	ProductVariantID string      `json:"product_variant_id"`
	ID               int32       `json:"id"`
	ImageID          pgtype.Int4 `json:"image_id"`
	DisplayOrder     int32       `json:"display_order"`
}

type ProductVariantOption struct {
	ProductVariantID string `json:"product_variant_id"`
	VariantOptionID  string `json:"variant_option_id"`
}

type Store struct {
	ID          string           `json:"id"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	Name        string           `json:"name"`
	Description pgtype.Text      `json:"description"`
	UserID      string           `json:"user_id"`
	Logo        pgtype.Text      `json:"logo"`
	ImageID     pgtype.Int4      `json:"image_id"`
}

type User struct {
	ID        string           `json:"id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	Role      UserRole         `json:"role"`
	StoreID   pgtype.Text      `json:"store_id"`
	ImageID   pgtype.Int4      `json:"image_id"`
}

type Variant struct {
	ID          string           `json:"id"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	Name        string           `json:"name"`
	Label       string           `json:"label"`
	Description pgtype.Text      `json:"description"`
	StoreID     string           `json:"store_id"`
}

type VariantOption struct {
	ID           string           `json:"id"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	VariantID    string           `json:"variant_id"`
	Value        string           `json:"value"`
	DisplayOrder int32            `json:"display_order"`
	Data         pgtype.Text      `json:"data"`
	ImageID      pgtype.Int4      `json:"image_id"`
}

type Wishlist struct {
	ID               string           `json:"id"`
	UserID           string           `json:"user_id"`
	ProductVariantID string           `json:"product_variant_id"`
	CreatedAt        pgtype.Timestamp `json:"created_at"`
	UpdatedAt        pgtype.Timestamp `json:"updated_at"`
}

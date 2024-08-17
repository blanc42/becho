package response

type ProductResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CategoryID  string    `json:"category_id"`
	StoreID     string    `json:"store_id"`
	Variants    []Variant `json:"variants"`
}

type Variant struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Options []string `json:"options"`
}

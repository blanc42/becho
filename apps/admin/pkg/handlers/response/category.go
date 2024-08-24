package response

type CategoryTreeNode struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Level            int32    `json:"level"`
	ParentID         *string  `json:"parent_id"`
	UniqueIdentifier string   `json:"unique_identifier"`
	Variants         []string `json:"variants"`
	// Children         []*CategoryTreeNode `json:"children,omitempty"`
}

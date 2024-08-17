package response

type CategoryTreeNode struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Level       int32               `json:"level"`
	Children    []*CategoryTreeNode `json:"children,omitempty"`
}

package requests

type UpdateFileTypeRequest struct {
	Type        *string `json:"type"`
	Description *string `json:"description"`
}

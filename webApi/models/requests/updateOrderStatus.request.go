package requests

type UpdateOrderStatusRequest struct {
	Type        *string `json:"type"`
	Description *string `json:"description"`
}

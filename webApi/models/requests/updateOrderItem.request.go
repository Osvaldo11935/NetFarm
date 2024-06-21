package requests

type UpdateOrderItemRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Quantity    *int    `json:"quantity"`
}

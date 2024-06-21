package requests

type UpdateImposedRequest struct {
	Name *string  `json:"name"`
	Rate *float64 `json:"rate"`
}

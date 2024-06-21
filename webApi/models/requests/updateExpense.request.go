package requests

type UpdateExpenseRequest struct {
	Name *string  `json:"name"`
	Rate *float64 `json:"rate"`
}

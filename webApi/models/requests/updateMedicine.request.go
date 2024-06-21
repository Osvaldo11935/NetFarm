package requests

type UpdateMedicineRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Quantity    *int     `json:"quantity"`
	Price       *float64 `json:"price"`
}

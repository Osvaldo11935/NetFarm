package responses

type OrderItemDetailResponse struct {
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
	MedicineName string  `json:"medicineName"`
	MedicineId   string  `json:"medicineId"`
}

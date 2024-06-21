package responses

type OrderDetailResponse struct {
	AmountPaid     float64                   `gorm:"column:AmountPaid" json:"amountPaid"`
	TaxAmount      float64                   `gorm:"column:TaxAmount" json:"taxAmount"`
	DeliveryAmount float64                   `gorm:"column:DeliveryAmount" json:"deliveryAmount"`
	ImposedId      string                    `gorm:"column:ImposedId" json:"imposedId"`
	ExpenseId      string                    `gorm:"column:ExpenseId" json:"expenseId"`
	OrderItems     []OrderItemDetailResponse `json:"orderItems"`
}

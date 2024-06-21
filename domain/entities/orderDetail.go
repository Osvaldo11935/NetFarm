package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderDetail struct {
	Id             string  `gorm:"column:Id;primaryKey" json:"id"`
	AmountPaid     float64 `gorm:"column:AmountPaid" json:"amountPaid"`
	TaxAmount      float64 `gorm:"column:TaxAmount" json:"taxAmount"`
	DeliveryAmount float64 `gorm:"column:DeliveryAmount" json:"deliveryAmount"`
	ImposedId      string  `gorm:"column:ImposedId" json:"imposedId"`
	ExpenseId      string  `gorm:"column:ExpenseId" json:"expenseId"`
	OrderId        string  `gorm:"column:OrderId" json:"orderId"`
}

func (orderDetail *OrderDetail) TableName() string {
	return "OrderDetail"
}

func (orderDetail *OrderDetail) BeforeCreate(*gorm.DB) (err error) {
	orderDetail.Id = uuid.New().String()
	return
}

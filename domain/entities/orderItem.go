package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	Id         string `gorm:"column:Id;primaryKey" json:"id"`
	Quantity   int    `gorm:"column:Quantity" json:"quantity"`
	MedicineId string `gorm:"column:MedicineId" json:"medicineId"`
	OrderId    string `gorm:"column:OrderId" json:"orderId"`
	IsActive   bool   `gorm:"column:IsActive" json:"isActive"`
}

func (orderItem *OrderItem) TableName() string {
	return "OrderItem"
}

func (orderItem *OrderItem) BeforeCreate(*gorm.DB) (err error) {
	orderItem.Id = uuid.New().String()
	orderItem.IsActive = true
	return
}

func (orderItem *OrderItem) Update(quantity *int) {
	
	if quantity != nil {
		orderItem.Quantity = *quantity
	}

}

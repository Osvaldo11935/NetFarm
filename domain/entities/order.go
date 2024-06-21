package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	Id          string      `gorm:"column:Id;primaryKey" json:"id"`
	CreatedAt   time.Time   `gorm:"column:CreatedAt" json:"createdAt"`
	StatusId    string      `gorm:"column:StatusId" json:"statusId"`
	UserId      string      `gorm:"column:UserId" json:"userId"`
	OrderFiles  []OrderFile `gorm:"foreignKey:OrderId;references:Id" json:"files"`
	OrderItems  []OrderItem `gorm:"foreignKey:OrderId;references:Id" json:"orderItems"`
	OrderDetail OrderItem   `gorm:"foreignKey:OrderId;references:Id" json:"orderDetail"`
}

func (order *Order) TableName() string {
	return "Order"
}

func (order *Order) BeforeCreate(*gorm.DB) (err error) {
	order.Id = uuid.New().String()
	order.CreatedAt = time.Now()
	return
}

func (order *Order) UpdateStatus(statusId *string) {

	if statusId != nil {
		order.StatusId = *statusId
	}
}

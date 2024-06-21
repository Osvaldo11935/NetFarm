package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type OrderStatus struct {
	Id          string  `gorm:"column:Id;primaryKey" json:"id"`
	Type        string  `gorm:"column:Type" json:"type"`
	Description string  `gorm:"column:Description" json:"description"`
	Orders      []Order `gorm:"foreignKey:StatusId;references:Id" json:"orders"`
}

func (orderStatus *OrderStatus) TableName() string {
	return "OrderStatus"
}

func (orderStatus *OrderStatus) BeforeCreate(*gorm.DB) (err error) {
	orderStatus.Id = uuid.New().String()
	return
}

func (orderStatus *OrderStatus) Update(name *string, description *string) {

	if name != nil {
		orderStatus.Type = *name
	}
	if description != nil {
		orderStatus.Description = *description
	}
}

func InitialValueStatus(db *gorm.DB) {
	defaultStatus := []OrderStatus{
		{Type: "Em processamento", Description: ""},
		{Type: "Cancelado", Description: ""},
		{Type: "Pendente", Description: ""},
		{Type: "Concluido", Description: ""},
	}
	for _, status := range defaultStatus {
		if err := db.FirstOrCreate(&status, OrderStatus{Type: status.Type}).Error; err != nil {
			log.Fatalf("Não foi possivel adicionar os valor default do status: %v", err)
		}
	}
	return
}

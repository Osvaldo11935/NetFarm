package entities

import (
	object_values "NetFarm/domain/objectValues"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
	if orderStatus.Id == "" {
		orderStatus.Id = uuid.New().String()
	}
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
		{Id: object_values.STATUS_TYPE_IN_PROCESSING_ID,  Type: "Em processamento", Description: ""},
		{Id: object_values.STATUS_TYPE_CANCELED_ID,  Type: "Cancelado", Description: ""},
		{Id: object_values.STATUS_TYPE_PENDING_ID,  Type: "Pendente", Description: ""},
		{Id: object_values.STATUS_TYPE_COMPLETED_ID,  Type: "Concluido", Description: ""},
	}
	for _, status := range defaultStatus {
		if err := db.FirstOrCreate(&status, OrderStatus{Id: status.Id, Type: status.Type}).Error; err != nil {
			log.Fatalf("Não foi possivel adicionar os valor default do status: %v", err)
		}
	}
	return
}

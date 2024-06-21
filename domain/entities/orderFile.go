package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderFile struct {
	Id         string `gorm:"column:Id;primaryKey" json:"id"`
	Url        string `gorm:"column:Url" json:"url"`
	FileTypeId string `gorm:"column:FileTypeId" json:"fileTypeId"`
	OrderId    string `gorm:"column:OrderId" json:"orderId"`
	IsActive   bool   `gorm:"column:IsActive" json:"isActive"`
}

func (orderFile *OrderFile) TableName() string {
	return "OrderFile"
}

func (orderFile *OrderFile) BeforeCreate(*gorm.DB) (err error) {
	orderFile.Id = uuid.New().String()
	return
}

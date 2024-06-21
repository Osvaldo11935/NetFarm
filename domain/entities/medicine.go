package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Medicine struct {
	Id               string             `gorm:"column:Id;primaryKey" json:"id"`
	Name             string             `gorm:"column:Name" json:"name"`
	Description      string             `gorm:"column:Description" json:"description"`
	Quantity         int                `gorm:"column:Quantity" json:"quantity"`
	Price            float64            `gorm:"column:Price" json:"price"`
	ProviderId       string             `gorm:"column:ProviderId" json:"providerId"`
	IsActive         bool               `gorm:"column:IsActive" json:"isActive"`
	MedicineCategory []MedicineCategory `gorm:"foreignKey:MedicineId;references:Id" json:"medicineCategory"`
}

func (medicine *Medicine) TableName() string {
	return "FileType"
}

func (medicine *Medicine) BeforeCreate(*gorm.DB) (err error) {
	medicine.Id = uuid.New().String()
	medicine.IsActive = true
	return
}

func (medicine *Medicine) Update(name *string, description *string, quantity *int, price *float64) {

	if name != nil {
		medicine.Name = *name
	}
	if description != nil {
		medicine.Description = *description
	}
	if quantity != nil {
		medicine.Quantity = *quantity
	}
	if price != nil {
		medicine.Price = *price
	}

}

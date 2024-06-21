package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Category struct {
	Id               string             `gorm:"column:Id;primaryKey" json:"id"`
	Name             string             `gorm:"column:Name;" json:"name"`
	Description      string             `gorm:"column:Description;" json:"description"`
	CreateAt         time.Time          `gorm:"column:CreateAt;"  json:"createAt"`
	UpdateAt         time.Time          `gorm:"column:UpdateAt;"  json:"updateAt"`
	MedicineCategory []MedicineCategory `gorm:"foreignKey:CategoryId;references:Id" json:"medicineCategory"`
}

func (category *Category) TableName() string {
	return "Category"
}

func (category *Category) BeforeCreate(*gorm.DB) (err error) {
	category.Id = uuid.New().String()
	return
}

func (category *Category) Update(name *string, description *string) {
	if name != nil {
		category.Name = *name
	}
	if description != nil {
		category.Description = *description
	}
}

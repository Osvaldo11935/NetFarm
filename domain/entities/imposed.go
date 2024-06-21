package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Imposed struct {
	Id   string  `gorm:"column:Id;primaryKey" json:"id"`
	Name string  `gorm:"column:Name" json:"name"`
	Rate float64 `gorm:"column:Rate" json:"rate"`
}

func (imposed *Imposed) TableName() string {
	return "Imposed"
}

func (imposed *Imposed) BeforeCreate(*gorm.DB) (err error) {
	imposed.Id = uuid.New().String()
	return
}

func (imposed *Imposed) Update(name *string, rate *float64) {
	if name != nil {
		imposed.Name = *name
	}
	if rate != nil {
		imposed.Rate = *rate
	}
}

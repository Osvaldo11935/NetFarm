package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Expense struct {
	Id   string  `gorm:"column:Id;primaryKey" json:"id"`
	Name string  `gorm:"column:Name" json:"name"`
	Rate float64 `gorm:"column:Rate" json:"rate"`
}

func (expense *Expense) TableName() string {
	return "Expense"
}

func (expense *Expense) BeforeCreate(*gorm.DB) (err error) {
	expense.Id = uuid.New().String()
	return
}

func (expense *Expense) Update(name *string, rate *float64) {
	if name != nil {
		expense.Name = *name
	}
	if rate != nil {
		expense.Rate = *rate
	}
}

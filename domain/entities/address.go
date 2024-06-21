package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	Id           string  `gorm:"column:Id;primaryKey" json:"id"`
	Country      string  `gorm:"column:Country" json:"country"`
	State        string  `gorm:"column:State" json:"state"`
	City         string  `gorm:"column:City" json:"city"`
	Neighborhood string  `gorm:"column:Neighborhood" json:"neighborhood"`
	Street       string  `gorm:"column:Street" json:"street"`
	Number       string  `gorm:"column:Number" json:"number"`
	Complement   string  `gorm:"column:Complement" json:"complement"`
	Latitude     float64 `gorm:"column:Latitude" json:"latitude"`
	Longitude    float64 `gorm:"column:Longitude" json:"longitude"`
	UserId       string  `gorm:"column:UserId" json:"userId"`
}

func (address *Address) TableName() string {
	return "Address"
}

func (address *Address) BeforeCreate(*gorm.DB) (err error) {
	address.Id = uuid.New().String()
	return
}

func (address *Address) Update(country *string, state *string, city *string, neighborhood *string,
	street *string, number *string, complement *string) {

	if country != nil {
		address.Country = *country
	}
	if state != nil {
		address.State = *state
	}
	if city != nil {
		address.City = *city
	}
	if neighborhood != nil {
		address.Neighborhood = *neighborhood
	}
	if street != nil {
		address.Street = *street
	}
	if number != nil {
		address.Number = *number
	}
	if complement != nil {
		address.Complement = *complement
	}
}

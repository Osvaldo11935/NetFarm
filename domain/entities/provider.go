package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Provider struct {
	Id        string     `gorm:"column:Id;primaryKey" json:"id"`
	Name      string     `gorm:"column:Name" json:"name"`
	IsActive  bool       `gorm:"column:IsActive" json:"isActive"`
	Medicines []Medicine `gorm:"foreignKey:ProviderId;references:Id" json:"medicines"`
}

func (provider *Provider) TableName() string {
	return "Provider"
}

func (provider *Provider) BeforeCreate(*gorm.DB) (err error) {
	provider.Id = uuid.New().String()
	provider.IsActive = true
	return
}

func (provider *Provider) Update(name *string) {

	if name != nil {
		provider.Name = *name
	}
}

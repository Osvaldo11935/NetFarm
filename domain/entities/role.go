package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type Role struct {
	Id    string `gorm:"column:Id;primaryKey" json:"id"`
	Name  string `gorm:"column:Name" json:"Name"`
	Users []User `gorm:"foreignKey:RoleId;references:Id" json:"users"`
}

func (role *Role) TableName() string {
	return "Role"
}

func (role *Role) BeforeCreate(*gorm.DB) (err error) {
	if role.Id == "" {
		role.Id = uuid.New().String()
	}

	return
}

func (role *Role) Update(name *string) {

	if name != nil {
		role.Name = *name
	}
}

func InitialValueRole(db *gorm.DB) {
	defaultRoles := []Role{
		{Id: "AB38D2FA-AA63-4460-B800-60E0867CE3CF", Name: "Cliente"},
		{Id: "B98F3022-2B94-44D1-A965-BE957B9434EA", Name: "Administrador"},
	}
	for _, role := range defaultRoles {
		if err := db.FirstOrCreate(&role, Role{Name: role.Name}).Error; err != nil {
			log.Fatalf("Não foi possivel adicionar os valor default do status: %v", err)
		}
	}
	return
}

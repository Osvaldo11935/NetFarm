package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Id          string  `gorm:"column:Id;primaryKey" json:"id"`
	Email       string  `gorm:"column:Email" json:"email"`
	UserName    string  `gorm:"column:UserName" json:"userName"`
	PhoneNumber string  `gorm:"column:PhoneNumber" json:"phoneNumber"`
	RoleId      string  `gorm:"column:RoleId" json:"roleId"`
	Password    string  `gorm:"column:Password" json:"password"`
	Orders      []Order `gorm:"foreignKey:UserId;references:Id" json:"orders"`
	Person      Person  `gorm:"foreignKey:UserId;references:Id" json:"person"`
	Address     Address `gorm:"foreignKey:UserId;references:Id" json:"address"`
}

func (user *User) TableName() string {
	return "User"
}

func (user *User) BeforeCreate(*gorm.DB) (err error) {
	user.Id = uuid.New().String()
	return
}

func (user *User) Update(email *string, userName *string, phoneNumber *string) {
	if email != nil {
		user.Email = *email
	}
	if userName != nil {
		user.UserName = *userName
	}
	if phoneNumber != nil {
		user.PhoneNumber = *phoneNumber
	}
}

func InitialValueUser(db *gorm.DB) {
	defaultUsers := []User{
		{Email: "userdetest1@gmail.com", UserName: "userdetest1", PhoneNumber: "943118944", RoleId: "AB38D2FA-AA63-4460-B800-60E0867CE3CF"},
		{Email: "userdetest2@gmail.com", UserName: "userdetest2", PhoneNumber: "943118944", RoleId: "AB38D2FA-AA63-4460-B800-60E0867CE3CF"},
	}
	for _, user := range defaultUsers {
		if err := db.FirstOrCreate(&user, User{Email: user.Email, UserName: user.UserName, PhoneNumber: user.PhoneNumber, RoleId: user.RoleId}).Error; err != nil {
			log.Fatalf("Não foi possivel adicionar os valor default do status: %v", err)
		}
	}
	return
}

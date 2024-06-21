package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Person struct {
	Id        string    `gorm:"column:Id;primaryKey" json:"id"`
	FullName  string    `gorm:"column:FullName" json:"fullName"`
	BirthDate time.Time `gorm:"column:BirthDate" json:"birthDate"`
	Gender    string    `gorm:"column:Gender" json:"gender"`
	UserId    string    `gorm:"column:UserId" json:"userId"`
}

func (person *Person) TableName() string {
	return "Person"
}

func (person *Person) BeforeCreate(*gorm.DB) (err error) {
	person.Id = uuid.New().String()
	return
}

func (person *Person) Update(fullName *string, birthDate *time.Time, gender *string) {

	if fullName != nil {
		person.FullName = *fullName
	}
	if birthDate != nil {
		person.BirthDate = *birthDate
	}
	if gender != nil {
		person.Gender = *gender
	}
}

package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileType struct {
	Id          string      `gorm:"column:Id;primaryKey" json:"id"`
	Type        string      `gorm:"column:Type" json:"type"`
	Description string      `gorm:"column:Description" json:"description"`
	OrderFiles  []OrderFile `gorm:"foreignKey:FileTypeId;references:Id" json:"orderFiles"`
}

func (fileType *FileType) TableName() string {
	return "FileType"
}

func (fileType *FileType) BeforeCreate(*gorm.DB) (err error) {
	fileType.Id = uuid.New().String()
	return
}

func (fileType *FileType) Update(name *string, description *string) {
	if name != nil {
		fileType.Type = *name
	}
	if description != nil {
		fileType.Description = *description
	}
}

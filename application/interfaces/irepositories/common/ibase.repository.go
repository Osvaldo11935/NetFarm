package common

import "gorm.io/gorm"

type IBaseRepository interface {
	Create(entity interface{}) error
	Query() *gorm.DB
	//FindAll(entity interface{}) error
	FindByID(entity interface{}, id uint) error
	FindByCondition(entity interface{}, condition string, args ...interface{}) error
	Update(entity interface{}) error
	Delete(entity interface{}) error
}

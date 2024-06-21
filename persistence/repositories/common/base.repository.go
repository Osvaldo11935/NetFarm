package common

import (
	"NetFarm/application/interfaces/irepositories/common"
	"NetFarm/persistence/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() common.IBaseRepository {
	db, _ := database.NewDatabase()

	return &Repository{db}
}

func (r *Repository) Create(entity interface{}) error {
	return r.db.Create(entity).Error
}
func (r *Repository) Query() *gorm.DB {
	return r.db
}

func (r *Repository) FindByID(entity interface{}, id uint) error {

	return r.db.First(entity, id).Error
}
func (r *Repository) FindByCondition(entity interface{}, condition string, args ...interface{}) error {

	return r.db.Where(condition, args...).Find(entity).Error
}
func (r *Repository) Update(entity interface{}) error {
	return r.db.Save(entity).Error
}

func (r *Repository) Delete(entity interface{}) error {
	return r.db.Delete(entity).Error
}

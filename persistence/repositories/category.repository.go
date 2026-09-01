package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	*common.Repository
}

func NewCategoryRepository(db *gorm.DB) irepositories.ICategoryRepository {
	return &CategoryRepository{
		Repository: common.NewRepository(db).(*common.Repository),
	}
}

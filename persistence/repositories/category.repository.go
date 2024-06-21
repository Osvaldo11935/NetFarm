package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
)

type CategoryRepository struct {
	*common.Repository
}

func NewCategoryRepository() irepositories.ICategoryRepository {
	return &CategoryRepository{
		Repository: common.NewRepository().(*common.Repository),
	}
}

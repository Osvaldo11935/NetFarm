package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
	"gorm.io/gorm"
)

type OrderFileRepository struct {
	*common.Repository
}

func NewOrderFileRepository(db *gorm.DB) irepositories.IOrderFileRepository {
	return &OrderFileRepository{
		Repository: common.NewRepository(db).(*common.Repository),
	}
}

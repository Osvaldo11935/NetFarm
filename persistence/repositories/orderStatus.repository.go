package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
	"gorm.io/gorm"
)

type OrderStatusRepository struct {
	*common.Repository
}

func NewOrderStatusRepository(db *gorm.DB) irepositories.IOrderStatusRepository {
	return &OrderStatusRepository{
		Repository: common.NewRepository(db).(*common.Repository),
	}
}

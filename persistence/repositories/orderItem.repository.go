package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
	"gorm.io/gorm"
)

type OrderItemRepository struct {
	*common.Repository
}

func NewOrderItemRepository(db *gorm.DB) irepositories.IOrderItemRepository {
	return &OrderItemRepository{
		Repository: common.NewRepository(db).(*common.Repository),
	}
}

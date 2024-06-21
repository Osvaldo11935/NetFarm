package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
)

type OrderItemRepository struct {
	*common.Repository
}

func NewOrderItemRepository() irepositories.IOrderItemRepository {
	return &OrderItemRepository{
		Repository: common.NewRepository().(*common.Repository),
	}
}

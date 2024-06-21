package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
)

type OrderStatusRepository struct {
	*common.Repository
}

func NewOrderStatusRepository() irepositories.IOrderStatusRepository {
	return &OrderStatusRepository{
		Repository: common.NewRepository().(*common.Repository),
	}
}

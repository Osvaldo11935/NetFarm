package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
)

type OrderRepository struct {
	*common.Repository
}

func NewOrderRepository() irepositories.IOrderRepository {
	return &OrderRepository{
		Repository: common.NewRepository().(*common.Repository),
	}
}

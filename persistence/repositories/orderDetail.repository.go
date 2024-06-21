package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
)

type OrderDetailRepository struct {
	*common.Repository
}

func NewOrderDetailRepository() irepositories.IOrderDetailRepository {
	return &OrderDetailRepository{
		Repository: common.NewRepository().(*common.Repository),
	}
}

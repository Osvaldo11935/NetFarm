package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
)

type OrderFileRepository struct {
	*common.Repository
}

func NewOrderFileRepository() irepositories.IOrderFileRepository {
	return &OrderFileRepository{
		Repository: common.NewRepository().(*common.Repository),
	}
}

package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
	"gorm.io/gorm"
)

type OrderRepository struct {
	*common.Repository
}

func NewOrderRepository(db *gorm.DB) irepositories.IOrderRepository {
	return &OrderRepository{
		Repository: common.NewRepository(db).(*common.Repository),
	}
}

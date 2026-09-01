package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
	"gorm.io/gorm"
)

type OrderDetailRepository struct {
	*common.Repository
}

func NewOrderDetailRepository(db *gorm.DB) irepositories.IOrderDetailRepository {
	return &OrderDetailRepository{
		Repository: common.NewRepository(db).(*common.Repository),
	}
}

package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"

	"gorm.io/gorm"
)

type AddressRepository struct {
	*common.Repository
}

func NewAddressRepository(db *gorm.DB) irepositories.IAddressRepository {
	return &AddressRepository{
		Repository: common.NewRepository(db).(*common.Repository),
	}
}

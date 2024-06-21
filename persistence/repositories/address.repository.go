package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
)

type AddressRepository struct {
	*common.Repository
}

func NewAddressRepository() irepositories.IAddressRepository {
	return &AddressRepository{
		Repository: common.NewRepository().(*common.Repository),
	}
}

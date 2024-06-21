package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
)

type ProviderRepository struct {
	*common.Repository
}

func NewProviderRepository() irepositories.IProviderRepository {
	return &ProviderRepository{
		Repository: common.NewRepository().(*common.Repository),
	}
}

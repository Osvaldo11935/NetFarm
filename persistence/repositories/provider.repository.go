package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
	"gorm.io/gorm"
)

type ProviderRepository struct {
	*common.Repository
}

func NewProviderRepository(db *gorm.DB) irepositories.IProviderRepository {
	return &ProviderRepository{
		Repository: common.NewRepository(db).(*common.Repository),
	}
}

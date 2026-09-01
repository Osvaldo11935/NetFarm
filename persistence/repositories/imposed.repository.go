package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
	"gorm.io/gorm"
)

type ImposedRepository struct {
	*common.Repository
}

func NewImposedRepository(db *gorm.DB) irepositories.IImposedRepository {
	return &ImposedRepository{
		Repository: common.NewRepository(db).(*common.Repository),
	}
}

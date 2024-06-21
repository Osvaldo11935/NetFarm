package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
)

type ImposedRepository struct {
	*common.Repository
}

func NewImposedRepository() irepositories.IImposedRepository {
	return &ImposedRepository{
		Repository: common.NewRepository().(*common.Repository),
	}
}

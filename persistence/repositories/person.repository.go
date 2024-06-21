package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
)

type PersonRepository struct {
	*common.Repository
}

func NewPersonRepository() irepositories.IPersonRepository {
	return &PersonRepository{
		Repository: common.NewRepository().(*common.Repository),
	}
}

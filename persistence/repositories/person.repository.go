package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
	"gorm.io/gorm"
)

type PersonRepository struct {
	*common.Repository
}

func NewPersonRepository(db *gorm.DB) irepositories.IPersonRepository {
	return &PersonRepository{
		Repository: common.NewRepository(db).(*common.Repository),
	}
}

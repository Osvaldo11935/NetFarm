package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
)

type UserRepository struct {
	*common.Repository
}

func NewUserRepository() irepositories.IUserRepository {
	return &UserRepository{
		Repository: common.NewRepository().(*common.Repository),
	}
}

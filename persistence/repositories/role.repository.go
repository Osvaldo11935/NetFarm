package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
)

type RoleRepository struct {
	*common.Repository
}

func NewRoleRepository() irepositories.IRoleRepository {
	return &RoleRepository{
		Repository: common.NewRepository().(*common.Repository),
	}
}

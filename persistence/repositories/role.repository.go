package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
	"gorm.io/gorm"
)

type RoleRepository struct {
	*common.Repository
}

func NewRoleRepository(db *gorm.DB) irepositories.IRoleRepository {
	return &RoleRepository{
		Repository: common.NewRepository(db).(*common.Repository),
	}
}

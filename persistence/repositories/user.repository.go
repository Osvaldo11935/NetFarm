package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
	"gorm.io/gorm"
)

type UserRepository struct {
	*common.Repository
}

func NewUserRepository(db *gorm.DB) irepositories.IUserRepository {
	return &UserRepository{
		Repository: common.NewRepository(db).(*common.Repository),
	}
}

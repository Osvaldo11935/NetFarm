package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
	"gorm.io/gorm"
)

type ExpenseRepository struct {
	*common.Repository
}

func NewExpenseRepository(db *gorm.DB) irepositories.IExpenseRepository {
	return &ExpenseRepository{
		Repository: common.NewRepository(db).(*common.Repository),
	}
}

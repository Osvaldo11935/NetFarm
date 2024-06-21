package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
)

type ExpenseRepository struct {
	*common.Repository
}

func NewExpenseRepository() irepositories.IExpenseRepository {
	return &ExpenseRepository{
		Repository: common.NewRepository().(*common.Repository),
	}
}

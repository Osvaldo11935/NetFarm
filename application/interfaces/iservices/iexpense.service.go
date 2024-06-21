package iservices

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
)

type IExpenseService interface {
	CreateExpense(expense *entities.Expense) *common.ErrorResponse
	FindAllExpense(pagination *common.Pagination) ([]entities.Expense, *common.ErrorResponse)
	FindExpenseById(expenseId string) (entities.Expense, *common.ErrorResponse)
	UpdateExpense(expenseId string, name *string, rate *float64) *common.ErrorResponse
	DeleteExpense(expenseId string) *common.ErrorResponse
}

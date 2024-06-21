package services

import (
	"NetFarm/application/extensions"
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/errors"
	"NetFarm/shared/models/common"
)

type ExpenseService struct {
	ExpenseRepo irepositories.IExpenseRepository
}

func NewExpenseService(ExpenseRepo irepositories.IExpenseRepository) iservices.IExpenseService {
	return &ExpenseService{
		ExpenseRepo: ExpenseRepo,
	}
}

func (s *ExpenseService) CreateExpense(expense *entities.Expense) *common.ErrorResponse {

	createdErr := s.ExpenseRepo.Create(expense)

	if createdErr != nil {
		return errors.NewExpenseUnknownError(createdErr.Error(), "falha ao cadastrar despesa")
	}

	return nil
}

func (s *ExpenseService) FindAllExpense(pagination *common.Pagination) ([]entities.Expense, *common.ErrorResponse) {

	var Expenses []entities.Expense

	query := s.ExpenseRepo.Query()

	query = query.Scopes(extensions.Paginate(&Expenses, pagination, query))

	findErr := query.Find(&Expenses).Error

	if findErr != nil {
		return Expenses, errors.NewExpenseUnknownError(findErr.Error(), "falha ao buscar despesa")
	}

	pagination.Data = Expenses

	return Expenses, nil
}

func (s *ExpenseService) FindExpenseById(expenseId string) (entities.Expense, *common.ErrorResponse) {

	var Expense entities.Expense

	findErr := s.ExpenseRepo.Query().First(&Expense, "Id", expenseId).Error

	if findErr != nil {
		return Expense, errors.NewExpenseUnknownError(findErr.Error(), "falha ao buscar despesa")
	}

	return Expense, nil
}

func (s *ExpenseService) UpdateExpense(expenseId string, name *string, rate *float64) *common.ErrorResponse {

	Expense, err := s.FindExpenseById(expenseId)

	if err != nil {
		return err
	}

	Expense.Update(name, rate)

	updatedErr := s.ExpenseRepo.Update(Expense)

	if updatedErr != nil {
		return errors.NewExpenseUnknownError(updatedErr.Error(), "falha ao atualizar despesa")
	}

	return nil
}

func (s *ExpenseService) DeleteExpense(ExpenseId string) *common.ErrorResponse {

	Expense, err := s.FindExpenseById(ExpenseId)

	if err != nil {
		return err
	}

	deletedErr := s.ExpenseRepo.Delete(Expense)

	if deletedErr != nil {
		return errors.NewExpenseUnknownError(deletedErr.Error(), "falha ao deletar despesa")
	}

	return nil
}

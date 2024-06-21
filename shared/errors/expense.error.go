package errors

import "NetFarm/shared/models/common"

func NewExpenseUnknownError(message string, description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"expenseUnknownError",
		message, description)
}

func NewExpenseAlreadyExists(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"expenseAlreadyExists",
		"despesa ja cadastrado", description)
}

func NewExpenseNotFound(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"expenseNotFound",
		"despesa não encontrado", description)
}

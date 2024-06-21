package errors

import "NetFarm/shared/models/common"

func NewOrderStatusUnknownError(message string, description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"orderStatusUnknownError",
		message, description)
}

func NewOrderStatusAlreadyExists(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"orderStatusAlreadyExists",
		"estado ja cadastrado", description)
}

func NewOrderStatusNotFound(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"orderStatusNotFound",
		"estado não encontrado", description)
}

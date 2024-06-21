package errors

import "NetFarm/shared/models/common"

func NewOrderUnknownError(message string, description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"orderUnknownError",
		message, description)
}

func NewOrderAlreadyExists(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"orderAlreadyExists",
		"pedido ja cadastrado", description)
}

func NewOrderNotFound(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"orderNotFound",
		"pedido não encontrado", description)
}

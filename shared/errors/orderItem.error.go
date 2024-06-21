package errors

import "NetFarm/shared/models/common"

func NewOrderItemUnknownError(message string, description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"orderItemUnknownError",
		message, description)
}

func NewOrderItemAlreadyExists(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"orderItemAlreadyExists",
		"item do pedido ja cadastrado", description)
}

func NewOrderItemNotFound(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"orderItemNotFound",
		"item do pedido não encontrado", description)
}

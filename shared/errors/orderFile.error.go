package errors

import "NetFarm/shared/models/common"

func NewOrderFileUnknownError(message string, description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"orderFileUnknownError",
		message, description)
}

func NewOrderFileAlreadyExists(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"orderFileAlreadyExists",
		"ficheiro do pedido ja cadastrado", description)
}

func NewOrderFileNotFound(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"orderFileNotFound",
		"ficheiro do pedido não encontrado", description)
}

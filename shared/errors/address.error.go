package errors

import "NetFarm/shared/models/common"

func NewAddressUnknownError(message string, description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"addressUnknownError",
		message, description)
}

func NewAddressAlreadyExists(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"addressAlreadyExists",
		"usuário já tem um endereço cadastrado", description)
}

func NewAddressNotFound(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"addressNotFound",
		"endereço do usuario não encontrado", description)
}

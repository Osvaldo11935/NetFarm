package errors

import "NetFarm/shared/models/common"

func NewPersonUnknownError(message string, description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"personUnknownError",
		message, description)
}

func NewPersonAlreadyExists(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"personAlreadyExists",
		"usuário já tem dados pessoais cadastrado", description)
}

func NewPersonNotFound(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"personNotFound",
		"dados pessoais do usuario não encontrado", description)
}

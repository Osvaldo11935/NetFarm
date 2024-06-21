package errors

import "NetFarm/shared/models/common"

func NewUserUnknownError(message string, description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"userUnknownError",
		message, description)
}

func NewUserAlreadyExists(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"userAlreadyExists",
		"usuário já existe", description)
}

func NewUserNotFound(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"userNotFound",
		"usuario não encontrado", description)
}

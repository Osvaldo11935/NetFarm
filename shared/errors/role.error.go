package errors

import "NetFarm/shared/models/common"

func NewRoleUnknownError(message string, description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"roleUnknownError",
		message, description)
}

func NewRoleAlreadyExists(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"roleAlreadyExists",
		"perfil já existe", description)
}

func NewRoleNotFound(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"roleNotFound",
		"perfil não encontrado", description)
}

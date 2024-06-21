package errors

import "NetFarm/shared/models/common"

func NewImposedUnknownError(message string, description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"imposedUnknownError",
		message, description)
}

func NewImposedAlreadyExists(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"imposedAlreadyExists",
		"imposto ja cadastrado", description)
}

func NewImposedNotFound(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"imposedNotFound",
		"imposto não encontrado", description)
}

package errors

import "NetFarm/shared/models/common"

func NewFileTypeUnknownError(message string, description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"fileTypeUnknownError",
		message, description)
}

func NewFileTypeAlreadyExists(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"fileTypeAlreadyExists",
		"tipo de ficheiro ja cadastrado", description)
}

func NewFileTypeNotFound(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"fileTypeNotFound",
		"tipo de ficheiro não encontrado", description)
}

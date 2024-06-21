package errors

import "NetFarm/shared/models/common"

func NewGoogleStorageUnknownError(message string, description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"uploadUnknownError",
		message, description)
}

func NewGoogleStorageAlreadyExists(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"uploadAlreadyExists",
		"remedio ja cadastrado", description)
}

func NewGoogleStorageNotFound(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"uploadNotFound",
		"arquivo keys.json não encontrado", description)
}

func NewGoogleStorageFileKeysNotFound() *common.ErrorResponse {
	return common.NewErrorResponse(
		"uploadNotFound",
		"arquivo keys.json não encontrado", "falha ao fazer upload do ficheiro")
}

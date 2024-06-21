package errors

import "NetFarm/shared/models/common"

func NewProviderUnknownError(message string, description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"providerUnknownError",
		message, description)
}

func NewProviderAlreadyExists(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"providerAlreadyExists",
		"fornecedor ja existente", description)
}

func NewProviderNotFound(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"providerNotFound",
		"fornecedor não encontrado", description)
}

package errors

import "NetFarm/shared/models/common"

func NewCategoryUnknownError(message string, description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"categoryUnknownError",
		message, description)
}

func NewCategoryAlreadyExists(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"categoryAlreadyExists",
		"categoria ja cadastrado", description)
}

func NewCategoryNotFound(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"categoryNotFound",
		"categoria não encontrado", description)
}

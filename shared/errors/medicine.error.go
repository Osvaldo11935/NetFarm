package errors

import "NetFarm/shared/models/common"

func NewMedicineUnknownError(message string, description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"medicineUnknownError",
		message, description)
}

func NewMedicineAlreadyExists(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"medicineAlreadyExists",
		"remedio ja cadastrado", description)
}

func NewMedicineNotFound(description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"medicineNotFound",
		"remedio não encontrado", description)
}

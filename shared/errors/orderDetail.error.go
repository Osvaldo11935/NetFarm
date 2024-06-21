package errors

import "NetFarm/shared/models/common"

func NewOrderDetailUnknownError(message string, description string) *common.ErrorResponse {
	return common.NewErrorResponse(
		"orderDetailUnknownError",
		message, description)
}

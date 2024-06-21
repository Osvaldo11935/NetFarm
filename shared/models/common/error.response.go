package common

type ErrorResponse struct {
	Code        string      `json:"code"`
	Message     string      `json:"message"`
	Description string      `json:"description"`
	Errors      interface{} `json:"errors"`
}

func NewErrorResponse(code string, message string, description string) *ErrorResponse {
	return &ErrorResponse{
		Code:        code,
		Message:     message,
		Description: description}
}

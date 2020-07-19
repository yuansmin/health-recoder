package controllers

var (
	BadRequestErr = "BadRequest"
	InternalErr   = "InternalErr"
)

type ApiError struct {
	ErrCode string
	Message string
}

func newApiError(code, message string) ApiError {
	return ApiError{
		ErrCode: code,
		Message: message,
	}
}

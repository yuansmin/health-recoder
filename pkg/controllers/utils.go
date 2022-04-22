package controllers

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

const (
	defaultLimit = 10
	maxLimit     = 100
)

var (
	CodeBadRequestErr = "BadRequest"
	CodeInternalErr   = "InternalErr"
	CodeNotFoundErr   = "NotFound"
)

type ApiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%s %s", e.Code, e.Message)
}

func newApiError(code, message string) ApiError {
	return ApiError{
		Code:    code,
		Message: message,
	}
}

func isNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func checkPageParameter(offset, limit int) *ApiError {
	if offset < 0 {
		return &ApiError{
			Code:    CodeBadRequestErr,
			Message: fmt.Sprintf("invalid offset: %d", offset),
		}
	}

	if limit < 0 || limit > maxLimit {
		return &ApiError{
			Code:    CodeBadRequestErr,
			Message: fmt.Sprintf("invalid limit: %d", limit),
		}
	}

	return nil
}

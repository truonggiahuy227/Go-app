package response

import (
	"akawork.io/constant"
	"akawork.io/gerror"
)

/**
 * Defines a response object
 */
type Response struct {
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

/**
 * Defines an error response object
 */
type ErrorResponse struct {
	ErrorCode uint32 `json:"ErrorCode"`
	Message   string `json:"Message"`
	Exception string `json:"Exception"`
}

/**
 * Returns a new error response object
 */
func NewErrorResponse(errorCode uint32, message string, exception string) (string, ErrorResponse) {
	msg := gerror.T(errorCode)
	// status := errorCode
	error := ErrorResponse{
		ErrorCode: errorCode,
		Message:   message,
	}

	return msg, error
}

func NewGerrorResponse(gerr *gerror.Error) (string, ErrorResponse) {
	msg := gerror.T(gerr.Code)
	error := ErrorResponse{
		ErrorCode: gerr.Code,
		Message:   constant.LogErrorPrefix + gerr.Error.Error(),
		Exception: gerr.Line,
	}
	return msg, error
}

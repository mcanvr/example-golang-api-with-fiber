package api

import "mcanvr/example-golang-api-with-fiber/internal/interfaces/common"

// ResponseModel is the standard structure used for all API responses.
// It is defined as a type alias to support Swagger documentation.
type ResponseModel = common.ResponseModel

// Re-export the common ResponseModel functions for backward compatibility

// NewSuccessResponse creates a success response with the given message and data.
func NewSuccessResponse(message string, data interface{}) common.ResponseModel {
	return common.NewSuccessResponse(message, data)
}

// NewErrorResponse creates an error response with the given message.
func NewErrorResponse(message, details string) common.ResponseModel {
	return common.NewErrorResponse(message, details)
}

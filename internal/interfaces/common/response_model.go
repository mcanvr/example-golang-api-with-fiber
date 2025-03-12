package common

// ResponseModel provides a standard structure for API responses
type ResponseModel struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewSuccessResponse creates successful API responses
func NewSuccessResponse(message string, data interface{}) ResponseModel {
	return ResponseModel{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// NewErrorResponse creates error API responses
func NewErrorResponse(message, details string) ResponseModel {
	// If details are provided, append them to the main message
	if details != "" {
		message = message + ": " + details
	}

	return ResponseModel{
		Success: false,
		Message: message,
		Data:    nil,
	}
}

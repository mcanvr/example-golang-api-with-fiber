package errors

import "fmt"

// Custom error types for application-specific error handling.
// These typed errors allow for more precise error handling and
// consistent error reporting throughout the application.

// ErrInvalidRequest represents validation errors for user input.
// It identifies which field has a problem and provides a specific message.
type ErrInvalidRequest struct {
	Field   string // The field that failed validation
	Message string // Description of the validation error
}

// Error implements the error interface for ErrInvalidRequest.
// Returns a formatted string containing the field name and error message.
func (e *ErrInvalidRequest) Error() string {
	return fmt.Sprintf("invalid %s: %s", e.Field, e.Message)
}

// ErrNotFound indicates that a requested resource couldn't be located.
// Used primarily for database lookup failures.
type ErrNotFound struct {
	Resource string // Type of resource (e.g., "user", "product")
	ID       any    // Identifier that was not found
}

// Error implements the error interface for ErrNotFound.
// Returns a descriptive message including the resource type and ID.
func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("%s with id %v not found", e.Resource, e.ID)
}

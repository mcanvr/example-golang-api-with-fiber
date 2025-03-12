package constants

// Message constants for consistent API responses.
// These provide a central place to manage all message texts.
const (
	// User operation success messages
	UsersFetched = "Users fetched successfully"
	UserFound    = "User found"
	UserCreated  = "User created successfully"
	UserUpdated  = "User updated successfully"
	UserDeleted  = "User deleted successfully"

	// User operation error messages - lowercase for error messages
	UserNotFound      = "user not found"
	CannotGetUsers    = "failed to retrieve users"
	CannotCreateUser  = "failed to create user"
	CannotUpdateUser  = "failed to update user"
	CannotDeleteUser  = "failed to delete user"
	InvalidIDFormat   = "invalid ID format"
	MissingIDParam    = "missing ID parameter"
	EmailAlreadyInUse = "email address is already in use"

	// General API messages
	InvalidRequestFormat = "Invalid request format"       // For UI display
	EndpointNotFound     = "Endpoint not found"           // For UI display
	InternalServerError  = "Internal server error"        // For UI display
	UnauthorizedAccess   = "Unauthorized access"          // For UI display
	ForbiddenAction      = "Forbidden action"             // For UI display
	ResourceNotFound     = "Requested resource not found" // For UI display
	GeneralError         = "Error"                        // For UI display
	RequestTimeout       = "Request timed out"            // For UI display
	TooManyRequests      = "Too many requests"            // For UI display
	RequestCanceled      = "Request canceled"             // For UI display

	// Rate limiter messages
	RateLimitExceeded       = "Rate limit exceeded for IP: %s"               // For logs
	RateLimitExceededUI     = "Rate limit exceeded. Please try again later." // For UI display
	RequestTimeoutMessageUI = "Request timed out. Please try again later."   // For UI display
	RequestCanceledUI       = "Request canceled."                            // For UI display
	RequestTimeoutLog       = "Request timed out: %s %s"                     // For logs - method, path

	// Authentication messages
	LoginSuccess         = "Login successful"      // For UI display
	LoginFailed          = "Login failed"          // For UI display
	AuthenticationFailed = "Authentication failed" // For UI display

	// Error messages (lowercase for use with errors.New/fmt.Errorf)
	InvalidCredentials    = "invalid username or password"
	TokenCreationFailed   = "failed to generate authentication token"
	TokenExpired          = "authentication token has expired"
	TokenInvalid          = "invalid authentication token"
	AccessDenied          = "access denied: insufficient permissions"
	MissingToken          = "authentication token not found"
	InvalidTokenFormat    = "invalid authentication format, use 'Bearer TOKEN' format"
	InvalidOrExpiredToken = "invalid or expired token: %s"

	// Validation messages - user facing, can be capitalized
	ValidationError = "Validation error: %s" // For UI display

	// Validation field errors - lowercase for consistent error formatting
	FieldRequired          = "field '%s' is required"
	FieldInvalidEmail      = "field '%s' must be a valid email address"
	FieldMinLength         = "field '%s' must be at least %s characters long"
	FieldMaxLength         = "field '%s' must be at most %s characters long"
	FieldMinValue          = "field '%s' must be greater than or equal to %s"
	FieldMaxValue          = "field '%s' must be less than or equal to %s"
	FieldGenericValidation = "field '%s' failed validation: %s"

	// Server messages - used in logs, can be capitalized
	ServerStarting          = "Server starting on %s"
	ServerStartFailed       = "Server failed to start: %v"
	UnexpectedError         = "Unexpected error: %v"
	UserFriendlyServerError = "An unexpected server error occurred. Please try again later." // For UI display
	PanicRecovered          = "Panic recovered: %v"
	SampleDataInitFailed    = "Failed to initialize sample users: %v"
)

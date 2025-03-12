package dto

import "mcanvr/example-golang-api-with-fiber/internal/domain/model"

// UserResponse represents the data structure returned to API clients.
// It translates domain entities to client-friendly format.
type UserResponse struct {
	ID    int    `json:"id"`    // User's unique identifier
	Name  string `json:"name"`  // User's full name
	Email string `json:"email"` // User's email address
	Age   int    `json:"age"`   // User's age
}

// UserRequest represents the expected input structure for user creation/update.
// It defines validation rules for incoming API data.
type UserRequest struct {
	Name  string `json:"name" validate:"required,min=2"`  // Name with minimum length validation
	Email string `json:"email" validate:"required,email"` // Email with format validation
	Age   int    `json:"age" validate:"gte=0,lte=120"`    // Age range validation
}

// ToUserResponse converts a domain user model to a response DTO.
func ToUserResponse(user *model.User) UserResponse {
	return UserResponse{
		ID:    user.ID(),
		Name:  user.Name(),
		Email: user.Email(),
		Age:   user.Age(),
	}
}

// ToUserResponseList converts a slice of domain user models to response DTOs.
func ToUserResponseList(users []*model.User) []UserResponse {
	result := make([]UserResponse, len(users))
	for i, user := range users {
		result[i] = ToUserResponse(user)
	}
	return result
}

package service

import (
	"errors"
	"mcanvr/example-golang-api-with-fiber/internal/domain/service"
	"mcanvr/example-golang-api-with-fiber/pkg/constants"
)

// AuthService handles user authentication and token issuance
type AuthService struct {
	userService *service.UserService
	jwtService  *JWTService
}

// NewAuthService creates a new authentication service
func NewAuthService(userService *service.UserService, jwtService *JWTService) *AuthService {
	return &AuthService{
		userService: userService,
		jwtService:  jwtService,
	}
}

// Login authenticates a user and issues a JWT token
func (s *AuthService) Login(username, password string) (string, error) {
	// Using simple authentication for demo purposes
	// In a real application, password comparison and database lookups would occur here

	// Demo static credentials
	if username == "admin" && password == "password" {
		// Generate token for admin user
		token, err := s.jwtService.GenerateToken(1, username, true)
		if err != nil {
			return "", errors.New(constants.TokenCreationFailed)
		}
		return token, nil
	}

	// In a real scenario, we would look up the user by username and validate the password
	// Example:
	// user, err := s.userService.FindByUsername(username)
	// if err != nil {
	//     return "", errors.New(constants.UserNotFound)
	// }
	// if !passwordMatches(password, user.Password()) {
	//     return "", errors.New(constants.InvalidCredentials)
	// }
	// token, err := s.jwtService.GenerateToken(user.ID(), user.Username(), user.IsAdmin())

	return "", errors.New(constants.InvalidCredentials)
}

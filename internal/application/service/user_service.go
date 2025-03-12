package service

import (
	"context"
	"mcanvr/example-golang-api-with-fiber/internal/application/dto"
	domainService "mcanvr/example-golang-api-with-fiber/internal/domain/service"
)

// UserApplicationService orchestrates the application flow for user operations.
// It coordinates domain logic and provides a use-case focused API for controllers.
type UserApplicationService struct {
	userDomainService *domainService.UserService
}

// NewUserApplicationService creates a new user application service instance.
func NewUserApplicationService(userDomainService *domainService.UserService) *UserApplicationService {
	return &UserApplicationService{
		userDomainService: userDomainService,
	}
}

// GetUserByID retrieves a user by ID and returns it as a DTO.
func (s *UserApplicationService) GetUserByID(ctx context.Context, id int) (*dto.UserResponse, error) {
	user, err := s.userDomainService.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	response := dto.ToUserResponse(user)
	return &response, nil
}

// GetAllUsers retrieves all users and returns them as DTOs.
func (s *UserApplicationService) GetAllUsers(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := s.userDomainService.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return dto.ToUserResponseList(users), nil
}

// CreateUser processes a user creation request.
func (s *UserApplicationService) CreateUser(ctx context.Context, request dto.UserRequest) (*dto.UserResponse, error) {
	// Delegate to domain service for core business logic
	user, err := s.userDomainService.CreateUser(ctx, request.Name, request.Email, request.Age)
	if err != nil {
		return nil, err
	}

	// Transform domain entity to DTO
	response := dto.ToUserResponse(user)
	return &response, nil
}

// UpdateUser processes a user update request.
func (s *UserApplicationService) UpdateUser(ctx context.Context, id int, request dto.UserRequest) (*dto.UserResponse, error) {
	// Delegate to domain service for core business logic
	user, err := s.userDomainService.UpdateUser(ctx, id, request.Name, request.Email, request.Age)
	if err != nil {
		return nil, err
	}

	// Transform domain entity to DTO
	response := dto.ToUserResponse(user)
	return &response, nil
}

// DeleteUser processes a user deletion request.
func (s *UserApplicationService) DeleteUser(ctx context.Context, id int) error {
	return s.userDomainService.DeleteUser(ctx, id)
}

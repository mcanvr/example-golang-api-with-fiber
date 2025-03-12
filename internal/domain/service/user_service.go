package service

import (
	"context"
	"errors"
	"fmt"
	"mcanvr/example-golang-api-with-fiber/internal/domain/model"
	"mcanvr/example-golang-api-with-fiber/internal/domain/repository"
	appErrors "mcanvr/example-golang-api-with-fiber/pkg/errors"
)

// Predefined domain errors
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user with this email already exists")
	ErrInvalidUserData   = errors.New("invalid user data")
	ErrRepositoryError   = errors.New("repository operation failed")
)

// UserService contains core domain logic for user operations.
// It enforces business rules that span multiple entities or repositories.
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new instance of the user domain service.
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// GetUserByID retrieves a user by ID, enforcing access rules if needed.
func (s *UserService) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: invalid ID value", ErrInvalidUserData)
	}

	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &appErrors.ErrNotFound{Resource: "user", ID: id}
	}
	return user, nil
}

// GetAllUsers retrieves all users, potentially with filtering in the future.
func (s *UserService) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	users, err := s.userRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRepositoryError, err)
	}
	return users, nil
}

// CreateUser handles the creation of a new user, enforcing uniqueness rules.
func (s *UserService) CreateUser(ctx context.Context, name, email string, age int) (*model.User, error) {
	// Check if email is already in use
	exists, err := s.userRepo.ExistsByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRepositoryError, err)
	}
	if exists {
		return nil, fmt.Errorf("%w: %s", ErrUserAlreadyExists, email)
	}

	// Create new user entity
	user, err := model.NewUser(name, email, age)
	if err != nil {
		return nil, &appErrors.ErrInvalidRequest{Field: "user data", Message: err.Error()}
	}

	// Persist the user
	if err := s.userRepo.Save(ctx, user); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRepositoryError, err)
	}

	return user, nil
}

// UpdateUser handles updating an existing user.
func (s *UserService) UpdateUser(ctx context.Context, id int, name, email string, age int) (*model.User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: invalid ID value", ErrInvalidUserData)
	}

	// Fetch existing user
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &appErrors.ErrNotFound{Resource: "user", ID: id}
	}

	// If email changed, verify it's not in use
	if user.Email() != email {
		exists, err := s.userRepo.ExistsByEmail(ctx, email)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrRepositoryError, err)
		}
		if exists {
			return nil, fmt.Errorf("%w: %s", ErrUserAlreadyExists, email)
		}
	}

	// Update the user properties with validation
	if err := user.SetName(name); err != nil {
		return nil, &appErrors.ErrInvalidRequest{Field: "name", Message: err.Error()}
	}

	if err := user.SetEmail(email); err != nil {
		return nil, &appErrors.ErrInvalidRequest{Field: "email", Message: err.Error()}
	}

	if err := user.SetAge(age); err != nil {
		return nil, &appErrors.ErrInvalidRequest{Field: "age", Message: err.Error()}
	}

	// Save changes
	if err := s.userRepo.Save(ctx, user); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRepositoryError, err)
	}

	return user, nil
}

// DeleteUser handles user deletion.
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("%w: invalid ID value", ErrInvalidUserData)
	}

	// Check if user exists
	_, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return &appErrors.ErrNotFound{Resource: "user", ID: id}
	}

	// Delete the user
	if err := s.userRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("%w: %v", ErrRepositoryError, err)
	}

	return nil
}

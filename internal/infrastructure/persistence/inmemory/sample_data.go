package inmemory

import (
	"errors"
	"mcanvr/example-golang-api-with-fiber/internal/domain/model"
	"mcanvr/example-golang-api-with-fiber/internal/domain/repository"
)

// GetSampleUsers creates a set of sample users for development and testing.
func GetSampleUsers() ([]*model.User, error) {
	// Create sample users
	user1, err := model.NewUserWithID(1, "John Doe", "john@example.com", 30)
	if err != nil {
		return nil, err
	}

	user2, err := model.NewUserWithID(2, "Jane Smith", "jane@example.com", 28)
	if err != nil {
		return nil, err
	}

	user3, err := model.NewUserWithID(3, "Bob Johnson", "bob@example.com", 45)
	if err != nil {
		return nil, err
	}

	return []*model.User{user1, user2, user3}, nil
}

// InitializeWithUsers initializes the repository with a given set of users.
// This is a helper function typically used for seeding data.
func InitializeWithUsers(repo repository.UserRepository, users []*model.User) error {
	// Try to cast to our concrete implementation
	inmemRepo, ok := repo.(*InMemoryUserRepository)
	if !ok {
		return errors.New("repository is not an InMemoryUserRepository")
	}

	// Use the internal Initialize method
	inmemRepo.Initialize(users)
	return nil
}

// InitializeWithSampleData is a convenience function that populates the repository
// with predefined sample data.
func InitializeWithSampleData(repo repository.UserRepository) error {
	users, err := GetSampleUsers()
	if err != nil {
		return err
	}
	return InitializeWithUsers(repo, users)
}

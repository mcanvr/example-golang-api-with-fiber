package repository

import (
	"context"
	"mcanvr/example-golang-api-with-fiber/internal/domain/model"
)

// UserRepository defines the contract for user persistence operations.
// This follows the Repository Pattern from DDD, which abstracts the data access layer.
type UserRepository interface {
	// FindByID retrieves a user by their unique identifier.
	FindByID(ctx context.Context, id int) (*model.User, error)

	// FindAll retrieves all users.
	FindAll(ctx context.Context) ([]*model.User, error)

	// Save persists a user entity (create or update).
	Save(ctx context.Context, user *model.User) error

	// Delete removes a user from the repository.
	Delete(ctx context.Context, id int) error

	// ExistsByEmail checks if a user with the given email exists.
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

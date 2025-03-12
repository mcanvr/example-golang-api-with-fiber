package inmemory

import (
	"context"
	"errors"
	"mcanvr/example-golang-api-with-fiber/internal/domain/model"
	"mcanvr/example-golang-api-with-fiber/internal/domain/repository"
	"sync"
)

// InMemoryUserRepository implements the UserRepository interface with an in-memory storage.
// This is primarily used for testing or small applications.
type InMemoryUserRepository struct {
	users  map[int]*model.User
	nextID int
	mu     sync.RWMutex
}

// NewInMemoryUserRepository creates a new instance of the in-memory user repository.
func NewInMemoryUserRepository() repository.UserRepository {
	return &InMemoryUserRepository{
		users:  make(map[int]*model.User),
		nextID: 1,
	}
}

// Initialize populates the repository with initial data.
func (r *InMemoryUserRepository) Initialize(users []*model.User) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range users {
		r.users[user.ID()] = user
		if user.ID() >= r.nextID {
			r.nextID = user.ID() + 1
		}
	}
}

// FindByID locates a user by their ID.
func (r *InMemoryUserRepository) FindByID(ctx context.Context, id int) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// FindAll retrieves all users.
func (r *InMemoryUserRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]*model.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

// Save creates or updates a user.
func (r *InMemoryUserRepository) Save(ctx context.Context, user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// If this is a new user (ID == 0), assign a new ID
	if user.ID() == 0 {
		// Create a new user with the assigned ID
		newUser, err := model.NewUserWithID(r.nextID, user.Name(), user.Email(), user.Age())
		if err != nil {
			return err
		}

		r.users[r.nextID] = newUser
		r.nextID++

		// Assuming User is mutable and the original reference is important,
		// we would need to update the original's ID. However, since our
		// domain model currently doesn't allow this, it would be handled differently
		// in a real application.

		return nil
	}

	// For existing users, just update the stored reference
	r.users[user.ID()] = user
	return nil
}

// Delete removes a user from the repository.
func (r *InMemoryUserRepository) Delete(ctx context.Context, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}

// ExistsByEmail checks if a user with the given email exists.
func (r *InMemoryUserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email() == email {
			return true, nil
		}
	}

	return false, nil
}

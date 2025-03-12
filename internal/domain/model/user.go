package model

import (
	"errors"
	"regexp"
)

// User represents a user entity in the domain model.
// It encapsulates user identity and enforces business rules for user data.
type User struct {
	id    int    // Private field, accessible via getter
	name  string // Private field, accessible via getter/setter
	email string // Private field, accessible via getter/setter
	age   int    // Private field, accessible via getter/setter
}

// NewUser is a factory function that creates a valid User entity.
// It enforces business rules during creation, returning errors if validation fails.
func NewUser(name, email string, age int) (*User, error) {
	u := &User{}

	if err := u.SetName(name); err != nil {
		return nil, err
	}

	if err := u.SetEmail(email); err != nil {
		return nil, err
	}

	if err := u.SetAge(age); err != nil {
		return nil, err
	}

	return u, nil
}

// NewUserWithID creates a user with a specified ID, typically used when
// reconstituting a user from persistent storage.
func NewUserWithID(id int, name, email string, age int) (*User, error) {
	u, err := NewUser(name, email, age)
	if err != nil {
		return nil, err
	}

	u.id = id
	return u, nil
}

// ID returns the user's unique identifier.
func (u *User) ID() int {
	return u.id
}

// Name returns the user's name.
func (u *User) Name() string {
	return u.name
}

// SetName updates the user's name, enforcing business rules.
func (u *User) SetName(name string) error {
	if len(name) < 2 {
		return errors.New("name must be at least 2 characters long")
	}
	u.name = name
	return nil
}

// Email returns the user's email address.
func (u *User) Email() string {
	return u.email
}

// SetEmail updates the user's email, enforcing validation rules.
func (u *User) SetEmail(email string) error {
	// Simple email validation using regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	u.email = email
	return nil
}

// Age returns the user's age.
func (u *User) Age() int {
	return u.age
}

// SetAge updates the user's age, enforcing business rules.
func (u *User) SetAge(age int) error {
	if age < 0 || age > 120 {
		return errors.New("age must be between 0 and 120")
	}
	u.age = age
	return nil
}

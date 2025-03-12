package model

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		email    string
		age      int
		wantErr  bool
	}{
		{
			name:     "Valid User",
			userName: "John Doe",
			email:    "john@example.com",
			age:      30,
			wantErr:  false,
		},
		{
			name:     "Invalid Name - Too Short",
			userName: "J",
			email:    "john@example.com",
			age:      30,
			wantErr:  true,
		},
		{
			name:     "Invalid Email",
			userName: "John Doe",
			email:    "invalid-email",
			age:      30,
			wantErr:  true,
		},
		{
			name:     "Invalid Age - Negative",
			userName: "John Doe",
			email:    "john@example.com",
			age:      -1,
			wantErr:  true,
		},
		{
			name:     "Invalid Age - Too High",
			userName: "John Doe",
			email:    "john@example.com",
			age:      150,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := NewUser(tt.userName, tt.email, tt.age)

			// Check if error is expected
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If no error is expected and user was created, check values
			if !tt.wantErr && user != nil {
				if user.Name() != tt.userName {
					t.Errorf("user.Name() = %v, want %v", user.Name(), tt.userName)
				}
				if user.Email() != tt.email {
					t.Errorf("user.Email() = %v, want %v", user.Email(), tt.email)
				}
				if user.Age() != tt.age {
					t.Errorf("user.Age() = %v, want %v", user.Age(), tt.age)
				}
			}
		})
	}
}

func TestNewUserWithID(t *testing.T) {
	id := 1
	name := "John Doe"
	email := "john@example.com"
	age := 30

	user, err := NewUserWithID(id, name, email, age)
	if err != nil {
		t.Fatalf("NewUserWithID() unexpected error = %v", err)
	}

	if user.ID() != id {
		t.Errorf("user.ID() = %v, want %v", user.ID(), id)
	}

	if user.Name() != name {
		t.Errorf("user.Name() = %v, want %v", user.Name(), name)
	}

	if user.Email() != email {
		t.Errorf("user.Email() = %v, want %v", user.Email(), email)
	}

	if user.Age() != age {
		t.Errorf("user.Age() = %v, want %v", user.Age(), age)
	}
}

func TestUserSetters(t *testing.T) {
	// Create initial user
	user, err := NewUser("Initial Name", "initial@example.com", 25)
	if err != nil {
		t.Fatalf("NewUser() unexpected error = %v", err)
	}

	// Setters test cases
	t.Run("SetName Valid", func(t *testing.T) {
		if err := user.SetName("New Name"); err != nil {
			t.Errorf("SetName() unexpected error = %v", err)
		}
		if user.Name() != "New Name" {
			t.Errorf("user.Name() = %v, want %v", user.Name(), "New Name")
		}
	})

	t.Run("SetName Invalid", func(t *testing.T) {
		if err := user.SetName(""); err == nil {
			t.Errorf("SetName() error = %v, want error", err)
		}
	})

	t.Run("SetEmail Valid", func(t *testing.T) {
		if err := user.SetEmail("new@example.com"); err != nil {
			t.Errorf("SetEmail() unexpected error = %v", err)
		}
		if user.Email() != "new@example.com" {
			t.Errorf("user.Email() = %v, want %v", user.Email(), "new@example.com")
		}
	})

	t.Run("SetEmail Invalid", func(t *testing.T) {
		if err := user.SetEmail("invalid-email"); err == nil {
			t.Errorf("SetEmail() error = %v, want error", err)
		}
	})

	t.Run("SetAge Valid", func(t *testing.T) {
		if err := user.SetAge(30); err != nil {
			t.Errorf("SetAge() unexpected error = %v", err)
		}
		if user.Age() != 30 {
			t.Errorf("user.Age() = %v, want %v", user.Age(), 30)
		}
	})

	t.Run("SetAge Invalid", func(t *testing.T) {
		if err := user.SetAge(150); err == nil {
			t.Errorf("SetAge() error = %v, want error", err)
		}
	})
}

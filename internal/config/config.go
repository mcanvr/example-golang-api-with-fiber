package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

// Config represents the application configuration loaded from environment variables.
// This centralized structure makes configuration management easier and more consistent.
type Config struct {
	ServerAddress      string `env:"SERVER_ADDRESS" envDefault:":8080"`    // HTTP server listening address and port
	Environment        string `env:"ENVIRONMENT" envDefault:"development"` // Runtime environment (development, staging, production)
	LogLevel           string `env:"LOG_LEVEL" envDefault:"info"`          // Logging verbosity level
	JWTSecret          string `env:"JWT_SECRET" envDefault:"mysecretkey"`  // Secret key for JWT token signing and verification
	JWTExpirationHours int    `env:"JWT_EXPIRATION_HOURS" envDefault:"24"` // JWT token expiration time in hours
}

// New creates a new application configuration by parsing environment variables.
// It returns a fully initialized Config struct with default values applied where needed.
// Exits the application with an error if environment variables can't be parsed.
func New() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err)
	}

	log.Printf("Configuration loaded: %+v", cfg)
	return cfg
}

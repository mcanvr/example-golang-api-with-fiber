package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

// Logger creates middleware that logs HTTP requests.
// It records request method, path, status code, and processing time.
func Logger() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Record start time
		start := time.Now()

		// Process request
		err := c.Next()

		// Calculate duration
		duration := time.Since(start)

		// Log details
		log.Printf("[%s] %s - %d (%v)",
			c.Method(),
			c.Path(),
			c.Response().StatusCode(),
			duration,
		)

		return err
	}
}

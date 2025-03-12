package middleware

import (
	"context"
	"time"

	"mcanvr/example-golang-api-with-fiber/internal/infrastructure/logger"
	"mcanvr/example-golang-api-with-fiber/internal/interfaces/common"
	"mcanvr/example-golang-api-with-fiber/pkg/constants"

	"github.com/gofiber/fiber/v3"
)

// RequestTimeout middleware adds a timeout to every request context
// If the request takes longer than the specified duration, it will be canceled
func RequestTimeout(timeout time.Duration) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Create a new context with timeout
		ctx, cancel := context.WithTimeout(c.Context(), timeout)
		defer cancel()

		// Replace the original context with the timeout context
		c.SetContext(ctx)

		// Create a channel to communicate when the request is complete
		done := make(chan error, 1)

		// Execute the request in a goroutine
		go func() {
			done <- c.Next()
		}()

		// Wait for either completion or timeout
		select {
		case err := <-done:
			// Request completed within the timeout
			return err
		case <-ctx.Done():
			// Context timed out
			if ctx.Err() == context.DeadlineExceeded {
				logger.Warn(constants.RequestTimeoutLog, c.Method(), c.Path())
				return c.Status(fiber.StatusRequestTimeout).JSON(common.NewErrorResponse(
					constants.RequestTimeout,
					constants.RequestTimeoutMessageUI,
				))
			}
			// Context was canceled for another reason
			return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(
				constants.RequestCanceled,
				constants.RequestCanceledUI,
			))
		}
	}
}

// DefaultRequestTimeout returns a timeout middleware with a sensible default timeout (10 seconds)
func DefaultRequestTimeout() fiber.Handler {
	return RequestTimeout(10 * time.Second)
}

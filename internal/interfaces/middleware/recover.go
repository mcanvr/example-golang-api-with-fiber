package middleware

import (
	"fmt"
	"mcanvr/example-golang-api-with-fiber/internal/infrastructure/logger"
	"mcanvr/example-golang-api-with-fiber/internal/interfaces/common"
	"mcanvr/example-golang-api-with-fiber/pkg/constants"

	"github.com/gofiber/fiber/v3"
)

// Recover middleware catches panic errors and prevents the application from crashing
// It returns a graceful error response to the client instead
func Recover() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Use defer to catch panics during request processing
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}

				// Log the panic
				logger.Error(constants.PanicRecovered, err)

				// Return user-friendly error response
				c.Status(fiber.StatusInternalServerError)
				_ = c.JSON(common.NewErrorResponse(
					constants.InternalServerError,
					constants.UserFriendlyServerError,
				))
			}
		}()

		// Process the request
		return c.Next()
	}
}

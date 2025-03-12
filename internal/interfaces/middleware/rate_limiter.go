package middleware

import (
	"time"

	"mcanvr/example-golang-api-with-fiber/internal/infrastructure/logger"
	"mcanvr/example-golang-api-with-fiber/internal/interfaces/common"
	"mcanvr/example-golang-api-with-fiber/pkg/constants"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

// ConfigureRateLimiter creates a rate limiter middleware with the specified configuration.
// It restricts the number of requests from a client based on IP address.
func ConfigureRateLimiter(max int, duration time.Duration) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:               max,      // Maximum number of requests
		Expiration:        duration, // Request counter reset duration
		LimiterMiddleware: limiter.SlidingWindow{},
		KeyGenerator: func(c fiber.Ctx) string {
			return c.IP() // IP based rate limiting
		},
		LimitReached: func(c fiber.Ctx) error {
			logger.Warn(constants.RateLimitExceeded, c.IP())
			return c.Status(fiber.StatusTooManyRequests).JSON(common.NewErrorResponse(
				constants.TooManyRequests,
				constants.RateLimitExceededUI,
			))
		},
	})
}

// ConfigureDefaultRateLimiter creates a rate limiter with sensible defaults:
// 50 requests per minute per IP address
func ConfigureDefaultRateLimiter() fiber.Handler {
	return ConfigureRateLimiter(50, 1*time.Minute)
}

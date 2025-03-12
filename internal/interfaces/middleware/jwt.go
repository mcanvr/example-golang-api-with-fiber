package middleware

import (
	"fmt"
	"mcanvr/example-golang-api-with-fiber/internal/application/service"
	"mcanvr/example-golang-api-with-fiber/internal/interfaces/common"
	"mcanvr/example-golang-api-with-fiber/pkg/constants"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v4"
)

// JWTProtected middleware for routes that require authentication
func JWTProtected(jwtService *service.JWTService) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Get auth header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(common.NewErrorResponse(
				constants.UnauthorizedAccess,
				constants.MissingToken,
			))
		}

		// Extract the Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(common.NewErrorResponse(
				constants.UnauthorizedAccess,
				constants.InvalidTokenFormat,
			))
		}

		tokenString := parts[1]

		// Validate the token
		token, claims, err := jwtService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(common.NewErrorResponse(
				constants.UnauthorizedAccess,
				fmt.Sprintf(constants.InvalidOrExpiredToken, err.Error()),
			))
		}

		// Store user information in locals for later use in the request lifecycle
		c.Locals("user_id", claims["user_id"])
		c.Locals("username", claims["username"])
		c.Locals("admin", claims["admin"])

		return c.Next()
	}
}

// ExtractTokenClaims extracts claims from the JWT token in the context
func ExtractTokenClaims(c fiber.Ctx) (jwt.MapClaims, error) {
	// First try to get from locals where JWTProtected middleware stores the claims
	if userId := c.Locals("user_id"); userId != nil {
		claims := jwt.MapClaims{
			"user_id":  c.Locals("user_id"),
			"username": c.Locals("username"),
			"admin":    c.Locals("admin"),
		}
		return claims, nil
	}

	// Fallback to token extraction if middleware didn't store claims
	token, ok := c.Locals("jwt").(*jwt.Token)
	if !ok {
		return nil, fmt.Errorf(constants.TokenInvalid)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf(constants.TokenInvalid)
	}

	return claims, nil
}

package api

import (
	"mcanvr/example-golang-api-with-fiber/internal/config"

	"github.com/gofiber/fiber/v3"
)

// SetupRoutes configures all API routes for the application
// It groups routes logically and applies appropriate middleware
func SetupRoutes(
	app *fiber.App,
	cfg *config.Config,
	userController *UserController,
	authController *AuthController,
	jwtMiddleware fiber.Handler,
) {
	// API group with version
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Authentication routes - public access
	v1.Post("/login", authController.Login)

	// User routes - protected with JWT authentication
	users := v1.Group("/users")
	users.Use(jwtMiddleware)

	// User CRUD operations
	users.Get("/", userController.GetUsers)
	users.Post("/", userController.CreateUser)
	users.Get("/:id", userController.GetUserByID)
	users.Put("/:id", userController.UpdateUser)
	users.Delete("/:id", userController.DeleteUser)
}

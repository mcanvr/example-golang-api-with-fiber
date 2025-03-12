package main

import (
	"mcanvr/example-golang-api-with-fiber/internal/application/service"
	"mcanvr/example-golang-api-with-fiber/internal/config"
	domainService "mcanvr/example-golang-api-with-fiber/internal/domain/service"
	"mcanvr/example-golang-api-with-fiber/internal/infrastructure/logger"
	"mcanvr/example-golang-api-with-fiber/internal/infrastructure/persistence/inmemory"
	"mcanvr/example-golang-api-with-fiber/internal/interfaces/api"
	"mcanvr/example-golang-api-with-fiber/internal/interfaces/common"
	"mcanvr/example-golang-api-with-fiber/internal/interfaces/middleware"
	"mcanvr/example-golang-api-with-fiber/pkg/constants"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v3"
)

// @title           Example Fiber API with DDD
// @version         1.0
// @description     This is an example REST API project using Domain-Driven Design principles with Fiber framework
// @termsOfService  http://swagger.io/terms/
// @contact.name    API Support
// @contact.email   support@example.com
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @host            localhost:8080
// @BasePath        /api/v1
// @schemes         http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT Authorization header using the Bearer scheme. Example: "Bearer {token}"

// @Security BearerAuth

// customErrorHandler handles all errors that occur during request processing
// It formats errors in a consistent way throughout the API
func customErrorHandler(c fiber.Ctx, err error) error {
	// Handle Fiber-specific errors
	if fiberErr, ok := err.(*fiber.Error); ok {
		return c.Status(fiberErr.Code).JSON(common.NewErrorResponse(
			constants.GeneralError,
			fiberErr.Message,
		))
	}

	// Handle other errors as general internal server errors
	logger.Error(constants.UnexpectedError, err)
	return c.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(
		constants.InternalServerError,
		constants.UserFriendlyServerError,
	))
}

func main() {
	// Setup logger
	logConfig := logger.DefaultConfig()
	if os.Getenv("ENV") == "development" {
		logConfig.Level = logger.DEBUG
	}
	appLogger := logger.NewLogger(logConfig)
	logger.SetDefaultLogger(appLogger)

	// Create configuration
	cfg := config.New()

	// Setup repositories
	userRepo := inmemory.NewInMemoryUserRepository()

	// Initialize with sample data
	if err := inmemory.InitializeWithSampleData(userRepo); err != nil {
		logger.Warn(constants.SampleDataInitFailed, err)
	}

	// Setup domain services
	userDomainService := domainService.NewUserService(userRepo)

	// Setup application services
	userAppService := service.NewUserApplicationService(userDomainService)

	// Setup JWT service
	jwtService := service.NewJWTService(cfg.JWTSecret, cfg.JWTExpirationHours)

	// Setup auth service
	authService := service.NewAuthService(userDomainService, jwtService)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "Golang Example API",
		ErrorHandler: customErrorHandler,
	})

	// Setup middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.ConfigureDefaultCORS())
	app.Use(middleware.ConfigureDefaultRateLimiter())
	app.Use(middleware.DefaultRequestTimeout())

	// Create JWT middleware
	jwtMiddleware := middleware.JWTProtected(jwtService)

	// Setup controllers
	userController := api.NewUserController(userAppService)
	authController := api.NewAuthController(authService)

	// Setup routes
	api.SetupRoutes(app, cfg, userController, authController, jwtMiddleware)

	// Serve Swagger documentation
	app.Get("/swagger/*", func(c fiber.Ctx) error {
		filename := c.Params("*")
		if filename == "" || filename == "/" {
			filename = "index.html"
		}
		return c.SendFile(filepath.Join("./static/swagger/", filename))
	})

	// Serve API documentation JSON
	app.Get("/api/swagger.json", func(c fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

	// Setup 404 handler
	app.Use(func(c fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(common.NewErrorResponse(
			constants.EndpointNotFound,
			constants.ResourceNotFound,
		))
	})

	// Start server
	logger.Info(constants.ServerStarting, cfg.ServerAddress)
	if err := app.Listen(cfg.ServerAddress); err != nil {
		logger.Fatal(constants.ServerStartFailed, err)
	}
}

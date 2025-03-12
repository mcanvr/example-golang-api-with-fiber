package api

import (
	"mcanvr/example-golang-api-with-fiber/internal/application/service"
	"mcanvr/example-golang-api-with-fiber/pkg/constants"

	"github.com/gofiber/fiber/v3"
)

// LoginRequest defines the expected format for login attempts.
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse defines the response structure for successful login.
type LoginResponse struct {
	Token string `json:"token"`
}

// AuthController handles authentication-related requests.
type AuthController struct {
	authService *service.AuthService
}

// NewAuthController creates a new instance of the auth controller.
func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Login authenticates a user and issues a JWT token.
// @Summary      User login
// @Description  Authenticates with username and password to receive a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login  body      api.LoginRequest  true  "User credentials"
// @Success      200    {object}  api.ResponseModel{data=api.LoginResponse}
// @Failure      400    {object}  api.ResponseModel
// @Failure      401    {object}  api.ResponseModel
// @Failure      500    {object}  api.ResponseModel
// @Router       /login [post]
func (c *AuthController) Login(ctx fiber.Ctx) error {
	var req LoginRequest

	// Parse and validate request body
	if err := ValidateRequest(ctx, &req); err != nil {
		return ctx.Status(err.Code).JSON(NewErrorResponse(
			constants.InvalidRequestFormat,
			err.Message,
		))
	}

	// Authenticate and get token
	token, err := c.authService.Login(req.Username, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(NewErrorResponse(
			constants.AuthenticationFailed,
			err.Error(),
		))
	}

	// Return successful response with token
	return ctx.Status(fiber.StatusOK).JSON(NewSuccessResponse(
		constants.LoginSuccess,
		LoginResponse{
			Token: token,
		},
	))
}

package api

import (
	"mcanvr/example-golang-api-with-fiber/internal/application/dto"
	"mcanvr/example-golang-api-with-fiber/internal/application/service"
	"mcanvr/example-golang-api-with-fiber/pkg/constants"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// UserController handles HTTP requests related to user management.
// It acts as a thin layer between the HTTP framework and application services.
type UserController struct {
	userAppService *service.UserApplicationService
}

// NewUserController creates a new instance of the user controller.
func NewUserController(userAppService *service.UserApplicationService) *UserController {
	return &UserController{
		userAppService: userAppService,
	}
}

// GetUsers handles the request to retrieve all users.
// @Summary      List all users
// @Description  Retrieves all users in the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  api.ResponseModel{data=[]dto.UserResponse}
// @Failure      401  {object}  api.ResponseModel  "Unauthorized"
// @Failure      500  {object}  api.ResponseModel  "Internal server error"
// @Router       /users [get]
func (c *UserController) GetUsers(ctx fiber.Ctx) error {
	users, err := c.userAppService.GetAllUsers(ctx.Context())
	if err != nil {
		return HandleDomainError(ctx, err, constants.CannotGetUsers)
	}

	return ctx.Status(fiber.StatusOK).JSON(NewSuccessResponse(
		constants.UsersFetched,
		users,
	))
}

// GetUserByID handles the request to retrieve a user by ID.
// @Summary      Show user details
// @Description  Retrieves a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  api.ResponseModel{data=dto.UserResponse}
// @Failure      400  {object}  api.ResponseModel  "Invalid ID format"
// @Failure      401  {object}  api.ResponseModel  "Unauthorized"
// @Failure      404  {object}  api.ResponseModel  "User not found"
// @Failure      500  {object}  api.ResponseModel  "Internal server error"
// @Router       /users/{id} [get]
func (c *UserController) GetUserByID(ctx fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewErrorResponse(
			constants.InvalidIDFormat,
			err.Error(),
		))
	}

	user, err := c.userAppService.GetUserByID(ctx.Context(), id)
	if err != nil {
		return HandleDomainError(ctx, err, constants.CannotGetUsers)
	}

	return ctx.Status(fiber.StatusOK).JSON(NewSuccessResponse(
		constants.UserFound,
		user,
	))
}

// CreateUser handles the request to create a new user.
// @Summary      Create new user
// @Description  Creates a new user record
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user  body      dto.UserRequest  true  "User information"
// @Success      201   {object}  api.ResponseModel{data=dto.UserResponse}
// @Failure      400   {object}  api.ResponseModel  "Invalid request or user already exists"
// @Failure      401   {object}  api.ResponseModel  "Unauthorized"
// @Failure      500   {object}  api.ResponseModel  "Internal server error"
// @Router       /users [post]
func (c *UserController) CreateUser(ctx fiber.Ctx) error {
	var userRequest dto.UserRequest

	if err := ValidateRequest(ctx, &userRequest); err != nil {
		return ctx.Status(err.Code).JSON(NewErrorResponse(
			constants.CannotCreateUser,
			err.Message,
		))
	}

	user, err := c.userAppService.CreateUser(ctx.Context(), userRequest)
	if err != nil {
		return HandleDomainError(ctx, err, constants.CannotCreateUser)
	}

	return ctx.Status(fiber.StatusCreated).JSON(NewSuccessResponse(
		constants.UserCreated,
		user,
	))
}

// UpdateUser handles the request to update an existing user.
// @Summary      Update user
// @Description  Updates a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int             true  "User ID"
// @Param        user  body      dto.UserRequest  true  "Updated user information"
// @Success      200   {object}  api.ResponseModel{data=dto.UserResponse}
// @Failure      400   {object}  api.ResponseModel  "Invalid request or email already used"
// @Failure      401   {object}  api.ResponseModel  "Unauthorized"
// @Failure      404   {object}  api.ResponseModel  "User not found"
// @Failure      500   {object}  api.ResponseModel  "Internal server error"
// @Router       /users/{id} [put]
func (c *UserController) UpdateUser(ctx fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewErrorResponse(
			constants.InvalidIDFormat,
			err.Error(),
		))
	}

	var userRequest dto.UserRequest
	if err := ValidateRequest(ctx, &userRequest); err != nil {
		return ctx.Status(err.Code).JSON(NewErrorResponse(
			constants.CannotUpdateUser,
			err.Message,
		))
	}

	user, err := c.userAppService.UpdateUser(ctx.Context(), id, userRequest)
	if err != nil {
		return HandleDomainError(ctx, err, constants.CannotUpdateUser)
	}

	return ctx.Status(fiber.StatusOK).JSON(NewSuccessResponse(
		constants.UserUpdated,
		user,
	))
}

// DeleteUser handles the request to delete a user.
// @Summary      Delete user
// @Description  Deletes a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Success      204  {object}  api.ResponseModel
// @Failure      400  {object}  api.ResponseModel  "Invalid ID format"
// @Failure      401  {object}  api.ResponseModel  "Unauthorized"
// @Failure      404  {object}  api.ResponseModel  "User not found"
// @Failure      500  {object}  api.ResponseModel  "Internal server error"
// @Router       /users/{id} [delete]
func (c *UserController) DeleteUser(ctx fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewErrorResponse(
			constants.InvalidIDFormat,
			err.Error(),
		))
	}

	err = c.userAppService.DeleteUser(ctx.Context(), id)
	if err != nil {
		return HandleDomainError(ctx, err, constants.CannotDeleteUser)
	}

	return ctx.Status(fiber.StatusNoContent).JSON(NewSuccessResponse(
		constants.UserDeleted,
		nil,
	))
}

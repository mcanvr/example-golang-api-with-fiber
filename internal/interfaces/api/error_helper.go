package api

import (
	"errors"
	domainService "mcanvr/example-golang-api-with-fiber/internal/domain/service"
	"mcanvr/example-golang-api-with-fiber/pkg/constants"
	appErrors "mcanvr/example-golang-api-with-fiber/pkg/errors"

	"github.com/gofiber/fiber/v3"
)

// HandleDomainError is a helper function that standardizes error handling for domain errors.
// It returns the appropriate HTTP status code and response based on the error type.
func HandleDomainError(c fiber.Ctx, err error, operationMsg string) error {
	// Check for specific error types first
	var notFoundErr *appErrors.ErrNotFound
	if errors.As(err, &notFoundErr) {
		return c.Status(fiber.StatusNotFound).JSON(NewErrorResponse(
			constants.UserNotFound,
			err.Error(),
		))
	}

	var invalidRequestErr *appErrors.ErrInvalidRequest
	if errors.As(err, &invalidRequestErr) {
		return c.Status(fiber.StatusBadRequest).JSON(NewErrorResponse(
			constants.InvalidRequestFormat,
			err.Error(),
		))
	}

	// Check for wrapped errors
	if errors.Is(err, domainService.ErrUserNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(NewErrorResponse(
			constants.UserNotFound,
			err.Error(),
		))
	}

	if errors.Is(err, domainService.ErrUserAlreadyExists) {
		return c.Status(fiber.StatusBadRequest).JSON(NewErrorResponse(
			constants.EmailAlreadyInUse,
			err.Error(),
		))
	}

	if errors.Is(err, domainService.ErrInvalidUserData) {
		return c.Status(fiber.StatusBadRequest).JSON(NewErrorResponse(
			constants.InvalidRequestFormat,
			err.Error(),
		))
	}

	if errors.Is(err, domainService.ErrRepositoryError) {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(
			constants.InternalServerError,
			"Your transaction cannot be processed at this time. Please try again later.",
		))
	}

	// Default case for unknown errors
	return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(
		operationMsg,
		err.Error(),
	))
}

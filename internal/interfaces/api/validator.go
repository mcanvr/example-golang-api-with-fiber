package api

import (
	"fmt"
	"mcanvr/example-golang-api-with-fiber/pkg/constants"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

// Create a single validator instance to be reused
var validate = validator.New()

// ValidateRequest handles validation of request models
// It returns a formatted error response if validation fails
func ValidateRequest(ctx fiber.Ctx, model interface{}) *fiber.Error {
	// Parse request body into model
	if err := ctx.Bind().Body(model); err != nil {
		return fiber.NewError(fiber.StatusBadRequest,
			fmt.Sprintf("%s: %s", constants.InvalidRequestFormat, err.Error()))
	}

	// Validate the model against its validation tags
	if err := validate.Struct(model); err != nil {
		// Format validation errors in a user-friendly way
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			switch err.Tag() {
			case "required":
				errorMessages = append(errorMessages, fmt.Sprintf(constants.FieldRequired, field))
			case "email":
				errorMessages = append(errorMessages, fmt.Sprintf(constants.FieldInvalidEmail, field))
			case "min":
				errorMessages = append(errorMessages, fmt.Sprintf(constants.FieldMinLength, field, err.Param()))
			case "max":
				errorMessages = append(errorMessages, fmt.Sprintf(constants.FieldMaxLength, field, err.Param()))
			case "gte":
				errorMessages = append(errorMessages, fmt.Sprintf(constants.FieldMinValue, field, err.Param()))
			case "lte":
				errorMessages = append(errorMessages, fmt.Sprintf(constants.FieldMaxValue, field, err.Param()))
			default:
				errorMessages = append(errorMessages, fmt.Sprintf(constants.FieldGenericValidation, field, err.Tag()))
			}
		}

		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf(constants.ValidationError, strings.Join(errorMessages, "; ")),
		)
	}

	return nil
}

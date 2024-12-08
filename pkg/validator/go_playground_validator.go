package validator

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

func ValidateStruct(ctx context.Context, s interface{}) error {
	return Validate.StructCtx(ctx, s)
}

func TranslateValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			// Customize error messages based on the tag
			var message string
			switch fieldError.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required", fieldError.Field())
			case "alpha":
				message = fmt.Sprintf("%s can only contain letters", fieldError.Field())
			case "min":
				message = fmt.Sprintf("%s must be at least %s characters long", fieldError.Field(), fieldError.Param())
			case "email":
				message = fmt.Sprintf("%s must be a valid email address", fieldError.Field())
			default:
				message = fmt.Sprintf("%s is invalid", fieldError.Field())
			}

			// Add to the error map
			errors[fieldError.Field()] = message
		}
	}

	return errors
}

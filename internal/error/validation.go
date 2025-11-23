package error

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

const ErrValidation ErrCode = "VALIDATION_ERROR"

var ValidationError = CustomError{
	Code:    ErrValidation,
	Message: "validation failed",
}

func HandleValidationError(err error) ErrResponse {
	var validationErrors []string
	result := ValidationError
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErr {
			validationErrors = append(validationErrors, getValidationErrorMessage(fieldErr))
		}
	} else {
		return ErrResponse{Error: result}
	}
	result.Message = validationErrors[0]
	return ErrResponse{Error: result}
}

func getValidationErrorMessage(fieldErr validator.FieldError) string {
	field := fieldErr.Field()
	tag := fieldErr.Tag()

	switch tag {
	case "required":
		return fmt.Sprintf("field '%s' is required", field)
	case "uuid":
		return fmt.Sprintf("field '%s' must be a valid UUID", field)
	case "min":
		return fmt.Sprintf("field '%s' must be at least %s characters long", field, fieldErr.Param())
	case "max":
		return fmt.Sprintf("field '%s' must be at most %s characters long", field, fieldErr.Param())
	case "gte":
		return fmt.Sprintf("field '%s' must be greater than or equal to %s", field, fieldErr.Param())
	case "lte":
		return fmt.Sprintf("field '%s' must be less than or equal to %s", field, fieldErr.Param())
	case "dive":
		return fmt.Sprintf("field '%s' contains invalid values", field)
	default:
		return fmt.Sprintf("field '%s' failed validation for tag '%s'", field, tag)
	}
}

package error

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var uuidRegex = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)

func GetValidator() *validator.Validate {
	validate := validator.New()
	err := validate.RegisterValidation("uuid", validateUUID)
	if err != nil {
		panic(err)
	}

	return validate
}

func validateUUID(fl validator.FieldLevel) bool {
	uuidStr := fl.Field().String()
	if uuidStr == "" {
		return false
	}
	return uuidRegex.MatchString(uuidStr)
}

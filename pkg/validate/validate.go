package validate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

func GetMessage(errField validator.FieldError) string {
	switch errField.Tag() {
	case "required":
		return fmt.Sprintf("Kolom '%s' harus diisi", errField.Field())
	case "email":
		return fmt.Sprintf("Kolom '%s' harus berupa email yang valid", errField.Field())
	case "min":
		return fmt.Sprintf("Kolom '%s' harus memiliki minimal %s karakter", errField.Field(), errField.Param())
	default:
		return fmt.Sprintf("Kolom '%s' tidak valid", errField.Field())
	}
}

func GetErrors(errs validator.ValidationErrors) ErrorResponse {
	var errors []Error
	for _, err := range errs {
		errors = append(errors, Error{
			Field:   err.Field(),
			Message: GetMessage(err),
		})
	}

	return ErrorResponse{
		Errors: errors,
	}
}

package validation

import (
	"strings"
	"unicode"

	jsonDTO "finpro-fenlie/model/dto/json"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/stoewer/go-strcase"
)

func GetValidationError(err error) []jsonDTO.ValidationField {
	var validationFields []jsonDTO.ValidationField

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, validationError := range ve {
			log.Debug().Msgf("validation error: %v", validationError)
			myField := convertFieldReuired(validationError.Namespace())
			validationFields = append(validationFields, jsonDTO.ValidationField{
				FieldName: myField,
				Message:   formatMessage(validationError),
			})
		}
	}

	return validationFields
}

func formatMessage(err validator.FieldError) string {
	var message string

	switch err.Tag() {
	case "required":
		message = "required"
	case "number":
		message = "must be number"
	case "email":
		message = "invalid format email"
	case "DateOnly":
		message = "invalid format date"
	case "min":
		message = "minimum value is not exceed"
	case "max":
		message = "maximum value is exceed"
	case "oneof":
		message = "value not accepted"
	case "isPassword":
		message = "the password must be a minimum of 8 characters, contain uppercase letters, lowercase letters and special characters"
	}

	return message
}

func convertFieldReuired(myValue string) string {
	log.Debug().Msgf("convertFieldRequired: %s", myValue)
	fieldSegmen := strings.Split(myValue, ".")

	var myField string
	length := len(fieldSegmen)
	i := 1
	for _, val := range fieldSegmen {
		if i == 1 {
			i++
			continue
		}

		if i == length {
			myField += strcase.SnakeCase(val)
			break
		}

		myField += strcase.LowerCamelCase(val) + "/"
		i++
	}

	return myField
}

var IsPassword validator.Func = func(fl validator.FieldLevel) bool {
	password, ok := fl.Field().Interface().(string)
	if ok {
		upper := false
		lower := false
		symbol := false

		for _, char := range password {
			if unicode.IsUpper(char) {
				upper = true
			}

			if unicode.IsNumber(char) {
				lower = true
			}

			if unicode.IsSymbol(char) || unicode.IsPunct(char) {
				symbol = true
			}
		}

		return upper && lower && symbol
	}

	return false
}

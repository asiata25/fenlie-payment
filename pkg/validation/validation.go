package validation

import (
	"errors"
	"strings"

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
	case "isPassword":
		message = "The password must be a minimum of 8 characters, contain uppercase letters, lowercase letters and special characters"
	case "oneof":
		message = "value not accepted"
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

func ValidateEmail(email, dbEmail string) error {
	if !isValidEmailFormat(email) {
		return errors.New("invalid email format")
	}

	if isEmailUsed(email, dbEmail) {
		return errors.New("email already used")
	}

	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	hasUppercase := false
	hasLowercase := false
	hasDigit := false
	hasSpecial := false

	for _, char := range password {
		if 'A' <= char && char <= 'Z' {
			hasUppercase = true
		}
		if 'a' <= char && char <= 'z' {
			hasLowercase = true
		}
		if '0' <= char && char <= '9' {
			hasDigit = true
		}
		if !('A' <= char && char <= 'Z') && !('a' <= char && char <= 'z') && !('0' <= char && char <= '9') {
			hasSpecial = true
		}
	}

	if !(hasUppercase && hasLowercase && hasDigit && hasSpecial) {
		return errors.New("password must contain uppercase, lowercase, digit, and special character")
	}

	return nil
}

func isValidEmailFormat(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func isEmailUsed(email, dbEmail string) bool {
	return email == dbEmail
}
 fb16f33 (feat: category endpoint for get, get by id, create, delete, and update)

package validation

import (
	"errors"
	"regexp"
)

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

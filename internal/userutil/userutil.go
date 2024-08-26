package userutil

import (
	"regexp"

	"github.com/xlzd/gotp"
)

func ValidateEmail(email string) bool {

	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return emailRegex.MatchString(email)
}



func ValidatePassword(password string) bool {
	return len(password) >= 8
}

func GenerateOTP() string {
	secretLength := 8
	return gotp.RandomSecret(secretLength)
}
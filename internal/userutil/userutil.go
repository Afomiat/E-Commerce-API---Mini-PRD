package userutil

import (
	"fmt"
	"regexp"
	"time"

	
	"golang.org/x/exp/rand"
)

func ValidateEmail(email string) bool {

	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return emailRegex.MatchString(email)
}



func ValidatePassword(password string) bool {
	return len(password) >= 8
}

func GenerateOTP() string {
    rand.Seed(uint64(time.Now().UnixNano())) // Seed the random number generator with int64
    otp := rand.Intn(900000) + 100000 // Generate a random number between 100000 and 999999
    return fmt.Sprintf("%06d", otp)   // Format the number as a 6-digit string
}
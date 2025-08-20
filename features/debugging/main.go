package main

import (
	"fmt"
	"unicode"
)

// verifyPassword checks if the given password meets basic rules:
// - At least 8 characters
// - Contains at least one digit
// - Contains at least one uppercase letter
// - Contains at least one lowercase letter
// - Contains at least one special character
func verifyPassword(password string) bool {
	var hasMinLen = len(password) >= 8
	var hasNumber, hasUpper, hasLower, hasSpecial bool

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			hasNumber = true
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	return hasMinLen && hasNumber && hasUpper && hasLower && hasSpecial
}

func main() {
	password := "Secure1@23!"
	fmt.Println(verifyPassword(password))

}

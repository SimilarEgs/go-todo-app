package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// this function returns bcrypted hash of the password
func GenHashPassword(passwordd string) (string, error) {

	hashed_password, err := bcrypt.GenerateFromPassword([]byte(passwordd), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}

	return string(hashed_password), nil
}

// this function checks if the given password is correct
func CheckPasswordMatch(password, hashed_password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
}

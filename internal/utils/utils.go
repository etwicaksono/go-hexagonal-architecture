package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Generate return a hashed password
func PasswordGenerate(raw string) (hashedPassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)

	if err != nil {
		return hashedPassword, err
	}

	return string(hash), nil
}

// PasswordVerify compares a hashed password with plaintext password
func PasswordVerify(hash string, raw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
}

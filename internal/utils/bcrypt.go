package utils

import (
	"fmt"

	"github.com/krobus00/auth-service/internal/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	password = fmt.Sprintf("%s%s", password, config.BcryptSalt())
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), config.BcryptCost())
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword string, password string) error {
	password = fmt.Sprintf("%s%s", password, config.BcryptSalt())
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

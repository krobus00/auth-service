package model

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrUsernameOrEmailAlreadyTaken = errors.New("username or email already taken")
	ErrWrongUsernameOrPassword     = errors.New("wrong username/email or password")
)

// MyClaims :nodoc:
type MyClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"userID"`
}

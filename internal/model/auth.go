package model

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrUsernameOrEmailAlreadyTaken = errors.New("username or email already taken")
	ErrWrongUsernameOrPassword     = errors.New("wrong username/email or password")
	ErrUnauthorizeAccess           = errors.New("unautohirze access")
)

// MyClaims :nodoc:
type MyClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"userID"`
}

type AuthUsecase interface {
	HasAccess(ctx context.Context, userID string, accessList []string) error

	// DI
	InjectUserAccessControlRepo(repo UserAccessControlRepository) error
}

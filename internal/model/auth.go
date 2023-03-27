//go:generate mockgen -destination=mock/mock_auth_usecase.go -package=mock github.com/krobus00/auth-service/internal/model AuthUsecase
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
	HasAccess(ctx context.Context, userID string, permissions []string) error

	// DI
	InjectUserGroupRepo(repo UserGroupRepository) error
}

//go:generate mockgen -destination=mock/mock_token_repository.go -package=mock github.com/krobus00/auth-service/internal/model TokenRepository

package model

import (
	"context"
	"errors"
	"fmt"

	goredis "github.com/go-redis/redis/v8"
)

type TokenType int

const (
	AccessToken TokenType = iota
	RefreshToken
)

var (
	ErrTokenInvalid     = errors.New("invalid token")
	ErrInvalidTokenType = errors.New("invalid token type")
)

type TokenRepository interface {
	Create(ctx context.Context, userID string, tokenID string, tokenType TokenType) (string, error)
	IsValidToken(ctx context.Context, userID string, tokenID string, tokenType TokenType) (bool, error)
	Revoke(ctx context.Context, userID string, tokenID string, tokenType TokenType) error

	// DI
	InjectRedisClient(client *goredis.Client) error
}

func RefreshTokenCacheKey(userID string, tokenID string) string {
	return fmt.Sprintf("refresh-token:%s:%s", userID, tokenID)
}

func AccessTokenCacheKey(userID string, tokenID string) string {
	return fmt.Sprintf("access-token:%s:%s", userID, tokenID)
}

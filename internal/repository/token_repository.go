package repository

import (
	"context"
	"time"

	goredis "github.com/go-redis/redis/v8"
	"github.com/krobus00/auth-service/internal/config"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	log "github.com/sirupsen/logrus"
)

type tokenRepository struct {
	redisClient *goredis.Client
}

// NewTokenRepository :nodoc:
func NewTokenRepository() model.TokenRepository {
	return new(tokenRepository)
}

// Create :nodoc:
func (r *tokenRepository) Create(ctx context.Context, userID string, tokenID string, tokenType model.TokenType) (string, error) {
	var (
		expDuration time.Duration
		cacheKey    string
		err         error
	)

	logger := log.WithFields(log.Fields{
		"userID": userID,
		"type":   tokenType,
	})

	switch tokenType {
	case model.AccessToken:
		expDuration = config.AccessTokenDuration()
		cacheKey = model.RefreshTokenCacheKey(userID, tokenID)
	case model.RefreshToken:
		expDuration = config.RefreshTokenDuration()
		cacheKey = model.AccessTokenCacheKey(userID, tokenID)
	default:
		err = model.ErrInvalidTokenType
	}
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	token, err := utils.GenerateToken(tokenID, userID, expDuration)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	err = SetWithExpiry(ctx, r.redisClient, cacheKey, token)

	if err != nil {
		logger.WithFields(log.Fields{
			"cacheKey": cacheKey,
		}).Error(err.Error())
		return "", err
	}

	return token, nil
}

// IsValidToken :nodoc:
func (r *tokenRepository) IsValidToken(ctx context.Context, userID string, tokenID string, tokenType model.TokenType) (bool, error) {
	var (
		cacheKey string
		token    string
		err      error
	)
	logger := log.WithFields(log.Fields{
		"userID": userID,
		"type":   tokenType,
	})
	switch tokenType {
	case model.RefreshToken:
		cacheKey = model.RefreshTokenCacheKey(userID, tokenID)
		cachedData, redisErr := Get(ctx, r.redisClient, cacheKey)
		err = redisErr
		token = string(cachedData)
	case model.AccessToken:
		cacheKey = model.AccessTokenCacheKey(userID, tokenID)
		cachedData, redisErr := Get(ctx, r.redisClient, cacheKey)
		err = redisErr
		token = string(cachedData)
	default:
		err = model.ErrInvalidTokenType
	}
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}
	if token == "" {
		return false, nil
	}
	return true, nil
}

// Revoke :nodoc:
func (r *tokenRepository) Revoke(ctx context.Context, userID string, tokenID string, tokenType model.TokenType) error {
	var (
		cacheKey string
		err      error
	)
	logger := log.WithFields(log.Fields{
		"userID": userID,
		"type":   tokenType,
	})
	switch tokenType {
	case model.RefreshToken:
		cacheKey = model.RefreshTokenCacheKey(userID, tokenID)
		err = DeleteByKeys(ctx, r.redisClient, []string{cacheKey})
	case model.AccessToken:
		cacheKey = model.AccessTokenCacheKey(userID, tokenID)
		err = DeleteByKeys(ctx, r.redisClient, []string{cacheKey})
	default:
		err = model.ErrInvalidTokenType
	}
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

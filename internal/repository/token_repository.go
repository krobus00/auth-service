package repository

import (
	"context"
	"errors"
	"time"

	goredis "github.com/go-redis/redis/v8"
	"github.com/krobus00/auth-service/internal/config"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	log "github.com/sirupsen/logrus"
)

type tokenRepository struct {
	redis *goredis.Client
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
	)

	logger := log.WithFields(log.Fields{
		"userID": userID,
		"type":   tokenType,
	})

	switch tokenType {
	case model.AccessToken:
		expDuration = config.AccessTokenDuration()
	case model.RefreshToken:
		expDuration = config.RefreshTokenDuration()
	default:
		expDuration = config.AccessTokenDuration()
	}

	token, err := utils.GenerateToken(tokenID, userID, expDuration)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	switch tokenType {
	case model.RefreshToken:
		cacheKey = model.RefreshTokenCacheKey(userID, tokenID)
		err = r.redis.HSet(ctx, cacheKey, map[string]interface{}{"token": token}).Err()
	case model.AccessToken:
		cacheKey = model.AccessTokenCacheKey(userID, tokenID)
		err = r.redis.HSet(ctx, cacheKey, map[string]interface{}{"token": token}).Err()
	default:
		err = model.ErrInvalidTokenType
	}

	if err != nil {
		logger.WithFields(log.Fields{
			"cacheKey": cacheKey,
		}).Error(err.Error())
		return "", err
	}

	err = r.redis.Expire(ctx, cacheKey, expDuration).Err()
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return token, nil
}

// IsValidToken :nodoc:
func (r *tokenRepository) IsValidToken(ctx context.Context, userID string, tokenID string, tokenType model.TokenType) (bool, error) {
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
		err = r.redis.HGet(ctx, cacheKey, "token").Err()
	case model.AccessToken:
		cacheKey = model.AccessTokenCacheKey(userID, tokenID)
		err = r.redis.HGet(ctx, cacheKey, "token").Err()
	default:
		err = model.ErrInvalidTokenType
	}
	if err != nil {
		if errors.Is(err, goredis.Nil) {
			return false, nil
		}
		logger.Error(err.Error())
		return false, err
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
		err = r.redis.Del(ctx, cacheKey).Err()
	case model.AccessToken:
		cacheKey = model.AccessTokenCacheKey(userID, tokenID)
		err = r.redis.Del(ctx, cacheKey).Err()
	default:
		err = model.ErrInvalidTokenType
	}
	if err != nil {
		if errors.Is(err, goredis.Nil) {
			return nil
		}
		logger.Error(err.Error())
		return err
	}
	return nil
}

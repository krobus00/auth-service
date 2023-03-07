package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/krobus00/auth-service/internal/config"
	"github.com/krobus00/auth-service/internal/model"
)

// GenerateToken :nodoc:
func GenerateToken(tokenID string, userID string, expDuration time.Duration) (string, error) {
	claims := model.MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expDuration)),
			Issuer:    "auth-service",
			ID:        tokenID,
		},
		UserID: userID,
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	accessToken, err := token.SignedString([]byte(config.TokenSecret()))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// ParseToken :nodoc:
func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, errors.New("signing method invalid")
		}

		return []byte(config.TokenSecret()), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, model.ErrTokenInvalid
	}

	return token, nil
}

// GetUserID :nodoc:
func GetUserID(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", model.ErrTokenInvalid
	}
	val, ok := claims["userID"]
	if !ok {
		return "", model.ErrTokenInvalid
	}
	return fmt.Sprintf("%v", val), nil

}

// GetTokenID :nodoc:
func GetTokenID(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", model.ErrTokenInvalid
	}
	val, ok := claims["jti"]
	if !ok {
		return "", model.ErrTokenInvalid
	}
	return fmt.Sprintf("%v", val), nil
}

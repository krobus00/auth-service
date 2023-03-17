package grpc

import (
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
)

func getUserIDFromAccessToken(accessToken string) (string, error) {
	token, err := utils.ParseToken(accessToken)
	if err != nil {
		return "", model.ErrTokenInvalid
	}
	userID, err := utils.GetUserID(token)
	if err != nil {
		return "", model.ErrTokenInvalid
	}

	return userID, nil

}

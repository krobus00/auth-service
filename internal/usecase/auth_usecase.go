package usecase

import (
	"context"

	"github.com/krobus00/auth-service/internal/constant"
	"github.com/krobus00/auth-service/internal/model"
)

type authUsecase struct {
	userAccessControlRepo model.UserAccessControlRepository
}

func NewAuthUsecase() model.AuthUsecase {
	return new(authUsecase)
}

func (uc *authUsecase) HasAccess(ctx context.Context, userID string, accessList []string) error {
	if userID == constant.SYSTEM_ID {
		return nil
	}
	err := uc.userAccessControlRepo.HasAccess(ctx, userID, accessList)
	if err != nil {
		return err
	}

	return nil
}

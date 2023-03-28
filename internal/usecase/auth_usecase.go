package usecase

import (
	"context"

	"github.com/krobus00/auth-service/internal/constant"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/sirupsen/logrus"
)

type authUsecase struct {
	userGroupRepo model.UserGroupRepository
}

func NewAuthUsecase() model.AuthUsecase {
	return new(authUsecase)
}

func (uc *authUsecase) HasAccess(ctx context.Context, userID string, permissions []string) error {
	logger := logrus.WithFields(logrus.Fields{
		"userID":      userID,
		"permissions": permissions,
	})

	if userID == constant.SystemID {
		return nil
	}

	userGroups, err := uc.userGroupRepo.FindByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if len(userGroups) == 0 {
		logger.Warn("user don't have any groups")
		return model.ErrUnauthorizeAccess
	}

	hasAccessCh := make(chan bool)

	for _, userGroup := range userGroups {
		go uc.userGroupHasPermissions(ctx, userGroup, permissions, hasAccessCh)
	}

	for range userGroups {
		if <-hasAccessCh {
			return nil
		}
	}

	return model.ErrUnauthorizeAccess
}

func (uc *authUsecase) userGroupHasPermissions(ctx context.Context, userGroup *model.UserGroup, permissions []string, ch chan bool) {
	for _, permission := range permissions {
		hasAccess, _ := uc.userGroupRepo.HasPermission(ctx, userGroup.GroupID, permission)
		if hasAccess {
			ch <- hasAccess
		}
	}
	ch <- false
}

package usecase

import (
	"context"
	"sync"

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

func (uc *authUsecase) HasAccess(ctx context.Context, payload *model.HasAccessPayload) error {
	logger := logrus.WithFields(logrus.Fields{
		"userID":      payload.UserID,
		"permissions": payload.Permissions,
	})

	if payload.UserID == constant.SystemID {
		return nil
	}

	userGroups, err := uc.userGroupRepo.FindByUserID(ctx, payload.UserID)
	if err != nil {
		return err
	}

	if len(userGroups) == 0 {
		logger.Warn("user don't have any groups")
		return model.ErrUnauthorizeAccess
	}

	hasAccessCh := make(chan bool)

	wg := sync.WaitGroup{}
	wg.Add(len(userGroups))
	for _, userGroup := range userGroups {
		go uc.userGroupHasPermissions(ctx, &wg, userGroup, payload.Permissions, hasAccessCh)
	}

	for range userGroups {
		if <-hasAccessCh {
			return nil
		}
	}

	wg.Wait()
	close(hasAccessCh)

	return model.ErrUnauthorizeAccess
}

func (uc *authUsecase) userGroupHasPermissions(ctx context.Context, wg *sync.WaitGroup, userGroup *model.UserGroup, permissions []string, ch chan bool) {
	defer wg.Done()
	for _, permission := range permissions {
		if permission == constant.PermissionAllowGuest {
			ch <- true
			break
		}
		hasAccess, _ := uc.userGroupRepo.HasPermission(ctx, userGroup.GroupID, permission)
		if hasAccess {
			ch <- hasAccess
		}
	}
	ch <- false
}

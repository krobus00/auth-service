package usecase

import (
	"context"
	"sync"

	"github.com/krobus00/auth-service/internal/constant"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	"github.com/sirupsen/logrus"
)

type authUsecase struct {
	userGroupRepo model.UserGroupRepository
}

func NewAuthUsecase() model.AuthUsecase {
	return new(authUsecase)
}

func (uc *authUsecase) HasAccess(ctx context.Context, payload *model.HasAccessPayload) error {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

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

	hasAccessCh := make(chan bool, len(payload.Permissions))

	wg := sync.WaitGroup{}
	for _, userGroup := range userGroups {
		wg.Add(1)
		go uc.userGroupHasPermissions(ctx, &wg, userGroup, payload.Permissions, hasAccessCh)
	}

	wg.Wait()
	close(hasAccessCh)

	for hasAccess := range hasAccessCh {
		if hasAccess {
			return nil
		}
	}

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
			break
		}
	}
	ch <- false
}

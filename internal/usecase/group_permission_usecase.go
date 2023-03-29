package usecase

import (
	"context"

	"github.com/krobus00/auth-service/internal/constant"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/sirupsen/logrus"
)

type groupPermissionUsecase struct {
	authUC              model.AuthUsecase
	groupRepo           model.GroupRepository
	permissionRepo      model.PermissionRepository
	groupPermissionRepo model.GroupPermissionRepository
}

func NewGroupPermissionUsecase() model.GroupPermissionUsecase {
	return new(groupPermissionUsecase)
}

func (uc *groupPermissionUsecase) Create(ctx context.Context, payload *model.CreateGroupPermissionPayload) (*model.GroupPermission, error) {
	logger := logrus.WithFields(logrus.Fields{
		"groupID":      payload.GroupID,
		"permissionID": payload.PermissionID,
	})

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionGroupPermissionAll,
			constant.PermissionGroupPermissionCreate,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	group, err := uc.groupRepo.FindByID(ctx, payload.GroupID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if group == nil {
		return nil, model.ErrGroupNotFound
	}

	permission, err := uc.permissionRepo.FindByID(ctx, payload.PermissionID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if permission == nil {
		return nil, model.ErrPermissionNotFound
	}

	groupPermission, err := uc.groupPermissionRepo.FindByGroupIDAndPermissionID(ctx, payload.GroupID, payload.PermissionID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if groupPermission != nil {
		return nil, model.ErrGroupPermissionAlreadyExist
	}

	data := &model.GroupPermission{
		GroupID:      payload.GroupID,
		PermissionID: payload.PermissionID,
	}

	err = uc.groupPermissionRepo.Create(ctx, data)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return data, nil
}

func (uc *groupPermissionUsecase) FindByGroupIDAndPermissionID(ctx context.Context, payload *model.FindGroupPermissionPayload) (*model.GroupPermission, error) {
	logger := logrus.WithFields(logrus.Fields{
		"groupID":      payload.GroupID,
		"permissionID": payload.PermissionID,
	})

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionGroupPermissionAll,
			constant.PermissionGroupPermissionRead,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	groupPermission, err := uc.groupPermissionRepo.FindByGroupIDAndPermissionID(ctx, payload.GroupID, payload.PermissionID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if groupPermission == nil {
		return nil, model.ErrGroupPermissionNotFound
	}

	return groupPermission, nil
}

func (uc *groupPermissionUsecase) DeleteByGroupIDAndPermissionID(ctx context.Context, payload *model.DeleteGroupPermissionPayload) error {
	logger := logrus.WithFields(logrus.Fields{
		"groupID":      payload.GroupID,
		"permissionID": payload.PermissionID,
	})

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionGroupPermissionAll,
			constant.PermissionGroupPermissionDelete,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	groupPermission, err := uc.groupPermissionRepo.FindByGroupIDAndPermissionID(ctx, payload.GroupID, payload.PermissionID)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if groupPermission == nil {
		return model.ErrGroupPermissionNotFound
	}

	err = uc.groupPermissionRepo.DeleteByGroupIDAndPermissionID(ctx, payload.GroupID, payload.PermissionID)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

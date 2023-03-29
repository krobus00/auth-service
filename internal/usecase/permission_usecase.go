package usecase

import (
	"context"

	"github.com/krobus00/auth-service/internal/constant"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	"github.com/sirupsen/logrus"
)

type permissionUsecase struct {
	authUC         model.AuthUsecase
	permissionRepo model.PermissionRepository
}

func NewPermissionUsecase() model.PermissionUsecase {
	return new(permissionUsecase)
}

func (uc *permissionUsecase) Create(ctx context.Context, payload *model.CreatePermissionPayload) (*model.Permission, error) {
	logger := logrus.WithFields(logrus.Fields{
		"name": payload.Name,
	})

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionPermissionAll,
			constant.PermissionPermissionCreate,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	existingPermission, err := uc.permissionRepo.FindByName(ctx, payload.Name)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if existingPermission != nil {
		return existingPermission, model.ErrGroupPermissionAlreadyExist
	}

	data := &model.Permission{
		ID:   utils.GenerateUUID(),
		Name: payload.Name,
	}

	err = uc.permissionRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *permissionUsecase) FindByID(ctx context.Context, payload *model.FindPermissionByIDPayload) (*model.Permission, error) {
	logger := logrus.WithFields(logrus.Fields{
		"id": payload.ID,
	})

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionPermissionAll,
			constant.PermissionPermissionRead,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	permission, err := uc.permissionRepo.FindByID(ctx, payload.ID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if permission == nil {
		return nil, model.ErrPermissionNotFound
	}

	return permission, nil
}

func (uc *permissionUsecase) FindByName(ctx context.Context, payload *model.FindPermissionByNamePayload) (*model.Permission, error) {
	logger := logrus.WithFields(logrus.Fields{
		"name": payload.Name,
	})

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionPermissionAll,
			constant.PermissionPermissionRead,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	permission, err := uc.permissionRepo.FindByName(ctx, payload.Name)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if permission == nil {
		return nil, model.ErrPermissionNotFound
	}

	return permission, nil
}

func (uc *permissionUsecase) Update(ctx context.Context, payload *model.UpdatePermissionPayload) (*model.Permission, error) {
	logger := logrus.WithFields(logrus.Fields{
		"id":   payload.ID,
		"name": payload.Name,
	})

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionPermissionAll,
			constant.PermissionPermissionUpdate,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	permission, err := uc.permissionRepo.FindByID(ctx, payload.ID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if permission == nil {
		return nil, model.ErrPermissionNotFound
	}

	err = uc.permissionRepo.Update(ctx, permission)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return permission, nil
}

func (uc *permissionUsecase) DeleteByID(ctx context.Context, payload *model.DeletePermissionByIDPayload) error {
	logger := logrus.WithFields(logrus.Fields{
		"id": payload.ID,
	})

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionPermissionAll,
			constant.PermissionPermissionDelete,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	permission, err := uc.permissionRepo.FindByID(ctx, payload.ID)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if permission == nil {
		return model.ErrPermissionNotFound
	}
	err = uc.permissionRepo.DeleteByID(ctx, payload.ID)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

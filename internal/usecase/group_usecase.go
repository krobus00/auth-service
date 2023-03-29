package usecase

import (
	"context"

	"github.com/krobus00/auth-service/internal/constant"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	"github.com/sirupsen/logrus"
)

type groupUsecase struct {
	authUC    model.AuthUsecase
	groupRepo model.GroupRepository
}

func NewGroupUsecase() model.GroupUsecase {
	return new(groupUsecase)
}

func (uc *groupUsecase) Create(ctx context.Context, payload *model.CreateGroupPayload) (*model.Group, error) {
	logger := logrus.WithFields(logrus.Fields{
		"name": payload.Name,
	})

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionGroupAll,
			constant.PermissionGroupCreate,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	existingGroup, err := uc.groupRepo.FindByName(ctx, payload.Name)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if existingGroup != nil {
		return nil, model.ErrGroupAlreadyExist
	}

	data := &model.Group{
		ID:   utils.GenerateUUID(),
		Name: payload.Name,
	}

	err = uc.groupRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *groupUsecase) FindByID(ctx context.Context, payload *model.FindGroupByIDPayload) (*model.Group, error) {
	logger := logrus.WithFields(logrus.Fields{
		"id": payload.ID,
	})

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionGroupAll,
			constant.PermissionGroupRead,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	group, err := uc.groupRepo.FindByID(ctx, payload.ID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if group == nil {
		return nil, model.ErrGroupNotFound
	}

	return group, nil
}

func (uc *groupUsecase) FindByName(ctx context.Context, payload *model.FindGroupByNamePayload) (*model.Group, error) {
	logger := logrus.WithFields(logrus.Fields{
		"name": payload.Name,
	})

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionGroupAll,
			constant.PermissionGroupRead,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	group, err := uc.groupRepo.FindByName(ctx, payload.Name)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if group == nil {
		return nil, model.ErrGroupNotFound
	}

	return group, nil
}

func (uc *groupUsecase) Update(ctx context.Context, payload *model.UpdateGroupPayload) (*model.Group, error) {
	logger := logrus.WithFields(logrus.Fields{
		"id":   payload.ID,
		"name": payload.Name,
	})

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionGroupAll,
			constant.PermissionGroupUpdate,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	group, err := uc.groupRepo.FindByID(ctx, payload.ID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if group == nil {
		return nil, model.ErrGroupNotFound
	}

	err = uc.groupRepo.Update(ctx, group)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return group, nil
}

func (uc *groupUsecase) DeleteByID(ctx context.Context, payload *model.DeleteGroupByIDPayload) error {
	logger := logrus.WithFields(logrus.Fields{
		"id": payload.ID,
	})

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionGroupAll,
			constant.PermissionGroupDelete,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	group, err := uc.groupRepo.FindByID(ctx, payload.ID)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if group == nil {
		return model.ErrGroupNotFound
	}

	err = uc.groupRepo.DeleteByID(ctx, payload.ID)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

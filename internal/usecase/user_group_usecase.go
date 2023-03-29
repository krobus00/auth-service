package usecase

import (
	"context"

	"github.com/krobus00/auth-service/internal/constant"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/sirupsen/logrus"
)

type userGroupUsecase struct {
	authUC        model.AuthUsecase
	userGroupRepo model.UserGroupRepository
	userRepo      model.UserRepository
	groupRepo     model.GroupRepository
}

func NewUserGroupUsecase() model.UserGroupUsecase {
	return new(userGroupUsecase)
}

func (uc *userGroupUsecase) Create(ctx context.Context, payload *model.CreateUserGroupPayload) (*model.UserGroup, error) {
	logger := logrus.WithFields(logrus.Fields{
		"userID":  payload.UserID,
		"groupID": payload.GroupID,
	})

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionUserGroupAll,
			constant.PermissionUserGroupCreate,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	user, err := uc.userRepo.FindByID(ctx, payload.UserID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if user == nil {
		return nil, model.ErrUserNotFound
	}

	group, err := uc.groupRepo.FindByID(ctx, payload.GroupID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if group == nil {
		return nil, model.ErrGroupNotFound
	}

	userGroup, err := uc.userGroupRepo.FindByUserIDAndGroupID(ctx, payload.UserID, payload.GroupID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if userGroup != nil {
		return nil, model.ErrUserGroupAlreadyExist
	}

	data := &model.UserGroup{
		UserID:  payload.UserID,
		GroupID: payload.GroupID,
	}

	err = uc.userGroupRepo.Create(ctx, data)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return data, nil
}

func (uc *userGroupUsecase) FindByUserIDAndGroupID(ctx context.Context, payload *model.FindUserGroupPayload) (*model.UserGroup, error) {
	logger := logrus.WithFields(logrus.Fields{
		"userID":  payload.UserID,
		"groupID": payload.GroupID,
	})

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionUserGroupAll,
			constant.PermissionUserGroupRead,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	userGroup, err := uc.userGroupRepo.FindByUserIDAndGroupID(ctx, payload.UserID, payload.GroupID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if userGroup == nil {
		return nil, model.ErrUserGroupNotFound
	}

	return userGroup, nil
}

func (uc *userGroupUsecase) DeleteByUserIDAndGroupID(ctx context.Context, payload *model.DeleteUserGroupPayload) error {
	logger := logrus.WithFields(logrus.Fields{
		"userID":  payload.UserID,
		"groupID": payload.GroupID,
	})

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionUserGroupAll,
			constant.PermissionUserGroupDelete,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	userGroup, err := uc.userGroupRepo.FindByUserIDAndGroupID(ctx, payload.UserID, payload.GroupID)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if userGroup == nil {
		return model.ErrUserGroupNotFound
	}

	err = uc.userGroupRepo.DeleteByUserIDAndGroupID(ctx, payload.UserID, payload.GroupID)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (uc *userGroupUsecase) FindByUserID(ctx context.Context, payload *model.FindUserGroupsByUserIDPayload) (model.UserGroups, error) {
	logger := logrus.WithFields(logrus.Fields{
		"userID": payload.UserID,
	})

	userGroups := make([]*model.UserGroup, 0)

	currentUserID := getUserIDFromCtx(ctx)

	err := uc.authUC.HasAccess(ctx, &model.HasAccessPayload{
		UserID: currentUserID,
		Permissions: []string{
			constant.PermissionFullAccess,
			constant.PermissionUserGroupAll,
			constant.PermissionUserGroupRead,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return userGroups, err
	}

	userGroups, err = uc.userGroupRepo.FindByUserID(ctx, payload.UserID)
	if err != nil {
		logger.Error(err.Error())
		return userGroups, nil
	}

	return userGroups, nil
}

package usecase

import (
	"errors"

	"github.com/krobus00/auth-service/internal/model"
)

func (uc *groupPermissionUsecase) InjectAuthUsecase(usecase model.AuthUsecase) error {
	if usecase == nil {
		return errors.New("invalid auth usecase")
	}
	uc.authUC = usecase
	return nil
}

func (uc *groupPermissionUsecase) InjectGroupPermissionRepo(repo model.GroupPermissionRepository) error {
	if repo == nil {
		return errors.New("invalid group permission repository")
	}
	uc.groupPermissionRepo = repo
	return nil
}

func (uc *groupPermissionUsecase) InjectGroupRepo(repo model.GroupRepository) error {
	if repo == nil {
		return errors.New("invalid group repository")
	}
	uc.groupRepo = repo
	return nil
}

func (uc *groupPermissionUsecase) InjectPermisisonRepo(repo model.PermissionRepository) error {
	if repo == nil {
		return errors.New("invalid permission repository")
	}
	uc.permissionRepo = repo
	return nil
}

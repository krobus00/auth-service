package usecase

import (
	"errors"

	"github.com/krobus00/auth-service/internal/model"
)

func (uc *permissionUsecase) InjectPermissionRepo(repo model.PermissionRepository) error {
	if repo == nil {
		return errors.New("invalid permission repository")
	}
	uc.permissionRepo = repo
	return nil
}

package usecase

import (
	"errors"

	"github.com/krobus00/auth-service/internal/model"
)

func (uc *authUsecase) InjectUserGroupRepo(repo model.UserGroupRepository) error {
	if repo == nil {
		return errors.New("invalid user group repository")
	}
	uc.userGroupRepo = repo
	return nil
}

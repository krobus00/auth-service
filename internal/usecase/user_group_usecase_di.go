package usecase

import (
	"errors"

	"github.com/krobus00/auth-service/internal/model"
)

func (uc *userGroupUsecase) InjectAuthUsecase(usecase model.AuthUsecase) error {
	if usecase == nil {
		return errors.New("invalid auth usecase")
	}
	uc.authUC = usecase
	return nil
}

func (uc *userGroupUsecase) InjectUserGroupRepo(repo model.UserGroupRepository) error {
	if repo == nil {
		return errors.New("invalid user group repository")
	}
	uc.userGroupRepo = repo
	return nil
}

func (uc *userGroupUsecase) InjectUserRepo(repo model.UserRepository) error {
	if repo == nil {
		return errors.New("invalid user repository")
	}
	uc.userRepo = repo
	return nil
}

func (uc *userGroupUsecase) InjectGroupRepo(repo model.GroupRepository) error {
	if repo == nil {
		return errors.New("invalid group repository")
	}
	uc.groupRepo = repo
	return nil
}

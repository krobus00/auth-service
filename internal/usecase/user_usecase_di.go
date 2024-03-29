package usecase

import (
	"errors"

	"github.com/krobus00/auth-service/internal/model"
	"gorm.io/gorm"
)

func (uc *userUsecase) InjectUserRepo(repo model.UserRepository) error {
	if repo == nil {
		return errors.New("invalid user repo")
	}
	uc.userRepo = repo
	return nil
}

func (uc *userUsecase) InjectTokenRepo(repo model.TokenRepository) error {
	if repo == nil {
		return errors.New("invalid token repo")
	}
	uc.tokenRepo = repo
	return nil
}

func (uc *userUsecase) InjectDB(db *gorm.DB) error {
	if db == nil {
		return errors.New("invalid db")
	}
	uc.db = db
	return nil
}

func (uc *userUsecase) InjectGroupRepo(repo model.GroupRepository) error {
	if repo == nil {
		return errors.New("invalid group repo")
	}
	uc.groupRepo = repo
	return nil
}

func (uc *userUsecase) InjectUserGroupRepo(repo model.UserGroupRepository) error {
	if repo == nil {
		return errors.New("invalid user group repo")
	}
	uc.userGroupRepo = repo
	return nil
}

package usecase

import (
	"errors"

	"github.com/krobus00/auth-service/internal/model"
)

func (uc *groupUsecase) InjectAuthUsecase(usecase model.AuthUsecase) error {
	if usecase == nil {
		return errors.New("invalid auth usecase")
	}
	uc.authUC = usecase
	return nil
}

func (uc *groupUsecase) InjectGroupRepo(repo model.GroupRepository) error {
	if repo == nil {
		return errors.New("invalid group repository")
	}
	uc.groupRepo = repo
	return nil
}

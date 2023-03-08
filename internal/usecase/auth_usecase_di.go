package usecase

import (
	"errors"

	"github.com/krobus00/auth-service/internal/model"
)

func (uc *authUsecase) InjectUserAccessControlRepo(repo model.UserAccessControlRepository) error {

	if repo == nil {
		return errors.New("invalid user access control repo")
	}
	uc.userAccessControlRepo = repo
	return nil
}

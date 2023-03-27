package usecase

import (
	"github.com/krobus00/auth-service/internal/model"
	"golang.org/x/net/context"
)

type userGroupUsecase struct {
	userGroupRepo model.UserGroupRepository
}

func (uc *userGroupUsecase) AddUserToGroup(ctx context.Context, userID string, groupID string) error {

	return nil
}

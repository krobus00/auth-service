package grpc

import (
	"errors"

	"github.com/krobus00/auth-service/internal/model"
)

func (t *Server) InjectUserUsecase(usecase model.UserUsecase) error {
	if usecase == nil {
		return errors.New("invalid user usecase")
	}
	t.userUC = usecase
	return nil
}

func (t *Server) InjectAuthUsecase(usecase model.AuthUsecase) error {
	if usecase == nil {
		return errors.New("invalid auth usecase")
	}
	t.authUC = usecase
	return nil
}

func (t *Server) InjectPermissionUsecase(usecase model.PermissionUsecase) error {
	if usecase == nil {
		return errors.New("invalid permission usecase")
	}
	t.permissionUC = usecase
	return nil
}

func (t *Server) InjectGroupUsecase(usecase model.GroupUsecase) error {
	if usecase == nil {
		return errors.New("invalid group usecase")
	}
	t.groupUC = usecase
	return nil
}

func (t *Server) InjectUserGroupUsecase(usecase model.UserGroupUsecase) error {
	if usecase == nil {
		return errors.New("invalid user group usecase")
	}
	t.userGroupUC = usecase
	return nil
}

func (t *Server) InjectGroupPermissionUsecase(usecase model.GroupPermissionUsecase) error {
	if usecase == nil {
		return errors.New("invalid group permission usecase")
	}
	t.groupPermissionUC = usecase
	return nil
}

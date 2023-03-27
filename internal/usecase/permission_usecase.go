package usecase

import (
	"context"

	"github.com/krobus00/auth-service/internal/model"
	"github.com/sirupsen/logrus"
)

type permissionUsecase struct {
	permissionRepo model.PermissionRepository
}

func NewPermissionUsecase() model.PermissionUsecase {
	return new(permissionUsecase)
}

func (uc *permissionUsecase) Create(ctx context.Context, p *model.Permission) (*model.Permission, error) {
	err := uc.permissionRepo.Create(ctx, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (uc *permissionUsecase) FindByID(ctx context.Context, id string) (*model.Permission, error) {
	permission, err := uc.permissionRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if permission == nil {
		return nil, model.ErrPermissionNotFound
	}

	return permission, nil
}

func (uc *permissionUsecase) FindByName(ctx context.Context, name string) (*model.Permission, error) {
	logger := logrus.WithFields(logrus.Fields{
		"name": name,
	})
	permission, err := uc.permissionRepo.FindByName(ctx, name)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if permission == nil {
		return nil, model.ErrPermissionNotFound
	}

	return permission, nil
}

func (uc *permissionUsecase) Update(ctx context.Context, p *model.Permission) (*model.Permission, error) {
	permission, err := uc.permissionRepo.FindByID(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	if permission == nil {
		return nil, model.ErrPermissionNotFound
	}

	err = uc.permissionRepo.Update(ctx, permission)
	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (uc *permissionUsecase) DeleteByID(ctx context.Context, id string) error {
	err := uc.permissionRepo.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

package usecase

import (
	"context"

	"github.com/krobus00/auth-service/internal/model"
)

type groupUsecase struct {
	groupRepo model.GroupRepository
}

func NewGroupUsecase() model.GroupUsecase {
	return new(groupUsecase)
}

func (uc *groupUsecase) Create(ctx context.Context, group *model.Group) (*model.Group, error) {
	err := uc.groupRepo.Create(ctx, group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (uc *groupUsecase) FindByID(ctx context.Context, id string) (*model.Group, error) {
	group, err := uc.groupRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, model.ErrGroupNotFound
	}

	return group, nil
}

func (uc *groupUsecase) FindByName(ctx context.Context, name string) (*model.Group, error) {
	group, err := uc.groupRepo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, model.ErrGroupNotFound
	}

	return group, nil
}

func (uc *groupUsecase) Update(ctx context.Context, p *model.Group) (*model.Group, error) {
	group, err := uc.groupRepo.FindByID(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, model.ErrGroupNotFound
	}

	err = uc.groupRepo.Update(ctx, group)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (uc *groupUsecase) DeleteByID(ctx context.Context, id string) error {
	err := uc.groupRepo.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

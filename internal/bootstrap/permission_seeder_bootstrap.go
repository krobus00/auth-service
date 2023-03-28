package bootstrap

import (
	"context"
	"fmt"

	"github.com/krobus00/auth-service/internal/constant"
	"github.com/krobus00/auth-service/internal/infrastructure"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/repository"
	"github.com/krobus00/auth-service/internal/usecase"
	"github.com/krobus00/auth-service/internal/utils"
	"github.com/sirupsen/logrus"
)

func StartPermissionSeeder() {
	var (
		err error
		ctx = context.Background()
	)
	// init infra
	infrastructure.InitializeDBConn()
	gormDB := infrastructure.DB

	redisClient, err := infrastructure.NewRedisClient()
	continueOrFatal(err)

	// init repo
	permissionRepo := repository.NewPermissionRepository()
	err = permissionRepo.InjectDB(gormDB)
	continueOrFatal(err)
	err = permissionRepo.InjectRedisClient(redisClient)
	continueOrFatal(err)

	groupRepo := repository.NewGroupRepository()
	err = groupRepo.InjectDB(gormDB)
	continueOrFatal(err)
	err = groupRepo.InjectRedisClient(redisClient)
	continueOrFatal(err)

	// init usecase
	permissionUsecase := usecase.NewPermissionUsecase()
	err = permissionUsecase.InjectPermissionRepo(permissionRepo)
	continueOrFatal(err)

	groupUsecase := usecase.NewGroupUsecase()
	err = groupUsecase.InjectGroupRepo(groupRepo)
	continueOrFatal(err)

	permissions := []*model.Permission{
		{
			ID:   utils.GenerateUUID(),
			Name: constant.PermissionFullAccess,
		},
	}

	for _, permission := range permissions {
		currentPermission, _ := permissionUsecase.FindByName(ctx, permission.Name)
		if currentPermission == nil {
			logrus.Info("created permission")
			_, err = permissionUsecase.Create(ctx, permission)
			continueOrFatal(err)
			logrus.Info(fmt.Sprintf("permission %s created", permission.Name))
		} else {
			logrus.Info(fmt.Sprintf("permission %s already exist", permission.Name))
		}
	}

	groups := []*model.Group{
		{
			ID:   utils.GenerateUUID(),
			Name: constant.GroupSuperUser,
		},
		{
			ID:   utils.GenerateUUID(),
			Name: constant.GroupDefault,
		},
	}

	for _, group := range groups {
		currentGroup, _ := groupUsecase.FindByName(ctx, group.Name)
		if currentGroup == nil {
			logrus.Info("created group")
			_, err = groupUsecase.Create(ctx, group)
			continueOrFatal(err)
			logrus.Info(fmt.Sprintf("group %s created", group.Name))
		} else {
			logrus.Info(fmt.Sprintf("group %s already exist", group.Name))
		}
	}
}

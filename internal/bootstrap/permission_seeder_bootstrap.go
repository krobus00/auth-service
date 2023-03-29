package bootstrap

import (
	"context"
	"fmt"

	"github.com/krobus00/auth-service/internal/constant"
	"github.com/krobus00/auth-service/internal/infrastructure"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/repository"
	"github.com/krobus00/auth-service/internal/usecase"
	"github.com/sirupsen/logrus"
)

func StartPermissionSeeder() {
	var (
		err error
		ctx = context.Background()
	)

	ctx = setUserIDCtx(ctx, constant.SystemID)
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

	userGroupRepo := repository.NewUserGroupRepository()
	err = userGroupRepo.InjectDB(infrastructure.DB)
	continueOrFatal(err)
	err = userGroupRepo.InjectRedisClient(redisClient)
	continueOrFatal(err)

	groupPermissionRepo := repository.NewGroupPermissionRepository()
	err = groupPermissionRepo.InjectDB(infrastructure.DB)
	continueOrFatal(err)
	err = groupPermissionRepo.InjectRedisClient(redisClient)
	continueOrFatal(err)

	// init usecase
	authUsecase := usecase.NewAuthUsecase()
	err = authUsecase.InjectUserGroupRepo(userGroupRepo)
	continueOrFatal(err)

	permissionUsecase := usecase.NewPermissionUsecase()
	err = permissionUsecase.InjectPermissionRepo(permissionRepo)
	continueOrFatal(err)
	err = permissionUsecase.InjectAuthUsecase(authUsecase)
	continueOrFatal(err)

	groupUsecase := usecase.NewGroupUsecase()
	err = groupUsecase.InjectGroupRepo(groupRepo)
	continueOrFatal(err)
	err = groupUsecase.InjectAuthUsecase(authUsecase)
	continueOrFatal(err)

	groupPermissionUsecase := usecase.NewGroupPermissionUsecase()
	err = groupPermissionUsecase.InjectGroupRepo(groupRepo)
	continueOrFatal(err)
	err = groupPermissionUsecase.InjectPermisisonRepo(permissionRepo)
	continueOrFatal(err)
	err = groupPermissionUsecase.InjectGroupPermissionRepo(groupPermissionRepo)
	continueOrFatal(err)
	err = groupPermissionUsecase.InjectAuthUsecase(authUsecase)
	continueOrFatal(err)

	for _, permission := range constant.SeedPermissions {
		currentPermission, _ := permissionUsecase.FindByName(ctx, &model.FindPermissionByNamePayload{
			Name: permission,
		})
		if currentPermission == nil {
			_, err = permissionUsecase.Create(ctx, &model.CreatePermissionPayload{
				Name: permission,
			})
			continueOrFatal(err)
			logrus.Info(fmt.Sprintf("permission %s created", permission))
		} else {
			logrus.Info(fmt.Sprintf("permission %s already exist", permission))
		}
	}

	for _, group := range constant.SeedGroups {
		currentGroup, _ := groupUsecase.FindByName(ctx, &model.FindGroupByNamePayload{
			Name: group,
		})
		if currentGroup == nil {
			currentGroup, err = groupUsecase.Create(ctx, &model.CreateGroupPayload{
				Name: group,
			})
			continueOrFatal(err)
			logrus.Info(fmt.Sprintf("group %s created", group))
		} else {
			logrus.Info(fmt.Sprintf("group %s already exist", group))
		}
		groupPermissions := constant.SeedGroupPermissios[group]
		for _, groupPermission := range groupPermissions {
			currentPermission, _ := permissionUsecase.FindByName(ctx, &model.FindPermissionByNamePayload{
				Name: groupPermission,
			})
			if currentPermission != nil {
				currentGroupPermission, _ := groupPermissionUsecase.FindByGroupIDAndPermissionID(ctx, &model.FindGroupPermissionPayload{
					GroupID:      currentGroup.ID,
					PermissionID: currentPermission.ID,
				})
				if currentGroupPermission == nil {
					_, err := groupPermissionUsecase.Create(ctx, &model.CreateGroupPermissionPayload{
						GroupID:      currentGroup.ID,
						PermissionID: currentPermission.ID,
					})
					continueOrFatal(err)
				}
			}
		}
	}
}

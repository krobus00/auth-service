package model

import (
	"context"
	"fmt"

	goredis "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserGroup struct {
	UserID  string
	GroupID string
}

func NewUserGroupCacheKeyByUserID(userID string) string {
	return fmt.Sprintf("user-groups:userID:%s", userID)
}

func NewUserGroupCacheKeyByUserIDAndGroupID(userID string, groupID string) string {
	return fmt.Sprintf("user-groups:userID:%s:groupID:%s", userID, groupID)
}

type GroupPermissionAccess struct {
	GroupID        string
	PermissionName string
}

func NewGroupPermissionCacheKey(groupID string) string {
	return fmt.Sprintf("user-groups:groupID:%s:permissions", groupID)
}

func GetUserGroupCacheKeys(userID string, groupID string) []string {
	return []string{
		NewUserGroupCacheKeyByUserID(userID),
		NewUserGroupCacheKeyByUserIDAndGroupID(userID, groupID),
		NewGroupPermissionCacheKey(groupID),
		"user-groups:*",
	}
}

type UserGroupRepository interface {
	Create(ctx context.Context, data *UserGroup) error
	FindByUserIDAndGroupID(ctx context.Context, userID, groupID string) (*UserGroup, error)
	DeleteByUserIDAndGroupID(ctx context.Context, userID, groupID string) error
	FindByUserID(ctx context.Context, userID string) ([]*UserGroup, error)

	HasPermission(ctx context.Context, groupID string, permission string) (bool, error)

	// DI
	InjectDB(db *gorm.DB) error
	InjectRedisClient(client *goredis.Client) error
}

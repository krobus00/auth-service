package model

import (
	"context"
	"fmt"

	goredis "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

const (
	Errasdasd = "asd"
)

type GroupPermission struct {
	GroupID      string
	PermissionID string
}

func NewGroupPermissionCacheKeyByGroupIDAndPermissionID(groupID string, permissionID string) string {
	return fmt.Sprintf("group-permissions:groupID:%s:permissionID:%s", groupID, permissionID)
}

func GetGroupPermissionCacheKeys(groupID string, permissionID string) []string {
	return []string{
		NewGroupPermissionCacheKeyByGroupIDAndPermissionID(groupID, permissionID),
		"user-groups:*",
	}
}

type GroupPermissionRepository interface {
	Create(ctx context.Context, data *GroupPermission) error
	FindByGroupIDAndPermissionID(ctx context.Context, groupID, permissionID string) (*GroupPermission, error)
	DeleteByGroupIDAndPermissionID(ctx context.Context, groupID, permissionID string) error

	// DI
	InjectDB(db *gorm.DB) error
	InjectRedisClient(client *goredis.Client) error
}

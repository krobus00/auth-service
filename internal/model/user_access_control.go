package model

import (
	"context"
	"fmt"

	goredis "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserAccessControl struct {
	UserID string
	ACID   string
}

func (UserAccessControl) TableName() string {
	return "user_access_controls"
}

type UserAccessControlRepository interface {
	Create(ctx context.Context, uac *UserAccessControl) error
	Revoke(ctx context.Context, uac *UserAccessControl) error
	HasAccess(ctx context.Context, userID string, accessList []string) error

	// DI
	InjectDB(db *gorm.DB) error
	InjectRedisClient(client *goredis.Client) error
}

func UserAccessControlCacheKey(userID string) string {
	return fmt.Sprintf("acl:%s", userID)
}

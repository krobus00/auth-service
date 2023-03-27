//go:generate mockgen -destination=mock/mock_access_control_repository.go -package=mock github.com/krobus00/auth-service/internal/model PermissionRepository
package model

import (
	"context"
	"errors"
	"fmt"

	goredis "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	ErrPermissionNotFound = errors.New("permission not found")
)

type Permission struct {
	ID   string
	Name string
}

func (Permission) TableName() string {
	return "permissions"
}

func NewPermissionCacheKeyByID(id string) string {
	return fmt.Sprintf("permission:id:%s", id)
}

func NewPermissionCacheKeyByName(name string) string {
	return fmt.Sprintf("permission:name:%s", name)
}

func GetPermissionCacheKeys(id string, name string) []string {
	return []string{
		NewPermissionCacheKeyByID(id),
		NewPermissionCacheKeyByName(name),
	}
}

type PermissionRepository interface {
	Create(ctx context.Context, permission *Permission) error
	FindByID(ctx context.Context, id string) (*Permission, error)
	FindByName(ctx context.Context, name string) (*Permission, error)
	Update(ctx context.Context, permission *Permission) error
	DeleteByID(ctx context.Context, id string) error

	// DI
	InjectDB(db *gorm.DB) error
	InjectRedisClient(client *goredis.Client) error
}

type PermissionUsecase interface {
	Create(ctx context.Context, permission *Permission) (*Permission, error)
	FindByID(ctx context.Context, id string) (*Permission, error)
	FindByName(ctx context.Context, name string) (*Permission, error)
	Update(ctx context.Context, permission *Permission) (*Permission, error)
	DeleteByID(ctx context.Context, id string) error

	// DI
	InjectPermissionRepo(repo PermissionRepository) error
}

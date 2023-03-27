package model

import (
	"context"
	"errors"
	"fmt"

	goredis "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	ErrGroupNotFound = errors.New("group not found")
)

type Group struct {
	ID   string
	Name string
}

func (Group) TableName() string {
	return "groups"
}

func NewGroupCacheKeyByID(id string) string {
	return fmt.Sprintf("groups:id:%s", id)
}

func NewGroupCacheKeyByName(name string) string {
	return fmt.Sprintf("groups:name:%s", name)
}

func GetGroupCacheKeys(id string, name string) []string {
	return []string{
		NewGroupCacheKeyByID(id),
		NewGroupCacheKeyByName(name),
	}
}

type GroupRepository interface {
	Create(ctx context.Context, group *Group) error
	FindByID(ctx context.Context, id string) (*Group, error)
	FindByName(ctx context.Context, name string) (*Group, error)
	Update(ctx context.Context, group *Group) error
	DeleteByID(ctx context.Context, id string) error

	// DI
	InjectDB(db *gorm.DB) error
	InjectRedisClient(client *goredis.Client) error
}

type GroupUsecase interface {
	Create(ctx context.Context, group *Group) (*Group, error)
	FindByID(ctx context.Context, id string) (*Group, error)
	FindByName(ctx context.Context, name string) (*Group, error)
	Update(ctx context.Context, group *Group) (*Group, error)
	DeleteByID(ctx context.Context, id string) error

	// DI
	InjectGroupRepo(repo GroupRepository) error
}

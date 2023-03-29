//go:generate mockgen -destination=mock/mock_permission_repository.go -package=mock github.com/krobus00/auth-service/internal/model PermissionRepository
//go:generate mockgen -destination=mock/mock_permission_usecase.go -package=mock github.com/krobus00/auth-service/internal/model PermissionUsecase

package model

import (
	"context"
	"errors"
	"fmt"

	goredis "github.com/go-redis/redis/v8"
	pb "github.com/krobus00/auth-service/pb/auth"
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

func (m *Permission) ToGRPCResponse() *pb.Permission {
	return &pb.Permission{
		Id:   m.ID,
		Name: m.Name,
	}
}

type CreatePermissionPayload struct {
	Name string
}

func (m *CreatePermissionPayload) ParseFromProto(req *pb.CreatePermissionRequest) {
	m.Name = req.GetName()
}

type FindPermissionByIDPayload struct {
	ID string
}

func (m *FindPermissionByIDPayload) ParseFromProto(req *pb.FindPermissionByIDRequest) {
	m.ID = req.GetId()
}

type FindPermissionByNamePayload struct {
	Name string
}

func (m *FindPermissionByNamePayload) ParseFromProto(req *pb.FindPermissionByNameRequest) {
	m.Name = req.GetName()
}

type UpdatePermissionPayload struct {
	ID   string
	Name string
}

type DeletePermissionByIDPayload struct {
	ID string
}

func (m *DeletePermissionByIDPayload) ParseFromProto(req *pb.DeletePermissionRequest) {
	m.ID = req.GetId()
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
	Create(ctx context.Context, payload *CreatePermissionPayload) (*Permission, error)
	FindByID(ctx context.Context, payload *FindPermissionByIDPayload) (*Permission, error)
	FindByName(ctx context.Context, payload *FindPermissionByNamePayload) (*Permission, error)
	Update(ctx context.Context, payload *UpdatePermissionPayload) (*Permission, error)
	DeleteByID(ctx context.Context, payload *DeletePermissionByIDPayload) error

	// DI
	InjectAuthUsecase(usecase AuthUsecase) error
	InjectPermissionRepo(repo PermissionRepository) error
}

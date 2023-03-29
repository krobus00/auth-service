//go:generate mockgen -destination=mock/mock_group_permission_repository.go -package=mock github.com/krobus00/auth-service/internal/model GroupPermissionRepository
//go:generate mockgen -destination=mock/mock_group_permission_usecase.go -package=mock github.com/krobus00/auth-service/internal/model GroupPermissionUsecase

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
	ErrGroupPermissionNotFound     = errors.New("group permission not found")
	ErrGroupPermissionAlreadyExist = errors.New("group permission already exist")
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

func (m *GroupPermission) ToGRPCResponse() *pb.GroupPermission {
	return &pb.GroupPermission{
		GroupId:      m.GroupID,
		PermissionId: m.PermissionID,
	}
}

type CreateGroupPermissionPayload struct {
	GroupID      string
	PermissionID string
}

func (m *CreateGroupPermissionPayload) ParseFromProto(req *pb.CreateGroupPermissionRequest) {
	m.GroupID = req.GetGroupId()
	m.PermissionID = req.GetPermissionId()
}

type FindGroupPermissionPayload struct {
	GroupID      string
	PermissionID string
}

func (m *FindGroupPermissionPayload) ParseFromProto(req *pb.FindGroupPermissionRequest) {
	m.GroupID = req.GetGroupId()
	m.PermissionID = req.GetPermissionId()
}

type DeleteGroupPermissionPayload struct {
	GroupID      string
	PermissionID string
}

func (m *DeleteGroupPermissionPayload) ParseFromProto(req *pb.DeleteGroupPermissionRequest) {
	m.GroupID = req.GetGroupId()
	m.PermissionID = req.GetPermissionId()
}

type GroupPermissionRepository interface {
	Create(ctx context.Context, data *GroupPermission) error
	FindByGroupIDAndPermissionID(ctx context.Context, groupID, permissionID string) (*GroupPermission, error)
	DeleteByGroupIDAndPermissionID(ctx context.Context, groupID, permissionID string) error

	// DI
	InjectDB(db *gorm.DB) error
	InjectRedisClient(client *goredis.Client) error
}

type GroupPermissionUsecase interface {
	Create(ctx context.Context, payload *CreateGroupPermissionPayload) (*GroupPermission, error)
	FindByGroupIDAndPermissionID(ctx context.Context, payload *FindGroupPermissionPayload) (*GroupPermission, error)
	DeleteByGroupIDAndPermissionID(ctx context.Context, payload *DeleteGroupPermissionPayload) error

	// DI
	InjectAuthUsecase(usecase AuthUsecase) error
	InjectGroupPermissionRepo(repo GroupPermissionRepository) error
	InjectGroupRepo(repo GroupRepository) error
	InjectPermisisonRepo(repo PermissionRepository) error
}

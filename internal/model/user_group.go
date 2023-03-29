//go:generate mockgen -destination=mock/mock_user_group_repository.go -package=mock github.com/krobus00/auth-service/internal/model UserGroupRepository
//go:generate mockgen -destination=mock/mock_user_group_usecase.go -package=mock github.com/krobus00/auth-service/internal/model UserGroupUsecase

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
	ErrUserGroupNotFound     = errors.New("user group not found")
	ErrUserGroupAlreadyExist = errors.New("user group already exist")
)

type UserGroup struct {
	UserID  string
	GroupID string
}

type UserGroups []*UserGroup

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

func (m *UserGroup) ToGRPCResponse() *pb.UserGroup {
	return &pb.UserGroup{
		UserId:  m.UserID,
		GroupId: m.GroupID,
	}
}

func (m UserGroups) ToGRPCResponse() *pb.FindAllUserGroupsResponse {
	res := make([]*pb.UserGroup, 0)
	for _, userGroup := range m {
		res = append(res, userGroup.ToGRPCResponse())
	}
	return &pb.FindAllUserGroupsResponse{
		UserGroups: res,
	}
}

type CreateUserGroupPayload struct {
	UserID  string
	GroupID string
}

func (m *CreateUserGroupPayload) ParseFromProto(req *pb.CreateUserGroupRequest) {
	m.UserID = req.GetUserId()
	m.GroupID = req.GetGroupId()
}

type FindUserGroupPayload struct {
	UserID  string
	GroupID string
}

func (m *FindUserGroupPayload) ParseFromProto(req *pb.FindUserGroupRequest) {
	m.UserID = req.GetUserId()
	m.GroupID = req.GetGroupId()
}

type DeleteUserGroupPayload struct {
	UserID  string
	GroupID string
}

func (m *DeleteUserGroupPayload) ParseFromProto(req *pb.DeleteUserGroupRequest) {
	m.UserID = req.GetUserId()
	m.GroupID = req.GetGroupId()
}

type FindUserGroupsByUserIDPayload struct {
	UserID string
}

func (m *FindUserGroupsByUserIDPayload) ParseFromProto(req *pb.FindAllUserGroupsRequest) {
	m.UserID = req.GetUserId()
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

type UserGroupUsecase interface {
	Create(ctx context.Context, payload *CreateUserGroupPayload) (*UserGroup, error)
	FindByUserIDAndGroupID(ctx context.Context, payload *FindUserGroupPayload) (*UserGroup, error)
	DeleteByUserIDAndGroupID(ctx context.Context, payload *DeleteUserGroupPayload) error
	FindByUserID(ctx context.Context, payload *FindUserGroupsByUserIDPayload) (UserGroups, error)

	// DI
	InjectAuthUsecase(usecase AuthUsecase) error
	InjectUserGroupRepo(repo UserGroupRepository) error
	InjectUserRepo(repo UserRepository) error
	InjectGroupRepo(repo GroupRepository) error
}

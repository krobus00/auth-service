//go:generate mockgen -destination=mock/mock_group_repository.go -package=mock github.com/krobus00/auth-service/internal/model GroupRepository
//go:generate mockgen -destination=mock/mock_group_usecase.go -package=mock github.com/krobus00/auth-service/internal/model GroupUsecase

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
	ErrGroupAlreadyExist = errors.New("group already exist")
	ErrGroupNotFound     = errors.New("group not found")
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

func (m *Group) ToGRPCResponse() *pb.Group {
	return &pb.Group{
		Id:   m.ID,
		Name: m.Name,
	}
}

type CreateGroupPayload struct {
	Name string
}

func (m *CreateGroupPayload) ParseFromProto(req *pb.CreateGroupRequest) {
	m.Name = req.GetName()
}

type FindGroupByIDPayload struct {
	ID string
}

func (m *FindGroupByIDPayload) ParseFromProto(req *pb.FindGroupByIDRequest) {
	m.ID = req.GetId()
}

type FindGroupByNamePayload struct {
	Name string
}

func (m *FindGroupByNamePayload) ParseFromProto(req *pb.FindGroupByNameRequest) {
	m.Name = req.GetName()
}

type UpdateGroupPayload struct {
	ID   string
	Name string
}

type DeleteGroupByIDPayload struct {
	ID string
}

func (m *DeleteGroupByIDPayload) ParseFromProto(req *pb.DeleteGroupRequest) {
	m.ID = req.GetId()
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
	Create(ctx context.Context, payload *CreateGroupPayload) (*Group, error)
	FindByID(ctx context.Context, payload *FindGroupByIDPayload) (*Group, error)
	FindByName(ctx context.Context, payload *FindGroupByNamePayload) (*Group, error)
	Update(ctx context.Context, payload *UpdateGroupPayload) (*Group, error)
	DeleteByID(ctx context.Context, payload *DeleteGroupByIDPayload) error

	// DI
	InjectAuthUsecase(usecase AuthUsecase) error
	InjectGroupRepo(repo GroupRepository) error
}

//go:generate mockgen -destination=mock/mock_user_repository.go -package=mock github.com/krobus00/auth-service/internal/model UserRepository
//go:generate mockgen -destination=mock/mock_user_usecase.go -package=mock github.com/krobus00/auth-service/internal/model UserUsecase
package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	goredis "github.com/go-redis/redis/v8"
	pb "github.com/krobus00/auth-service/pb/auth"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	ID        string
	FullName  string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewUserCacheKeyByID(id string) string {
	return fmt.Sprintf("users:id:%s", id)
}

func NewUserCacheKeyByUsername(username string) string {
	return fmt.Sprintf("users:username:%s", username)
}

func NewUserCacheKeyByEmail(email string) string {
	return fmt.Sprintf("users:email:%s", email)
}

func GetUserCacheKeys(id string, username string, email string) []string {
	return []string{
		NewUserCacheKeyByID(id),
		NewUserCacheKeyByUsername(username),
		NewUserCacheKeyByEmail(email),
	}
}

// Usecase payload

type UserRegistrationPayload struct {
	FullName string
	Username string
	Email    string
	Password string
}

func (m *UserRegistrationPayload) ParseFromProto(req *pb.RegisterRequest) {
	m.FullName = req.GetEmail()
	m.Username = req.GetUsername()
	m.Email = req.GetEmail()
	m.Password = req.GetPassword()
}

type UserLoginPayload struct {
	Username string
	Password string
}

func (m *UserLoginPayload) ParseFromProto(req *pb.LoginRequest) {
	m.Username = req.GetUsername()
	m.Password = req.GetPassword()
}

type GetUserInfoPayload struct {
	ID string
}

type AuthResponse struct {
	AccessToken  string
	RefreshToken string
}

func (m *AuthResponse) ToGRPCResponse() *pb.AuthResponse {
	return &pb.AuthResponse{
		AccessToken:  m.AccessToken,
		RefreshToken: m.RefreshToken,
	}
}

type UserInfoResponse struct {
	ID        string
	FullName  string
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (m *UserInfoResponse) ToGRPCResponse() *pb.User {
	createdAt := m.CreatedAt.Format(time.RFC3339Nano)
	updatedAt := m.UpdatedAt.Format(time.RFC3339Nano)
	return &pb.User{
		Id:        m.ID,
		FullName:  m.FullName,
		Username:  m.Username,
		Email:     m.Email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type RefreshTokenPayload struct {
	UserID  string
	TokenID string
}

func (m *RefreshTokenPayload) ParseFromProto(req *pb.RefreshTokenRequest) {
	m.UserID = req.GetSessionUserId()
	m.TokenID = req.GetTokenId()
}

type UserLogoutPayload struct {
	UserID  string
	TokenID string
}

func (m *UserLogoutPayload) ParseFromProto(req *pb.LogoutRequest) {
	m.UserID = req.GetSessionUserId()
	m.TokenID = req.GetTokenId()
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	UpdateByID(ctx context.Context, id string) (*User, error)
	DeleteByID(ctx context.Context, id string) error

	// DI
	InjectDB(db *gorm.DB) error
	InjectRedisClient(client *goredis.Client) error
}

type UserUsecase interface {
	Register(ctx context.Context, payload *UserRegistrationPayload) (*AuthResponse, error)
	Login(ctx context.Context, payload *UserLoginPayload) (*AuthResponse, error)
	GetUserInfo(ctx context.Context, payload *GetUserInfoPayload) (*UserInfoResponse, error)
	RefreshToken(ctx context.Context, payload *RefreshTokenPayload) (*AuthResponse, error)
	Logout(ctx context.Context, payload *UserLogoutPayload) error

	// DI
	InjectDB(db *gorm.DB) error
	InjectTokenRepo(repo TokenRepository) error
	InjectUserRepo(repo UserRepository) error
	InjectGroupRepo(repo GroupRepository) error
	InjectUserGroupRepo(repo UserGroupRepository) error
}

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

// HTTP DTO

type HTTPUserRegistrationRequest struct {
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"Password"`
}

func (req HTTPUserRegistrationRequest) ToPayload() *UserRegistrationPayload {
	return &UserRegistrationPayload{
		FullName: req.FullName,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

type HTTPUserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req HTTPUserLoginRequest) ToPayload() *UserLoginPayload {
	return &UserLoginPayload{
		Username: req.Username,
		Password: req.Password,
	}
}

type HTTPAuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type HTTPUserInfoResponse struct {
	ID        string     `json:"id"`
	FullName  string     `json:"fullName"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `josn:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// Usecase payload

type UserRegistrationPayload struct {
	FullName string
	Username string
	Email    string
	Password string
}

type UserLoginPayload struct {
	Username string
	Password string
}

type GetUserInfoPayload struct {
	ID string
}

type AuthResponse struct {
	AccessToken  string
	RefreshToken string
}

func (res *AuthResponse) ToHTTPResponse() *HTTPAuthResponse {
	return &HTTPAuthResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}
}

func (res *AuthResponse) ToGRPCResponse() *pb.AuthResponse {
	return &pb.AuthResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
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

func (res *UserInfoResponse) ToHTTPResponse() *HTTPUserInfoResponse {
	return &HTTPUserInfoResponse{
		ID:        res.ID,
		FullName:  res.FullName,
		Username:  res.Username,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		DeletedAt: res.DeletedAt,
	}
}

func (res *UserInfoResponse) ToGRPCResponse() *pb.User {
	createdAt := res.CreatedAt.Format(time.RFC3339Nano)
	updatedAt := res.UpdatedAt.Format(time.RFC3339Nano)
	return &pb.User{
		Id:        res.ID,
		FullName:  res.FullName,
		Username:  res.Username,
		Email:     res.Email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type RefreshTokenPayload struct {
	UserID  string
	TokenID string
}

type LogoutPayload struct {
	UserID  string
	TokenID string
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
	Logout(ctx context.Context, payload *LogoutPayload) error

	// DI
	InjectDB(db *gorm.DB) error
	InjectTokenRepo(repo TokenRepository) error
	InjectUserRepo(repo UserRepository) error
}

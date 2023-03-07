package model

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// HTTP DTO
// HTTPUserRegistrationRequest :nodoc:
type HTTPUserRegistrationRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"Password"`
}

func (req HTTPUserRegistrationRequest) ToPayload() *UserRegistrationPayload {
	return &UserRegistrationPayload{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

// HTTPUserLoginRequest :nodoc:
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

// HTTPAuthResponse :nodoc:
type HTTPAuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// HTTPUserInfoResponse :nodoc:
type HTTPUserInfoResponse struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `josn:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// Usecase payload
// UserRegistrationPayload :nodoc:
type UserRegistrationPayload struct {
	Username string
	Email    string
	Password string
}

// UserLoginPayload :nodoc:
type UserLoginPayload struct {
	Username string
	Password string
}

type GetUserInfoPayload struct {
	ID string
}

// AuthResponse :nodoc:
type AuthResponse struct {
	AccessToken  string
	RefreshToken string
}

// ToHTTPResponse :nodoc:
func (res *AuthResponse) ToHTTPResponse() *HTTPAuthResponse {
	return &HTTPAuthResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}
}

// UserInfoResponse :nodoc:
type UserInfoResponse struct {
	ID        string
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// ToHTTPResponse :nodoc:
func (res *UserInfoResponse) ToHTTPResponse() *HTTPUserInfoResponse {
	return &HTTPUserInfoResponse{
		ID:        res.ID,
		Username:  res.Username,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		DeletedAt: res.DeletedAt,
	}
}

// RefreshTokenPayload :nodoc:
type RefreshTokenPayload struct {
	UserID  string
	TokenID string
}

// LogoutPayload :nodoc:
type LogoutPayload struct {
	UserID  string
	TokenID string
}

// UserRepository :nodoc:
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	UpdateByID(ctx context.Context, id string) (*User, error)
	DeleteByID(ctx context.Context, id string) error

	// DI
	InjectDB(db *gorm.DB) error
}

// UserUsecase :nodoc:
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

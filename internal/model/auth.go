//go:generate mockgen -destination=mock/mock_auth_usecase.go -package=mock github.com/krobus00/auth-service/internal/model AuthUsecase
package model

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v4"
	pb "github.com/krobus00/auth-service/pb/auth"
)

var (
	ErrUsernameOrEmailAlreadyTaken = errors.New("username or email already taken")
	ErrWrongUsernameOrPassword     = errors.New("wrong username/email or password")
	ErrUnauthorizeAccess           = errors.New("unautohirze access")
)

type JWTClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"userID"`
}

type HasAccessPayload struct {
	UserID      string
	Permissions []string
}

func (m *HasAccessPayload) ParseFromProto(req *pb.HasAccessRequest) {
	m.UserID = req.GetUserId()
	m.Permissions = req.GetPermissions()
}

type AuthUsecase interface {
	HasAccess(ctx context.Context, payload *HasAccessPayload) error

	// DI
	InjectUserGroupRepo(repo UserGroupRepository) error
}

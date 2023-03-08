package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/krobus00/auth-service/internal/model"
	pb "github.com/krobus00/auth-service/pb/auth"
)

// Server :nodoc:
type Server struct {
	userUC model.UserUsecase
	authUC model.AuthUsecase
	pb.UnimplementedAuthServiceServer
}

// NewGRPCServer :nodoc:
func NewGRPCServer() *Server {
	return new(Server)
}

func (d *Server) GetUserInfo(ctx context.Context, req *pb.AuthRequest) (*pb.User, error) {
	userID, err := getUserIDFromAccessToken(req.AccessToken)
	if err != nil {
		return nil, err
	}
	res, err := d.userUC.GetUserInfo(ctx, &model.GetUserInfoPayload{
		ID: userID,
	})
	if err != nil {
		return nil, err
	}
	return res.ToGRPCResponse(), nil
}

func (d *Server) HasAccess(ctx context.Context, req *pb.HasAccessRequest) (*wrappers.BoolValue, error) {
	err := d.authUC.HasAccess(ctx, req.GetUserId(), req.GetAccessNames())
	return &wrappers.BoolValue{
		Value: err == nil,
	}, err

}

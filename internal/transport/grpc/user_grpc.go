package grpc

import (
	"context"

	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	pb "github.com/krobus00/auth-service/pb/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (t *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	payload := new(model.UserLoginPayload)
	payload.ParseFromProto(req)

	result, err := t.userUC.Login(ctx, payload)
	switch err {
	case nil:
	case model.ErrWrongUsernameOrPassword:
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
	return result.ToGRPCResponse(), nil
}

func (t *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	payload := new(model.UserRegistrationPayload)
	payload.ParseFromProto(req)

	result, err := t.userUC.Register(ctx, payload)
	switch err {
	case nil:
	case model.ErrUsernameOrEmailAlreadyTaken:
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
	return result.ToGRPCResponse(), nil
}

func (t *Server) Logout(ctx context.Context, req *pb.LogoutRequest) (*emptypb.Empty, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	payload := new(model.UserLogoutPayload)
	payload.ParseFromProto(req)

	err := t.userUC.Logout(ctx, payload)
	switch err {
	case nil:
	case model.ErrUserNotFound:
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

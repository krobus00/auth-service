package grpc

import (
	"context"
	"time"

	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	pb "github.com/krobus00/auth-service/pb/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type Server struct {
	userUC model.UserUsecase
	authUC model.AuthUsecase
	pb.UnimplementedAuthServiceServer
}

func NewGRPCServer() *Server {
	return new(Server)
}

func (t *Server) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.User, error) {
	defer func(tn time.Time) {
		_, _, fn := utils.Trace()
		utils.TimeTrack(tn, fn)
	}(time.Now())

	res, err := t.userUC.GetUserInfo(ctx, &model.GetUserInfoPayload{
		ID: req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	return res.ToGRPCResponse(), nil
}

func (t *Server) HasAccess(ctx context.Context, req *pb.HasAccessRequest) (*wrapperspb.BoolValue, error) {
	defer func(tn time.Time) {
		_, _, fn := utils.Trace()
		utils.TimeTrack(tn, fn)
	}(time.Now())

	err := t.authUC.HasAccess(ctx, req.GetUserId(), req.GetAccessNames())
	return &wrapperspb.BoolValue{
		Value: err == nil,
	}, err
}

func (t *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	defer func(tn time.Time) {
		_, _, fn := utils.Trace()
		utils.TimeTrack(tn, fn)
	}(time.Now())

	result, err := t.userUC.Login(ctx, &model.UserLoginPayload{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	})
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
	defer func(tn time.Time) {
		_, _, fn := utils.Trace()
		utils.TimeTrack(tn, fn)
	}(time.Now())

	result, err := t.userUC.Register(ctx, &model.UserRegistrationPayload{
		FullName: req.GetFullName(),
		Email:    req.GetEmail(),
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	})
	switch err {
	case nil:
	case model.ErrUsernameOrEmailAlreadyTaken:
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
	return result.ToGRPCResponse(), nil
}

func (t *Server) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.AuthResponse, error) {
	defer func(tn time.Time) {
		_, _, fn := utils.Trace()
		utils.TimeTrack(tn, fn)
	}(time.Now())

	token, err := utils.ParseToken(req.GetRefreshToken())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	userID, err := utils.GetUserID(token)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	tokenID, err := utils.GetTokenID(token)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	result, err := t.userUC.RefreshToken(ctx, &model.RefreshTokenPayload{
		UserID:  userID,
		TokenID: tokenID,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return result.ToGRPCResponse(), nil
}

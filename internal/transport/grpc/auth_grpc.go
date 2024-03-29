package grpc

import (
	"context"

	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	pb "github.com/krobus00/auth-service/pb/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (t *Server) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.User, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	res, err := t.userUC.GetUserInfo(ctx, &model.GetUserInfoPayload{
		ID: req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	return res.ToGRPCResponse(), nil
}

func (t *Server) HasAccess(ctx context.Context, req *pb.HasAccessRequest) (*wrapperspb.BoolValue, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	payload := new(model.HasAccessPayload)
	payload.ParseFromProto(req)

	err := t.authUC.HasAccess(ctx, payload)
	return &wrapperspb.BoolValue{
		Value: err == nil,
	}, err
}

func (t *Server) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.AuthResponse, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	payload := new(model.RefreshTokenPayload)
	payload.ParseFromProto(req)

	result, err := t.userUC.RefreshToken(ctx, payload)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return result.ToGRPCResponse(), nil
}

package grpc

import (
	"context"

	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	pb "github.com/krobus00/auth-service/pb/auth"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (t *Server) FindPermissionByID(ctx context.Context, req *pb.FindPermissionByIDRequest) (*pb.Permission, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"sessionUserID": req.GetSessionUserId(),
		"permissionID":  req.GetId(),
	})

	ctx = setUserIDCtx(ctx, req.GetSessionUserId())

	payload := new(model.FindPermissionByIDPayload)
	payload.ParseFromProto(req)

	permission, err := t.permissionUC.FindByID(ctx, payload)
	switch err {
	case nil:
	case model.ErrPermissionNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrUnauthorizeAccess:
		return nil, status.Error(codes.Unauthenticated, err.Error())
	default:
		logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return permission.ToGRPCResponse(), nil
}

func (t *Server) FindPermissionByName(ctx context.Context, req *pb.FindPermissionByNameRequest) (*pb.Permission, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"sessionUserID":  req.GetSessionUserId(),
		"permissionName": req.GetName(),
	})

	ctx = setUserIDCtx(ctx, req.GetSessionUserId())

	payload := new(model.FindPermissionByNamePayload)
	payload.ParseFromProto(req)

	permission, err := t.permissionUC.FindByName(ctx, payload)
	switch err {
	case nil:
	case model.ErrPermissionNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrUnauthorizeAccess:
		return nil, status.Error(codes.Unauthenticated, err.Error())
	default:
		logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return permission.ToGRPCResponse(), nil
}

func (t *Server) CreatePermission(ctx context.Context, req *pb.CreatePermissionRequest) (*pb.Permission, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"sessionUserID":  req.GetSessionUserId(),
		"permissionName": req.GetName(),
	})

	ctx = setUserIDCtx(ctx, req.GetSessionUserId())

	payload := new(model.CreatePermissionPayload)
	payload.ParseFromProto(req)

	permission, err := t.permissionUC.Create(ctx, payload)
	switch err {
	case nil:
	case model.ErrGroupPermissionAlreadyExist:
		return nil, status.Error(codes.AlreadyExists, err.Error())
	case model.ErrUnauthorizeAccess:
		return nil, status.Error(codes.Unauthenticated, err.Error())
	default:
		logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return permission.ToGRPCResponse(), nil
}

func (t *Server) DeletePermission(ctx context.Context, req *pb.DeletePermissionRequest) (*emptypb.Empty, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"sessionUserID": req.GetSessionUserId(),
		"permissionID":  req.GetId(),
	})

	ctx = setUserIDCtx(ctx, req.GetSessionUserId())

	payload := new(model.DeletePermissionByIDPayload)
	payload.ParseFromProto(req)

	err := t.permissionUC.DeleteByID(ctx, payload)
	switch err {
	case nil:
	case model.ErrPermissionNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrUnauthorizeAccess:
		return nil, status.Error(codes.Unauthenticated, err.Error())
	default:
		logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

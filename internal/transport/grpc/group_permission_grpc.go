package grpc

import (
	"context"
	"time"

	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	pb "github.com/krobus00/auth-service/pb/auth"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (t *Server) FindGroupPermission(ctx context.Context, req *pb.FindGroupPermissionRequest) (*pb.GroupPermission, error) {
	defer func(tn time.Time) {
		_, _, fn := utils.Trace()
		utils.TimeTrack(tn, fn)
	}(time.Now())

	logger := logrus.WithFields(logrus.Fields{
		"sessionUserID": req.GetSessionUserId(),
		"groupID":       req.GetGroupId(),
		"permissionID":  req.GetPermissionId(),
	})

	ctx = setUserIDCtx(ctx, req.GetSessionUserId())

	payload := new(model.FindGroupPermissionPayload)
	payload.ParseFromProto(req)

	groupPermission, err := t.groupPermissionUC.FindByGroupIDAndPermissionID(ctx, payload)
	switch err {
	case nil:
	case model.ErrUserNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrGroupNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrUserGroupNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrUnauthorizeAccess:
		return nil, status.Error(codes.Unauthenticated, err.Error())
	default:
		logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return groupPermission.ToGRPCResponse(), nil
}

func (t *Server) CreateGroupPermission(ctx context.Context, req *pb.CreateGroupPermissionRequest) (*pb.GroupPermission, error) {
	defer func(tn time.Time) {
		_, _, fn := utils.Trace()
		utils.TimeTrack(tn, fn)
	}(time.Now())

	logger := logrus.WithFields(logrus.Fields{
		"sessionUserID": req.GetSessionUserId(),
		"groupID":       req.GetGroupId(),
		"permissionID":  req.GetPermissionId(),
	})

	ctx = setUserIDCtx(ctx, req.GetSessionUserId())

	payload := new(model.CreateGroupPermissionPayload)
	payload.ParseFromProto(req)

	groupPermission, err := t.groupPermissionUC.Create(ctx, payload)
	switch err {
	case nil:
	case model.ErrUserNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrGroupNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrUserGroupNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrGroupPermissionAlreadyExist:
		return nil, status.Error(codes.AlreadyExists, err.Error())
	case model.ErrUnauthorizeAccess:
		return nil, status.Error(codes.Unauthenticated, err.Error())
	default:
		logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return groupPermission.ToGRPCResponse(), nil
}

func (t *Server) DeleteGroupPermission(ctx context.Context, req *pb.DeleteGroupPermissionRequest) (*emptypb.Empty, error) {
	defer func(tn time.Time) {
		_, _, fn := utils.Trace()
		utils.TimeTrack(tn, fn)
	}(time.Now())

	logger := logrus.WithFields(logrus.Fields{
		"sessionUserID": req.GetSessionUserId(),
		"groupID":       req.GetGroupId(),
		"permissionID":  req.GetPermissionId(),
	})

	ctx = setUserIDCtx(ctx, req.GetSessionUserId())

	payload := new(model.DeleteGroupPermissionPayload)
	payload.ParseFromProto(req)

	err := t.groupPermissionUC.DeleteByGroupIDAndPermissionID(ctx, payload)
	switch err {
	case nil:
	case model.ErrUserNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrGroupNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrUserGroupNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrUnauthorizeAccess:
		return nil, status.Error(codes.Unauthenticated, err.Error())
	default:
		logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

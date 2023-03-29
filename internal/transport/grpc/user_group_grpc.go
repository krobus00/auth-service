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

func (t *Server) FindAllUserGroups(ctx context.Context, req *pb.FindAllUserGroupsRequest) (*pb.FindAllUserGroupsResponse, error) {
	defer func(tn time.Time) {
		_, _, fn := utils.Trace()
		utils.TimeTrack(tn, fn)
	}(time.Now())

	logger := logrus.WithFields(logrus.Fields{
		"sessionUserID": req.GetSessionUserId(),
		"userID":        req.GetUserId(),
	})

	ctx = setUserIDCtx(ctx, req.GetSessionUserId())

	payload := new(model.FindUserGroupsByUserIDPayload)
	payload.ParseFromProto(req)

	userGroups, err := t.userGroupUC.FindByUserID(ctx, payload)
	switch err {
	case nil:
	case model.ErrUserNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrUnauthorizeAccess:
		return nil, status.Error(codes.Unauthenticated, err.Error())
	default:
		logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return userGroups.ToGRPCResponse(), nil
}

func (t *Server) FindUserGroup(ctx context.Context, req *pb.FindUserGroupRequest) (*pb.UserGroup, error) {
	defer func(tn time.Time) {
		_, _, fn := utils.Trace()
		utils.TimeTrack(tn, fn)
	}(time.Now())

	logger := logrus.WithFields(logrus.Fields{
		"sessionUserID": req.GetSessionUserId(),
		"userID":        req.GetUserId(),
		"groupID":       req.GetGroupId(),
	})

	ctx = setUserIDCtx(ctx, req.GetSessionUserId())

	payload := new(model.FindUserGroupPayload)
	payload.ParseFromProto(req)

	userGroup, err := t.userGroupUC.FindByUserIDAndGroupID(ctx, payload)
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

	return userGroup.ToGRPCResponse(), nil
}

func (t *Server) CreateUserGroup(ctx context.Context, req *pb.CreateUserGroupRequest) (*pb.UserGroup, error) {
	defer func(tn time.Time) {
		_, _, fn := utils.Trace()
		utils.TimeTrack(tn, fn)
	}(time.Now())

	logger := logrus.WithFields(logrus.Fields{
		"sessionUserID": req.GetSessionUserId(),
		"userID":        req.GetUserId(),
		"groupID":       req.GetGroupId(),
	})

	ctx = setUserIDCtx(ctx, req.GetSessionUserId())

	payload := new(model.CreateUserGroupPayload)
	payload.ParseFromProto(req)

	userGroup, err := t.userGroupUC.Create(ctx, payload)
	switch err {
	case nil:
	case model.ErrUserNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrGroupNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrUserGroupNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrUserGroupAlreadyExist:
		return nil, status.Error(codes.AlreadyExists, err.Error())
	case model.ErrUnauthorizeAccess:
		return nil, status.Error(codes.Unauthenticated, err.Error())
	default:
		logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return userGroup.ToGRPCResponse(), nil
}

func (t *Server) DeleteUserGroup(ctx context.Context, req *pb.DeleteUserGroupRequest) (*emptypb.Empty, error) {
	defer func(tn time.Time) {
		_, _, fn := utils.Trace()
		utils.TimeTrack(tn, fn)
	}(time.Now())

	logger := logrus.WithFields(logrus.Fields{
		"sessionUserID": req.GetSessionUserId(),
		"userID":        req.GetUserId(),
		"groupID":       req.GetGroupId(),
	})

	ctx = setUserIDCtx(ctx, req.GetSessionUserId())

	payload := new(model.DeleteUserGroupPayload)
	payload.ParseFromProto(req)

	err := t.userGroupUC.DeleteByUserIDAndGroupID(ctx, payload)
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

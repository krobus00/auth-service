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

func (t *Server) FindGroupByID(ctx context.Context, req *pb.FindGroupByIDRequest) (*pb.Group, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"sessionUserID": req.GetSessionUserId(),
		"groupID":       req.GetId(),
	})

	ctx = setUserIDCtx(ctx, req.GetSessionUserId())

	payload := new(model.FindGroupByIDPayload)
	payload.ParseFromProto(req)

	group, err := t.groupUC.FindByID(ctx, payload)
	switch err {
	case nil:
	case model.ErrGroupNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrUnauthorizeAccess:
		return nil, status.Error(codes.Unauthenticated, err.Error())
	default:
		logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return group.ToGRPCResponse(), nil
}

func (t *Server) FindGroupByName(ctx context.Context, req *pb.FindGroupByNameRequest) (*pb.Group, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"sessionUserID": req.GetSessionUserId(),
		"groupName":     req.GetName(),
	})

	ctx = setUserIDCtx(ctx, req.GetSessionUserId())

	payload := new(model.FindGroupByNamePayload)
	payload.ParseFromProto(req)

	group, err := t.groupUC.FindByName(ctx, payload)
	switch err {
	case nil:
	case model.ErrGroupNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrUnauthorizeAccess:
		return nil, status.Error(codes.Unauthenticated, err.Error())
	default:
		logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return group.ToGRPCResponse(), nil
}

func (t *Server) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.Group, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"sessionUserID": req.GetSessionUserId(),
		"groupName":     req.GetName(),
	})

	ctx = setUserIDCtx(ctx, req.GetSessionUserId())

	payload := new(model.CreateGroupPayload)
	payload.ParseFromProto(req)

	group, err := t.groupUC.Create(ctx, payload)
	switch err {
	case nil:
	case model.ErrGroupAlreadyExist:
		return nil, status.Error(codes.AlreadyExists, err.Error())
	case model.ErrUnauthorizeAccess:
		return nil, status.Error(codes.Unauthenticated, err.Error())
	default:
		logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return group.ToGRPCResponse(), nil
}

func (t *Server) DeleteGroupByID(ctx context.Context, req *pb.DeleteGroupRequest) (*emptypb.Empty, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"sessionUserID": req.GetSessionUserId(),
		"groupID":       req.GetId(),
	})

	ctx = setUserIDCtx(ctx, req.GetSessionUserId())

	payload := new(model.DeleteGroupByIDPayload)
	payload.ParseFromProto(req)

	err := t.groupUC.DeleteByID(ctx, payload)
	switch err {
	case nil:
	case model.ErrGroupNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case model.ErrUnauthorizeAccess:
		return nil, status.Error(codes.Unauthenticated, err.Error())
	default:
		logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

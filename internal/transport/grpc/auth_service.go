package grpc

import (
	"github.com/krobus00/auth-service/internal/model"
	pb "github.com/krobus00/auth-service/pb/auth"
)

type Server struct {
	userUC            model.UserUsecase
	authUC            model.AuthUsecase
	permissionUC      model.PermissionUsecase
	groupUC           model.GroupUsecase
	userGroupUC       model.UserGroupUsecase
	groupPermissionUC model.GroupPermissionUsecase
	pb.UnimplementedAuthServiceServer
}

func NewGRPCServer() *Server {
	return new(Server)
}

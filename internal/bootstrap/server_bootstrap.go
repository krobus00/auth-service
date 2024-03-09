package bootstrap

import (
	"context"
	"fmt"
	"net"

	"net/http"

	"github.com/krobus00/auth-service/internal/config"
	"github.com/krobus00/auth-service/internal/infrastructure"
	"github.com/krobus00/auth-service/internal/repository"
	grpcTransport "github.com/krobus00/auth-service/internal/transport/grpc"
	"github.com/krobus00/auth-service/internal/usecase"
	pb "github.com/krobus00/auth-service/pb/auth"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartServer() {
	infrastructure.InitializeDBConn()

	// init infra
	db, err := infrastructure.DB.DB()
	continueOrFatal(err)

	redisClient, err := infrastructure.NewRedisClient()
	continueOrFatal(err)

	tp, err := infrastructure.JaegerTraceProvider()
	continueOrFatal(err)

	// init repository
	userRepo := repository.NewUserRepository()
	err = userRepo.InjectDB(infrastructure.DB)
	continueOrFatal(err)
	err = userRepo.InjectRedisClient(redisClient)
	continueOrFatal(err)

	tokenRepo := repository.NewTokenRepository()
	err = tokenRepo.InjectRedisClient(redisClient)
	continueOrFatal(err)

	groupRepo := repository.NewGroupRepository()
	err = groupRepo.InjectDB(infrastructure.DB)
	continueOrFatal(err)
	err = groupRepo.InjectRedisClient(redisClient)
	continueOrFatal(err)

	userGroupRepo := repository.NewUserGroupRepository()
	err = userGroupRepo.InjectDB(infrastructure.DB)
	continueOrFatal(err)
	err = userGroupRepo.InjectRedisClient(redisClient)
	continueOrFatal(err)

	permissionRepo := repository.NewPermissionRepository()
	err = permissionRepo.InjectDB(infrastructure.DB)
	continueOrFatal(err)
	err = permissionRepo.InjectRedisClient(redisClient)
	continueOrFatal(err)

	groupPermissionRepo := repository.NewGroupPermissionRepository()
	err = groupPermissionRepo.InjectDB(infrastructure.DB)
	continueOrFatal(err)
	err = groupPermissionRepo.InjectRedisClient(redisClient)
	continueOrFatal(err)

	// init usecase
	userUsecase := usecase.NewUserUsecase()
	err = userUsecase.InjectDB(infrastructure.DB)
	continueOrFatal(err)
	err = userUsecase.InjectUserRepo(userRepo)
	continueOrFatal(err)
	err = userUsecase.InjectTokenRepo(tokenRepo)
	continueOrFatal(err)
	err = userUsecase.InjectGroupRepo(groupRepo)
	continueOrFatal(err)
	err = userUsecase.InjectUserGroupRepo(userGroupRepo)
	continueOrFatal(err)

	authUsecase := usecase.NewAuthUsecase()
	err = authUsecase.InjectUserGroupRepo(userGroupRepo)
	continueOrFatal(err)

	permissionUsecase := usecase.NewPermissionUsecase()
	err = permissionUsecase.InjectPermissionRepo(permissionRepo)
	continueOrFatal(err)
	err = permissionUsecase.InjectAuthUsecase(authUsecase)
	continueOrFatal(err)

	groupUsecase := usecase.NewGroupUsecase()
	err = groupUsecase.InjectGroupRepo(groupRepo)
	continueOrFatal(err)
	err = groupUsecase.InjectAuthUsecase(authUsecase)
	continueOrFatal(err)

	userGroupUsecase := usecase.NewUserGroupUsecase()
	err = userGroupUsecase.InjectAuthUsecase(authUsecase)
	continueOrFatal(err)
	err = userGroupUsecase.InjectGroupRepo(groupRepo)
	continueOrFatal(err)
	err = userGroupUsecase.InjectUserRepo(userRepo)
	continueOrFatal(err)
	err = userGroupUsecase.InjectUserGroupRepo(userGroupRepo)
	continueOrFatal(err)

	groupPermissionUsecase := usecase.NewGroupPermissionUsecase()
	err = groupPermissionUsecase.InjectAuthUsecase(authUsecase)
	continueOrFatal(err)
	err = groupPermissionUsecase.InjectGroupPermissionRepo(groupPermissionRepo)
	continueOrFatal(err)
	err = groupPermissionUsecase.InjectGroupRepo(groupRepo)
	continueOrFatal(err)
	err = groupPermissionUsecase.InjectPermisisonRepo(permissionRepo)
	continueOrFatal(err)

	grpcDelivery := grpcTransport.NewGRPCServer()
	err = grpcDelivery.InjectUserUsecase(userUsecase)
	continueOrFatal(err)
	err = grpcDelivery.InjectAuthUsecase(authUsecase)
	continueOrFatal(err)
	err = grpcDelivery.InjectPermissionUsecase(permissionUsecase)
	continueOrFatal(err)
	err = grpcDelivery.InjectGroupUsecase(groupUsecase)
	continueOrFatal(err)
	err = grpcDelivery.InjectUserGroupUsecase(userGroupUsecase)
	continueOrFatal(err)
	err = grpcDelivery.InjectGroupPermissionUsecase(groupPermissionUsecase)
	continueOrFatal(err)

	authGrpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(authGrpcServer, grpcDelivery)
	if config.Env() == "development" {
		reflection.Register(authGrpcServer)
	}
	lis, err := net.Listen("tcp", ":"+config.PortGRPC())
	continueOrFatal(err)

	go func() {
		err = authGrpcServer.Serve(lis)
		if err != nil {
			continueOrFatal(err)
		}
	}()
	logrus.Info(fmt.Sprintf("grpc server started on :%s", config.PortGRPC()))

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		err = http.ListenAndServe(fmt.Sprintf(":%s", config.PortMetrics()), nil)
		if err != nil {
			continueOrFatal(err)
		}
	}()
	logrus.Info(fmt.Sprintf("metrics server started on :%s", config.PortMetrics()))

	wait := gracefulShutdown(context.Background(), config.GracefulShutdownTimeOut(), map[string]operation{
		"redis connection": func(ctx context.Context) error {
			return redisClient.Close()
		},
		"database connection": func(ctx context.Context) error {
			infrastructure.StopTickerCh <- true
			return db.Close()
		},
		"grpc": func(ctx context.Context) error {
			return lis.Close()
		},
		"trace provider": func(ctx context.Context) error {
			if config.DisableTracing() {
				return nil
			}
			return tp.Shutdown(ctx)
		},
	})
	<-wait
}

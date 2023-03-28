package bootstrap

import (
	"context"
	"fmt"
	"net"

	"github.com/krobus00/auth-service/internal/config"
	"github.com/krobus00/auth-service/internal/infrastructure"
	"github.com/krobus00/auth-service/internal/repository"
	grpcServer "github.com/krobus00/auth-service/internal/transport/grpc"
	"github.com/krobus00/auth-service/internal/usecase"
	pb "github.com/krobus00/auth-service/pb/auth"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartServer() {
	infrastructure.InitializeDBConn()

	// init infra
	db, err := infrastructure.DB.DB()
	continueOrFatal(err)

	redisClient, err := infrastructure.NewRedisClient()
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

	userGroupRepo := repository.NewUserGroupRepository()
	err = userGroupRepo.InjectDB(infrastructure.DB)
	continueOrFatal(err)
	err = userGroupRepo.InjectRedisClient(redisClient)
	continueOrFatal(err)

	// init usecase
	userUsecase := usecase.NewUserUsecase()
	err = userUsecase.InjectUserRepo(userRepo)
	continueOrFatal(err)
	err = userUsecase.InjectDB(infrastructure.DB)
	continueOrFatal(err)
	err = userUsecase.InjectTokenRepo(tokenRepo)
	continueOrFatal(err)

	authUsecase := usecase.NewAuthUsecase()
	err = authUsecase.InjectUserGroupRepo(userGroupRepo)
	continueOrFatal(err)

	grpcDelivery := grpcServer.NewGRPCServer()
	err = grpcDelivery.InjectUserUsecase(userUsecase)
	continueOrFatal(err)
	err = grpcDelivery.InjectAuthUsecase(authUsecase)
	continueOrFatal(err)

	authGrpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(authGrpcServer, grpcDelivery)
	if config.Env() == "development" {
		reflection.Register(authGrpcServer)
	}
	lis, _ := net.Listen("tcp", ":"+config.GRPCport())

	go func() {
		_ = authGrpcServer.Serve(lis)
	}()
	log.Info(fmt.Sprintf("grpc server started on :%s", config.GRPCport()))

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
	})
	<-wait
}

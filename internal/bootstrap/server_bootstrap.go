package bootstrap

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/krobus00/auth-service/internal/config"
	grpcServer "github.com/krobus00/auth-service/internal/delivery/grpc"
	"github.com/krobus00/auth-service/internal/delivery/http"
	"github.com/krobus00/auth-service/internal/infrastructure"
	"github.com/krobus00/auth-service/internal/repository"
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

	echo := infrastructure.NewEcho()

	// init repository
	userRepo := repository.NewUserRepository()
	err = userRepo.InjectDB(infrastructure.DB)
	continueOrFatal(err)

	tokenRepo := repository.NewTokenRepository()
	err = tokenRepo.InjectRedisClient(redisClient)
	continueOrFatal(err)

	userAccessControlRepo := repository.NewUserAccessControlRepository()
	err = userAccessControlRepo.InjectDB(infrastructure.DB)
	continueOrFatal(err)
	err = userAccessControlRepo.InjectRedisClient(redisClient)
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
	err = authUsecase.InjectUserAccessControlRepo(userAccessControlRepo)
	continueOrFatal(err)

	userCtrl := http.NewUserController()
	err = userCtrl.InjectUserUsecase(userUsecase)
	continueOrFatal(err)

	httpMiddleware := http.NewHTTPMiddleware()
	err = httpMiddleware.InjectTokenRepo(tokenRepo)
	continueOrFatal(err)

	httpDelivery := http.NewHTTPDelivery()
	err = httpDelivery.InjectEcho(echo)
	continueOrFatal(err)
	err = httpDelivery.InjectHTTPMiddleware(httpMiddleware)
	continueOrFatal(err)
	err = httpDelivery.InjectUserController(userCtrl)
	continueOrFatal(err)
	httpDelivery.InitRoutes()

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
		_ = echo.Start(":" + config.HTTPPort())
	}()
	log.Info(fmt.Sprintf("http server started on :%s", config.HTTPPort()))

	go func() {
		_ = authGrpcServer.Serve(lis)
	}()
	log.Info(fmt.Sprintf("grpc server started on :%s", config.GRPCport()))

	wait := gracefulShutdown(context.Background(), 30*time.Second, map[string]operation{
		"redis connection": func(ctx context.Context) error {
			return redisClient.Close()
		},
		"database connection": func(ctx context.Context) error {
			infrastructure.StopTickerCh <- true
			return db.Close()
		},
		"http": func(ctx context.Context) error {
			return echo.Shutdown(ctx)
		},
		"grpc": func(ctx context.Context) error {
			return lis.Close()
		},
	})

	<-wait
}

package bootstrap

import (
	"context"
	"time"

	"github.com/krobus00/auth-service/internal/config"
	"github.com/krobus00/auth-service/internal/delivery/http"
	"github.com/krobus00/auth-service/internal/infrastructure"
	"github.com/krobus00/auth-service/internal/repository"
	"github.com/krobus00/auth-service/internal/usecase"
)

func StartServer() {
	infrastructure.InitializeDBConn()

	// init infra
	db, err := infrastructure.DB.DB()
	continueOrFatal(err)

	redisClient, err := infrastructure.NewRedisClient()
	continueOrFatal(err)

	echo := infrastructure.NewEcho()

	userRepo := repository.NewUserRepository()
	err = userRepo.InjectDB(infrastructure.DB)
	continueOrFatal(err)

	tokenRepo := repository.NewTokenRepository()
	err = tokenRepo.InjectRedisClient(redisClient)
	continueOrFatal(err)

	// init usecase
	userUsecase := usecase.NewUserUsecase()
	err = userUsecase.InjectUserRepo(userRepo)
	continueOrFatal(err)
	err = userUsecase.InjectDB(infrastructure.DB)
	continueOrFatal(err)
	err = userUsecase.InjectTokenRepo(tokenRepo)
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

	go func() {
		_ = echo.Start(":" + config.HTTPPort())

	}()

	wait := gracefulShutdown(context.Background(), 30*time.Second, map[string]operation{
		"redis connection": func(ctx context.Context) error {
			return redisClient.Close()
		},
		"database connection": func(ctx context.Context) error {
			infrastructure.StopTickerCh <- true
			return db.Close()
		},
		"echo": func(ctx context.Context) error {
			return echo.Shutdown(ctx)
		},
	})

	<-wait
}

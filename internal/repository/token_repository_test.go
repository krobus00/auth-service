package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/krobus00/auth-service/internal/infrastructure"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	"github.com/spf13/viper"
)

func newTokenRepoMock(t *testing.T) (model.TokenRepository, *miniredis.Miniredis) {
	miniRedis := miniredis.RunT(t)
	viper.Set("redis.cache_host", fmt.Sprintf("redis://%s", miniRedis.Addr()))
	redisClient, err := infrastructure.NewRedisClient()
	utils.ContinueOrFatal(err)
	tokenRepo := NewTokenRepository()
	err = tokenRepo.InjectRedisClient(redisClient)
	utils.ContinueOrFatal(err)

	return tokenRepo, miniRedis
}

func Test_tokenRepository_Create(t *testing.T) {
	type args struct {
		userID    string
		tokenID   string
		tokenType model.TokenType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success create new access token",
			args: args{
				userID:    utils.GenerateUUID(),
				tokenID:   utils.GenerateUUID(),
				tokenType: model.AccessToken,
			},
			wantErr: false,
		},
		{
			name: "success create new refresh token",
			args: args{
				userID:    utils.GenerateUUID(),
				tokenID:   utils.GenerateUUID(),
				tokenType: model.RefreshToken,
			},
			wantErr: false,
		},
		{
			name: "error invalid token type",
			args: args{
				userID:    utils.GenerateUUID(),
				tokenID:   utils.GenerateUUID(),
				tokenType: model.TokenType(3),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := newTokenRepoMock(t)

			_, err := r.Create(context.TODO(), tt.args.userID, tt.args.tokenID, tt.args.tokenType)
			if (err != nil) != tt.wantErr {
				t.Errorf("tokenRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_tokenRepository_IsValidToken(t *testing.T) {
	var (
		tokenID = utils.GenerateUUID()
		userID  = utils.GenerateUUID()
	)
	type args struct {
		userID    string
		tokenID   string
		tokenType model.TokenType
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "success get valid access token",
			args: args{
				userID:    userID,
				tokenID:   tokenID,
				tokenType: model.AccessToken,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "success get valid refresh token",
			args: args{
				userID:    userID,
				tokenID:   tokenID,
				tokenType: model.RefreshToken,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "success access token not found",
			args: args{
				userID:    utils.GenerateUUID(),
				tokenID:   utils.GenerateUUID(),
				tokenType: model.AccessToken,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "success refresh token not found",
			args: args{
				userID:    utils.GenerateUUID(),
				tokenID:   utils.GenerateUUID(),
				tokenType: model.RefreshToken,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "error invalid token type",
			args: args{
				userID:    utils.GenerateUUID(),
				tokenID:   utils.GenerateUUID(),
				tokenType: model.TokenType(3),
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, redisMock := newTokenRepoMock(t)

			if tt.args.tokenType == model.AccessToken {
				if tt.want {
					_ = redisMock.Set(model.AccessTokenCacheKey(tt.args.userID, tt.args.tokenID), "access-token")
				} else {
					_ = redisMock.Del(model.AccessTokenCacheKey(tt.args.userID, tt.args.tokenID))
				}
			}
			if tt.args.tokenType == model.RefreshToken {
				if tt.want {
					_ = redisMock.Set(model.RefreshTokenCacheKey(tt.args.userID, tt.args.tokenID), "refresh-token")
				} else {
					_ = redisMock.Del(model.RefreshTokenCacheKey(tt.args.userID, tt.args.tokenID))
				}
			}

			got, err := r.IsValidToken(context.TODO(), tt.args.userID, tt.args.tokenID, tt.args.tokenType)
			if (err != nil) != tt.wantErr {
				t.Errorf("tokenRepository.IsValidToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("tokenRepository.IsValidToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokenRepository_Revoke(t *testing.T) {
	type args struct {
		userID    string
		tokenID   string
		tokenType model.TokenType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success revoke access token",
			args: args{
				userID:    utils.GenerateUUID(),
				tokenID:   utils.GenerateUUID(),
				tokenType: model.AccessToken,
			},
			wantErr: false,
		},
		{
			name: "success revoke refresh token",
			args: args{
				userID:    utils.GenerateUUID(),
				tokenID:   utils.GenerateUUID(),
				tokenType: model.RefreshToken,
			},
			wantErr: false,
		},
		{
			name: "error invalid token type",
			args: args{
				userID:    utils.GenerateUUID(),
				tokenID:   utils.GenerateUUID(),
				tokenType: model.TokenType(3),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := newTokenRepoMock(t)

			if err := r.Revoke(context.TODO(), tt.args.userID, tt.args.tokenID, tt.args.tokenType); (err != nil) != tt.wantErr {
				t.Errorf("tokenRepository.Revoke() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/krobus00/auth-service/internal/constant"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/model/mock"
	"github.com/krobus00/auth-service/internal/utils"
)

func Test_userUsecase_Register(t *testing.T) {
	var (
		userID    = utils.GenerateUUID()
		groupID   = utils.GenerateUUID()
		userEmail = "user@gmail.com"
		username  = "user1"
	)
	type mockFindByUsername struct {
		res *model.User
		err error
	}
	type mockFindByEmail struct {
		res *model.User
		err error
	}
	type mockCreateUser struct {
		res *model.User
		err error
	}
	type mockFindGroupByName struct {
		res *model.Group
		err error
	}
	type mockCreateUserGroup struct {
		err error
	}
	type mockCreateToken struct {
		res string
		err error
	}
	type args struct {
		payload *model.UserRegistrationPayload
	}
	tests := []struct {
		name                   string
		args                   args
		mockFindByUsername     *mockFindByUsername
		mockFindByEmail        *mockFindByEmail
		mockCreateUser         *mockCreateUser
		mockFindGroupByName    *mockFindGroupByName
		mockCreateUserGroup    *mockCreateUserGroup
		mockCreateAccessToken  *mockCreateToken
		mockCreateRefreshToken *mockCreateToken
		wantCommit             bool
		want                   *model.AuthResponse
		wantErr                bool
	}{
		{
			name: "success",
			args: args{
				payload: &model.UserRegistrationPayload{
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
			},
			mockFindByUsername: &mockFindByUsername{
				res: nil,
				err: nil,
			},
			mockFindByEmail: &mockFindByEmail{
				res: nil,
				err: nil,
			},
			mockCreateUser: &mockCreateUser{
				res: &model.User{
					ID:       userID,
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
				err: nil,
			},
			mockFindGroupByName: &mockFindGroupByName{
				res: &model.Group{
					ID:   groupID,
					Name: constant.GroupDefault,
				},
				err: nil,
			},
			mockCreateUserGroup: &mockCreateUserGroup{
				err: nil,
			},
			mockCreateAccessToken: &mockCreateToken{
				res: "access-token",
				err: nil,
			},
			mockCreateRefreshToken: &mockCreateToken{
				res: "refresh-token",
				err: nil,
			},
			wantCommit: true,
			want: &model.AuthResponse{
				AccessToken:  "access-token",
				RefreshToken: "refresh-token",
			},
			wantErr: false,
		},
		{
			name: "error username already taken",
			args: args{
				payload: &model.UserRegistrationPayload{
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
			},
			mockFindByUsername: &mockFindByUsername{
				res: &model.User{
					ID:       userID,
					FullName: "old user",
					Username: username,
				},
				err: nil,
			},
			wantCommit: false,
			wantErr:    true,
		},
		{
			name: "error find user by username",
			args: args{
				payload: &model.UserRegistrationPayload{
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
			},
			mockFindByUsername: &mockFindByUsername{
				res: nil,
				err: errors.New("db error"),
			},
			wantCommit: false,
			wantErr:    true,
		},
		{
			name: "error email already taken",
			args: args{
				payload: &model.UserRegistrationPayload{
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
			},
			mockFindByUsername: &mockFindByUsername{
				res: nil,
				err: nil,
			},
			mockFindByEmail: &mockFindByEmail{
				res: &model.User{
					ID:       userID,
					FullName: "old user",
					Username: "olduser17",
					Email:    userEmail,
				},
				err: nil,
			},
			wantCommit: false,
			wantErr:    true,
		},
		{
			name: "error find user by email",
			args: args{
				payload: &model.UserRegistrationPayload{
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
			},
			mockFindByUsername: &mockFindByUsername{
				res: nil,
				err: nil,
			},
			mockFindByEmail: &mockFindByEmail{
				res: nil,
				err: errors.New("db error"),
			},
			wantCommit: false,
			wantErr:    true,
		},
		{
			name: "error create user",
			args: args{
				payload: &model.UserRegistrationPayload{
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
			},
			mockFindByUsername: &mockFindByUsername{
				res: nil,
				err: nil,
			},
			mockFindByEmail: &mockFindByEmail{
				res: nil,
				err: nil,
			},
			mockCreateUser: &mockCreateUser{
				res: &model.User{},
				err: errors.New("db error"),
			},
			wantCommit: false,
			wantErr:    true,
		},
		{
			name: "error find group",
			args: args{
				payload: &model.UserRegistrationPayload{
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
			},
			mockFindByUsername: &mockFindByUsername{
				res: nil,
				err: nil,
			},
			mockFindByEmail: &mockFindByEmail{
				res: nil,
				err: nil,
			},
			mockCreateUser: &mockCreateUser{
				res: &model.User{
					ID:       userID,
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
				err: nil,
			},
			mockFindGroupByName: &mockFindGroupByName{
				res: nil,
				err: errors.New("db error"),
			},
			wantCommit: false,
			wantErr:    true,
		},
		{
			name: "error group not found",
			args: args{
				payload: &model.UserRegistrationPayload{
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
			},
			mockFindByUsername: &mockFindByUsername{
				res: nil,
				err: nil,
			},
			mockFindByEmail: &mockFindByEmail{
				res: nil,
				err: nil,
			},
			mockCreateUser: &mockCreateUser{
				res: &model.User{
					ID:       userID,
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
				err: nil,
			},
			mockFindGroupByName: &mockFindGroupByName{
				res: nil,
				err: nil,
			},
			wantCommit: false,
			wantErr:    true,
		},
		{
			name: "error add user group",
			args: args{
				payload: &model.UserRegistrationPayload{
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
			},
			mockFindByUsername: &mockFindByUsername{
				res: nil,
				err: nil,
			},
			mockFindByEmail: &mockFindByEmail{
				res: nil,
				err: nil,
			},
			mockCreateUser: &mockCreateUser{
				res: &model.User{
					ID:       userID,
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
				err: nil,
			},
			mockFindGroupByName: &mockFindGroupByName{
				res: &model.Group{
					ID:   groupID,
					Name: constant.GroupDefault,
				},
				err: nil,
			},
			mockCreateUserGroup: &mockCreateUserGroup{
				err: errors.New("db error"),
			},
			wantCommit: false,
			wantErr:    true,
		},
		{
			name: "error generate access token",
			args: args{
				payload: &model.UserRegistrationPayload{
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
			},
			mockFindByUsername: &mockFindByUsername{
				res: nil,
				err: nil,
			},
			mockFindByEmail: &mockFindByEmail{
				res: nil,
				err: nil,
			},
			mockCreateUser: &mockCreateUser{
				res: &model.User{
					ID:       userID,
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
				err: nil,
			},
			mockFindGroupByName: &mockFindGroupByName{
				res: &model.Group{
					ID:   groupID,
					Name: constant.GroupDefault,
				},
				err: nil,
			},
			mockCreateUserGroup: &mockCreateUserGroup{
				err: nil,
			},
			mockCreateAccessToken: &mockCreateToken{
				res: "",
				err: errors.New("redis error"),
			},
			wantCommit: false,
			wantErr:    true,
		},
		{
			name: "error generate refresh token",
			args: args{
				payload: &model.UserRegistrationPayload{
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
			},
			mockFindByUsername: &mockFindByUsername{
				res: nil,
				err: nil,
			},
			mockFindByEmail: &mockFindByEmail{
				res: nil,
				err: nil,
			},
			mockCreateUser: &mockCreateUser{
				res: &model.User{
					ID:       userID,
					FullName: "new user",
					Username: username,
					Email:    userEmail,
					Password: "strongpassword",
				},
				err: nil,
			},
			mockFindGroupByName: &mockFindGroupByName{
				res: &model.Group{
					ID:   groupID,
					Name: constant.GroupDefault,
				},
				err: nil,
			},
			mockCreateUserGroup: &mockCreateUserGroup{
				err: nil,
			},
			mockCreateAccessToken: &mockCreateToken{
				res: "access-token",
				err: nil,
			},
			mockCreateRefreshToken: &mockCreateToken{
				res: "",
				err: errors.New("redis error"),
			},
			wantCommit: false,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.TODO()

			dbConn, dbMock := utils.NewDBMock()
			userRepo := mock.NewMockUserRepository(ctrl)
			tokenRepo := mock.NewMockTokenRepository(ctrl)
			groupRepo := mock.NewMockGroupRepository(ctrl)
			userGroupRepo := mock.NewMockUserGroupRepository(ctrl)

			dbMock.ExpectBegin()
			if tt.wantCommit {
				dbMock.ExpectCommit()
			} else {
				dbMock.ExpectRollback()
			}

			if tt.mockFindByUsername != nil {
				userRepo.EXPECT().FindByUsername(gomock.Any(), tt.args.payload.Username).
					Times(1).
					Return(tt.mockFindByUsername.res, tt.mockFindByUsername.err)
			}

			if tt.mockFindByEmail != nil {
				userRepo.EXPECT().FindByEmail(gomock.Any(), tt.args.payload.Email).
					Times(1).
					Return(tt.mockFindByEmail.res, tt.mockFindByEmail.err)
			}

			if tt.mockCreateUser != nil {
				userRepo.EXPECT().Create(gomock.Any(), gomock.Any()).
					Times(1).
					DoAndReturn(func(ctx context.Context, user *model.User) error {
						user.ID = userID
						user.FullName = tt.args.payload.FullName
						user.FullName = tt.args.payload.FullName
						user.Email = tt.args.payload.Email
						user.Password = tt.mockCreateUser.res.Password
						return tt.mockCreateUser.err
					})
			}

			if tt.mockFindGroupByName != nil {
				groupRepo.EXPECT().
					FindByName(gomock.Any(), constant.GroupDefault).
					Times(1).
					Return(tt.mockFindGroupByName.res, tt.mockFindGroupByName.err)
			}

			if tt.mockCreateUserGroup != nil {
				userGroupRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(tt.mockCreateUserGroup.err)
			}

			if tt.mockCreateAccessToken != nil {
				tokenRepo.EXPECT().Create(gomock.Any(), userID, gomock.Any(), model.AccessToken).Times(1).DoAndReturn(func(ctx context.Context, userID string, tokenID string, tokenType model.TokenType) (string, error) {
					return tt.mockCreateAccessToken.res, tt.mockCreateAccessToken.err
				})
			}

			if tt.mockCreateRefreshToken != nil {
				tokenRepo.EXPECT().Create(gomock.Any(), userID, gomock.Any(), model.RefreshToken).Times(1).DoAndReturn(func(ctx context.Context, userID string, tokenID string, tokenType model.TokenType) (string, error) {
					return tt.mockCreateRefreshToken.res, tt.mockCreateRefreshToken.err
				})
			}

			uc := NewUserUsecase()
			err := uc.InjectDB(dbConn)
			utils.ContinueOrFatal(err)
			err = uc.InjectGroupRepo(groupRepo)
			utils.ContinueOrFatal(err)
			err = uc.InjectUserRepo(userRepo)
			utils.ContinueOrFatal(err)
			err = uc.InjectTokenRepo(tokenRepo)
			utils.ContinueOrFatal(err)
			err = uc.InjectUserGroupRepo(userGroupRepo)
			utils.ContinueOrFatal(err)

			got, err := uc.Register(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_Login(t *testing.T) {
	var (
		userID    = utils.GenerateUUID()
		userEmail = "user@gmail.com"
		username  = "user1"
	)
	userPassword, _ := utils.HashPassword("strongpassword")
	type mockFindByUsername struct {
		res *model.User
		err error
	}
	type mockFindByEmail struct {
		res *model.User
		err error
	}
	type mockCreateToken struct {
		res string
		err error
	}
	type args struct {
		payload *model.UserLoginPayload
	}
	tests := []struct {
		name                   string
		args                   args
		mockFindByUsername     *mockFindByUsername
		mockFindByEmail        *mockFindByEmail
		mockCreateAccessToken  *mockCreateToken
		mockCreateRefreshToken *mockCreateToken
		want                   *model.AuthResponse
		wantErr                bool
	}{
		{
			name: "success get user by username",
			args: args{
				payload: &model.UserLoginPayload{
					Username: username,
					Password: "strongpassword",
				},
			},
			mockFindByUsername: &mockFindByUsername{
				res: &model.User{
					ID:       userID,
					FullName: "user",
					Username: username,
					Email:    userEmail,
					Password: userPassword,
				},
				err: nil,
			},
			mockCreateAccessToken: &mockCreateToken{
				res: "access-token",
				err: nil,
			},
			mockCreateRefreshToken: &mockCreateToken{
				res: "refresh-token",
				err: nil,
			},
			want: &model.AuthResponse{
				AccessToken:  "access-token",
				RefreshToken: "refresh-token",
			},
			wantErr: false,
		},
		{
			name: "success get user by email",
			args: args{
				payload: &model.UserLoginPayload{
					Username: userEmail,
					Password: "strongpassword",
				},
			},
			mockFindByUsername: &mockFindByUsername{
				res: nil,
				err: nil,
			},
			mockFindByEmail: &mockFindByEmail{
				res: &model.User{
					ID:       userID,
					FullName: "user",
					Username: username,
					Email:    userEmail,
					Password: userPassword,
				},
				err: nil,
			},
			mockCreateAccessToken: &mockCreateToken{
				res: "access-token",
				err: nil,
			},
			mockCreateRefreshToken: &mockCreateToken{
				res: "refresh-token",
				err: nil,
			},
			want: &model.AuthResponse{
				AccessToken:  "access-token",
				RefreshToken: "refresh-token",
			},
			wantErr: false,
		},
		{
			name: "error user not found",
			args: args{
				payload: &model.UserLoginPayload{
					Username: userEmail,
					Password: "strongpassword",
				},
			},
			mockFindByUsername: &mockFindByUsername{
				res: nil,
				err: nil,
			},
			mockFindByEmail: &mockFindByEmail{
				res: nil,
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error password not match",
			args: args{
				payload: &model.UserLoginPayload{
					Username: username,
					Password: "strongpassword",
				},
			},
			mockFindByUsername: &mockFindByUsername{
				res: &model.User{
					ID:       userID,
					FullName: "user",
					Username: username,
					Email:    userEmail,
					Password: "randompassword",
				},
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error generate token",
			args: args{
				payload: &model.UserLoginPayload{
					Username: username,
					Password: "strongpassword",
				},
			},
			mockFindByUsername: &mockFindByUsername{
				res: &model.User{
					ID:       userID,
					FullName: "user",
					Username: username,
					Email:    userEmail,
					Password: userPassword,
				},
				err: nil,
			},
			mockCreateAccessToken: &mockCreateToken{
				res: "",
				err: errors.New("redis error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.TODO()

			userRepo := mock.NewMockUserRepository(ctrl)
			tokenRepo := mock.NewMockTokenRepository(ctrl)

			if tt.mockFindByUsername != nil {
				userRepo.EXPECT().FindByUsername(gomock.Any(), tt.args.payload.Username).
					Times(1).
					Return(tt.mockFindByUsername.res, tt.mockFindByUsername.err)
			}

			if tt.mockFindByEmail != nil {
				userRepo.EXPECT().FindByEmail(gomock.Any(), tt.args.payload.Username).
					Times(1).
					Return(tt.mockFindByEmail.res, tt.mockFindByEmail.err)
			}

			if tt.mockCreateAccessToken != nil {
				tokenRepo.EXPECT().Create(gomock.Any(), userID, gomock.Any(), model.AccessToken).Times(1).DoAndReturn(func(ctx context.Context, userID string, tokenID string, tokenType model.TokenType) (string, error) {
					return tt.mockCreateAccessToken.res, tt.mockCreateAccessToken.err
				})
			}

			if tt.mockCreateRefreshToken != nil {
				tokenRepo.EXPECT().Create(gomock.Any(), userID, gomock.Any(), model.RefreshToken).Times(1).DoAndReturn(func(ctx context.Context, userID string, tokenID string, tokenType model.TokenType) (string, error) {
					return tt.mockCreateRefreshToken.res, tt.mockCreateRefreshToken.err
				})
			}

			uc := NewUserUsecase()
			err := uc.InjectUserRepo(userRepo)
			utils.ContinueOrFatal(err)
			err = uc.InjectTokenRepo(tokenRepo)
			utils.ContinueOrFatal(err)

			got, err := uc.Login(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_GetUserInfo(t *testing.T) {
	var (
		userID = utils.GenerateUUID()
	)
	type mockFindByID struct {
		res *model.User
		err error
	}
	type args struct {
		payload *model.GetUserInfoPayload
	}
	tests := []struct {
		name         string
		args         args
		mockFindByID *mockFindByID
		want         *model.UserInfoResponse
		wantErr      bool
	}{
		{
			name: "success",
			args: args{
				payload: &model.GetUserInfoPayload{
					ID: userID,
				},
			},
			mockFindByID: &mockFindByID{
				res: &model.User{
					ID:       userID,
					FullName: "user",
					Username: "username",
					Email:    "user@gmail.com",
				},
				err: nil,
			},
			want: &model.UserInfoResponse{
				ID:       userID,
				FullName: "user",
				Username: "username",
				Email:    "user@gmail.com",
			},
			wantErr: false,
		},
		{
			name: "error find user",
			args: args{
				payload: &model.GetUserInfoPayload{
					ID: userID,
				},
			},
			mockFindByID: &mockFindByID{
				res: nil,
				err: errors.New("db error"),
			},
			wantErr: true,
		},
		{
			name: "error user not found",
			args: args{
				payload: &model.GetUserInfoPayload{
					ID: userID,
				},
			},
			mockFindByID: &mockFindByID{
				res: nil,
				err: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.TODO()

			userRepo := mock.NewMockUserRepository(ctrl)

			if tt.mockFindByID != nil {
				userRepo.EXPECT().
					FindByID(gomock.Any(), tt.args.payload.ID).
					Times(1).
					Return(tt.mockFindByID.res, tt.mockFindByID.err)
			}

			uc := NewUserUsecase()
			err := uc.InjectUserRepo(userRepo)
			utils.ContinueOrFatal(err)

			got, err := uc.GetUserInfo(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.GetUserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_RefreshToken(t *testing.T) {
	var (
		userID  = utils.GenerateUUID()
		tokenID = utils.GenerateUUID()
	)
	type mockIsValidToken struct {
		res bool
		err error
	}
	type mockRevokeToken struct {
		err error
	}
	type mockCreateToken struct {
		res string
		err error
	}
	type args struct {
		payload *model.RefreshTokenPayload
	}
	tests := []struct {
		name                   string
		args                   args
		mockIsValidToken       *mockIsValidToken
		mockRevokeAccessToken  *mockRevokeToken
		mockRevokeRefreshToken *mockRevokeToken
		mockCreateAccessToken  *mockCreateToken
		mockCreateRefreshToken *mockCreateToken
		want                   *model.AuthResponse
		wantErr                bool
	}{
		{
			name: "success",
			args: args{
				payload: &model.RefreshTokenPayload{
					UserID:  userID,
					TokenID: tokenID,
				},
			},
			mockIsValidToken: &mockIsValidToken{
				res: true,
				err: nil,
			},
			mockRevokeAccessToken: &mockRevokeToken{
				err: nil,
			},
			mockRevokeRefreshToken: &mockRevokeToken{
				err: nil,
			},
			mockCreateAccessToken: &mockCreateToken{
				res: "access-token",
				err: nil,
			},
			mockCreateRefreshToken: &mockCreateToken{
				res: "refresh-token",
				err: nil,
			},
			want: &model.AuthResponse{
				AccessToken:  "access-token",
				RefreshToken: "refresh-token",
			},
			wantErr: false,
		},
		{
			name: "error invalid token",
			args: args{
				payload: &model.RefreshTokenPayload{
					UserID:  userID,
					TokenID: tokenID,
				},
			},
			mockIsValidToken: &mockIsValidToken{
				res: false,
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error validate token",
			args: args{
				payload: &model.RefreshTokenPayload{
					UserID:  userID,
					TokenID: tokenID,
				},
			},
			mockIsValidToken: &mockIsValidToken{
				res: false,
				err: errors.New("error"),
			},
			wantErr: true,
		},
		{
			name: "error revoke access token",
			args: args{
				payload: &model.RefreshTokenPayload{
					UserID:  userID,
					TokenID: tokenID,
				},
			},
			mockIsValidToken: &mockIsValidToken{
				res: true,
				err: nil,
			},
			mockRevokeAccessToken: &mockRevokeToken{
				err: errors.New("error"),
			},
			wantErr: true,
		},
		{
			name: "error revoke refresh token",
			args: args{
				payload: &model.RefreshTokenPayload{
					UserID:  userID,
					TokenID: tokenID,
				},
			},
			mockIsValidToken: &mockIsValidToken{
				res: true,
				err: nil,
			},
			mockRevokeAccessToken: &mockRevokeToken{
				err: nil,
			},
			mockRevokeRefreshToken: &mockRevokeToken{
				err: errors.New("error"),
			},
			wantErr: true,
		},
		{
			name: "error generate new token",
			args: args{
				payload: &model.RefreshTokenPayload{
					UserID:  userID,
					TokenID: tokenID,
				},
			},
			mockIsValidToken: &mockIsValidToken{
				res: true,
				err: nil,
			},
			mockRevokeAccessToken: &mockRevokeToken{
				err: nil,
			},
			mockRevokeRefreshToken: &mockRevokeToken{
				err: nil,
			},
			mockCreateAccessToken: &mockCreateToken{
				res: "",
				err: errors.New("redis error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.TODO()

			tokenRepo := mock.NewMockTokenRepository(ctrl)

			if tt.mockIsValidToken != nil {
				tokenRepo.EXPECT().
					IsValidToken(gomock.Any(), tt.args.payload.UserID, tt.args.payload.TokenID, model.RefreshToken).
					Times(1).
					Return(tt.mockIsValidToken.res, tt.mockIsValidToken.err)
			}

			if tt.mockRevokeAccessToken != nil {
				tokenRepo.EXPECT().
					Revoke(gomock.Any(), tt.args.payload.UserID, tt.args.payload.TokenID, model.AccessToken).AnyTimes().
					Times(1).
					Return(tt.mockRevokeAccessToken.err)
			}

			if tt.mockRevokeRefreshToken != nil {
				tokenRepo.EXPECT().
					Revoke(gomock.Any(), tt.args.payload.UserID, tt.args.payload.TokenID, model.RefreshToken).AnyTimes().
					Times(1).
					Return(tt.mockRevokeRefreshToken.err)
			}

			if tt.mockCreateAccessToken != nil {
				tokenRepo.EXPECT().
					Create(gomock.Any(), tt.args.payload.UserID, gomock.Any(), model.AccessToken).
					Times(1).
					Return(tt.mockCreateAccessToken.res, tt.mockCreateAccessToken.err)
			}

			if tt.mockCreateRefreshToken != nil {
				tokenRepo.EXPECT().
					Create(gomock.Any(), tt.args.payload.UserID, gomock.Any(), model.RefreshToken).
					Times(1).
					Return(tt.mockCreateRefreshToken.res, tt.mockCreateRefreshToken.err)
			}

			uc := NewUserUsecase()
			err := uc.InjectTokenRepo(tokenRepo)
			utils.ContinueOrFatal(err)

			got, err := uc.RefreshToken(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.RefreshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_Logout(t *testing.T) {
	var (
		userID  = utils.GenerateUUID()
		tokenID = utils.GenerateUUID()
	)
	type mockRevokeToken struct {
		err error
	}
	type args struct {
		payload *model.UserLogoutPayload
	}
	tests := []struct {
		name                   string
		args                   args
		mockRevokeAccessToken  *mockRevokeToken
		mockRevokeRefreshToken *mockRevokeToken
		wantErr                bool
	}{
		{
			name: "success",
			args: args{
				payload: &model.UserLogoutPayload{
					UserID:  userID,
					TokenID: tokenID,
				},
			},
			mockRevokeAccessToken: &mockRevokeToken{
				err: nil,
			},
			mockRevokeRefreshToken: &mockRevokeToken{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "error revoke access token",
			args: args{
				payload: &model.UserLogoutPayload{
					UserID:  userID,
					TokenID: tokenID,
				},
			},
			mockRevokeAccessToken: &mockRevokeToken{
				err: errors.New("error"),
			},
			wantErr: true,
		},
		{
			name: "error revoke refresh token",
			args: args{
				payload: &model.UserLogoutPayload{
					UserID:  userID,
					TokenID: tokenID,
				},
			},
			mockRevokeAccessToken: &mockRevokeToken{
				err: nil,
			},
			mockRevokeRefreshToken: &mockRevokeToken{
				err: errors.New("error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.TODO()

			tokenRepo := mock.NewMockTokenRepository(ctrl)

			if tt.mockRevokeAccessToken != nil {
				tokenRepo.EXPECT().
					Revoke(gomock.Any(), tt.args.payload.UserID, tt.args.payload.TokenID, model.AccessToken).AnyTimes().
					Times(1).
					Return(tt.mockRevokeAccessToken.err)
			}

			if tt.mockRevokeRefreshToken != nil {
				tokenRepo.EXPECT().
					Revoke(gomock.Any(), tt.args.payload.UserID, tt.args.payload.TokenID, model.RefreshToken).AnyTimes().
					Times(1).
					Return(tt.mockRevokeRefreshToken.err)
			}

			uc := NewUserUsecase()
			err := uc.InjectTokenRepo(tokenRepo)
			utils.ContinueOrFatal(err)

			if err := uc.Logout(ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.Logout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

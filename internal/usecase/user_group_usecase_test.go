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

func Test_userGroupUsecase_Create(t *testing.T) {
	var (
		userID  = utils.GenerateUUID()
		groupID = utils.GenerateUUID()
	)
	type mockHasAccess struct {
		err error
	}
	type mockFindGroupByID struct {
		res *model.Group
		err error
	}
	type mockFindUserByID struct {
		res *model.User
		err error
	}
	type mockFindUserGroup struct {
		res *model.UserGroup
		err error
	}
	type mockCreate struct {
		res *model.UserGroup
		err error
	}
	type args struct {
		userID  string
		payload *model.CreateUserGroupPayload
	}
	tests := []struct {
		name              string
		args              args
		mockHasAccess     *mockHasAccess
		mockFindGroupByID *mockFindGroupByID
		mockFindUserByID  *mockFindUserByID
		mockFindUserGroup *mockFindUserGroup
		mockCreate        *mockCreate
		want              *model.UserGroup
		wantErr           bool
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				payload: &model.CreateUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindUserByID: &mockFindUserByID{
				res: &model.User{
					ID:       userID,
					FullName: "user",
				},
				err: nil,
			},
			mockFindGroupByID: &mockFindGroupByID{
				res: &model.Group{
					ID:   groupID,
					Name: "group",
				},
				err: nil,
			},
			mockFindUserGroup: &mockFindUserGroup{
				res: nil,
				err: nil,
			},
			mockCreate: &mockCreate{
				res: &model.UserGroup{
					UserID:  userID,
					GroupID: groupID,
				},
				err: nil,
			},
			want: &model.UserGroup{
				UserID:  userID,
				GroupID: groupID,
			},
			wantErr: false,
		},
		{
			name: "error unauthorized access",
			args: args{
				userID: userID,
				payload: &model.CreateUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: model.ErrUnauthorizeAccess,
			},
			wantErr: true,
		},
		{
			name: "error user not found",
			args: args{
				userID: userID,
				payload: &model.CreateUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindUserByID: &mockFindUserByID{
				res: nil,
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error find user",
			args: args{
				userID: userID,
				payload: &model.CreateUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindUserByID: &mockFindUserByID{
				res: nil,
				err: errors.New("db error"),
			},
			wantErr: true,
		},
		{
			name: "error group not found",
			args: args{
				userID: userID,
				payload: &model.CreateUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindUserByID: &mockFindUserByID{
				res: &model.User{
					ID:       userID,
					FullName: "user",
				},
				err: nil,
			},
			mockFindGroupByID: &mockFindGroupByID{
				res: nil,
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error find group",
			args: args{
				userID: userID,
				payload: &model.CreateUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindUserByID: &mockFindUserByID{
				res: &model.User{
					ID:       userID,
					FullName: "user",
				},
				err: nil,
			},
			mockFindGroupByID: &mockFindGroupByID{
				res: nil,
				err: errors.New("db error"),
			},
			wantErr: true,
		},
		{
			name: "error check exisiting data",
			args: args{
				userID: userID,
				payload: &model.CreateUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindUserByID: &mockFindUserByID{
				res: &model.User{
					ID:       userID,
					FullName: "user",
				},
				err: nil,
			},
			mockFindGroupByID: &mockFindGroupByID{
				res: &model.Group{
					ID:   groupID,
					Name: "group",
				},
				err: nil,
			},
			mockFindUserGroup: &mockFindUserGroup{
				res: nil,
				err: errors.New("db error"),
			},
			wantErr: true,
		},
		{
			name: "error user group already exist",
			args: args{
				userID: userID,
				payload: &model.CreateUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindUserByID: &mockFindUserByID{
				res: &model.User{
					ID:       userID,
					FullName: "user",
				},
				err: nil,
			},
			mockFindGroupByID: &mockFindGroupByID{
				res: &model.Group{
					ID:   groupID,
					Name: "group",
				},
				err: nil,
			},
			mockFindUserGroup: &mockFindUserGroup{
				res: &model.UserGroup{
					UserID:  userID,
					GroupID: groupID,
				},
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error create user group",
			args: args{
				userID: userID,
				payload: &model.CreateUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindUserByID: &mockFindUserByID{
				res: &model.User{
					ID:       userID,
					FullName: "user",
				},
				err: nil,
			},
			mockFindGroupByID: &mockFindGroupByID{
				res: &model.Group{
					ID:   groupID,
					Name: "group",
				},
				err: nil,
			},
			mockFindUserGroup: &mockFindUserGroup{
				res: nil,
				err: nil,
			},
			mockCreate: &mockCreate{
				res: nil,
				err: errors.New("db error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.TODO()
			ctx = context.WithValue(ctx, constant.KeyUserIDCtx, tt.args.userID)

			groupRepo := mock.NewMockGroupRepository(ctrl)
			userRepo := mock.NewMockUserRepository(ctrl)
			userGroupRepo := mock.NewMockUserGroupRepository(ctrl)
			authUsecase := mock.NewMockAuthUsecase(ctrl)

			if tt.mockHasAccess != nil {
				authUsecase.EXPECT().HasAccess(gomock.Any(), gomock.Any()).Times(1).Return(tt.mockHasAccess.err)
			}

			if tt.mockFindGroupByID != nil {
				groupRepo.EXPECT().FindByID(gomock.Any(), tt.args.payload.GroupID).
					Times(1).
					Return(
						tt.mockFindGroupByID.res,
						tt.mockFindGroupByID.err,
					)
			}

			if tt.mockFindUserByID != nil {
				userRepo.EXPECT().FindByID(gomock.Any(), tt.args.payload.UserID).
					Times(1).
					Return(
						tt.mockFindUserByID.res,
						tt.mockFindUserByID.err,
					)
			}

			if tt.mockFindUserGroup != nil {
				userGroupRepo.EXPECT().
					FindByUserIDAndGroupID(gomock.Any(), tt.args.payload.UserID, tt.args.payload.GroupID).
					Times(1).
					Return(tt.mockFindUserGroup.res, tt.mockFindUserGroup.err)
			}

			if tt.mockCreate != nil {
				userGroupRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).DoAndReturn(func(ctx context.Context, data *model.UserGroup) error {
					data.GroupID = tt.args.payload.GroupID
					data.UserID = tt.args.payload.UserID
					return tt.mockCreate.err
				})
			}

			uc := NewUserGroupUsecase()
			err := uc.InjectAuthUsecase(authUsecase)
			utils.ContinueOrFatal(err)
			err = uc.InjectUserRepo(userRepo)
			utils.ContinueOrFatal(err)
			err = uc.InjectGroupRepo(groupRepo)
			utils.ContinueOrFatal(err)
			err = uc.InjectUserGroupRepo(userGroupRepo)
			utils.ContinueOrFatal(err)

			got, err := uc.Create(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("userGroupUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userGroupUsecase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userGroupUsecase_FindByUserIDAndGroupID(t *testing.T) {
	var (
		userID  = utils.GenerateUUID()
		groupID = utils.GenerateUUID()
	)
	type mockHasAccess struct {
		err error
	}
	type mockFindUserGroup struct {
		res *model.UserGroup
		err error
	}
	type args struct {
		userID  string
		payload *model.FindUserGroupPayload
	}
	tests := []struct {
		name              string
		args              args
		mockHasAccess     *mockHasAccess
		mockFindUserGroup *mockFindUserGroup
		want              *model.UserGroup
		wantErr           bool
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				payload: &model.FindUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindUserGroup: &mockFindUserGroup{
				res: &model.UserGroup{
					UserID:  userID,
					GroupID: groupID,
				},
				err: nil,
			},
			want: &model.UserGroup{
				UserID:  userID,
				GroupID: groupID,
			},
			wantErr: false,
		},
		{
			name: "error unauthorized access",
			args: args{
				userID: userID,
				payload: &model.FindUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: model.ErrUnauthorizeAccess,
			},
			wantErr: true,
		},
		{
			name: "error find user group",
			args: args{
				userID: userID,
				payload: &model.FindUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindUserGroup: &mockFindUserGroup{
				res: nil,
				err: errors.New("db error"),
			},
			wantErr: true,
		},
		{
			name: "error user group not found",
			args: args{
				userID: userID,
				payload: &model.FindUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindUserGroup: &mockFindUserGroup{
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
			ctx = context.WithValue(ctx, constant.KeyUserIDCtx, tt.args.userID)

			groupRepo := mock.NewMockGroupRepository(ctrl)
			userRepo := mock.NewMockUserRepository(ctrl)
			userGroupRepo := mock.NewMockUserGroupRepository(ctrl)
			authUsecase := mock.NewMockAuthUsecase(ctrl)

			if tt.mockHasAccess != nil {
				authUsecase.EXPECT().HasAccess(gomock.Any(), gomock.Any()).Times(1).Return(tt.mockHasAccess.err)
			}

			if tt.mockFindUserGroup != nil {
				userGroupRepo.EXPECT().
					FindByUserIDAndGroupID(gomock.Any(), tt.args.payload.UserID, tt.args.payload.GroupID).
					Times(1).
					Return(tt.mockFindUserGroup.res, tt.mockFindUserGroup.err)
			}

			uc := NewUserGroupUsecase()
			err := uc.InjectAuthUsecase(authUsecase)
			utils.ContinueOrFatal(err)
			err = uc.InjectUserRepo(userRepo)
			utils.ContinueOrFatal(err)
			err = uc.InjectGroupRepo(groupRepo)
			utils.ContinueOrFatal(err)
			err = uc.InjectUserGroupRepo(userGroupRepo)
			utils.ContinueOrFatal(err)

			got, err := uc.FindByUserIDAndGroupID(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("userGroupUsecase.FindByUserIDAndGroupID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userGroupUsecase.FindByUserIDAndGroupID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userGroupUsecase_DeleteByUserIDAndGroupID(t *testing.T) {
	var (
		userID  = utils.GenerateUUID()
		groupID = utils.GenerateUUID()
	)
	type mockHasAccess struct {
		err error
	}
	type mockFindUserGroup struct {
		res *model.UserGroup
		err error
	}
	type mockDelete struct {
		err error
	}
	type args struct {
		userID  string
		payload *model.DeleteUserGroupPayload
	}
	tests := []struct {
		name              string
		args              args
		mockHasAccess     *mockHasAccess
		mockFindUserGroup *mockFindUserGroup
		mockDelete        *mockDelete
		wantErr           bool
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				payload: &model.DeleteUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindUserGroup: &mockFindUserGroup{
				res: &model.UserGroup{
					UserID:  userID,
					GroupID: groupID,
				},
				err: nil,
			},
			mockDelete: &mockDelete{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "error unauthorized access",
			args: args{
				userID: userID,
				payload: &model.DeleteUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: model.ErrUnauthorizeAccess,
			},
			wantErr: true,
		},
		{
			name: "error find user group",
			args: args{
				userID: userID,
				payload: &model.DeleteUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindUserGroup: &mockFindUserGroup{
				res: nil,
				err: errors.New("db error"),
			},
			wantErr: true,
		},
		{
			name: "error user group not found",
			args: args{
				userID: userID,
				payload: &model.DeleteUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindUserGroup: &mockFindUserGroup{
				res: nil,
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error delete user group",
			args: args{
				userID: userID,
				payload: &model.DeleteUserGroupPayload{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindUserGroup: &mockFindUserGroup{
				res: &model.UserGroup{
					UserID:  userID,
					GroupID: groupID,
				},
				err: nil,
			},
			mockDelete: &mockDelete{
				err: errors.New("db error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.TODO()
			ctx = context.WithValue(ctx, constant.KeyUserIDCtx, tt.args.userID)

			groupRepo := mock.NewMockGroupRepository(ctrl)
			userRepo := mock.NewMockUserRepository(ctrl)
			userGroupRepo := mock.NewMockUserGroupRepository(ctrl)
			authUsecase := mock.NewMockAuthUsecase(ctrl)

			if tt.mockHasAccess != nil {
				authUsecase.EXPECT().HasAccess(gomock.Any(), gomock.Any()).Times(1).Return(tt.mockHasAccess.err)
			}

			if tt.mockFindUserGroup != nil {
				userGroupRepo.EXPECT().
					FindByUserIDAndGroupID(gomock.Any(), tt.args.payload.UserID, tt.args.payload.GroupID).
					Times(1).
					Return(tt.mockFindUserGroup.res, tt.mockFindUserGroup.err)
			}

			if tt.mockDelete != nil {
				userGroupRepo.EXPECT().
					DeleteByUserIDAndGroupID(gomock.Any(), tt.args.payload.UserID, tt.args.payload.GroupID).
					Times(1).
					Return(tt.mockDelete.err)
			}

			uc := NewUserGroupUsecase()
			err := uc.InjectAuthUsecase(authUsecase)
			utils.ContinueOrFatal(err)
			err = uc.InjectUserRepo(userRepo)
			utils.ContinueOrFatal(err)
			err = uc.InjectGroupRepo(groupRepo)
			utils.ContinueOrFatal(err)
			err = uc.InjectUserGroupRepo(userGroupRepo)
			utils.ContinueOrFatal(err)

			if err := uc.DeleteByUserIDAndGroupID(ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("userGroupUsecase.DeleteByUserIDAndGroupID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

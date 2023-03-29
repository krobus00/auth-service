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

func Test_groupPermissionUsecase_Create(t *testing.T) {
	var (
		userID       = utils.GenerateUUID()
		groupID      = utils.GenerateUUID()
		permissionID = utils.GenerateUUID()
	)
	type args struct {
		userID  string
		payload *model.CreateGroupPermissionPayload
	}
	type mockHasAccess struct {
		err error
	}
	type mockFindGroupByID struct {
		res *model.Group
		err error
	}
	type mockFindPermissionByID struct {
		res *model.Permission
		err error
	}
	type mockFindGroupPermission struct {
		res *model.GroupPermission
		err error
	}
	type mockCreate struct {
		err error
	}
	tests := []struct {
		name                    string
		args                    args
		mockHasAccess           *mockHasAccess
		mockFindGroupByID       *mockFindGroupByID
		mockFindPermissionByID  *mockFindPermissionByID
		mockFindGroupPermission *mockFindGroupPermission
		mockCreate              *mockCreate
		want                    *model.GroupPermission
		wantErr                 bool
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				payload: &model.CreateGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindGroupByID: &mockFindGroupByID{
				res: &model.Group{
					ID:   groupID,
					Name: "group1",
				},
				err: nil,
			},
			mockFindPermissionByID: &mockFindPermissionByID{
				res: &model.Permission{
					ID:   permissionID,
					Name: "TEST_READ",
				},
				err: nil,
			},
			mockFindGroupPermission: &mockFindGroupPermission{
				res: nil,
				err: nil,
			},
			mockCreate: &mockCreate{
				err: nil,
			},
			want: &model.GroupPermission{
				GroupID:      groupID,
				PermissionID: permissionID,
			},
			wantErr: false,
		},
		{
			name: "error unauthorized access",
			args: args{
				userID: userID,
				payload: &model.CreateGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: model.ErrUnauthorizeAccess,
			},
			wantErr: true,
		},
		{
			name: "error group not found",
			args: args{
				userID: userID,
				payload: &model.CreateGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
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
				payload: &model.CreateGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindGroupByID: &mockFindGroupByID{
				res: nil,
				err: errors.New("db error"),
			},
			wantErr: true,
		},
		{
			name: "error permission not found",
			args: args{
				userID: userID,
				payload: &model.CreateGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindGroupByID: &mockFindGroupByID{
				res: &model.Group{
					ID:   groupID,
					Name: "group1",
				},
				err: nil,
			},
			mockFindPermissionByID: &mockFindPermissionByID{
				res: nil,
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error find permission",
			args: args{
				userID: userID,
				payload: &model.CreateGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindGroupByID: &mockFindGroupByID{
				res: &model.Group{
					ID:   groupID,
					Name: "group1",
				},
				err: nil,
			},
			mockFindPermissionByID: &mockFindPermissionByID{
				res: nil,
				err: errors.New("db error"),
			},
			wantErr: true,
		},
		{
			name: "error check exisiting data",
			args: args{
				userID: userID,
				payload: &model.CreateGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindGroupByID: &mockFindGroupByID{
				res: &model.Group{
					ID:   groupID,
					Name: "group1",
				},
				err: nil,
			},
			mockFindPermissionByID: &mockFindPermissionByID{
				res: &model.Permission{
					ID:   permissionID,
					Name: "TEST_READ",
				},
				err: nil,
			},
			mockFindGroupPermission: &mockFindGroupPermission{
				res: nil,
				err: errors.New("db error"),
			},
			wantErr: true,
		},
		{
			name: "error group permission already exist",
			args: args{
				userID: userID,
				payload: &model.CreateGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindGroupByID: &mockFindGroupByID{
				res: &model.Group{
					ID:   groupID,
					Name: "group1",
				},
				err: nil,
			},
			mockFindPermissionByID: &mockFindPermissionByID{
				res: &model.Permission{
					ID:   permissionID,
					Name: "TEST_READ",
				},
				err: nil,
			},
			mockFindGroupPermission: &mockFindGroupPermission{
				res: &model.GroupPermission{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error create new data",
			args: args{
				userID: userID,
				payload: &model.CreateGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindGroupByID: &mockFindGroupByID{
				res: &model.Group{
					ID:   groupID,
					Name: "group1",
				},
				err: nil,
			},
			mockFindPermissionByID: &mockFindPermissionByID{
				res: &model.Permission{
					ID:   permissionID,
					Name: "TEST_READ",
				},
				err: nil,
			},
			mockFindGroupPermission: &mockFindGroupPermission{
				res: nil,
				err: nil,
			},
			mockCreate: &mockCreate{
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
			permissionRepo := mock.NewMockPermissionRepository(ctrl)
			groupPermissionRepo := mock.NewMockGroupPermissionRepository(ctrl)
			authUsecase := mock.NewMockAuthUsecase(ctrl)

			if tt.mockHasAccess != nil {
				authUsecase.EXPECT().HasAccess(ctx, gomock.Any()).Times(1).Return(tt.mockHasAccess.err)
			}

			if tt.mockFindGroupByID != nil {
				groupRepo.EXPECT().FindByID(ctx, tt.args.payload.GroupID).
					Times(1).
					Return(
						tt.mockFindGroupByID.res,
						tt.mockFindGroupByID.err,
					)
			}

			if tt.mockFindPermissionByID != nil {
				permissionRepo.EXPECT().FindByID(ctx, tt.args.payload.PermissionID).
					Times(1).
					Return(
						tt.mockFindPermissionByID.res,
						tt.mockFindPermissionByID.err,
					)
			}

			if tt.mockFindGroupPermission != nil {
				groupPermissionRepo.EXPECT().FindByGroupIDAndPermissionID(ctx, tt.args.payload.GroupID, tt.args.payload.PermissionID).
					Times(1).
					Return(tt.mockFindGroupPermission.res, tt.mockFindGroupPermission.err)
			}

			if tt.mockCreate != nil {
				groupPermissionRepo.EXPECT().Create(ctx, &model.GroupPermission{
					GroupID:      tt.args.payload.GroupID,
					PermissionID: tt.args.payload.PermissionID,
				}).
					Times(1).
					Return(tt.mockCreate.err)
			}

			uc := NewGroupPermissionUsecase()
			err := uc.InjectAuthUsecase(authUsecase)
			utils.ContinueOrFatal(err)
			err = uc.InjectGroupPermissionRepo(groupPermissionRepo)
			utils.ContinueOrFatal(err)
			err = uc.InjectGroupRepo(groupRepo)
			utils.ContinueOrFatal(err)
			err = uc.InjectPermisisonRepo(permissionRepo)
			utils.ContinueOrFatal(err)

			got, err := uc.Create(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("groupPermissionUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupPermissionUsecase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupPermissionUsecase_FindByGroupIDAndPermissionID(t *testing.T) {
	var (
		userID       = utils.GenerateUUID()
		groupID      = utils.GenerateUUID()
		permissionID = utils.GenerateUUID()
	)
	type mockHasAccess struct {
		err error
	}
	type mockFindGroupPermission struct {
		res *model.GroupPermission
		err error
	}
	type args struct {
		userID  string
		payload *model.FindGroupPermissionPayload
	}
	tests := []struct {
		name                    string
		args                    args
		mockHasAccess           *mockHasAccess
		mockFindGroupPermission *mockFindGroupPermission
		want                    *model.GroupPermission
		wantErr                 bool
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				payload: &model.FindGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindGroupPermission: &mockFindGroupPermission{
				res: &model.GroupPermission{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
				err: nil,
			},
			want: &model.GroupPermission{
				GroupID:      groupID,
				PermissionID: permissionID,
			},
			wantErr: false,
		},
		{
			name: "error unauthorized access",
			args: args{
				userID: userID,
				payload: &model.FindGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: model.ErrUnauthorizeAccess,
			},
			wantErr: true,
		},
		{
			name: "error group permission not found",
			args: args{
				userID: userID,
				payload: &model.FindGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindGroupPermission: &mockFindGroupPermission{
				res: nil,
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error find group permission",
			args: args{
				userID: userID,
				payload: &model.FindGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindGroupPermission: &mockFindGroupPermission{
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

			groupPermissionRepo := mock.NewMockGroupPermissionRepository(ctrl)
			authUsecase := mock.NewMockAuthUsecase(ctrl)

			if tt.mockHasAccess != nil {
				authUsecase.EXPECT().HasAccess(ctx, gomock.Any()).Times(1).Return(tt.mockHasAccess.err)
			}

			if tt.mockFindGroupPermission != nil {
				groupPermissionRepo.EXPECT().FindByGroupIDAndPermissionID(ctx, tt.args.payload.GroupID, tt.args.payload.PermissionID).
					Times(1).
					Return(tt.mockFindGroupPermission.res, tt.mockFindGroupPermission.err)
			}

			uc := NewGroupPermissionUsecase()
			err := uc.InjectAuthUsecase(authUsecase)
			utils.ContinueOrFatal(err)
			err = uc.InjectGroupPermissionRepo(groupPermissionRepo)
			utils.ContinueOrFatal(err)

			got, err := uc.FindByGroupIDAndPermissionID(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("groupPermissionUsecase.FindByGroupIDAndPermissionID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupPermissionUsecase.FindByGroupIDAndPermissionID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupPermissionUsecase_DeleteByGroupIDAndPermissionID(t *testing.T) {
	var (
		userID       = utils.GenerateUUID()
		groupID      = utils.GenerateUUID()
		permissionID = utils.GenerateUUID()
	)
	type mockHasAccess struct {
		err error
	}
	type mockFindGroupPermission struct {
		res *model.GroupPermission
		err error
	}
	type mockDeleteGroupPermission struct {
		err error
	}
	type args struct {
		userID  string
		payload *model.DeleteGroupPermissionPayload
	}
	tests := []struct {
		name                      string
		args                      args
		mockHasAccess             *mockHasAccess
		mockFindGroupPermission   *mockFindGroupPermission
		mockDeleteGroupPermission *mockDeleteGroupPermission
		wantErr                   bool
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				payload: &model.DeleteGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindGroupPermission: &mockFindGroupPermission{
				res: &model.GroupPermission{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
				err: nil,
			},
			mockDeleteGroupPermission: &mockDeleteGroupPermission{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "error unauthorized access",
			args: args{
				userID: userID,
				payload: &model.DeleteGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: model.ErrUnauthorizeAccess,
			},
			wantErr: true,
		},
		{
			name: "error group permission not found",
			args: args{
				userID: userID,
				payload: &model.DeleteGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindGroupPermission: &mockFindGroupPermission{
				res: nil,
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error find group permission",
			args: args{
				userID: userID,
				payload: &model.DeleteGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindGroupPermission: &mockFindGroupPermission{
				res: nil,
				err: errors.New("db error"),
			},
			wantErr: true,
		},
		{
			name: "error delete group permission",
			args: args{
				userID: userID,
				payload: &model.DeleteGroupPermissionPayload{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindGroupPermission: &mockFindGroupPermission{
				res: &model.GroupPermission{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
				err: nil,
			},
			mockDeleteGroupPermission: &mockDeleteGroupPermission{
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

			groupPermissionRepo := mock.NewMockGroupPermissionRepository(ctrl)
			authUsecase := mock.NewMockAuthUsecase(ctrl)

			if tt.mockHasAccess != nil {
				authUsecase.EXPECT().HasAccess(ctx, gomock.Any()).Times(1).Return(tt.mockHasAccess.err)
			}

			if tt.mockFindGroupPermission != nil {
				groupPermissionRepo.EXPECT().FindByGroupIDAndPermissionID(ctx, tt.args.payload.GroupID, tt.args.payload.PermissionID).
					Times(1).
					Return(tt.mockFindGroupPermission.res, tt.mockFindGroupPermission.err)
			}

			if tt.mockDeleteGroupPermission != nil {
				groupPermissionRepo.EXPECT().DeleteByGroupIDAndPermissionID(ctx, tt.args.payload.GroupID, tt.args.payload.PermissionID).
					Times(1).
					Return(tt.mockDeleteGroupPermission.err)
			}

			uc := NewGroupPermissionUsecase()
			err := uc.InjectAuthUsecase(authUsecase)
			utils.ContinueOrFatal(err)
			err = uc.InjectGroupPermissionRepo(groupPermissionRepo)
			utils.ContinueOrFatal(err)

			if err := uc.DeleteByGroupIDAndPermissionID(ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("groupPermissionUsecase.DeleteByGroupIDAndPermissionID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

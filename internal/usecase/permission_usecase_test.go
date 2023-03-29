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

func Test_permissionUsecase_Create(t *testing.T) {
	var (
		userID       = utils.GenerateUUID()
		permissionID = utils.GenerateUUID()
	)
	type mockHasAccess struct {
		err error
	}
	type mockFindByName struct {
		res *model.Permission
		err error
	}
	type mockCreate struct {
		res *model.Permission
		err error
	}
	type args struct {
		userID  string
		payload *model.CreatePermissionPayload
	}
	tests := []struct {
		name           string
		args           args
		mockHasAccess  *mockHasAccess
		mockFindByName *mockFindByName
		mockCreate     *mockCreate
		want           *model.Permission
		wantErr        bool
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				payload: &model.CreatePermissionPayload{
					Name: "permission1",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByName: &mockFindByName{
				res: nil,
				err: nil,
			},
			mockCreate: &mockCreate{
				res: &model.Permission{
					ID:   permissionID,
					Name: "permission1",
				},
				err: nil,
			},
			want: &model.Permission{
				ID:   permissionID,
				Name: "permission1",
			},
			wantErr: false,
		},
		{
			name: "error unauthorized access",
			args: args{
				userID: userID,
				payload: &model.CreatePermissionPayload{
					Name: "permission1",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: model.ErrUnauthorizeAccess,
			},
			wantErr: true,
		},
		{
			name: "error permission already exist",
			args: args{
				userID: userID,
				payload: &model.CreatePermissionPayload{
					Name: "permission1",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByName: &mockFindByName{
				res: &model.Permission{
					ID:   permissionID,
					Name: "permission1",
				},
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error find permission",
			args: args{
				userID: userID,
				payload: &model.CreatePermissionPayload{
					Name: "permission1",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByName: &mockFindByName{
				res: nil,
				err: errors.New("db error"),
			},
			wantErr: true,
		},
		{
			name: "error create permission",
			args: args{
				userID: userID,
				payload: &model.CreatePermissionPayload{
					Name: "permission1",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByName: &mockFindByName{
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

			permissionRepo := mock.NewMockPermissionRepository(ctrl)
			authUsecase := mock.NewMockAuthUsecase(ctrl)

			if tt.mockHasAccess != nil {
				authUsecase.EXPECT().HasAccess(ctx, gomock.Any()).Times(1).Return(tt.mockHasAccess.err)
			}

			if tt.mockFindByName != nil {
				permissionRepo.EXPECT().FindByName(ctx, tt.args.payload.Name).
					Times(1).
					Return(tt.mockFindByName.res, tt.mockFindByName.err)
			}

			if tt.mockCreate != nil {
				permissionRepo.EXPECT().Create(ctx, gomock.Any()).Times(1).DoAndReturn(func(ctx context.Context, permission *model.Permission) error {
					permission.ID = permissionID
					permission.Name = tt.args.payload.Name
					return tt.mockCreate.err
				})
			}

			uc := NewPermissionUsecase()
			err := uc.InjectAuthUsecase(authUsecase)
			utils.ContinueOrFatal(err)
			err = uc.InjectPermissionRepo(permissionRepo)
			utils.ContinueOrFatal(err)

			got, err := uc.Create(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("permissionUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permissionUsecase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permissionUsecase_FindByID(t *testing.T) {
	var (
		userID       = utils.GenerateUUID()
		permissionID = utils.GenerateUUID()
	)
	type mockHasAccess struct {
		err error
	}
	type mockFindByID struct {
		res *model.Permission
		err error
	}
	type args struct {
		userID  string
		payload *model.FindPermissionByIDPayload
	}
	tests := []struct {
		name          string
		args          args
		mockHasAccess *mockHasAccess
		mockFindByID  *mockFindByID
		want          *model.Permission
		wantErr       bool
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				payload: &model.FindPermissionByIDPayload{
					ID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByID: &mockFindByID{
				res: &model.Permission{
					ID:   permissionID,
					Name: "permission1",
				},
				err: nil,
			},
			want: &model.Permission{
				ID:   permissionID,
				Name: "permission1",
			},
			wantErr: false,
		},
		{
			name: "error unauthorized access",
			args: args{
				userID: userID,
				payload: &model.FindPermissionByIDPayload{
					ID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: model.ErrUnauthorizeAccess,
			},
			wantErr: true,
		},
		{
			name: "error permission not found",
			args: args{
				userID: userID,
				payload: &model.FindPermissionByIDPayload{
					ID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByID: &mockFindByID{
				res: nil,
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error find permission",
			args: args{
				userID: userID,
				payload: &model.FindPermissionByIDPayload{
					ID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByID: &mockFindByID{
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

			permissionRepo := mock.NewMockPermissionRepository(ctrl)
			authUsecase := mock.NewMockAuthUsecase(ctrl)

			if tt.mockHasAccess != nil {
				authUsecase.EXPECT().HasAccess(ctx, gomock.Any()).Times(1).Return(tt.mockHasAccess.err)
			}

			if tt.mockFindByID != nil {
				permissionRepo.EXPECT().FindByID(ctx, tt.args.payload.ID).
					Times(1).
					Return(tt.mockFindByID.res, tt.mockFindByID.err)
			}

			uc := NewPermissionUsecase()
			err := uc.InjectAuthUsecase(authUsecase)
			utils.ContinueOrFatal(err)
			err = uc.InjectPermissionRepo(permissionRepo)
			utils.ContinueOrFatal(err)

			got, err := uc.FindByID(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("permissionUsecase.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permissionUsecase.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permissionUsecase_FindByName(t *testing.T) {
	var (
		userID       = utils.GenerateUUID()
		permissionID = utils.GenerateUUID()
	)
	type mockHasAccess struct {
		err error
	}
	type mockFindByName struct {
		res *model.Permission
		err error
	}
	type args struct {
		userID  string
		payload *model.FindPermissionByNamePayload
	}
	tests := []struct {
		name           string
		args           args
		mockHasAccess  *mockHasAccess
		mockFindByName *mockFindByName
		want           *model.Permission
		wantErr        bool
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				payload: &model.FindPermissionByNamePayload{
					Name: "permision1",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByName: &mockFindByName{
				res: &model.Permission{
					ID:   permissionID,
					Name: "permission1",
				},
				err: nil,
			},
			want: &model.Permission{
				ID:   permissionID,
				Name: "permission1",
			},
			wantErr: false,
		},
		{
			name: "error unauthorized access",
			args: args{
				userID: userID,
				payload: &model.FindPermissionByNamePayload{
					Name: "permision1",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: model.ErrUnauthorizeAccess,
			},
			wantErr: true,
		},
		{
			name: "error permission not found",
			args: args{
				userID: userID,
				payload: &model.FindPermissionByNamePayload{
					Name: "permision1",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByName: &mockFindByName{
				res: nil,
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error find permission",
			args: args{
				userID: userID,
				payload: &model.FindPermissionByNamePayload{
					Name: "permision1",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByName: &mockFindByName{
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

			permissionRepo := mock.NewMockPermissionRepository(ctrl)
			authUsecase := mock.NewMockAuthUsecase(ctrl)

			if tt.mockHasAccess != nil {
				authUsecase.EXPECT().HasAccess(ctx, gomock.Any()).Times(1).Return(tt.mockHasAccess.err)
			}

			if tt.mockFindByName != nil {
				permissionRepo.EXPECT().FindByName(ctx, tt.args.payload.Name).
					Times(1).
					Return(tt.mockFindByName.res, tt.mockFindByName.err)
			}

			uc := NewPermissionUsecase()
			err := uc.InjectAuthUsecase(authUsecase)
			utils.ContinueOrFatal(err)
			err = uc.InjectPermissionRepo(permissionRepo)
			utils.ContinueOrFatal(err)

			got, err := uc.FindByName(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("permissionUsecase.FindByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permissionUsecase.FindByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permissionUsecase_Update(t *testing.T) {
	var (
		userID       = utils.GenerateUUID()
		permissionID = utils.GenerateUUID()
	)
	type mockHasAccess struct {
		err error
	}
	type mockFindByID struct {
		res *model.Permission
		err error
	}
	type mockUpdate struct {
		res *model.Permission
		err error
	}
	type args struct {
		userID  string
		payload *model.UpdatePermissionPayload
	}
	tests := []struct {
		name          string
		args          args
		mockHasAccess *mockHasAccess
		mockFindByID  *mockFindByID
		mockUpdate    *mockUpdate
		want          *model.Permission
		wantErr       bool
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				payload: &model.UpdatePermissionPayload{
					ID:   permissionID,
					Name: "updated",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByID: &mockFindByID{
				res: &model.Permission{
					ID:   permissionID,
					Name: "permission1",
				},
				err: nil,
			},
			mockUpdate: &mockUpdate{
				res: &model.Permission{
					ID:   permissionID,
					Name: "updated",
				},
				err: nil,
			},
			want: &model.Permission{
				ID:   permissionID,
				Name: "updated",
			},
			wantErr: false,
		},
		{
			name: "error unauthorized access",
			args: args{
				userID: userID,
				payload: &model.UpdatePermissionPayload{
					ID:   permissionID,
					Name: "updated",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: model.ErrUnauthorizeAccess,
			},
			wantErr: true,
		},
		{
			name: "error permission not found",
			args: args{
				userID: userID,
				payload: &model.UpdatePermissionPayload{
					ID:   permissionID,
					Name: "updated",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByID: &mockFindByID{
				res: nil,
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error find permission",
			args: args{
				userID: userID,
				payload: &model.UpdatePermissionPayload{
					ID:   permissionID,
					Name: "updated",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByID: &mockFindByID{
				res: nil,
				err: errors.New("db error"),
			},
			wantErr: true,
		},
		{
			name: "error update permission",
			args: args{
				userID: userID,
				payload: &model.UpdatePermissionPayload{
					ID:   permissionID,
					Name: "updated",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByID: &mockFindByID{
				res: &model.Permission{
					ID:   permissionID,
					Name: "permission1",
				},
				err: nil,
			},
			mockUpdate: &mockUpdate{
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

			permissionRepo := mock.NewMockPermissionRepository(ctrl)
			authUsecase := mock.NewMockAuthUsecase(ctrl)

			if tt.mockHasAccess != nil {
				authUsecase.EXPECT().HasAccess(ctx, gomock.Any()).Times(1).Return(tt.mockHasAccess.err)
			}

			if tt.mockFindByID != nil {
				permissionRepo.EXPECT().FindByID(ctx, tt.args.payload.ID).
					Times(1).
					Return(tt.mockFindByID.res, tt.mockFindByID.err)
			}

			if tt.mockUpdate != nil {
				permissionRepo.EXPECT().Update(ctx, &model.Permission{
					ID:   tt.args.payload.ID,
					Name: tt.mockFindByID.res.Name,
				}).Times(1).DoAndReturn(func(ctx context.Context, permission *model.Permission) error {
					permission.ID = tt.args.payload.ID
					permission.Name = tt.args.payload.Name
					return tt.mockUpdate.err
				})
			}

			uc := NewPermissionUsecase()
			err := uc.InjectAuthUsecase(authUsecase)
			utils.ContinueOrFatal(err)
			err = uc.InjectPermissionRepo(permissionRepo)
			utils.ContinueOrFatal(err)

			got, err := uc.Update(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("permissionUsecase.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permissionUsecase.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permissionUsecase_DeleteByID(t *testing.T) {
	var (
		userID       = utils.GenerateUUID()
		permissionID = utils.GenerateUUID()
	)
	type mockHasAccess struct {
		err error
	}
	type mockFindByID struct {
		res *model.Permission
		err error
	}
	type mockDelete struct {
		err error
	}
	type args struct {
		userID  string
		payload *model.DeletePermissionByIDPayload
	}
	tests := []struct {
		name          string
		args          args
		mockHasAccess *mockHasAccess
		mockFindByID  *mockFindByID
		mockDelete    *mockDelete
		wantErr       bool
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				payload: &model.DeletePermissionByIDPayload{
					ID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByID: &mockFindByID{
				res: &model.Permission{
					ID:   permissionID,
					Name: "permission1",
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
				payload: &model.DeletePermissionByIDPayload{
					ID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: model.ErrUnauthorizeAccess,
			},
			wantErr: true,
		},
		{
			name: "error permission not found",
			args: args{
				userID: userID,
				payload: &model.DeletePermissionByIDPayload{
					ID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByID: &mockFindByID{
				res: nil,
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error find permission",
			args: args{
				userID: userID,
				payload: &model.DeletePermissionByIDPayload{
					ID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByID: &mockFindByID{
				res: nil,
				err: errors.New("db error"),
			},
			wantErr: true,
		},
		{
			name: "error delete permission",
			args: args{
				userID: userID,
				payload: &model.DeletePermissionByIDPayload{
					ID: permissionID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByID: &mockFindByID{
				res: &model.Permission{
					ID:   permissionID,
					Name: "permission1",
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

			permissionRepo := mock.NewMockPermissionRepository(ctrl)
			authUsecase := mock.NewMockAuthUsecase(ctrl)

			if tt.mockHasAccess != nil {
				authUsecase.EXPECT().HasAccess(ctx, gomock.Any()).Times(1).Return(tt.mockHasAccess.err)
			}

			if tt.mockFindByID != nil {
				permissionRepo.EXPECT().FindByID(ctx, tt.args.payload.ID).
					Times(1).
					Return(tt.mockFindByID.res, tt.mockFindByID.err)
			}

			if tt.mockDelete != nil {
				permissionRepo.EXPECT().DeleteByID(ctx, tt.args.payload.ID).
					Times(1).
					Return(tt.mockDelete.err)
			}

			uc := NewPermissionUsecase()
			err := uc.InjectAuthUsecase(authUsecase)
			utils.ContinueOrFatal(err)
			err = uc.InjectPermissionRepo(permissionRepo)
			utils.ContinueOrFatal(err)

			if err := uc.DeleteByID(ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("permissionUsecase.DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

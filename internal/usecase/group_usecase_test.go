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

func Test_groupUsecase_Create(t *testing.T) {
	var (
		userID  = utils.GenerateUUID()
		groupID = utils.GenerateUUID()
	)
	type mockHasAccess struct {
		err error
	}
	type mockFindByName struct {
		res *model.Group
		err error
	}
	type mockCreate struct {
		res *model.Group
		err error
	}
	type args struct {
		userID  string
		payload *model.CreateGroupPayload
	}
	tests := []struct {
		name           string
		args           args
		mockHasAccess  *mockHasAccess
		mockFindByName *mockFindByName
		mockCreate     *mockCreate
		want           *model.Group
		wantErr        bool
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				payload: &model.CreateGroupPayload{
					Name: "group1",
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
				res: &model.Group{
					ID:   groupID,
					Name: "group1",
				},
				err: nil,
			},
			want: &model.Group{
				ID:   groupID,
				Name: "group1",
			},
			wantErr: false,
		},
		{
			name: "error unauthorized access",
			args: args{
				userID: userID,
				payload: &model.CreateGroupPayload{
					Name: "group1",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: model.ErrUnauthorizeAccess,
			},
			wantErr: true,
		},
		{
			name: "error find group",
			args: args{
				userID: userID,
				payload: &model.CreateGroupPayload{
					Name: "group1",
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
			name: "error group already exist",
			args: args{
				userID: userID,
				payload: &model.CreateGroupPayload{
					Name: "group1",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByName: &mockFindByName{
				res: &model.Group{
					ID:   groupID,
					Name: "group1",
				},
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "error create group",
			args: args{
				userID: userID,
				payload: &model.CreateGroupPayload{
					Name: "group1",
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

			groupRepo := mock.NewMockGroupRepository(ctrl)
			authUsecase := mock.NewMockAuthUsecase(ctrl)

			if tt.mockHasAccess != nil {
				authUsecase.EXPECT().HasAccess(gomock.Any(), gomock.Any()).Times(1).Return(tt.mockHasAccess.err)
			}

			if tt.mockFindByName != nil {
				groupRepo.EXPECT().FindByName(gomock.Any(), tt.args.payload.Name).
					Times(1).
					Return(tt.mockFindByName.res, tt.mockFindByName.err)
			}

			if tt.mockCreate != nil {
				groupRepo.EXPECT().Create(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, group *model.Group) error {
					group.ID = groupID
					group.Name = tt.args.payload.Name
					return tt.mockCreate.err
				})
			}

			uc := NewGroupUsecase()
			err := uc.InjectAuthUsecase(authUsecase)
			utils.ContinueOrFatal(err)
			err = uc.InjectGroupRepo(groupRepo)
			utils.ContinueOrFatal(err)

			got, err := uc.Create(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("groupUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupUsecase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupUsecase_FindByID(t *testing.T) {
	var (
		userID  = utils.GenerateUUID()
		groupID = utils.GenerateUUID()
	)
	type mockHasAccess struct {
		err error
	}
	type mockFindByID struct {
		res *model.Group
		err error
	}
	type args struct {
		userID  string
		payload *model.FindGroupByIDPayload
	}
	tests := []struct {
		name          string
		args          args
		mockHasAccess *mockHasAccess
		mockFindByID  *mockFindByID
		want          *model.Group
		wantErr       bool
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				payload: &model.FindGroupByIDPayload{
					ID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByID: &mockFindByID{
				res: &model.Group{
					ID:   groupID,
					Name: "group1",
				},
				err: nil,
			},
			want: &model.Group{
				ID:   groupID,
				Name: "group1",
			},
			wantErr: false,
		},
		{
			name: "error unauthorized access",
			args: args{
				userID: userID,
				payload: &model.FindGroupByIDPayload{
					ID: groupID,
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
				payload: &model.FindGroupByIDPayload{
					ID: groupID,
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
			name: "error find group",
			args: args{
				userID: userID,
				payload: &model.FindGroupByIDPayload{
					ID: groupID,
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

			groupRepo := mock.NewMockGroupRepository(ctrl)
			authUsecase := mock.NewMockAuthUsecase(ctrl)

			if tt.mockHasAccess != nil {
				authUsecase.EXPECT().HasAccess(gomock.Any(), gomock.Any()).Times(1).Return(tt.mockHasAccess.err)
			}

			if tt.mockFindByID != nil {
				groupRepo.EXPECT().FindByID(gomock.Any(), tt.args.payload.ID).
					Times(1).
					Return(tt.mockFindByID.res, tt.mockFindByID.err)
			}

			uc := NewGroupUsecase()
			err := uc.InjectAuthUsecase(authUsecase)
			utils.ContinueOrFatal(err)
			err = uc.InjectGroupRepo(groupRepo)
			utils.ContinueOrFatal(err)

			got, err := uc.FindByID(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("groupUsecase.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupUsecase.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupUsecase_FindByName(t *testing.T) {
	var (
		userID  = utils.GenerateUUID()
		groupID = utils.GenerateUUID()
	)
	type mockHasAccess struct {
		err error
	}
	type mockFindByName struct {
		res *model.Group
		err error
	}
	type args struct {
		userID  string
		payload *model.FindGroupByNamePayload
	}
	tests := []struct {
		name           string
		args           args
		mockHasAccess  *mockHasAccess
		mockFindByName *mockFindByName
		want           *model.Group
		wantErr        bool
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				payload: &model.FindGroupByNamePayload{
					Name: "group1",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByName: &mockFindByName{
				res: &model.Group{
					ID:   groupID,
					Name: "group1",
				},
				err: nil,
			},
			want: &model.Group{
				ID:   groupID,
				Name: "group1",
			},
			wantErr: false,
		},
		{
			name: "error unauthorized access",
			args: args{
				userID: userID,
				payload: &model.FindGroupByNamePayload{
					Name: "group1",
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
				payload: &model.FindGroupByNamePayload{
					Name: "group1",
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
			name: "error find group",
			args: args{
				userID: userID,
				payload: &model.FindGroupByNamePayload{
					Name: "group1",
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

			groupRepo := mock.NewMockGroupRepository(ctrl)
			authUsecase := mock.NewMockAuthUsecase(ctrl)

			if tt.mockHasAccess != nil {
				authUsecase.EXPECT().HasAccess(gomock.Any(), gomock.Any()).Times(1).Return(tt.mockHasAccess.err)
			}

			if tt.mockFindByName != nil {
				groupRepo.EXPECT().FindByName(gomock.Any(), tt.args.payload.Name).
					Times(1).
					Return(tt.mockFindByName.res, tt.mockFindByName.err)
			}

			uc := NewGroupUsecase()
			err := uc.InjectAuthUsecase(authUsecase)
			utils.ContinueOrFatal(err)
			err = uc.InjectGroupRepo(groupRepo)
			utils.ContinueOrFatal(err)

			got, err := uc.FindByName(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("groupUsecase.FindByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupUsecase.FindByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupUsecase_Update(t *testing.T) {
	var (
		userID  = utils.GenerateUUID()
		groupID = utils.GenerateUUID()
	)
	type mockHasAccess struct {
		err error
	}
	type mockFindByID struct {
		res *model.Group
		err error
	}
	type mockUpdate struct {
		res *model.Group
		err error
	}
	type args struct {
		userID  string
		payload *model.UpdateGroupPayload
	}
	tests := []struct {
		name          string
		args          args
		mockHasAccess *mockHasAccess
		mockFindByID  *mockFindByID
		mockUpdate    *mockUpdate
		want          *model.Group
		wantErr       bool
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				payload: &model.UpdateGroupPayload{
					ID:   groupID,
					Name: "updated",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByID: &mockFindByID{
				res: &model.Group{
					ID:   groupID,
					Name: "group1",
				},
				err: nil,
			},
			mockUpdate: &mockUpdate{
				res: &model.Group{
					ID:   groupID,
					Name: "updated",
				},
				err: nil,
			},
			want: &model.Group{
				ID:   groupID,
				Name: "updated",
			},
			wantErr: false,
		},
		{
			name: "error unauthorized access",
			args: args{
				userID: userID,
				payload: &model.UpdateGroupPayload{
					ID:   groupID,
					Name: "updated",
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
				payload: &model.UpdateGroupPayload{
					ID:   groupID,
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
			name: "error find group",
			args: args{
				userID: userID,
				payload: &model.UpdateGroupPayload{
					ID:   groupID,
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
			name: "error update group",
			args: args{
				userID: userID,
				payload: &model.UpdateGroupPayload{
					ID:   groupID,
					Name: "updated",
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByID: &mockFindByID{
				res: &model.Group{
					ID:   groupID,
					Name: "group1",
				},
				err: nil,
			},
			mockUpdate: &mockUpdate{
				res: &model.Group{
					ID:   groupID,
					Name: "updated",
				},
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
			authUsecase := mock.NewMockAuthUsecase(ctrl)

			if tt.mockHasAccess != nil {
				authUsecase.EXPECT().HasAccess(gomock.Any(), gomock.Any()).Times(1).Return(tt.mockHasAccess.err)
			}

			if tt.mockFindByID != nil {
				groupRepo.EXPECT().FindByID(gomock.Any(), tt.args.payload.ID).
					Times(1).
					Return(tt.mockFindByID.res, tt.mockFindByID.err)
			}

			if tt.mockUpdate != nil {
				groupRepo.EXPECT().Update(gomock.Any(), tt.mockFindByID.res).
					Times(1).
					DoAndReturn(func(ctx context.Context, group *model.Group) error {
						group.ID = tt.args.payload.ID
						group.Name = tt.args.payload.Name
						return tt.mockUpdate.err
					})
			}

			uc := NewGroupUsecase()
			err := uc.InjectAuthUsecase(authUsecase)
			utils.ContinueOrFatal(err)
			err = uc.InjectGroupRepo(groupRepo)
			utils.ContinueOrFatal(err)

			got, err := uc.Update(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("groupUsecase.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupUsecase.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupUsecase_DeleteByID(t *testing.T) {
	var (
		userID  = utils.GenerateUUID()
		groupID = utils.GenerateUUID()
	)
	type mockHasAccess struct {
		err error
	}
	type mockFindByID struct {
		res *model.Group
		err error
	}
	type mockDelete struct {
		err error
	}
	type args struct {
		userID  string
		payload *model.DeleteGroupByIDPayload
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
				payload: &model.DeleteGroupByIDPayload{
					ID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByID: &mockFindByID{
				res: &model.Group{
					ID:   groupID,
					Name: "group1",
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
				payload: &model.DeleteGroupByIDPayload{
					ID: groupID,
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
				payload: &model.DeleteGroupByIDPayload{
					ID: groupID,
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
			name: "error find group",
			args: args{
				userID: userID,
				payload: &model.DeleteGroupByIDPayload{
					ID: groupID,
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
			name: "error delete group",
			args: args{
				userID: userID,
				payload: &model.DeleteGroupByIDPayload{
					ID: groupID,
				},
			},
			mockHasAccess: &mockHasAccess{
				err: nil,
			},
			mockFindByID: &mockFindByID{
				res: &model.Group{
					ID:   groupID,
					Name: "group1",
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
			authUsecase := mock.NewMockAuthUsecase(ctrl)

			if tt.mockHasAccess != nil {
				authUsecase.EXPECT().HasAccess(gomock.Any(), gomock.Any()).Times(1).Return(tt.mockHasAccess.err)
			}

			if tt.mockFindByID != nil {
				groupRepo.EXPECT().FindByID(gomock.Any(), tt.args.payload.ID).
					Times(1).
					Return(tt.mockFindByID.res, tt.mockFindByID.err)
			}

			if tt.mockDelete != nil {
				groupRepo.EXPECT().DeleteByID(gomock.Any(), tt.args.payload.ID).
					Times(1).
					Return(tt.mockDelete.err)
			}

			uc := NewGroupUsecase()
			err := uc.InjectAuthUsecase(authUsecase)
			utils.ContinueOrFatal(err)
			err = uc.InjectGroupRepo(groupRepo)
			utils.ContinueOrFatal(err)
			if err := uc.DeleteByID(ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("groupUsecase.DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

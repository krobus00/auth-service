package usecase

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/krobus00/auth-service/internal/constant"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/model/mock"
	"github.com/krobus00/auth-service/internal/utils"
)

func Test_authUsecase_HasAccess(t *testing.T) {
	var (
		userID  = utils.GenerateUUID()
		groupID = utils.GenerateUUID()
	)
	type args struct {
		payload *model.HasAccessPayload
	}
	type mockFindByUserID struct {
		userGroups []*model.UserGroup
		err        error
	}
	type mockHasAccess struct {
		hasAccess bool
		err       error
	}
	tests := []struct {
		name             string
		args             args
		mockFindByUserID *mockFindByUserID
		mockHasAccess    *mockHasAccess
		wantErr          bool
	}{
		{
			name: "success",
			args: args{
				payload: &model.HasAccessPayload{
					UserID:      userID,
					Permissions: []string{"TEST_READ"},
				},
			},
			mockFindByUserID: &mockFindByUserID{
				userGroups: []*model.UserGroup{
					{
						UserID:  userID,
						GroupID: groupID,
					},
				},
				err: nil,
			},
			mockHasAccess: &mockHasAccess{
				hasAccess: true,
			},
			wantErr: false,
		},
		{
			name: "success system id",
			args: args{
				payload: &model.HasAccessPayload{
					UserID:      constant.SystemID,
					Permissions: []string{"TEST_READ"},
				},
			},
			mockFindByUserID: nil,
			mockHasAccess:    nil,
			wantErr:          false,
		},
		{
			name: "success allow guest",
			args: args{
				payload: &model.HasAccessPayload{
					UserID:      userID,
					Permissions: []string{constant.PermissionAllowGuest},
				},
			},
			mockFindByUserID: &mockFindByUserID{
				userGroups: []*model.UserGroup{
					{
						UserID:  userID,
						GroupID: groupID,
					},
				},
				err: nil,
			},
			mockHasAccess: nil,
			wantErr:       false,
		},
		{
			name: "error user groups not found",
			args: args{
				payload: &model.HasAccessPayload{
					UserID:      userID,
					Permissions: []string{"TEST_READ"},
				},
			},
			mockFindByUserID: &mockFindByUserID{
				userGroups: []*model.UserGroup{},
				err:        nil,
			},
			mockHasAccess: nil,
			wantErr:       true,
		},
		{
			name: "error when find user groups",
			args: args{
				payload: &model.HasAccessPayload{
					UserID:      userID,
					Permissions: []string{"TEST_READ"},
				},
			},
			mockFindByUserID: &mockFindByUserID{
				userGroups: nil,
				err:        errors.New("db error"),
			},
			mockHasAccess: nil,
			wantErr:       true,
		},
		{
			name: "error unauthorized access",
			args: args{
				payload: &model.HasAccessPayload{
					UserID:      userID,
					Permissions: []string{"TEST_READ"},
				},
			},
			mockFindByUserID: &mockFindByUserID{
				userGroups: []*model.UserGroup{
					{
						UserID:  userID,
						GroupID: groupID,
					},
				},
				err: nil,
			},
			mockHasAccess: &mockHasAccess{
				hasAccess: false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.TODO()
			ctx = context.WithValue(ctx, constant.KeyUserIDCtx, tt.args.payload.UserID)
			wg := sync.WaitGroup{}

			userGroupRepo := mock.NewMockUserGroupRepository(ctrl)
			if tt.mockFindByUserID != nil {
				userGroupRepo.EXPECT().FindByUserID(gomock.Any(), tt.args.payload.UserID).
					Times(1).
					Return(
						tt.mockFindByUserID.userGroups,
						tt.mockFindByUserID.err,
					)
			}

			if tt.mockHasAccess != nil {
				for _, permission := range tt.args.payload.Permissions {
					if permission != constant.PermissionAllowGuest {
						wg.Add(1)
						userGroupRepo.EXPECT().HasPermission(gomock.Any(), groupID, permission).
							Times(1).
							DoAndReturn(func(ctx context.Context, groupID string, permission string) (bool, error) {
								defer wg.Done()
								return tt.mockHasAccess.hasAccess, tt.mockHasAccess.err
							})
					}
				}
			}

			uc := NewAuthUsecase()
			err := uc.InjectUserGroupRepo(userGroupRepo)
			utils.ContinueOrFatal(err)

			if err := uc.HasAccess(ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("authUsecase.HasAccess() error = %v, wantErr %v", err, tt.wantErr)
			}
			wg.Wait()
		})
	}
}

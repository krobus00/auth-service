package repository

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/goccy/go-json"
	"github.com/krobus00/auth-service/internal/infrastructure"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func newUserGroupRepoMock(t *testing.T) (model.UserGroupRepository, sqlmock.Sqlmock, *miniredis.Miniredis) {
	dbConn, dbMock := utils.NewDBMock()
	miniRedis := miniredis.RunT(t)
	viper.Set("redis.cache_host", fmt.Sprintf("redis://%s", miniRedis.Addr()))
	redisClient, err := infrastructure.NewRedisClient()
	utils.ContinueOrFatal(err)
	userGroupRepo := NewUserGroupRepository()
	err = userGroupRepo.InjectDB(dbConn)
	utils.ContinueOrFatal(err)
	err = userGroupRepo.InjectRedisClient(redisClient)
	utils.ContinueOrFatal(err)

	return userGroupRepo, dbMock, miniRedis
}

func Test_userGroupRepository_Create(t *testing.T) {
	type args struct {
		data *model.UserGroup
	}
	tests := []struct {
		name    string
		args    args
		mockErr error
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				data: &model.UserGroup{
					UserID:  utils.GenerateUUID(),
					GroupID: utils.GenerateUUID(),
				},
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "db error",
			args: args{
				data: &model.UserGroup{
					UserID:  utils.GenerateUUID(),
					GroupID: utils.GenerateUUID(),
				},
			},
			mockErr: errors.New("db error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, _ := newUserGroupRepoMock(t)
			dbMock.ExpectBegin()
			dbMock.ExpectExec("INSERT INTO \"user_groups\"").
				WithArgs(
					tt.args.data.UserID,
					tt.args.data.GroupID,
				).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tt.mockErr)

			if tt.wantErr {
				dbMock.ExpectRollback()
			} else {
				dbMock.ExpectCommit()
			}
			if err := r.Create(context.TODO(), tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("userGroupRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userGroupRepository_FindByUserIDAndGroupID(t *testing.T) {
	var (
		userID  = utils.GenerateUUID()
		groupID = utils.GenerateUUID()
	)
	type args struct {
		userID  string
		groupID string
	}
	type mockSelect struct {
		userGroup *model.UserGroup
		err       error
	}
	type mockCache struct {
		userGroup *model.UserGroup
	}
	tests := []struct {
		name       string
		args       args
		mockSelect *mockSelect
		mockCache  *mockCache
		want       *model.UserGroup
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				userID:  userID,
				groupID: groupID,
			},
			mockSelect: &mockSelect{
				userGroup: &model.UserGroup{
					UserID:  userID,
					GroupID: groupID,
				},
				err: nil,
			},
			mockCache: nil,
			want: &model.UserGroup{
				UserID:  userID,
				GroupID: groupID,
			},
			wantErr: false,
		},
		{
			name: "success found in cache",
			args: args{
				userID:  userID,
				groupID: groupID,
			},
			mockSelect: nil,
			mockCache: &mockCache{
				userGroup: &model.UserGroup{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			want: &model.UserGroup{
				UserID:  userID,
				GroupID: groupID,
			},
			wantErr: false,
		},
		{
			name: "error cache not found",
			args: args{
				userID:  userID,
				groupID: groupID,
			},
			mockSelect: nil,
			mockCache: &mockCache{
				userGroup: nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "error db not found",
			args: args{
				userID:  userID,
				groupID: groupID,
			},
			mockSelect: &mockSelect{
				userGroup: nil,
				err:       gorm.ErrRecordNotFound,
			},
			mockCache: nil,
			want:      nil,
			wantErr:   false,
		},
		{
			name: "error db not found",
			args: args{
				userID:  userID,
				groupID: groupID,
			},
			mockSelect: &mockSelect{
				userGroup: nil,
				err:       errors.New("db error"),
			},
			mockCache: nil,
			want:      nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, redisMock := newUserGroupRepoMock(t)
			cacheKey := model.NewUserGroupCacheKeyByUserIDAndGroupID(tt.args.userID, tt.args.groupID)
			if tt.mockSelect != nil {
				row := sqlmock.NewRows([]string{"user_id", "group_id"})
				if tt.mockSelect.userGroup != nil {
					userGroup := tt.mockSelect.userGroup
					row.AddRow(userGroup.UserID, userGroup.GroupID)
				}

				dbMock.ExpectQuery("^SELECT .+ FROM \"user_groups\"").
					WithArgs(tt.args.userID, tt.args.groupID).
					WillReturnRows(row).
					WillReturnError(tt.mockSelect.err)
			}
			if tt.mockCache != nil {
				cacheData, err := json.Marshal(tt.mockCache.userGroup)
				if err != nil {
					utils.ContinueOrFatal(err)
				}
				_ = redisMock.Set(cacheKey, string(cacheData))
			}

			got, err := r.FindByUserIDAndGroupID(context.TODO(), tt.args.userID, tt.args.groupID)
			if (err != nil) != tt.wantErr {
				t.Errorf("userGroupRepository.FindByUserIDAndGroupID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userGroupRepository.FindByUserIDAndGroupID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userGroupRepository_FindByUserID(t *testing.T) {
	var (
		userID  = utils.GenerateUUID()
		groupID = utils.GenerateUUID()
	)
	type args struct {
		userID string
	}
	type mockSelect struct {
		userGroup *model.UserGroup
		err       error
	}
	type mockCache struct {
		userGroup []*model.UserGroup
	}
	tests := []struct {
		name       string
		args       args
		mockSelect *mockSelect
		mockCache  *mockCache
		want       []*model.UserGroup
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				userID: userID,
			},
			mockSelect: &mockSelect{
				userGroup: &model.UserGroup{
					UserID:  userID,
					GroupID: groupID,
				},
				err: nil,
			},
			mockCache: nil,
			want: []*model.UserGroup{
				{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			wantErr: false,
		},
		{
			name: "success found in cache",
			args: args{
				userID: userID,
			},
			mockSelect: nil,
			mockCache: &mockCache{
				userGroup: []*model.UserGroup{
					{
						UserID:  userID,
						GroupID: groupID,
					},
				},
			},
			want: []*model.UserGroup{
				{
					UserID:  userID,
					GroupID: groupID,
				},
			},
			wantErr: false,
		},
		{
			name: "error cache not found",
			args: args{
				userID: userID,
			},
			mockSelect: nil,
			mockCache: &mockCache{
				userGroup: make([]*model.UserGroup, 0),
			},
			want:    make([]*model.UserGroup, 0),
			wantErr: false,
		},
		{
			name: "error db not found",
			args: args{
				userID: userID,
			},
			mockSelect: &mockSelect{
				userGroup: nil,
				err:       gorm.ErrRecordNotFound,
			},
			mockCache: nil,
			want:      make([]*model.UserGroup, 0),
			wantErr:   false,
		},
		{
			name: "error db not found",
			args: args{
				userID: userID,
			},
			mockSelect: &mockSelect{
				userGroup: nil,
				err:       errors.New("db error"),
			},
			mockCache: nil,
			want:      make([]*model.UserGroup, 0),
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, redisMock := newUserGroupRepoMock(t)
			cacheKey := model.NewUserGroupCacheKeyByUserID(tt.args.userID)
			if tt.mockSelect != nil {
				row := sqlmock.NewRows([]string{"user_id", "group_id"})
				if tt.mockSelect.userGroup != nil {
					userGroup := tt.mockSelect.userGroup
					row.AddRow(userGroup.UserID, userGroup.GroupID)
				}

				dbMock.ExpectQuery("^SELECT .+ FROM \"user_groups\"").
					WithArgs(tt.args.userID).
					WillReturnRows(row).
					WillReturnError(tt.mockSelect.err)
			}
			if tt.mockCache != nil {
				cacheData, err := json.Marshal(tt.mockCache.userGroup)
				if err != nil {
					utils.ContinueOrFatal(err)
				}
				_ = redisMock.Set(cacheKey, string(cacheData))
			}

			got, err := r.FindByUserID(context.TODO(), tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("userGroupRepository.FindByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, got, tt.want) {
				t.Errorf("userGroupRepository.FindByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userGroupRepository_DeleteByUserIDAndGroupID(t *testing.T) {
	type args struct {
		userID  string
		groupID string
	}
	tests := []struct {
		name    string
		args    args
		mockErr error
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				userID:  utils.GenerateUUID(),
				groupID: utils.GenerateUUID(),
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "db error",
			args: args{
				userID:  utils.GenerateUUID(),
				groupID: utils.GenerateUUID(),
			},
			mockErr: errors.New("db error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, _ := newUserGroupRepoMock(t)

			row := sqlmock.NewRows([]string{"user_id", "group_id"})

			row.AddRow(tt.args.userID, tt.args.groupID)
			dbMock.ExpectBegin()
			dbMock.ExpectQuery("DELETE FROM \"user_groups\"").
				WithArgs(tt.args.userID, tt.args.groupID).
				WillReturnRows(row).
				WillReturnError(tt.mockErr)

			if tt.wantErr {
				dbMock.ExpectRollback()
			} else {
				dbMock.ExpectCommit()
			}
			if err := r.DeleteByUserIDAndGroupID(context.TODO(), tt.args.userID, tt.args.groupID); (err != nil) != tt.wantErr {
				t.Errorf("userGroupRepository.DeleteByUserIDAndGroupID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userGroupRepository_HasPermission(t *testing.T) {
	var (
		groupID    = utils.GenerateUUID()
		permission = "FULL_ACCESS"
	)
	type args struct {
		groupID    string
		permission string
	}
	type mockSelect struct {
		groupPermissionAccess *model.GroupPermissionAccess
		err                   error
	}

	tests := []struct {
		name       string
		args       args
		mockSelect *mockSelect
		want       bool
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				groupID:    groupID,
				permission: permission,
			},
			mockSelect: &mockSelect{
				groupPermissionAccess: &model.GroupPermissionAccess{
					GroupID:        groupID,
					PermissionName: permission,
				},
				err: nil,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "error not found",
			args: args{
				groupID:    groupID,
				permission: permission,
			},
			mockSelect: &mockSelect{
				groupPermissionAccess: nil,
				err:                   gorm.ErrRecordNotFound,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "db error",
			args: args{
				groupID:    groupID,
				permission: permission,
			},
			mockSelect: &mockSelect{
				groupPermissionAccess: nil,
				err:                   errors.New("db error"),
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, _ := newUserGroupRepoMock(t)
			if tt.mockSelect != nil {
				row := sqlmock.NewRows([]string{"group_id", "permission_name"})
				if tt.mockSelect.groupPermissionAccess != nil {
					groupPermissionAccess := tt.mockSelect.groupPermissionAccess
					row.AddRow(groupPermissionAccess.GroupID, groupPermissionAccess.PermissionName)
				}

				dbMock.ExpectQuery("SELECT .+ FROM user_groups .+ JOIN group_permissions .+ JOIN permissions").
					WithArgs(tt.args.groupID, tt.args.permission).
					WillReturnRows(row).
					WillReturnError(tt.mockSelect.err)
			}

			got, err := r.HasPermission(context.TODO(), tt.args.groupID, tt.args.permission)
			if (err != nil) != tt.wantErr {
				t.Errorf("userGroupRepository.HasPermission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userGroupRepository.HasPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}

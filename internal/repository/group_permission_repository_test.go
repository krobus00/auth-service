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
)

func newGroupPermissionRepoMock(t *testing.T) (model.GroupPermissionRepository, sqlmock.Sqlmock, *miniredis.Miniredis) {
	dbConn, dbMock := utils.NewDBMock()
	miniRedis := miniredis.RunT(t)
	viper.Set("redis.cache_host", fmt.Sprintf("redis://%s", miniRedis.Addr()))
	redisClient, err := infrastructure.NewRedisClient()
	utils.ContinueOrFatal(err)
	groupPermissionRepo := NewGroupPermissionRepository()
	err = groupPermissionRepo.InjectDB(dbConn)
	utils.ContinueOrFatal(err)
	err = groupPermissionRepo.InjectRedisClient(redisClient)
	utils.ContinueOrFatal(err)

	return groupPermissionRepo, dbMock, miniRedis
}

func Test_groupPermissionRepo_Create(t *testing.T) {
	var (
		groupID      = utils.GenerateUUID()
		permissionID = utils.GenerateUUID()
	)
	type args struct {
		data *model.GroupPermission
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
				data: &model.GroupPermission{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "db error",
			args: args{
				data: &model.GroupPermission{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			mockErr: errors.New("db error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, _ := newGroupPermissionRepoMock(t)

			dbMock.ExpectBegin()
			dbMock.ExpectExec("INSERT INTO \"group_permissions\"").
				WithArgs(
					tt.args.data.GroupID,
					tt.args.data.PermissionID,
				).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tt.mockErr)

			if tt.wantErr {
				dbMock.ExpectRollback()
			} else {
				dbMock.ExpectCommit()
			}
			if err := r.Create(context.TODO(), tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("groupPermissionRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_groupPermissionRepo_FindByGroupIDAndPermissionID(t *testing.T) {
	var (
		groupID      = utils.GenerateUUID()
		permissionID = utils.GenerateUUID()
	)
	type args struct {
		groupID      string
		permissionID string
	}
	type mockSelect struct {
		groupPermission *model.GroupPermission
		err             error
	}
	type mockCache struct {
		groupPermission *model.GroupPermission
	}
	tests := []struct {
		name       string
		args       args
		mockSelect *mockSelect
		mockCache  *mockCache
		want       *model.GroupPermission
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				groupID:      groupID,
				permissionID: permissionID,
			},
			mockSelect: &mockSelect{
				groupPermission: &model.GroupPermission{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
				err: nil,
			},
			mockCache: nil,
			want: &model.GroupPermission{
				GroupID:      groupID,
				PermissionID: permissionID,
			},
			wantErr: false,
		},
		{
			name: "success found in cache",
			args: args{
				groupID:      groupID,
				permissionID: permissionID,
			},
			mockSelect: nil,
			mockCache: &mockCache{
				groupPermission: &model.GroupPermission{
					GroupID:      groupID,
					PermissionID: permissionID,
				},
			},
			want: &model.GroupPermission{
				GroupID:      groupID,
				PermissionID: permissionID,
			},
			wantErr: false,
		},
		{
			name: "error not found",
			args: args{
				groupID:      groupID,
				permissionID: permissionID,
			},
			mockSelect: &mockSelect{
				groupPermission: nil,
				err:             nil,
			},
			mockCache: nil,
			want:      nil,
			wantErr:   false,
		},
		{
			name: "db error",
			args: args{
				groupID:      groupID,
				permissionID: permissionID,
			},
			mockSelect: &mockSelect{
				groupPermission: nil,
				err:             errors.New("db error"),
			},
			mockCache: nil,
			want:      nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, redisMock := newGroupPermissionRepoMock(t)
			cacheKey := model.NewGroupPermissionCacheKeyByGroupIDAndPermissionID(tt.args.groupID, tt.args.permissionID)
			if tt.mockSelect != nil {
				row := sqlmock.NewRows([]string{"group_id", "permission_id"})
				if tt.mockSelect.groupPermission != nil {
					groupPermission := tt.mockSelect.groupPermission
					row.AddRow(groupPermission.GroupID, groupPermission.PermissionID)
				}

				dbMock.ExpectQuery("^SELECT .+ FROM \"group_permissions\"").
					WithArgs(tt.args.groupID, tt.args.permissionID).
					WillReturnRows(row).
					WillReturnError(tt.mockSelect.err)
			}
			if tt.mockCache != nil {
				cacheData, err := json.Marshal(tt.mockCache.groupPermission)
				if err != nil {
					utils.ContinueOrFatal(err)
				}
				_ = redisMock.Set(cacheKey, string(cacheData))
			}
			got, err := r.FindByGroupIDAndPermissionID(context.TODO(), tt.args.groupID, tt.args.permissionID)
			if (err != nil) != tt.wantErr {
				t.Errorf("groupPermissionRepo.FindByGroupIDAndPermissionID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupPermissionRepo.FindByGroupIDAndPermissionID() = %v, want %v", got, tt.want)
			}
			if tt.mockCache != nil {
				cachedData, _ := redisMock.Get(cacheKey)
				if cachedData == "" {
					t.Errorf("permissionRepository.FindByID() cache not found")
				}
			}
		})
	}
}

func Test_groupPermissionRepo_DeleteByGroupIDAndPermissionID(t *testing.T) {
	type args struct {
		groupID      string
		permissionID string
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
				groupID:      utils.GenerateUUID(),
				permissionID: utils.GenerateUUID(),
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "db error",
			args: args{
				groupID:      utils.GenerateUUID(),
				permissionID: utils.GenerateUUID(),
			},
			mockErr: errors.New("db error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, _ := newGroupPermissionRepoMock(t)
			dbMock.ExpectBegin()
			row := sqlmock.NewRows([]string{"group_id", "permission_id"})

			row.AddRow(tt.args.groupID, tt.args.permissionID)

			dbMock.ExpectQuery("DELETE FROM \"group_permissions\"").
				WithArgs(tt.args.groupID, tt.args.permissionID).
				WillReturnRows(row).
				WillReturnError(tt.mockErr)

			if tt.wantErr {
				dbMock.ExpectRollback()
			} else {
				dbMock.ExpectCommit()
			}
			if err := r.DeleteByGroupIDAndPermissionID(context.TODO(), tt.args.groupID, tt.args.permissionID); (err != nil) != tt.wantErr {
				t.Errorf("groupPermissionRepo.DeleteByGroupIDAndPermissionID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

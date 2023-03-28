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
	"gorm.io/gorm"
)

func newPermissionRepoMock(t *testing.T) (model.PermissionRepository, sqlmock.Sqlmock, *miniredis.Miniredis) {
	dbConn, dbMock := utils.NewDBMock()
	miniRedis := miniredis.RunT(t)
	viper.Set("redis.cache_host", fmt.Sprintf("redis://%s", miniRedis.Addr()))
	redisClient, err := infrastructure.NewRedisClient()
	utils.ContinueOrFatal(err)
	permissionRepo := NewPermissionRepository()
	err = permissionRepo.InjectDB(dbConn)
	utils.ContinueOrFatal(err)
	err = permissionRepo.InjectRedisClient(redisClient)
	utils.ContinueOrFatal(err)

	return permissionRepo, dbMock, miniRedis
}

func Test_permissionRepository_Create(t *testing.T) {
	type args struct {
		permission *model.Permission
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
				permission: &model.Permission{
					ID:   utils.GenerateUUID(),
					Name: "new permission",
				},
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "db error",
			args: args{
				permission: &model.Permission{
					ID:   utils.GenerateUUID(),
					Name: "new permission",
				},
			},
			mockErr: errors.New("db error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, _ := newPermissionRepoMock(t)
			dbMock.ExpectBegin()
			dbMock.ExpectExec("INSERT INTO \"permissions\"").
				WithArgs(
					tt.args.permission.ID,
					tt.args.permission.Name,
				).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tt.mockErr)

			if tt.wantErr {
				dbMock.ExpectRollback()
			} else {
				dbMock.ExpectCommit()
			}
			if err := r.Create(context.TODO(), tt.args.permission); (err != nil) != tt.wantErr {
				t.Errorf("permissionRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_permissionRepository_FindByID(t *testing.T) {
	var (
		permissionID = utils.GenerateUUID()
	)
	type args struct {
		id string
	}
	type mockSelect struct {
		permission *model.Permission
		err        error
	}
	type mockCache struct {
		permission *model.Permission
	}
	tests := []struct {
		name       string
		args       args
		mockSelect *mockSelect
		mockCache  *mockCache
		want       *model.Permission
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				id: permissionID,
			},
			mockSelect: &mockSelect{
				permission: &model.Permission{
					ID:   permissionID,
					Name: "permission1",
				},
				err: nil,
			},
			mockCache: nil,
			want: &model.Permission{
				ID:   permissionID,
				Name: "permission1",
			},
			wantErr: false,
		},
		{
			name: "success found in cache",
			args: args{
				id: permissionID,
			},
			mockSelect: nil,
			mockCache: &mockCache{
				permission: &model.Permission{
					ID:   permissionID,
					Name: "permission1",
				},
			},
			want: &model.Permission{
				ID:   permissionID,
				Name: "permission1",
			},
			wantErr: false,
		},
		{
			name: "error not found",
			args: args{
				id: permissionID,
			},
			mockSelect: &mockSelect{
				permission: nil,
				err:        gorm.ErrRecordNotFound,
			},
			mockCache: nil,
			want:      nil,
			wantErr:   false,
		},
		{
			name: "db error",
			args: args{
				id: permissionID,
			},
			mockSelect: &mockSelect{
				permission: nil,
				err:        errors.New("db error"),
			},
			mockCache: nil,
			want:      nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, redisMock := newPermissionRepoMock(t)
			cacheKey := model.NewPermissionCacheKeyByID(tt.args.id)
			if tt.mockSelect != nil {
				row := sqlmock.NewRows([]string{"id", "name"})
				if tt.mockSelect.permission != nil {
					permission := tt.mockSelect.permission
					row.AddRow(permission.ID, permission.Name)
				}

				dbMock.ExpectQuery("^SELECT .+ FROM \"permissions\"").
					WithArgs(tt.args.id).
					WillReturnRows(row).
					WillReturnError(tt.mockSelect.err)
			}
			if tt.mockCache != nil {
				cacheData, err := json.Marshal(tt.mockCache.permission)
				if err != nil {
					utils.ContinueOrFatal(err)
				}
				_ = redisMock.Set(cacheKey, string(cacheData))
			}
			got, err := r.FindByID(context.TODO(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("permissionRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permissionRepository.FindByID() = %v, want %v", got, tt.want)
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

func Test_permissionRepository_FindByName(t *testing.T) {
	var (
		permissionID = utils.GenerateUUID()
	)
	type args struct {
		name string
	}
	type mockSelect struct {
		permission *model.Permission
		err        error
	}
	type mockCache struct {
		permission *model.Permission
	}
	tests := []struct {
		name       string
		args       args
		mockSelect *mockSelect
		mockCache  *mockCache
		want       *model.Permission
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				name: "permission-name",
			},
			mockSelect: &mockSelect{
				permission: &model.Permission{
					ID:   permissionID,
					Name: "permission-name",
				},
				err: nil,
			},
			mockCache: nil,
			want: &model.Permission{
				ID:   permissionID,
				Name: "permission-name",
			},
			wantErr: false,
		},
		{
			name: "success found in cache",
			args: args{
				name: "permission-name",
			},
			mockSelect: nil,
			mockCache: &mockCache{
				permission: &model.Permission{
					ID:   permissionID,
					Name: "permission-name",
				},
			},
			want: &model.Permission{
				ID:   permissionID,
				Name: "permission-name",
			},
			wantErr: false,
		},
		{
			name: "error not found",
			args: args{
				name: "permission-name-1",
			},
			mockSelect: &mockSelect{
				permission: nil,
				err:        gorm.ErrRecordNotFound,
			},
			mockCache: nil,
			want:      nil,
			wantErr:   false,
		},
		{
			name: "db error",
			args: args{
				name: "groun-name",
			},
			mockSelect: &mockSelect{
				permission: nil,
				err:        errors.New("db error"),
			},
			mockCache: nil,
			want:      nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, redisMock := newPermissionRepoMock(t)
			cacheKey := model.NewPermissionCacheKeyByName(tt.args.name)
			if tt.mockSelect != nil {
				row := sqlmock.NewRows([]string{"id", "name"})
				if tt.mockSelect.permission != nil {
					permission := tt.mockSelect.permission
					row.AddRow(permission.ID, permission.Name)
				}

				dbMock.ExpectQuery("^SELECT .+ FROM \"permissions\"").
					WithArgs(tt.args.name).
					WillReturnRows(row).
					WillReturnError(tt.mockSelect.err)
			}
			if tt.mockCache != nil {
				cacheData, err := json.Marshal(tt.mockCache.permission)
				if err != nil {
					utils.ContinueOrFatal(err)
				}
				_ = redisMock.Set(cacheKey, string(cacheData))
			}
			got, err := r.FindByName(context.TODO(), tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("permissionRepository.FindByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permissionRepository.FindByName() = %v, want %v", got, tt.want)
			}
			if tt.mockCache != nil {
				cachedData, _ := redisMock.Get(cacheKey)
				if cachedData == "" {
					t.Errorf("permissionRepository.FindByName() cache not found")
				}
			}
		})
	}
}

func Test_permissionRepository_Update(t *testing.T) {
	type args struct {
		permission *model.Permission
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
				permission: &model.Permission{
					ID:   utils.GenerateUUID(),
					Name: "update permission name",
				},
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "db error",
			args: args{
				permission: &model.Permission{
					ID:   utils.GenerateUUID(),
					Name: "update permission name",
				},
			},
			mockErr: errors.New("db error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, _ := newPermissionRepoMock(t)

			dbMock.ExpectBegin()
			dbMock.ExpectExec("UPDATE \"permissions\"").
				WithArgs(
					tt.args.permission.Name,
					tt.args.permission.ID,
				).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tt.mockErr)
			if tt.wantErr {
				dbMock.ExpectRollback()
			} else {
				dbMock.ExpectCommit()
			}
			if err := r.Update(context.TODO(), tt.args.permission); (err != nil) != tt.wantErr {
				t.Errorf("permissionRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_permissionRepository_DeleteByID(t *testing.T) {
	type args struct {
		id string
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
				id: utils.GenerateUUID(),
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "db error",
			args: args{
				id: utils.GenerateUUID(),
			},
			mockErr: errors.New("db error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, _ := newPermissionRepoMock(t)

			dbMock.ExpectBegin()
			row := sqlmock.NewRows([]string{"id", "name"})

			row.AddRow(tt.args.id, "permission-name")

			dbMock.ExpectQuery("DELETE FROM \"permissions\"").
				WithArgs(tt.args.id).
				WillReturnRows(row).
				WillReturnError(tt.mockErr)

			if tt.wantErr {
				dbMock.ExpectRollback()
			} else {
				dbMock.ExpectCommit()
			}
			if err := r.DeleteByID(context.TODO(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("permissionRepository.DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

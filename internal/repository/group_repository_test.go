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

func newGroupRepoMock(t *testing.T) (model.GroupRepository, sqlmock.Sqlmock, *miniredis.Miniredis) {
	dbConn, dbMock := utils.NewDBMock()
	miniRedis := miniredis.RunT(t)
	viper.Set("redis.cache_host", fmt.Sprintf("redis://%s", miniRedis.Addr()))
	redisClient, err := infrastructure.NewRedisClient()
	utils.ContinueOrFatal(err)
	groupRepo := NewGroupRepository()
	err = groupRepo.InjectDB(dbConn)
	utils.ContinueOrFatal(err)
	err = groupRepo.InjectRedisClient(redisClient)
	utils.ContinueOrFatal(err)

	return groupRepo, dbMock, miniRedis
}

func Test_groupRepository_Create(t *testing.T) {
	type args struct {
		group *model.Group
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
				group: &model.Group{
					ID:   utils.GenerateUUID(),
					Name: "new group",
				},
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "db error",
			args: args{
				group: &model.Group{
					ID:   utils.GenerateUUID(),
					Name: "new group",
				},
			},
			mockErr: errors.New("db error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, _ := newGroupRepoMock(t)
			dbMock.ExpectBegin()
			dbMock.ExpectExec("INSERT INTO \"groups\"").
				WithArgs(
					tt.args.group.ID,
					tt.args.group.Name,
				).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tt.mockErr)

			if tt.wantErr {
				dbMock.ExpectRollback()
			} else {
				dbMock.ExpectCommit()
			}
			if err := r.Create(context.TODO(), tt.args.group); (err != nil) != tt.wantErr {
				t.Errorf("groupRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_groupRepository_FindByID(t *testing.T) {
	var (
		groupID = utils.GenerateUUID()
	)
	type args struct {
		id string
	}
	type mockSelect struct {
		group *model.Group
		err   error
	}
	type mockCache struct {
		group *model.Group
	}
	tests := []struct {
		name       string
		args       args
		mockSelect *mockSelect
		mockCache  *mockCache
		want       *model.Group
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				id: groupID,
			},
			mockSelect: &mockSelect{
				group: &model.Group{
					ID:   groupID,
					Name: "group1",
				},
				err: nil,
			},
			mockCache: nil,
			want: &model.Group{
				ID:   groupID,
				Name: "group1",
			},
			wantErr: false,
		},
		{
			name: "success found in cache",
			args: args{
				id: groupID,
			},
			mockSelect: nil,
			mockCache: &mockCache{
				group: &model.Group{
					ID:   groupID,
					Name: "group1",
				},
			},
			want: &model.Group{
				ID:   groupID,
				Name: "group1",
			},
			wantErr: false,
		},
		{
			name: "error not found",
			args: args{
				id: groupID,
			},
			mockSelect: &mockSelect{
				group: nil,
				err:   gorm.ErrRecordNotFound,
			},
			mockCache: nil,
			want:      nil,
			wantErr:   false,
		},
		{
			name: "db error",
			args: args{
				id: groupID,
			},
			mockSelect: &mockSelect{
				group: nil,
				err:   errors.New("db error"),
			},
			mockCache: nil,
			want:      nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, redisMock := newGroupRepoMock(t)
			cacheKey := model.NewGroupCacheKeyByID(tt.args.id)
			if tt.mockSelect != nil {
				row := sqlmock.NewRows([]string{"id", "name"})
				if tt.mockSelect.group != nil {
					group := tt.mockSelect.group
					row.AddRow(group.ID, group.Name)
				}

				dbMock.ExpectQuery("^SELECT .+ FROM \"groups\"").
					WithArgs(tt.args.id).
					WillReturnRows(row).
					WillReturnError(tt.mockSelect.err)
			}
			if tt.mockCache != nil {
				cacheData, err := json.Marshal(tt.mockCache.group)
				if err != nil {
					utils.ContinueOrFatal(err)
				}
				_ = redisMock.Set(cacheKey, string(cacheData))
			}
			got, err := r.FindByID(context.TODO(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("groupRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupRepository.FindByID() = %v, want %v", got, tt.want)
			}
			if tt.mockCache != nil {
				cachedData, _ := redisMock.Get(cacheKey)
				if cachedData == "" {
					t.Errorf("groupRepository.FindByID() cache not found")
				}
			}
		})
	}
}

func Test_groupRepository_FindByName(t *testing.T) {
	var (
		groupID = utils.GenerateUUID()
	)
	type args struct {
		name string
	}
	type mockSelect struct {
		group *model.Group
		err   error
	}
	type mockCache struct {
		group *model.Group
	}
	tests := []struct {
		name       string
		args       args
		mockSelect *mockSelect
		mockCache  *mockCache
		want       *model.Group
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				name: "group-name",
			},
			mockSelect: &mockSelect{
				group: &model.Group{
					ID:   groupID,
					Name: "group-name",
				},
				err: nil,
			},
			mockCache: nil,
			want: &model.Group{
				ID:   groupID,
				Name: "group-name",
			},
			wantErr: false,
		},
		{
			name: "success found in cache",
			args: args{
				name: "group-name",
			},
			mockSelect: nil,
			mockCache: &mockCache{
				group: &model.Group{
					ID:   groupID,
					Name: "group-name",
				},
			},
			want: &model.Group{
				ID:   groupID,
				Name: "group-name",
			},
			wantErr: false,
		},
		{
			name: "error not found",
			args: args{
				name: "group-name-1",
			},
			mockSelect: &mockSelect{
				group: nil,
				err:   gorm.ErrRecordNotFound,
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
				group: nil,
				err:   errors.New("db error"),
			},
			mockCache: nil,
			want:      nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, redisMock := newGroupRepoMock(t)
			cacheKey := model.NewGroupCacheKeyByName(tt.args.name)
			if tt.mockSelect != nil {
				row := sqlmock.NewRows([]string{"id", "name"})
				if tt.mockSelect.group != nil {
					group := tt.mockSelect.group
					row.AddRow(group.ID, group.Name)
				}

				dbMock.ExpectQuery("^SELECT .+ FROM \"groups\"").
					WithArgs(tt.args.name).
					WillReturnRows(row).
					WillReturnError(tt.mockSelect.err)
			}
			if tt.mockCache != nil {
				cacheData, err := json.Marshal(tt.mockCache.group)
				if err != nil {
					utils.ContinueOrFatal(err)
				}
				_ = redisMock.Set(cacheKey, string(cacheData))
			}
			got, err := r.FindByName(context.TODO(), tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("groupRepository.FindByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupRepository.FindByName() = %v, want %v", got, tt.want)
			}
			if tt.mockCache != nil {
				cachedData, _ := redisMock.Get(cacheKey)
				if cachedData == "" {
					t.Errorf("groupRepository.FindByName() cache not found")
				}
			}
		})
	}
}

func Test_groupRepository_Update(t *testing.T) {
	type args struct {
		group *model.Group
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
				group: &model.Group{
					ID:   utils.GenerateUUID(),
					Name: "update group name",
				},
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "db error",
			args: args{
				group: &model.Group{
					ID:   utils.GenerateUUID(),
					Name: "update group name",
				},
			},
			mockErr: errors.New("db error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, _ := newGroupRepoMock(t)

			dbMock.ExpectBegin()
			dbMock.ExpectExec("UPDATE \"groups\"").
				WithArgs(
					tt.args.group.Name,
					tt.args.group.ID,
				).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tt.mockErr)
			if tt.wantErr {
				dbMock.ExpectRollback()
			} else {
				dbMock.ExpectCommit()
			}
			if err := r.Update(context.TODO(), tt.args.group); (err != nil) != tt.wantErr {
				t.Errorf("groupRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_groupRepository_DeleteByID(t *testing.T) {

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
			r, dbMock, _ := newGroupRepoMock(t)

			dbMock.ExpectBegin()
			row := sqlmock.NewRows([]string{"id", "name"})

			row.AddRow(tt.args.id, "group-name")

			dbMock.ExpectQuery("DELETE FROM \"groups\"").
				WithArgs(tt.args.id).
				WillReturnRows(row).
				WillReturnError(tt.mockErr)

			if tt.wantErr {
				dbMock.ExpectRollback()
			} else {
				dbMock.ExpectCommit()
			}
			if err := r.DeleteByID(context.TODO(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("groupRepository.DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

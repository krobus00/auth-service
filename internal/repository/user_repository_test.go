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

func newUserRepoMock(t *testing.T) (model.UserRepository, sqlmock.Sqlmock, *miniredis.Miniredis) {
	dbConn, dbMock := utils.NewDBMock()
	miniRedis := miniredis.RunT(t)
	viper.Set("redis.cache_host", fmt.Sprintf("redis://%s", miniRedis.Addr()))
	redisClient, err := infrastructure.NewRedisClient()
	utils.ContinueOrFatal(err)
	userRepo := NewUserRepository()
	err = userRepo.InjectDB(dbConn)
	utils.ContinueOrFatal(err)
	err = userRepo.InjectRedisClient(redisClient)
	utils.ContinueOrFatal(err)

	return userRepo, dbMock, miniRedis
}

func Test_userRepository_Create(t *testing.T) {
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		args    args
		mockErr error
		wantErr bool
	}{
		{
			name: "success create new user",
			args: args{
				user: &model.User{
					ID:       utils.GenerateUUID(),
					FullName: "full name",
					Username: "username",
					Email:    "user@gmail.com",
					Password: "hashed-password",
				},
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "db error",
			args: args{
				user: &model.User{
					ID:       utils.GenerateUUID(),
					FullName: "full name",
					Username: "username",
					Email:    "user@gmail.com",
					Password: "hashed-password",
				},
			},
			mockErr: errors.New("db error"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, _ := newUserRepoMock(t)
			dbMock.ExpectBegin()
			dbMock.ExpectExec("INSERT INTO \"users\"").
				WithArgs(
					tt.args.user.ID,
					tt.args.user.FullName,
					tt.args.user.Username,
					tt.args.user.Email,
					tt.args.user.Password,
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
				).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tt.mockErr)

			if tt.wantErr {
				dbMock.ExpectRollback()
			} else {
				dbMock.ExpectCommit()
			}
			if err := r.Create(context.TODO(), tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userRepository_FindByID(t *testing.T) {
	var (
		userID = utils.GenerateUUID()
	)
	type args struct {
		id string
	}

	type mockSelect struct {
		user *model.User
		err  error
	}
	type mockCache struct {
		user *model.User
	}
	tests := []struct {
		name       string
		args       args
		mockSelect *mockSelect
		mockCache  *mockCache
		want       *model.User
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				id: userID,
			},
			mockSelect: &mockSelect{
				user: &model.User{
					ID:       userID,
					FullName: "full name",
					Username: "username",
					Email:    "user@gmail.com",
					Password: "hashed-password",
				},
				err: nil,
			},
			want: &model.User{
				ID:       userID,
				FullName: "full name",
				Username: "username",
				Email:    "user@gmail.com",
				Password: "hashed-password",
			},
			wantErr: false,
		},
		{
			name: "success found in cache",
			args: args{
				id: userID,
			},
			mockCache: &mockCache{
				user: &model.User{
					ID:       userID,
					FullName: "full name",
					Username: "username",
					Email:    "user@gmail.com",
					Password: "hashed-password",
				},
			},
			want: &model.User{
				ID:       userID,
				FullName: "full name",
				Username: "username",
				Email:    "user@gmail.com",
				Password: "hashed-password",
			},
			wantErr: false,
		},
		{
			name: "user not found",
			args: args{
				id: userID,
			},
			mockSelect: &mockSelect{
				user: nil,
				err:  gorm.ErrRecordNotFound,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "db error",
			args: args{
				id: userID,
			},
			mockSelect: &mockSelect{
				user: nil,
				err:  errors.New("db error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, redisMock := newUserRepoMock(t)
			cacheKey := model.NewUserCacheKeyByID(tt.args.id)
			if tt.mockSelect != nil {
				row := sqlmock.NewRows([]string{"id", "full_name", "username", "email", "password", "created_at", "updated_at", "deleted_at"})
				if tt.mockSelect.user != nil {
					user := tt.mockSelect.user
					row.AddRow(user.ID, user.FullName, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
				}

				dbMock.ExpectQuery("^SELECT .+ FROM \"users\"").
					WithArgs(tt.args.id).
					WillReturnRows(row).
					WillReturnError(tt.mockSelect.err)
			}
			if tt.mockCache != nil {
				cacheData, err := json.Marshal(tt.mockCache.user)
				if err != nil {
					utils.ContinueOrFatal(err)
				}
				_ = redisMock.Set(cacheKey, string(cacheData))
			}
			got, err := r.FindByID(context.TODO(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.FindByID() = %v, want %v", got, tt.want)
			}
			if tt.mockCache != nil {
				cachedData, _ := redisMock.Get(cacheKey)
				if cachedData == "" {
					t.Errorf("userRepository.FindByID() cache not found")
				}
			}
		})
	}
}

func Test_userRepository_FindByUsername(t *testing.T) {
	var (
		userID   = utils.GenerateUUID()
		username = "username"
	)
	type args struct {
		username string
	}

	type mockSelect struct {
		user *model.User
		err  error
	}
	type mockCache struct {
		user *model.User
	}
	tests := []struct {
		name       string
		args       args
		mockSelect *mockSelect
		mockCache  *mockCache
		want       *model.User
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				username: username,
			},
			mockSelect: &mockSelect{
				user: &model.User{
					ID:       userID,
					FullName: "full name",
					Username: username,
					Email:    "user@gmail.com",
					Password: "hashed-password",
				},
				err: nil,
			},
			want: &model.User{
				ID:       userID,
				FullName: "full name",
				Username: username,
				Email:    "user@gmail.com",
				Password: "hashed-password",
			},
			wantErr: false,
		},
		{
			name: "success found in cache",
			args: args{
				username: username,
			},
			mockCache: &mockCache{
				user: &model.User{
					ID:       userID,
					FullName: "full name",
					Username: username,
					Email:    "user@gmail.com",
					Password: "hashed-password",
				},
			},
			want: &model.User{
				ID:       userID,
				FullName: "full name",
				Username: username,
				Email:    "user@gmail.com",
				Password: "hashed-password",
			},
			wantErr: false,
		},
		{
			name: "success found nil in cache",
			args: args{
				username: username,
			},
			mockCache: &mockCache{
				user: nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "user not found",
			args: args{
				username: "uname",
			},
			mockSelect: &mockSelect{
				user: nil,
				err:  gorm.ErrRecordNotFound,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "db error",
			args: args{
				username: username,
			},
			mockSelect: &mockSelect{
				user: nil,
				err:  errors.New("db error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, redisMock := newUserRepoMock(t)
			cacheKey := model.NewUserCacheKeyByUsername(tt.args.username)
			if tt.mockSelect != nil {
				row := sqlmock.NewRows([]string{"id", "full_name", "username", "email", "password", "created_at", "updated_at", "deleted_at"})
				if tt.mockSelect.user != nil {
					user := tt.mockSelect.user
					row.AddRow(user.ID, user.FullName, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
				}

				dbMock.ExpectQuery("^SELECT .+ FROM \"users\"").
					WithArgs(tt.args.username).
					WillReturnRows(row).
					WillReturnError(tt.mockSelect.err)
			}
			if tt.mockCache != nil {
				cacheData, err := json.Marshal(tt.mockCache.user)
				if err != nil {
					utils.ContinueOrFatal(err)
				}
				_ = redisMock.Set(cacheKey, string(cacheData))
			}
			got, err := r.FindByUsername(context.TODO(), tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.FindByUsername() = %v, want %v", got, tt.want)
			}
			if tt.mockCache != nil {
				cachedData, _ := redisMock.Get(cacheKey)
				if cachedData == "" {
					t.Errorf("userRepository.FindByUsername() cache not found")
				}
			}
		})
	}
}

func Test_userRepository_FindByEmail(t *testing.T) {
	var (
		userID = utils.GenerateUUID()
		email  = "user@gmail.com"
	)
	type args struct {
		email string
	}

	type mockSelect struct {
		user *model.User
		err  error
	}
	type mockCache struct {
		user *model.User
	}
	tests := []struct {
		name       string
		args       args
		mockSelect *mockSelect
		mockCache  *mockCache
		want       *model.User
		wantErr    bool
	}{
		{
			name: "success",
			args: args{
				email: email,
			},
			mockSelect: &mockSelect{
				user: &model.User{
					ID:       userID,
					FullName: "full name",
					Username: "username",
					Email:    email,
					Password: "hashed-password",
				},
				err: nil,
			},
			want: &model.User{
				ID:       userID,
				FullName: "full name",
				Username: "username",
				Email:    email,
				Password: "hashed-password",
			},
			wantErr: false,
		},
		{
			name: "success found in cache",
			args: args{
				email: email,
			},
			mockCache: &mockCache{
				user: &model.User{
					ID:       userID,
					FullName: "full name",
					Username: "username",
					Email:    email,
					Password: "hashed-password",
				},
			},
			want: &model.User{
				ID:       userID,
				FullName: "full name",
				Username: "username",
				Email:    email,
				Password: "hashed-password",
			},
			wantErr: false,
		},
		{
			name: "user not found",
			args: args{
				email: "user1@gmail.com",
			},
			mockSelect: &mockSelect{
				user: nil,
				err:  gorm.ErrRecordNotFound,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "db error",
			args: args{
				email: email,
			},
			mockSelect: &mockSelect{
				user: nil,
				err:  errors.New("db error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, dbMock, redisMock := newUserRepoMock(t)
			cacheKey := model.NewUserCacheKeyByEmail(tt.args.email)
			if tt.mockSelect != nil {
				row := sqlmock.NewRows([]string{"id", "full_name", "username", "email", "password", "created_at", "updated_at", "deleted_at"})
				if tt.mockSelect.user != nil {
					user := tt.mockSelect.user
					row.AddRow(user.ID, user.FullName, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
				}

				dbMock.ExpectQuery("^SELECT .+ FROM \"users\"").
					WithArgs(tt.args.email).
					WillReturnRows(row).
					WillReturnError(tt.mockSelect.err)
			}
			if tt.mockCache != nil {
				cacheData, err := json.Marshal(tt.mockCache.user)
				if err != nil {
					utils.ContinueOrFatal(err)
				}
				_ = redisMock.Set(cacheKey, string(cacheData))
			}
			got, err := r.FindByEmail(context.TODO(), tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.FindByEmail() = %v, want %v", got, tt.want)
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

func Test_userRepository_UpdateByID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "unimplemented",
			args: args{
				id: "id",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _, _ := newUserRepoMock(t)
			got, err := r.UpdateByID(context.TODO(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.UpdateByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.UpdateByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_DeleteByID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "unimplemented",
			args: args{
				id: "id",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _, _ := newUserRepoMock(t)
			err := r.DeleteByID(context.TODO(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

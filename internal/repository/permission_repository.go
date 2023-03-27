package repository

import (
	"context"
	"encoding/json"
	"errors"

	goredis "github.com/go-redis/redis/v8"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type permissionRepository struct {
	db          *gorm.DB
	redisClient *goredis.Client
}

func NewPermissionRepository() model.PermissionRepository {
	return new(permissionRepository)
}

func (r *permissionRepository) Create(ctx context.Context, permission *model.Permission) error {
	logger := logrus.WithFields(logrus.Fields{
		"id":   permission.ID,
		"name": permission.Name,
	})

	db := utils.GetTxFromContext(ctx, r.db)

	err := db.WithContext(ctx).Create(permission).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_ = DeleteByKeys(ctx, r.redisClient, model.GetPermissionCacheKeys(permission.ID, permission.Name))

	return nil
}

func (r *permissionRepository) FindByID(ctx context.Context, id string) (*model.Permission, error) {
	logger := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	db := utils.GetTxFromContext(ctx, r.db)
	permission := new(model.Permission)
	cacheKey := model.NewPermissionCacheKeyByID(id)

	cachedData, err := Get(ctx, r.redisClient, cacheKey)
	if err != nil {
		logger.Error(err.Error())
	}
	err = json.Unmarshal(cachedData, &permission)
	if err == nil {
		return permission, nil
	}

	permission = new(model.Permission)

	err = db.WithContext(ctx).Where("id = ?", id).First(permission).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = SetWithExpiry(ctx, r.redisClient, cacheKey, nil)
			if err != nil {
				logger.Error(err.Error())
			}
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, err
	}

	err = SetWithExpiry(ctx, r.redisClient, cacheKey, permission)
	if err != nil {
		logger.Error(err.Error())
	}
	return permission, nil
}

func (r *permissionRepository) FindByName(ctx context.Context, name string) (*model.Permission, error) {
	logger := logrus.WithFields(logrus.Fields{
		"name": name,
	})

	db := utils.GetTxFromContext(ctx, r.db)
	permission := new(model.Permission)
	cacheKey := model.NewPermissionCacheKeyByName(name)

	cachedData, err := Get(ctx, r.redisClient, cacheKey)
	if err != nil {
		logger.Error(err.Error())
	}
	err = json.Unmarshal(cachedData, &permission)
	if err == nil {
		return permission, nil
	}

	permission = new(model.Permission)

	err = db.WithContext(ctx).Where("name = ?", name).First(permission).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = SetWithExpiry(ctx, r.redisClient, cacheKey, nil)
			if err != nil {
				logger.Error(err.Error())
			}
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, err
	}

	err = SetWithExpiry(ctx, r.redisClient, cacheKey, permission)
	if err != nil {
		logger.Error(err.Error())
	}
	return permission, nil
}

func (r *permissionRepository) Update(ctx context.Context, permission *model.Permission) error {
	logger := logrus.WithFields(logrus.Fields{
		"id": permission.ID,
	})

	db := utils.GetTxFromContext(ctx, r.db)

	_ = DeleteByKeys(ctx, r.redisClient, model.GetPermissionCacheKeys(permission.ID, permission.Name))

	err := db.WithContext(ctx).Updates(permission).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *permissionRepository) DeleteByID(ctx context.Context, id string) error {
	logger := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	db := utils.GetTxFromContext(ctx, r.db)
	permission := new(model.Permission)

	err := db.WithContext(ctx).Clauses(clause.Returning{}).
		Where("id = ?", id).
		Delete(permission).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_ = DeleteByKeys(ctx, r.redisClient, model.GetPermissionCacheKeys(permission.ID, permission.Name))

	return nil
}

package repository

import (
	"context"
	"errors"
	"fmt"

	goredis "github.com/go-redis/redis/v8"
	"github.com/krobus00/auth-service/internal/config"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userAccessControlRepository struct {
	db    *gorm.DB
	redis *goredis.Client
}

func NewUserAccessControlRepository() model.UserAccessControlRepository {
	return new(userAccessControlRepository)
}

func (r *userAccessControlRepository) Create(ctx context.Context, uac *model.UserAccessControl) error {
	logger := log.WithFields(log.Fields{
		"UserID": uac.UserID,
		"ACID":   uac.ACID,
	})

	db := utils.GetTxFromContext(ctx, r.db)
	err := db.WithContext(ctx).Create(uac).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
func (r *userAccessControlRepository) Revoke(ctx context.Context, uac *model.UserAccessControl) error {
	logger := log.WithFields(log.Fields{
		"UserID": uac.UserID,
		"ACID":   uac.ACID,
	})

	db := utils.GetTxFromContext(ctx, r.db)
	err := db.WithContext(ctx).
		Where("user_id = ? ", uac.UserID).
		Where("ac_id = ? ", uac.ACID).
		Delete(uac).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
func (r *userAccessControlRepository) HasAccess(ctx context.Context, userID string, accessList []string) error {
	logger := log.WithFields(log.Fields{
		"userID":     userID,
		"accessList": accessList,
	})

	result := new(model.AccessControl)
	cacheKey := model.UserAccessControlCacheKey(userID)

	hasAccess := r.hasAccessFromRedis(ctx, userID, accessList)
	if hasAccess {
		return nil
	}

	err := r.db.Table("access_controls ac").Select("id", "name", fmt.Sprintf(`CASE 
    WHEN  uac.user_id = '%s' THEN true
    	ELSE false
  	END AS has_access`, userID)).Joins("left join user_access_controls uac on ac.id = uac.ac_id").Where("ac.name in ?", accessList).Order("has_access DESC").First(result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ErrUnauthorizeAccess
		}
		logger.Error(err.Error())
		return err
	}

	err = r.redis.HSet(ctx, cacheKey, map[string]interface{}{result.Name: true}).Err()

	if err != nil {
		// ignore error
		logger.Error(err.Error())
	}
	err = r.redis.Expire(ctx, cacheKey, config.RedisCacheTTL()).Err()
	if err != nil {
		logger.Error(err.Error())
		return model.ErrUnauthorizeAccess
	}
	if !result.HasAccess {
		return model.ErrUnauthorizeAccess
	}
	return nil
}

func (r *userAccessControlRepository) hasAccessFromRedis(ctx context.Context, userID string, accessList []string) bool {
	cacheKey := model.UserAccessControlCacheKey(userID)
	results, _ := r.redis.HGetAll(ctx, cacheKey).Result()
	if len(results) > 0 {
		for _, v := range accessList {
			if val, ok := results[v]; ok && val == "true" {
				return true
			}
		}
	}
	return false
}

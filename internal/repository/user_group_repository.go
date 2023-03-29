package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/goccy/go-json"

	goredis "github.com/go-redis/redis/v8"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userGroupRepository struct {
	db          *gorm.DB
	redisClient *goredis.Client
}

func NewUserGroupRepository() model.UserGroupRepository {
	return new(userGroupRepository)
}

func (r *userGroupRepository) Create(ctx context.Context, data *model.UserGroup) error {
	logger := logrus.WithFields(logrus.Fields{
		"userID":  data.UserID,
		"groupID": data.GroupID,
	})

	db := utils.GetTxFromContext(ctx, r.db)

	err := db.WithContext(ctx).Create(data).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_ = DeleteByKeys(ctx, r.redisClient, model.GetUserGroupCacheKeys(data.UserID, data.GroupID))

	return nil
}

func (r *userGroupRepository) FindByUserIDAndGroupID(ctx context.Context, userID, groupID string) (*model.UserGroup, error) {
	logger := logrus.WithFields(logrus.Fields{
		"userID":  userID,
		"groupID": groupID,
	})

	db := utils.GetTxFromContext(ctx, r.db)
	userGroup := new(model.UserGroup)

	cacheKey := model.NewUserGroupCacheKeyByUserIDAndGroupID(userID, groupID)
	cachedData, err := Get(ctx, r.redisClient, cacheKey)
	if err != nil {
		logger.Error(err.Error())
	}
	err = json.Unmarshal(cachedData, &userGroup)
	if err == nil {
		return userGroup, nil
	}

	userGroup = new(model.UserGroup)

	err = db.WithContext(ctx).
		Where("user_id = ? AND group_id = ?", userID, groupID).
		First(userGroup).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, err
	}

	err = SetWithExpiry(ctx, r.redisClient, cacheKey, userGroup)
	if err != nil {
		logger.Error(err.Error())
	}

	return userGroup, nil
}

func (r *userGroupRepository) FindByUserID(ctx context.Context, userID string) ([]*model.UserGroup, error) {
	logger := logrus.WithFields(logrus.Fields{
		"userID": userID,
	})

	db := utils.GetTxFromContext(ctx, r.db)
	userGroups := make([]*model.UserGroup, 0)

	cacheKey := model.NewUserGroupCacheKeyByUserID(userID)
	cachedData, err := Get(ctx, r.redisClient, cacheKey)
	if err != nil {
		logger.Error(err.Error())
	}
	err = json.Unmarshal(cachedData, &userGroups)
	if err == nil {
		return userGroups, nil
	}

	userGroups = make([]*model.UserGroup, 0)

	err = db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&userGroups).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = SetWithExpiry(ctx, r.redisClient, cacheKey, userGroups)
			if err != nil {
				logger.Error(err.Error())
			}
			return userGroups, nil
		}
		logger.Error(err.Error())
		return userGroups, err
	}

	err = SetWithExpiry(ctx, r.redisClient, cacheKey, userGroups)
	if err != nil {
		logger.Error(err.Error())
	}
	return userGroups, nil
}

func (r *userGroupRepository) DeleteByUserIDAndGroupID(ctx context.Context, userID, groupID string) error {
	logger := logrus.WithFields(logrus.Fields{
		"userID":  userID,
		"groupID": groupID,
	})

	db := utils.GetTxFromContext(ctx, r.db)
	userGroup := new(model.UserGroup)

	err := db.WithContext(ctx).
		Clauses(clause.Returning{}).
		Where("user_id = ? AND group_id = ?", userID, groupID).
		Delete(userGroup).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_ = DeleteByKeys(ctx, r.redisClient, model.GetUserGroupCacheKeys(userID, groupID))

	return nil
}

func (r *userGroupRepository) HasPermission(ctx context.Context, groupID string, permission string) (bool, error) {
	db := utils.GetTxFromContext(ctx, r.db)
	cacheBucketKey := utils.NewBucketKey(model.NewGroupPermissionCacheKey(groupID), permission)
	found, hasAccess, _ := r.getHasPermissionCache(ctx, cacheBucketKey, permission)
	if found {
		return hasAccess, nil
	}

	result := new(model.GroupPermissionAccess)

	err := db.WithContext(ctx).
		Table("user_groups ug").
		Select("ug.group_id as group_id", "p.name as permission_name").
		Joins("JOIN group_permissions gp ON ug.group_id = gp.group_id AND ug.group_id = ?", groupID).
		Joins("JOIN permissions p ON gp.permission_id = p.id AND p.name = ?", permission).
		First(result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			_ = HSetWithExpiry(ctx, r.redisClient, cacheBucketKey, permission, fmt.Sprintf("%v", false))
			return false, nil
		}
		return false, err
	}

	_ = HSetWithExpiry(ctx, r.redisClient, cacheBucketKey, permission, fmt.Sprintf("%v", true))

	return true, nil
}

func (r *userGroupRepository) getHasPermissionCache(ctx context.Context, bucketKey string, permission string) (found bool, hasAccess bool, err error) {
	result, err := r.redisClient.HGet(ctx, bucketKey, permission).Result()
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return false, false, nil
		}
		return false, false, err
	}

	err = json.Unmarshal([]byte(strings.Replace(result, "\"", "", 2)), &hasAccess)
	if err != nil {
		return false, false, err
	}

	return true, hasAccess, nil
}

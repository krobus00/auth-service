package repository

import (
	"context"
	"errors"

	goredis "github.com/go-redis/redis/v8"
	"github.com/goccy/go-json"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type groupPermissionRepo struct {
	db          *gorm.DB
	redisClient *goredis.Client
}

func NewGroupPermissionRepository() model.GroupPermissionRepository {
	return new(groupPermissionRepo)
}

func (r *groupPermissionRepo) Create(ctx context.Context, data *model.GroupPermission) error {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"groupID":      data.GroupID,
		"permissionID": data.PermissionID,
	})

	db := utils.GetTxFromContext(ctx, r.db)

	err := db.WithContext(ctx).Create(data).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_ = DeleteByKeys(ctx, r.redisClient, model.GetGroupPermissionCacheKeys(data.GroupID, data.PermissionID))

	return nil
}

func (r *groupPermissionRepo) FindByGroupIDAndPermissionID(ctx context.Context, groupID string, permissionID string) (*model.GroupPermission, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"groupID":      groupID,
		"permissionID": permissionID,
	})

	db := utils.GetTxFromContext(ctx, r.db)
	groupPermission := new(model.GroupPermission)
	cacheKey := model.NewGroupPermissionCacheKeyByGroupIDAndPermissionID(groupID, permissionID)

	cachedData, err := Get(ctx, r.redisClient, cacheKey)
	if err != nil {
		logger.Error(err.Error())
	}
	err = json.Unmarshal(cachedData, &groupPermission)
	if err == nil {
		return groupPermission, nil
	}

	groupPermission = new(model.GroupPermission)

	err = db.WithContext(ctx).
		Where("group_id = ? AND permission_id = ?", groupID, permissionID).
		First(groupPermission).Error
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

	err = SetWithExpiry(ctx, r.redisClient, cacheKey, groupPermission)
	if err != nil {
		logger.Error(err.Error())
	}
	return groupPermission, nil
}

func (r *groupPermissionRepo) DeleteByGroupIDAndPermissionID(ctx context.Context, groupID string, permissionID string) error {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"groupID":      groupID,
		"permissionID": permissionID,
	})

	db := utils.GetTxFromContext(ctx, r.db)
	groupPermission := new(model.GroupPermission)

	err := db.WithContext(ctx).Clauses(clause.Returning{}).
		Where("group_id = ? AND permission_id = ?", groupID, permissionID).
		Delete(groupPermission).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_ = DeleteByKeys(ctx, r.redisClient, model.GetGroupPermissionCacheKeys(groupID, permissionID))

	return nil
}

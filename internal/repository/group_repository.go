package repository

import (
	"context"
	"errors"

	"github.com/goccy/go-json"

	goredis "github.com/go-redis/redis/v8"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type groupRepository struct {
	db          *gorm.DB
	redisClient *goredis.Client
}

func NewGroupRepository() model.GroupRepository {
	return new(groupRepository)
}

func (r *groupRepository) Create(ctx context.Context, group *model.Group) error {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"id":   group.ID,
		"name": group.Name,
	})

	db := utils.GetTxFromContext(ctx, r.db)

	err := db.WithContext(ctx).Create(group).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_ = DeleteByKeys(ctx, r.redisClient, model.GetGroupCacheKeys(group.ID, group.Name))

	return nil
}

func (r *groupRepository) FindByID(ctx context.Context, id string) (*model.Group, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	db := utils.GetTxFromContext(ctx, r.db)
	group := new(model.Group)
	cacheKey := model.NewGroupCacheKeyByID(id)
	cachedData, err := Get(ctx, r.redisClient, cacheKey)
	if err != nil {
		logger.Error(err.Error())
	}
	err = json.Unmarshal(cachedData, &group)
	if err == nil {
		return group, nil
	}

	group = new(model.Group)
	err = db.WithContext(ctx).Where("id = ?", id).First(group).Error
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

	err = SetWithExpiry(ctx, r.redisClient, cacheKey, group)
	if err != nil {
		logger.Error(err.Error())
	}
	return group, nil
}

func (r *groupRepository) FindByName(ctx context.Context, name string) (*model.Group, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"name": name,
	})

	db := utils.GetTxFromContext(ctx, r.db)
	group := new(model.Group)
	cacheKey := model.NewGroupCacheKeyByName(name)
	cachedData, err := Get(ctx, r.redisClient, cacheKey)
	if err != nil {
		logger.Error(err.Error())
	}
	err = json.Unmarshal(cachedData, &group)
	if err == nil {
		return group, nil
	}

	group = new(model.Group)
	err = db.WithContext(ctx).Where("name = ?", name).First(group).Error
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

	err = SetWithExpiry(ctx, r.redisClient, cacheKey, group)
	if err != nil {
		logger.Error(err.Error())
	}
	return group, nil
}

func (r *groupRepository) Update(ctx context.Context, group *model.Group) error {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"id": group.ID,
	})

	db := utils.GetTxFromContext(ctx, r.db)

	_ = DeleteByKeys(ctx, r.redisClient, model.GetGroupCacheKeys(group.ID, group.Name))

	err := db.WithContext(ctx).Updates(group).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *groupRepository) DeleteByID(ctx context.Context, id string) error {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"id": id,
	})

	db := utils.GetTxFromContext(ctx, r.db)
	group := new(model.Group)

	err := db.WithContext(ctx).Clauses(clause.Returning{}).
		Where("id = ?", id).
		Delete(group).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_ = DeleteByKeys(ctx, r.redisClient, model.GetGroupCacheKeys(group.ID, group.Name))

	return nil
}

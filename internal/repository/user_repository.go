package repository

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/goccy/go-json"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepository struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewUserRepository() model.UserRepository {
	return new(userRepository)
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := log.WithFields(log.Fields{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})

	db := utils.GetTxFromContext(ctx, r.db)

	err := db.WithContext(ctx).Create(user).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_ = DeleteByKeys(ctx, r.redisClient, model.GetUserCacheKeys(user.ID, user.Username, user.Email))

	return nil
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := log.WithFields(log.Fields{
		"id": id,
	})
	user := new(model.User)
	cacheKey := model.NewUserCacheKeyByID(id)
	cachedData, err := Get(ctx, r.redisClient, cacheKey)
	if err != nil {
		logger.Error(err.Error())
	}
	err = json.Unmarshal(cachedData, &user)
	if err == nil {
		return user, nil
	}

	user = new(model.User)

	db := utils.GetTxFromContext(ctx, r.db)
	err = db.WithContext(ctx).Take(user, "id = ?", id).Error

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

	err = SetWithExpiry(ctx, r.redisClient, cacheKey, user)
	if err != nil {
		logger.Error(err.Error())
	}

	return user, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := log.WithFields(log.Fields{
		"username": username,
	})
	user := new(model.User)

	cacheKey := model.NewUserCacheKeyByUsername(username)
	cachedData, err := Get(ctx, r.redisClient, cacheKey)
	if err != nil {
		logger.Error(err.Error())
	}
	err = json.Unmarshal(cachedData, &user)
	if err == nil {
		return user, nil
	}

	user = new(model.User)

	db := utils.GetTxFromContext(ctx, r.db)
	err = db.WithContext(ctx).Take(user, "username = ?", username).Error
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

	err = SetWithExpiry(ctx, r.redisClient, cacheKey, user)
	if err != nil {
		logger.Error(err.Error())
	}
	return user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	_, _, fn := utils.Trace()
	ctx, span := utils.NewSpan(ctx, fn)
	defer span.End()

	logger := log.WithFields(log.Fields{
		"email": email,
	})
	user := new(model.User)
	cacheKey := model.NewUserCacheKeyByEmail(email)
	cachedData, err := Get(ctx, r.redisClient, cacheKey)
	if err != nil {
		logger.Error(err.Error())
	}
	err = json.Unmarshal(cachedData, &user)
	if err == nil {
		return user, nil
	}

	user = new(model.User)

	db := utils.GetTxFromContext(ctx, r.db)
	err = db.WithContext(ctx).Take(user, "email = ?", email).Error
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

	err = SetWithExpiry(ctx, r.redisClient, cacheKey, user)
	if err != nil {
		logger.Error(err.Error())
	}
	return user, nil
}

func (r *userRepository) UpdateByID(ctx context.Context, id string) (*model.User, error) {
	return nil, errors.New("unimplemented")
}

func (r *userRepository) DeleteByID(ctx context.Context, id string) error {
	return errors.New("unimplemented")
}

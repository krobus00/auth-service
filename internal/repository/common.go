package repository

import (
	"context"
	"time"

	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v8"
)

func HSetWithExpiry(ctx context.Context, redisClient *redis.Client, bucketCacheKey string, field string, data any) error {
	cacheData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = redisClient.HSet(ctx, bucketCacheKey, field, cacheData).Err()
	if err != nil {
		return err
	}
	err = redisClient.ExpireNX(ctx, bucketCacheKey, time.Minute*15).Err()
	if err != nil {
		return err
	}
	return nil
}

func SetWithExpiry(ctx context.Context, redisClient *redis.Client, cacheKey string, data any) error {

	cacheData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = redisClient.Set(ctx, cacheKey, cacheData, time.Minute*15).Err()
	if err != nil {
		return err
	}
	return nil
}

func Get(ctx context.Context, redisClient *redis.Client, cacheKey string) (data []byte, err error) {
	cachedData, err := redisClient.Get(ctx, cacheKey).Bytes()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return cachedData, nil
}

func DeleteByKeys(ctx context.Context, redisClient *redis.Client, cacheKeys []string) error {
	for _, cacheKey := range cacheKeys {
		err := redisClient.Del(ctx, cacheKey).Err()
		if err != nil && err != redis.Nil {
			logrus.WithField("cacheKey", cacheKey).Error(err.Error())
			return err
		}
	}
	return nil
}

func HGet(ctx context.Context, redisClient *redis.Client, bucketCacheKey string, field string) (data []byte, err error) {
	cachedData, err := redisClient.HGet(ctx, bucketCacheKey, field).Bytes()
	if err != nil {
		return nil, err
	}
	return cachedData, nil
}

package infrastructure

import (
	goredis "github.com/go-redis/redis/v8"
	"github.com/krobus00/auth-service/internal/config"
)

// NewRedisClient create redis db connection
func NewRedisClient() (*goredis.Client, error) {
	redisURL, err := goredis.ParseURL(config.RedisCacheHost())
	if err != nil {
		return nil, err
	}
	rdb := goredis.NewClient(redisURL)
	return rdb, nil
}

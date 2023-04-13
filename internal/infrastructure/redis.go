package infrastructure

import (
	"github.com/go-redis/redis/extra/redisotel/v8"
	goredis "github.com/go-redis/redis/v8"
	"github.com/krobus00/auth-service/internal/config"
)

// NewRedisClient create redis db connection.
func NewRedisClient() (*goredis.Client, error) {
	redisURL, err := goredis.ParseURL(config.RedisCacheHost())
	if err != nil {
		return nil, err
	}
	redisOpts := &goredis.Options{
		Network:      redisURL.Network,
		Addr:         redisURL.Addr,
		DB:           redisURL.DB,
		Username:     redisURL.Username,
		Password:     redisURL.Password,
		DialTimeout:  config.RedisDialTimeout(),
		WriteTimeout: config.RedisWriteTimeout(),
		ReadTimeout:  config.RedisReadTimeout(),
	}
	rdb := goredis.NewClient(redisOpts)
	rdb.AddHook(redisotel.NewTracingHook())

	return rdb, nil
}

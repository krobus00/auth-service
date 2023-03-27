package repository

import (
	"errors"

	goredis "github.com/go-redis/redis/v8"
)

// InjectRedisClient :nodoc:
func (r *tokenRepository) InjectRedisClient(client *goredis.Client) error {
	if client == nil {
		return errors.New("invalid redis client")
	}
	r.redisClient = client
	return nil
}

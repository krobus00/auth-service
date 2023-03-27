package repository

import (
	"errors"

	goredis "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func (r *groupPermissionRepo) InjectDB(db *gorm.DB) error {
	if db == nil {
		return errors.New("invalid db")
	}
	r.db = db
	return nil
}

func (r *groupPermissionRepo) InjectRedisClient(client *goredis.Client) error {
	if client == nil {
		return errors.New("invalid redis client")
	}
	r.redisClient = client
	return nil
}

package redis

import (
	"time"

	"github.com/xiaoenai/tp-micro/v6/model/redis"
)

var (
	redisClient *redis.Client
	cacheExpire time.Duration
)

func Init(redisConfig *redis.Config, _cacheExpire time.Duration) error {
	cacheExpire = _cacheExpire
	var err error
	if redisConfig != nil {
		redisClient, err = redis.NewClient(redisConfig)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetRedis returns the redis client.
func GetRedis() *redis.Client {
	return redisClient
}

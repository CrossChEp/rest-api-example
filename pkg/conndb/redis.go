package conndb

import (
	"github.com/redis/go-redis/v9"
	"rest-api-example/config"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Cache.Session.Host,
		Password: cfg.Cache.Session.Password,
		DB:       cfg.Cache.Session.DB,
	})
}

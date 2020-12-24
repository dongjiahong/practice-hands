package gredis

import (
	"fmt"

	"github.com/go-redis/redis"

	conf "gin-machinery/config"
)

// Client Global client
var Client *redis.Client

// InitRedisClient Establish a connection pool
func InitRedisClient() {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Cfg.TaskRedisHost, conf.Cfg.TaskRedisPort),
		Password: conf.Cfg.TaskRedisPassword,
		DB:       conf.Cfg.TaskRedisDB,
		PoolSize: conf.Cfg.TaskRedisPoolSize,
	})
}

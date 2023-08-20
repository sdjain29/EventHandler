package init

import (
	"EventHandler/config"
	"context"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func SetupRedis() *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{Addr: config.Config.Redis.RedisHost[0], Password: config.Config.Redis.RedisPassword, DB: config.Config.Redis.RedisDb})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return RedisClient
}

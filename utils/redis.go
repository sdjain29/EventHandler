package utils

import (
	core "EventHandler/init"
	"EventHandler/logger"
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RedisSetInt(c *gin.Context, key string, value int, expiry time.Duration) {
	core.RedisClient.Set(context.Background(), key, fmt.Sprint(value), expiry)
}

func RedisSetZAdd(c *gin.Context, key string, score float64, value interface{}) {
	result := core.RedisClient.ZAdd(context.Background(), key, redis.Z{Score: score, Member: value})
	logger.Info(c, result)
}

func RedisSetZAddWithoutContext(key string, score float64, value interface{}) {
	core.RedisClient.ZAdd(context.Background(), key, redis.Z{Score: score, Member: value})
}

func RedisSetZDel(key string, value interface{}) {
	core.RedisClient.ZRem(context.Background(), key, value)
}

func RedisZGet(key, max string) []string {
	var result []string
	result, _ = core.RedisClient.ZRangeByScore(context.Background(), key, &redis.ZRangeBy{Min: "-inf", Max: max, Offset: 0, Count: 100}).Result()
	return result
}

func RedisGet(key string) (string, error) {
	return core.RedisClient.Get(context.Background(), key).Result()
}

func RedisSet(key string, value string, expiry time.Duration) {
	core.RedisClient.Set(context.Background(), key, value, expiry)
}

func RedisDel(c *gin.Context, key string) {
	core.RedisClient.Del(context.Background(), key)
}

func RedisPing(c *gin.Context) {
	_, err := core.RedisClient.Ping(context.Background()).Result()
	CheckErrAndCrash(c, 500, err, fmt.Sprint(err))
}

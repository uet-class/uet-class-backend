package database

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/uet-class/uet-class-backend/config"
)

var rdb *redis.Client
var ctx context.Context

func InitRedis() {
	redisConfig := config.GetConfig()
	ctx = context.Background()

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.GetString("REDIS_HOST"), redisConfig.GetString("REDIS_PORT")),
		Password: redisConfig.GetString("REDIS_PASSWORD"),
		DB:       redisConfig.GetInt("REDIS_DATABASE"),
	})
}

func GetRedis() *redis.Client {
	return rdb
}

func GetRedisContext() context.Context {
	return ctx
}

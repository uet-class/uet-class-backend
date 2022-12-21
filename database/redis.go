package database

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx context.Context

func InitRedis() {
	ctx = context.Background()

	databaseId, _ := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       databaseId,
	})
}

func GetRedis() *redis.Client {
	return rdb
}

func GetRedisContext() context.Context {
	return ctx
}

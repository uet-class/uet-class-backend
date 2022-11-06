package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx context.Context

func InitRedis() {
	// config := config.GetConfig()
	ctx = context.Background()

	opt, err := redis.ParseURL("redis://:@localhost:6379/0?dial_timeout=5s")
	if err != nil {
		log.Fatal(err)
	}

	rdb = redis.NewClient(opt)
}

func GetRedis() *redis.Client {
	return rdb
}

func GetRedisContext() context.Context {
	return ctx
}

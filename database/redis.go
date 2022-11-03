package database

import (
	"context"
	"fmt"

	"github.com/uet-class/uet-class-backend/config"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func InitRedis() {
	config := config.GetConfig()

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.GetString("REDIS_HOST"), config.GetString("REDIS_PORT")),
		Password: config.GetString("REDIS_PASSWORD"),
		DB:       config.GetInt("REDIS_DATABASE"),
	})

	err := rdb.Set(ctx, "authorized", "true", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "authorized").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func GetRedis() *redis.Client {
	return rdb
}

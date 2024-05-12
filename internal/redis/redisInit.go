package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Rdb *redis.Client

func RedisInit(){
	Rdb = redis.NewClient(&redis.Options{
		Addr:	 "127.0.0.1:6379",
		Password: "Pa$$w0rd",
		DB:		  0,
	})
}
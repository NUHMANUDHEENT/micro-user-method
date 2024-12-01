package helper

import (
	"os"

	"github.com/go-redis/redis"
)

var Rdb = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_PORT"),
	Password: "",
	DB:       0,
})


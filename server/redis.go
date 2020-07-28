package server

import (
	"os"

	"github.com/go-redis/redis/v8"
)

func connectRedis() *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})
	return db
}

package server

import "github.com/go-redis/redis"

func connectRedis() *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return db
}

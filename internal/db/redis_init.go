package db

import (
	"context"
	c "go-sso/internal/config"
	"log"

	"github.com/redis/go-redis/v9"
)

var REDIS_CLIENT *redis.Client

func InitRedis() {
	Redis_URL := c.GetConfig().Redis_URL
	REDIS_CLIENT = redis.NewClient(&redis.Options{
		Addr: Redis_URL, // Adjust host and port as needed
		DB:   0,         // Default DB
	})

	if err := REDIS_CLIENT.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

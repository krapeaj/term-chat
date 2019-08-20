package main

import (
	"chat-server/cache/redis"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "chat-server", 0)

	redisService := redis.NewRedisService(":6379", "", 0, logger)

	handler := NewHandler(redisService, logger)
	handler.ServeHTTP(":3000")
}


package main

import (
	"chat-server/cache/redis"
	"chat-server/service"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "chat-server", 0)

	redisService := redis.NewRedisService(":6379", "", 0, logger)

	sessionService := service.NewSessionService(redisService, logger)
	chatService := service.NewChatService(redisService, logger)

	handler := NewHandler(sessionService, chatService, logger)
	handler.ServeHTTP(":3000")
}


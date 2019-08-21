package main

import (
	"chat-server/cache/redis"
	"chat-server/service"
	"log"
	"os"

	r "github.com/go-redis/redis"
)

func main() {
	logger := log.New(os.Stdout, "chat-server: ", 0)

	client := r.NewClient(&r.Options{
		Addr:     "0.0.0.0:7379",
		Password: "",
		DB:       0,
	})
	redisService := redis.NewRedisService(client, logger)

	sessionService := service.NewSessionService(redisService, logger)
	chatService := service.NewChatService(redisService, logger)

	handler := NewHandler(sessionService, chatService, logger)
	handler.ServeHTTP(":3000")
}
